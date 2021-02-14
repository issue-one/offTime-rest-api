package ws

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	ws "github.com/issue-one/offTime-rest-api/internal/delivery/ws/go_playground_ws"
	uuid "github.com/satori/go.uuid"
)

// Message struct codifes the format for sending messages
type Message struct {
	Event string      `json:"event,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// ErrorData is the form Data from Message takes when replying erronous responses
type ErrorData struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// ErrRoomNotFound is obivous
var ErrRoomNotFound = errors.New("ws: room not found")

// ConnectedEvent is an event emitted whenever a new client joins.
var ConnectedEvent = "connected"

// CloseEvent is an event emitted whenever a new client line closes
var CloseEvent = "close"

// ListenerFn is the interface for event listeners
type ListenerFn func(client *Client, msg interface{})

// Hub is the what manages the websockets and provides room functionality
type Hub struct {
	ws.Hub
	mutex            sync.RWMutex
	listeners        map[string]ListenerFn
	Rooms            map[string]map[string]*Client
	MesssageListener func(client *Client, event string, msg interface{})
}

// Emit emits the specified event to all the sockets found in the given room
func (hub *Hub) Emit(roomName string, event string, msg interface{}) error {
	hub.mutex.RLock()
	defer hub.mutex.RUnlock()
	if room, ok := hub.Rooms[roomName]; ok {
		for _, client := range room {
			client.Emit(event, msg)
		}
		return nil
	}
	return ErrRoomNotFound
}

// GetListner find the listener registered for the specified event
func (hub *Hub) GetListner(eventName string) (listner ListenerFn, ok bool) {
	listner, ok = hub.listeners[eventName]
	return
}

// Listen registers a listener for the specifed event
func (hub *Hub) Listen(eventName string, listner ListenerFn) {
	hub.listeners[eventName] = listner
}

func (hub *Hub) handleMessage(client *Client, msg interface{}) {
	log.Println(fmt.Sprintf("message recieved: %v", msg))
	msgMap, ok := msg.(map[string]interface{})
	if !ok {
		client.Emit("invalidProtocol", "Requires JSON with `{ event: string, data: any }` format.")
		return
	}
	eventInterface, ok := msgMap["event"]
	if !ok {
		client.Emit("invalidProtocol", "Requires JSON with `{ event: string, data: any }` format.")
		return
	}
	event, ok := eventInterface.(string)
	if !ok {
		client.Emit("invalidProtocol", "Requires JSON with `{ event: string, data: any }` format.")
		return
	}
	message := Message{
		Event: event,
		Data:  msgMap["data"],
	}

	eventName := message.Event
	listner, ok := hub.GetListner(eventName)
	if ok {
		listner(client, message.Data)
	}
	if hub.MesssageListener != nil {
		hub.MesssageListener(client, eventName, message.Data)
	}
}

func (hub *Hub) handleClose(client *Client) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	delete(hub.Rooms, client.ID)
	for _, rooms := range hub.Rooms {
		delete(rooms, client.ID)
	}
	closeListener, ok := hub.GetListner(CloseEvent)
	if ok {
		closeListener(client, nil)
	}
}

func (hub *Hub) joinRoom(c *Client, roomName string) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	if room, ok := c.hub.Rooms[roomName]; ok {
		room[c.ID] = c
		c.Rooms[roomName] = struct{}{}
	} else {
		c.hub.Rooms[roomName] = map[string]*Client{c.ID: c}
	}
}

func (hub *Hub) leaveRoom(c *Client, roomName string) {
	hub.mutex.Lock()
	defer hub.mutex.Unlock()
	if room, ok := c.hub.Rooms[roomName]; ok {
		delete(room, c.ID)
	}
	delete(c.Rooms, roomName)
}

// NewHub returns a new WsHub instance and the handler that accepts and upgrades
// connections to WebSockets
func NewHub() (*Hub, http.Handler) {

	var err error

	if err != nil {
		log.Fatal(err)
	}

	hub := &Hub{
		Rooms:     make(map[string]map[string]*Client),
		listeners: make(map[string]ListenerFn),
	}
	hub.Hub = *ws.New(upgrader, func(h *ws.Hub, conn *websocket.Conn, r *http.Request) ws.Client {
		id := uuid.NewV4().String()
		client := &Client{
			ID:    id,
			Rooms: map[string]struct{}{id: {}},
		}
		client.Client = ws.NewClient(h, conn,
			func(b interface{}) { hub.handleMessage(client, b) },
			func() { hub.handleClose(client) },
		)
		hub.Rooms[id] = map[string]*Client{id: client}

		connectedListener, ok := hub.GetListner(ConnectedEvent)
		if ok {
			connectedListener(client, nil)
		}

		return client
	})

	return hub, http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		err := hub.Accept(rw, r)
		if err != nil {
			log.Printf("error accepting WebSocket connection: '%s'\n", err)
		}
	})
}

// Client represents a single client connection
type Client struct {
	ws.Client
	ID    string
	Rooms map[string]struct{}
	hub   *Hub
}

// Join adds the client into the specified room
func (c *Client) Join(roomName string) {
	c.hub.joinRoom(c, roomName)
}

// Leave removes the client from the specified room
func (c *Client) Leave(roomName string) {
	c.hub.leaveRoom(c, roomName)
}

// Emit send the specified message under the given event
func (c *Client) Emit(event string, msg interface{}) {
	c.Write(Message{
		Event: event,
		Data:  msg,
	})
}

/*
func copyConn(readConn, writeConn *websocket.Conn, readClosed, writeClosed chan struct{}) {
	var rerr error
	for {
		var r io.Reader
		var messageType int

		messageType, r, rerr = readConn.NextReader()
		if rerr != nil {
			break
		}
		w, err := writeConn.NextWriter(messageType)
		if err != nil {
			break
		}
		if _, err := io.Copy(w, r); err != nil {
			break
		}
		if err := w.Close(); err != nil {
			break
		}
	}

	// Close the reading connection. If we broke out of the loop because of a
	// normal close, then NextReader echoed the close message and we should now
	// close the connection.  If it's an abnormal close, then we should give up
	// and close the connection.
	readConn.Close()

	// Tell the other goroutine that readConn was closed.
	close(readClosed)

	// Did we break out of the loop because we received a close message?
	if e, ok := rerr.(*websocket.CloseError); ok && e.Code != websocket.CloseAbnormalClosure {

		// Forward the close message to writeConn.
		var m []byte
		if e.Code != websocket.CloseNoStatusReceived {
			m = websocket.FormatCloseMessage(e.Code, e.Text)
		}
		err := writeConn.WriteMessage(websocket.CloseMessage, m)

		// Did we successfully send the close message?
		if err == nil {
			// Wait with timeout for the other goroutine to handle the handshake.
			select {
			case <-writeClosed:
				// The other goroutine closed writeConn.
				return
			case <-time.After(10 * time.Second):
			}
		}
	}

	// A blocked reader returns with an error when the connection is closed.
	writeConn.Close()
}
*/
