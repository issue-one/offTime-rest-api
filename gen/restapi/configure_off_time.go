// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	uuid "github.com/satori/go.uuid"

	"github.com/issue-one/offTime-rest-api/gen/restapi/operations"
	"github.com/issue-one/offTime-rest-api/internal/delivery/ws"
	"github.com/issue-one/offTime-rest-api/internal/repositories"
	"github.com/issue-one/offTime-rest-api/internal/repositories/mock"
)

//go:generate swagger generate server --skip-main --target ../../gen --name OffTime --spec ../../swagger.yaml --principal interface{}

func configureFlags(api *operations.OffTimeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }

}

func configureAPI(api *operations.OffTimeAPI) http.Handler {

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.New(os.Stdout, "", log.Lmicroseconds|log.Lshortfile).Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// operations.PutUsersUsernamePictureMaxParseMemory = 32 << 20

	userRepo := mock.NewMockUserRepository()
	{
		var ok bool
		imageStoragePath, ok = os.LookupEnv("IMAGE_STORAGE_PATH")
		if !ok {
			imageStoragePath = "data/images/"
		}
		imageServingRoute, ok = os.LookupEnv("IMAGE_SERVING_ROUTE")
		if !ok {
			imageServingRoute = "/images/"
		}
	}
	roomRepo := mock.NewMockRoomRepository()

	// USER handlers
	api.PutUsersUsernameHandler = operations.PutUsersUsernameHandlerFunc(
		func(params operations.PutUsersUsernameParams) middleware.Responder {
			ctx := params.HTTPRequest.Context()
			// TODO: check if min length is auto enforced
			/* if len(params.Body.Password.String()) < 8 {
				operations.NewPutUsersUsernameBadRequest().WithPayload(
					&operations.PutUsersUsernameBadRequestBody{
						Message: ,
					}
				)
			} */

			occupied, err := userRepo.IsUsernameOccupied(ctx, params.Username)
			if err != nil {
				return operations.NewPatchUsersUsernameInternalServerError().WithPayload(
					&operations.PatchUsersUsernameInternalServerErrorBody{Message: err.Error()},
				)
			}
			if occupied {
				return operations.NewPutUsersUsernameConflict().WithPayload(
					&operations.PutUsersUsernameConflictBody{Field: "Username"},
				)
			}
			occupied, err = userRepo.IsEmailOccupied(ctx, params.Body.Email.String(), "")
			if err != nil {
				return operations.NewPutUsersUsernameInternalServerError().WithPayload(
					&operations.PutUsersUsernameInternalServerErrorBody{Message: err.Error()},
				)
			}
			if occupied {
				return operations.NewPutUsersUsernameConflict().WithPayload(
					&operations.PutUsersUsernameConflictBody{Field: "Email"},
				)
			}
			user, err := userRepo.CreateUser(ctx, params.Username, params.Body)
			if err != nil {
				return operations.NewPutUsersUsernameInternalServerError().WithPayload(
					&operations.PutUsersUsernameInternalServerErrorBody{Message: err.Error()},
				)
			}
			if user.PictureURL != "" {
				user.PictureURL = urlFromFilename(user.PictureURL)
			}
			return operations.NewPutUsersUsernameOK().WithPayload(user)
		})

	api.GetUsersUsernameHandler = operations.GetUsersUsernameHandlerFunc(
		func(params operations.GetUsersUsernameParams) middleware.Responder {
			user, err := userRepo.GetUser(params.HTTPRequest.Context(), params.Username)
			switch err {
			case nil:

				if user.PictureURL != "" {
					user.PictureURL = urlFromFilename(user.PictureURL)
				}
				return operations.NewGetUsersUsernameOK().WithPayload(user)
			case repositories.ErrUserNotFound:
				return operations.NewGetUsersUsernameNotFound()
			default:
				return operations.NewGetUsersUsernameInternalServerError().WithPayload(
					&operations.GetUsersUsernameInternalServerErrorBody{Message: err.Error()},
				)
			}
		})

	api.GetUsersHandler = operations.GetUsersHandlerFunc(
		func(params operations.GetUsersParams) middleware.Responder {
			// TODO: check if min limit and offset is auto enforced
			users, totalCount, err := userRepo.GetAllUsers(params.HTTPRequest.Context(), *params.Limit, *params.Offset)
			if err != nil {
				return operations.NewGetUsersInternalServerError().WithPayload(
					&operations.GetUsersInternalServerErrorBody{Message: err.Error()},
				)
			}
			for _, user := range users {
				if user.PictureURL != "" {
					user.PictureURL = urlFromFilename(user.PictureURL)
				}
			}
			return operations.NewGetUsersOK().WithPayload(
				&operations.GetUsersOKBody{
					Items:      users,
					TotalCount: int64(totalCount),
				},
			)
		})

	api.PatchUsersUsernameHandler = operations.PatchUsersUsernameHandlerFunc(
		func(params operations.PatchUsersUsernameParams) middleware.Responder {
			ctx := params.HTTPRequest.Context()
			// TODO: check if min length is auto enforced
			/* if len(params.Body.Password.String()) < 8 {
				operations.NewPutUsersUsernameBadRequest().WithPayload(
					&operations.PutUsersUsernameBadRequestBody{
						Message: ,
					}
				)
			} */
			if params.Body.Email != "" {
				occupied, err := userRepo.IsEmailOccupied(ctx, params.Body.Email.String(), "")
				if err != nil {
					return operations.NewPatchUsersUsernameInternalServerError().WithPayload(
						&operations.PatchUsersUsernameInternalServerErrorBody{Message: err.Error()},
					)
				}
				if occupied {
					return operations.NewPatchUsersUsernameConflict().WithPayload(
						&operations.PatchUsersUsernameConflictBody{Field: "Email"},
					)
				}
			}
			user, err := userRepo.UpdateUser(ctx, params.Username, params.Body)
			switch err {
			case nil:
				if user.PictureURL != "" {
					user.PictureURL = urlFromFilename(user.PictureURL)
				}
				return operations.NewPatchUsersUsernameOK().WithPayload(user)
			case repositories.ErrUserNotFound:
				return operations.NewPatchUsersUsernameNotFound()
			default:
				return operations.NewPatchUsersUsernameInternalServerError().WithPayload(
					&operations.PatchUsersUsernameInternalServerErrorBody{Message: err.Error()},
				)
			}
		})

	api.PutUsersUsernamePictureHandler = operations.PutUsersUsernamePictureHandlerFunc(
		func(params operations.PutUsersUsernamePictureParams) middleware.Responder {
			// save the image temporarily
			tempFile, extension, err := saveImageFromRequest(params.Image)
			defer func() {
				_ = os.Remove(tempFile.Name())
				_ = tempFile.Close()
			}()
			switch err {
			case nil:
			case errReadingFromImage:
			case errUnacceptedType:
				return operations.NewPutUsersUsernamePictureBadRequest().WithPayload(
					&operations.PutUsersUsernamePictureBadRequestBody{
						Message: err.Error(),
					},
				)
			default:
				return operations.NewPutUsersUsernamePictureInternalServerError().WithPayload(
					&operations.PutUsersUsernamePictureInternalServerErrorBody{
						Message: err.Error(),
					},
				)
			}
			// update user
			user, err := userRepo.SetImage(
				params.HTTPRequest.Context(),
				params.Username,
				generateFileNameForStorage(params.Username+"."+extension, "user"),
			)
			switch err {
			case nil:
				err = saveTempImagePermanentlyToPath(tempFile, imageStoragePath+user.PictureURL)
				if err != nil {
					return operations.NewPutUsersUsernamePictureInternalServerError().WithPayload(
						&operations.PutUsersUsernamePictureInternalServerErrorBody{
							Message: err.Error(),
						},
					)
				}
				return operations.NewPutUsersUsernamePictureOK().WithPayload(
					urlFromFilename(user.PictureURL),
				)
			case repositories.ErrUserNotFound:
				return operations.NewPutUsersUsernamePictureNotFound()
			default:
				return operations.NewPutUsersUsernamePictureInternalServerError().WithPayload(
					&operations.PutUsersUsernamePictureInternalServerErrorBody{
						Message: err.Error(),
					},
				)
			}
		})

	api.DeleteUsersUsernameHandler = operations.DeleteUsersUsernameHandlerFunc(
		func(params operations.DeleteUsersUsernameParams) middleware.Responder {
			err := userRepo.DeleteUser(params.HTTPRequest.Context(), params.Username)
			if err != nil {
				return operations.NewDeleteUsersUsernameInternalServerError().WithPayload(
					&operations.DeleteUsersUsernameInternalServerErrorBody{Message: err.Error()},
				)
			}
			return operations.NewDeleteUsersUsernameOK()
		})

	// USAGE handlers
	if api.GetUsersUsernameUsageHistoryHandler == nil {
		api.GetUsersUsernameUsageHistoryHandler = operations.GetUsersUsernameUsageHistoryHandlerFunc(func(params operations.GetUsersUsernameUsageHistoryParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetUsersUsernameUsageHistory has not yet been implemented")
		})
	}

	if api.PostUsersUsernameUsageHistoryHandler == nil {
		api.PostUsersUsernameUsageHistoryHandler = operations.PostUsersUsernameUsageHistoryHandlerFunc(func(params operations.PostUsersUsernameUsageHistoryParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.PostUsersUsernameUsageHistory has not yet been implemented")
		})
	}

	if api.DeleteUsersUsernameUsageHistoryHandler == nil {
		api.DeleteUsersUsernameUsageHistoryHandler = operations.DeleteUsersUsernameUsageHistoryHandlerFunc(func(params operations.DeleteUsersUsernameUsageHistoryParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteUsersUsernameUsageHistory has not yet been implemented")
		})
	}

	// ROOM handlers
	api.GetRoomsHandler = operations.GetRoomsHandlerFunc(
		func(params operations.GetRoomsParams) middleware.Responder {
			// TODO: check if min limit and offset is auto enforced
			rooms, totalCount, err := roomRepo.GetAllRooms(params.HTTPRequest.Context(), *params.Limit, *params.Offset)
			if err != nil {
				return operations.NewGetRoomsInternalServerError().WithPayload(
					&operations.GetRoomsInternalServerErrorBody{
						Message: err.Error(),
					},
				)
			}
			return operations.NewGetRoomsOK().WithPayload(
				&operations.GetRoomsOKBody{
					Items:      rooms,
					TotalCount: int64(totalCount),
				},
			)
		})

	api.GetRoomsRoomIDHandler = operations.GetRoomsRoomIDHandlerFunc(
		func(params operations.GetRoomsRoomIDParams) middleware.Responder {
			room, err := roomRepo.GetRoom(params.HTTPRequest.Context(), params.RoomID)
			switch err {
			case nil:
				return operations.NewGetRoomsRoomIDOK().WithPayload(room)
			case repositories.ErrRoomNotFound:
				return operations.NewGetRoomsRoomIDNotFound()
			default:
				return operations.NewGetRoomsRoomIDInternalServerError().WithPayload(
					&operations.GetRoomsRoomIDInternalServerErrorBody{
						Message: err.Error(),
					},
				)
			}
		},
	)

	api.GetUsersUsernameRoomHistoryHandler = operations.GetUsersUsernameRoomHistoryHandlerFunc(
		func(params operations.GetUsersUsernameRoomHistoryParams) middleware.Responder {
			user, err := userRepo.GetUser(params.HTTPRequest.Context(), params.Username)
			switch err {
			case nil:
			case repositories.ErrUserNotFound:
				return operations.NewGetUsersUsernameRoomHistoryNotFound().WithPayload(
					&operations.GetUsersUsernameRoomHistoryNotFoundBody{
						Entity:    "User",
						Identifer: params.Username,
					},
				)
			default:
				return operations.NewGetUsersUsernameRoomHistoryInternalServerError().WithPayload(
					&operations.GetUsersUsernameRoomHistoryInternalServerErrorBody{Message: err.Error()},
				)
			}
			rooms, err := roomRepo.GetMultipleRooms(params.HTTPRequest.Context(), user.RoomHistory)
			switch err {
			case nil:
				return operations.NewGetUsersUsernameRoomHistoryOK().WithPayload(rooms)
			case repositories.ErrRoomNotFound:
				return operations.NewGetUsersUsernameRoomHistoryNotFound().WithPayload(
					&operations.GetUsersUsernameRoomHistoryNotFoundBody{
						Entity: "Room",
					},
				)
			default:
				return operations.NewGetUsersUsernameRoomHistoryInternalServerError().WithPayload(
					&operations.GetUsersUsernameRoomHistoryInternalServerErrorBody{Message: err.Error()},
				)
			}
		})

	hub, wsHandler := ws.NewHub()

	var wsMiddleware = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/ws":
				wsHandler.ServeHTTP(w, r)
				// api.Logger(fmt.Sprintf("wsMiddleware: %v", hub.Rooms))
				break
			case "/":
				s := struct {
					Addr string
				}{
					Addr: "localhost:8080",
				}
				err := doc.ExecuteTemplate(w, "root", s)
				if err != nil {
					log.Fatal(err)
				}
				break
			default:
				next.ServeHTTP(w, r)
			}
		})
	}

	hub.Listen(ws.ConnectedEvent, func(client *ws.Client, msg interface{}) {
		api.Logger(fmt.Sprintf("client connected at id: %v", client.ID))
	})
	hub.Listen(ws.CloseEvent, func(client *ws.Client, msg interface{}) {
		api.Logger(fmt.Sprintf("client at id: %v left", client.ID))
	})

	hub.MesssageListener = func(client *ws.Client, event string, msg interface{}) {
		api.Logger(fmt.Sprintf("client at id: (%v) sent message: %v", client.ID, msg))
	}

	hub.Listen("echo", func(client *ws.Client, msg interface{}) {
		api.Logger(fmt.Sprintf("echo: %v", msg))
		client.Emit("echo", msg)
	})
	/* gosf.Listen("createRoom", func(client *gosf.Client, request *gosf.Request) *gosf.Message {
		createRoomMessage := struct {
			username string `json:"username,omitempty"`
			roomName string `json:"room_name,omitempty"`
		}{}
		type response struct {
			code    int         `json:"code,omitempty"`
			message interface{} `json:"message,omitempty"`
		}
		gosf.MapToStruct(request.Message.Body, createRoomMessage)
		if createRoomMessage.username == "" {
			return gosf.NewFailureMessage("errorCreatingRoom", gosf.StructToMap(response{
				code:    400,
				message: "No username field found in request.",
			}))
		}
		if createRoomMessage.username == "" {
			return gosf.NewFailureMessage("errorCreatingRoom", gosf.StructToMap(response{
				code:    400,
				message: "No username field found in request.",
			}))
		}
		found, err := userRepo.IsUsernameOccupied(context.TODO(), createRoomMessage.username)
		if err != nil {
			return gosf.NewFailureMessage("errorCreatingRoom", gosf.StructToMap(response{
				code:    500,
				message: "Internal server error: " + err.Error(),
			}))
		}
		if !found {
			return gosf.NewFailureMessage("errorCreatingRoom", gosf.StructToMap(response{
				code:    404,
				message: "No user found under given username: " + createRoomMessage.username,
			}))
		}

		room, err := roomRepo.CreateRoom(context.TODO(), createRoomMessage.username, createRoomMessage.roomName)
		if err != nil {
			return gosf.NewFailureMessage("errorCreatingRoom", gosf.StructToMap(response{
				code:    500,
				message: "Internal server error: " + err.Error(),
			}))
		}
		client.Join(room.ID.String())
		return gosf.NewSuccessMessage("successCreatingRoom", map[string]interface{}{
			"code": 200,
			"data": fmt.Sprint(room),
		})
	}) */

	api.PreServerShutdown = func() {
		hub.Shutdown()
	}

	api.ServerShutdown = func() {}

	return loggerMiddleware(api,
		fileServerMiddleware(
			wsMiddleware(
				api.Serve(func(next http.Handler) http.Handler {
					return next
				}),
			),
		),
	)
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

var serverAddress string

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	serverAddress = scheme + "://" + addr
}
func loggerMiddleware(api *operations.OffTimeAPI, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request) {
			api.Logger(fmt.Sprintf("%v", r.RequestURI))
			next.ServeHTTP(rw, r)
		},
	)
}

var imageStoragePath string
var imageServingRoute string

func fileServerMiddleware(next http.Handler) http.Handler {

	fileServer := http.StripPrefix(imageServingRoute, http.FileServer(http.Dir(imageStoragePath)))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, imageServingRoute) {
			fileServer.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// TODO: research if saving to temp file is better than holding in memory
func saveImageFromRequest(file io.ReadCloser) (*os.File, string, error) {
	newFile, err := ioutil.TempFile("", "tempIMG-*.jpg")
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(newFile, file)
	if err != nil {
		return nil, "", err
	}
	extension, err := checkIfFileIsAcceptedType(newFile)
	if err != nil {
		return nil, "", err
	}
	return newFile, extension, nil
}

var errUnacceptedType = fmt.Errorf("file mime type not accepted")
var errReadingFromImage = fmt.Errorf("err reading image file from request")

func checkIfFileIsAcceptedType(file *os.File) (string, error) { // this block checks if image is of accepted types
	acceptedTypes := map[string]string{
		"image/jpeg": "jpg",
		"image/png":  "png",
	}
	tempBuffer := make([]byte, 512)
	_, err := file.ReadAt(tempBuffer, 0)
	if err != nil {
		return "", errReadingFromImage
	}
	contentType := http.DetectContentType(tempBuffer)
	extension, ok := acceptedTypes[contentType]
	if !ok {
		return "", errUnacceptedType
	}
	return extension, nil
}

func saveTempImagePermanentlyToPath(tmpFile *os.File, path string) error {
	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = tmpFile.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = io.Copy(newFile, tmpFile)
	if err != nil {
		return err
	}
	return nil
}

func generateFileNameForStorage(fileName, prefix string) string {
	return prefix + "." + uuid.NewV4().String() + "." + fileName
}

func urlFromFilename(fileName string) string {
	return serverAddress + imageServingRoute + url.PathEscape(fileName)
}

var doc = template.Must(template.New("root").Parse(html))

const (
	html = `
<!DOCTYPE html>
<html lang="en">

<head>
    <title>Chat Example</title>
    <script type="text/javascript">
        window.onload = function () {
            var conn;
            var msg = document.getElementById("msg");
            var evt = document.getElementById("evt");
            var log = document.getElementById("log");
            if (!msg || !evt || !log) {
                console.error("elements not found")
            }

            function appendLog(item) {
                var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
                log.appendChild(item);
                if (doScroll) {
                    log.scrollTop = log.scrollHeight - log.clientHeight;
                }
            }

            document.getElementById("form").onsubmit = function () {
                try {
                    if (!conn) {
                        console.log("no connection found");
                        return false;
                    }
                    if (!msg.value) {
                        console.log("msg box empty");
                        return false;
                    }
                    if (!evt.value) {
                        console.log("evt box empty");
                        return false;
                    }
                    conn.send(JSON.stringify({ event: evt.value, data: msg.value }));
                } catch (e) {
                    console.error("error")
                    console.error(e)
                }
                return false;
            };

            if (window["WebSocket"]) {
                conn = new WebSocket("ws://" + document.location.host + "/ws");
                conn.onclose = function (evt) {
                    var item = document.createElement("div");
                    item.innerHTML = "<b>Connection closed.</b>";
                    appendLog(item);
                };
                conn.onmessage = function (evt) {
                    // var message = JSON.parse(evt);
                    console.log("msg recieved");
                    console.log(evt)
                    var item = document.createElement("div");
                    item.innerText = evt.data;
                    appendLog(item);
                };
            } else {
                var item = document.createElement("div");
                item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
                appendLog(item);
            }
        };
    </script>
    <style type="text/css">
        html {
            overflow: hidden;
        }

        body {
            overflow: hidden;
            padding: 0;
            margin: 0;
            width: 100%;
            height: 100%;
            background: gray;
        }

        #log {
            background: white;
            margin: 0;
            padding: 0.5em 0.5em 0.5em 0.5em;
            position: absolute;
            top: 0.5em;
            left: 0.5em;
            right: 0.5em;
            bottom: 3em;
            overflow: auto;
        }

        #form {
            padding: 0 0.5em 0 0.5em;
            margin: 0;
            position: absolute;
            bottom: 1em;
            left: 0px;
            width: 100%;
            overflow: hidden;
        }
    </style>
</head>

<body>
    <div id="log"></div>
    <form id="form">
        <input type="submit" value="Send" />
        <input style="display: inline-block;" type="text" id="evt" size="64" autofocus />
        <input style="display: inline-block;" type="text" id="msg" size="64" autofocus />
    </form>
</body>

</html>
`
)
