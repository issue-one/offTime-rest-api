// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

// GetRoomsRoomIDOKCode is the HTTP code returned for type GetRoomsRoomIDOK
const GetRoomsRoomIDOKCode int = 200

/*GetRoomsRoomIDOK Success GETting room under roomID.

swagger:response getRoomsRoomIdOK
*/
type GetRoomsRoomIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Room `json:"body,omitempty"`
}

// NewGetRoomsRoomIDOK creates GetRoomsRoomIDOK with default headers values
func NewGetRoomsRoomIDOK() *GetRoomsRoomIDOK {

	return &GetRoomsRoomIDOK{}
}

// WithPayload adds the payload to the get rooms room Id o k response
func (o *GetRoomsRoomIDOK) WithPayload(payload *models.Room) *GetRoomsRoomIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rooms room Id o k response
func (o *GetRoomsRoomIDOK) SetPayload(payload *models.Room) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRoomsRoomIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRoomsRoomIDBadRequestCode is the HTTP code returned for type GetRoomsRoomIDBadRequest
const GetRoomsRoomIDBadRequestCode int = 400

/*GetRoomsRoomIDBadRequest Illegal input for operation.

swagger:response getRoomsRoomIdBadRequest
*/
type GetRoomsRoomIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetRoomsRoomIDBadRequestBody `json:"body,omitempty"`
}

// NewGetRoomsRoomIDBadRequest creates GetRoomsRoomIDBadRequest with default headers values
func NewGetRoomsRoomIDBadRequest() *GetRoomsRoomIDBadRequest {

	return &GetRoomsRoomIDBadRequest{}
}

// WithPayload adds the payload to the get rooms room Id bad request response
func (o *GetRoomsRoomIDBadRequest) WithPayload(payload *GetRoomsRoomIDBadRequestBody) *GetRoomsRoomIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rooms room Id bad request response
func (o *GetRoomsRoomIDBadRequest) SetPayload(payload *GetRoomsRoomIDBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRoomsRoomIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRoomsRoomIDInternalServerErrorCode is the HTTP code returned for type GetRoomsRoomIDInternalServerError
const GetRoomsRoomIDInternalServerErrorCode int = 500

/*GetRoomsRoomIDInternalServerError Server error.

swagger:response getRoomsRoomIdInternalServerError
*/
type GetRoomsRoomIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetRoomsRoomIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetRoomsRoomIDInternalServerError creates GetRoomsRoomIDInternalServerError with default headers values
func NewGetRoomsRoomIDInternalServerError() *GetRoomsRoomIDInternalServerError {

	return &GetRoomsRoomIDInternalServerError{}
}

// WithPayload adds the payload to the get rooms room Id internal server error response
func (o *GetRoomsRoomIDInternalServerError) WithPayload(payload *GetRoomsRoomIDInternalServerErrorBody) *GetRoomsRoomIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rooms room Id internal server error response
func (o *GetRoomsRoomIDInternalServerError) SetPayload(payload *GetRoomsRoomIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRoomsRoomIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}