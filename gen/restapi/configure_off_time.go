// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
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
	"github.com/google/uuid"

	"github.com/issue-one/offTime-rest-api/gen/restapi/operations"
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

	userRepo := mock.NewMockUserRepositories()
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
				return operations.NewPutUsersUsernameInternalServerError().WithPayload(
					&operations.PutUsersUsernameInternalServerErrorBody{Message: err.Error()},
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
			if err != nil {
				return operations.NewGetUsersUsernameInternalServerError().WithPayload(
					&operations.GetUsersUsernameInternalServerErrorBody{Message: err.Error()},
				)
			}
			if user.PictureURL != "" {
				user.PictureURL = urlFromFilename(user.PictureURL)
			}
			return operations.NewGetUsersUsernameOK().WithPayload(user)
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
					return operations.NewPatchUsersUsernameBadRequest().WithPayload(
						&operations.PatchUsersUsernameBadRequestBody{Message: err.Error()},
					)
				}
				if occupied {
					return operations.NewPatchUsersUsernameConflict().WithPayload(
						&operations.PatchUsersUsernameConflictBody{Field: "Email"},
					)
				}
			}
			user, err := userRepo.UpdateUser(ctx, params.Username, params.Body)
			if err != nil {
				return operations.NewPatchUsersUsernameInternalServerError().WithPayload(
					&operations.PatchUsersUsernameInternalServerErrorBody{Message: err.Error()},
				)
			}
			if user.PictureURL != "" {
				user.PictureURL = urlFromFilename(user.PictureURL)
			}
			return operations.NewPatchUsersUsernameOK().WithPayload(user)
		})

	api.PutUsersUsernamePictureHandler = operations.PutUsersUsernamePictureHandlerFunc(
		func(params operations.PutUsersUsernamePictureParams) middleware.Responder {
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
			user, err := userRepo.SetImage(
				params.HTTPRequest.Context(),
				params.Username,
				generateFileNameForStorage(params.Username+"."+extension, "user"),
			)
			if err != nil {
				return operations.NewPutUsersUsernamePictureInternalServerError().WithPayload(
					&operations.PutUsersUsernamePictureInternalServerErrorBody{
						Message: err.Error(),
					},
				)
			}
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
	if api.GetRoomsHandler == nil {
		api.GetRoomsHandler = operations.GetRoomsHandlerFunc(func(params operations.GetRoomsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetRooms has not yet been implemented")
		})
	}
	if api.GetRoomsRoomIDHandler == nil {
		api.GetRoomsRoomIDHandler = operations.GetRoomsRoomIDHandlerFunc(func(params operations.GetRoomsRoomIDParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetRoomsRoomID has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api)(api.Serve(setupMiddlewares(api)))
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

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(api *operations.OffTimeAPI) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler { return next }
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

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(api *operations.OffTimeAPI) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return loggerMiddleware(api, fileServerMiddleware(next))
	}
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
	return prefix + "." + uuid.New().String() + "." + fileName
}

func urlFromFilename(fileName string) string {
	return serverAddress + imageServingRoute + url.PathEscape(fileName)
}
