package go_playground_ws

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var (
	empty = []byte{}
)

type client struct {
	hub     *Hub
	conn    *websocket.Conn
	readFn  ReadFn
	closeFn CloseFn
	send    chan interface{}
	close   chan struct{}
	once    sync.Once
}

// ReadFn is the function to be called directly after a read operation
// by the default client.
type ReadFn func(interface{})

// CloseFn is the function to be after the connectin is closed
// and all cleaning up has happened
type CloseFn func( /* interface{} */ )

// NewClient returns a new instance of the default client.
func NewClient(h *Hub, conn *websocket.Conn, readFn ReadFn, closeFn CloseFn) Client {
	return &client{
		hub:     h,
		conn:    conn,
		readFn:  readFn,
		closeFn: closeFn,
		close:   make(chan struct{}),
		send:    make(chan interface{}),
	}
}

func (c *client) Listen() error {
	go c.write()
	return c.read()
}

// Close closes the connection
func (c *client) Close() {
	c.once.Do(func() {
		c.conn.Close()
		close(c.close)
		c.hub.Remove(c)
		c.closeFn()
	})
}

func (c *client) read() error {

	defer func() {
		c.Close()
	}()

	c.conn.SetReadLimit(c.hub.ReadLimit())

	err := c.conn.SetReadDeadline(c.hub.ReadDeadline())
	if err != nil {
		log.Printf("read deadline reached '%s'\n", err)
		return err
	}

	c.conn.SetPongHandler(func(string) error {
		err := c.conn.SetReadDeadline(c.hub.ReadDeadline())
		if err != nil {
			log.Printf("error in pong handler '%s'\n", err)
		}
		return err
	})

	for {
		var msg interface{}
		err := c.conn.ReadJSON(&msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v\n", err)
			}
			/* if closeErr, ok := err.(*websocket.CloseError); ok {
				closeErr.Code
			} */
			return err
		}
		c.readFn(msg)
	}
}

func (c *client) write() {

	ticker := time.NewTicker(c.hub.PingInterval())
	defer func() {
		ticker.Stop()
		close(c.send)
		c.Close()
	}()

FOR:
	for {

		select {
		case <-c.close:
			break FOR
		case <-ticker.C:

			err := c.conn.SetWriteDeadline(c.hub.WriteDeadline())
			if err != nil {
				log.Printf("error setting write deadline '%s'\n", err)
				break FOR
			}

			if err := c.conn.WriteMessage(websocket.PingMessage, empty); err != nil {
				log.Printf("error sending ping message '%s'\n", err)
				break FOR
			}

		case msg := <-c.send:

			err := c.conn.SetWriteDeadline(c.hub.WriteDeadline())
			if err != nil {
				log.Printf("error setting write deadline '%s'\n", err)
				break FOR
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("error retrieving next writer '%s'\n", err)
				break FOR
			}
			err = json.NewEncoder(w).Encode(msg)
			if err != nil {
				log.Printf("error writing message '%s'\n", err)
			}
			/* // Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			} */

			if err = w.Close(); err != nil {
				break FOR
			}
		}
	}
}

func (c *client) Write(msg interface{}) {
	c.send <- msg
	/* select {
	case <-c.close:
	default:
		c.send <- msg
	} */
}
