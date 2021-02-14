// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

// GetUsersUsernameRoomHistoryOKCode is the HTTP code returned for type GetUsersUsernameRoomHistoryOK
const GetUsersUsernameRoomHistoryOKCode int = 200

/*GetUsersUsernameRoomHistoryOK Success gettin all Room history of User.

swagger:response getUsersUsernameRoomHistoryOK
*/
type GetUsersUsernameRoomHistoryOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Room `json:"body,omitempty"`
}

// NewGetUsersUsernameRoomHistoryOK creates GetUsersUsernameRoomHistoryOK with default headers values
func NewGetUsersUsernameRoomHistoryOK() *GetUsersUsernameRoomHistoryOK {

	return &GetUsersUsernameRoomHistoryOK{}
}

// WithPayload adds the payload to the get users username room history o k response
func (o *GetUsersUsernameRoomHistoryOK) WithPayload(payload []*models.Room) *GetUsersUsernameRoomHistoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users username room history o k response
func (o *GetUsersUsernameRoomHistoryOK) SetPayload(payload []*models.Room) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersUsernameRoomHistoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Room, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetUsersUsernameRoomHistoryNotFoundCode is the HTTP code returned for type GetUsersUsernameRoomHistoryNotFound
const GetUsersUsernameRoomHistoryNotFoundCode int = 404

/*GetUsersUsernameRoomHistoryNotFound Entity not found.

swagger:response getUsersUsernameRoomHistoryNotFound
*/
type GetUsersUsernameRoomHistoryNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetUsersUsernameRoomHistoryNotFoundBody `json:"body,omitempty"`
}

// NewGetUsersUsernameRoomHistoryNotFound creates GetUsersUsernameRoomHistoryNotFound with default headers values
func NewGetUsersUsernameRoomHistoryNotFound() *GetUsersUsernameRoomHistoryNotFound {

	return &GetUsersUsernameRoomHistoryNotFound{}
}

// WithPayload adds the payload to the get users username room history not found response
func (o *GetUsersUsernameRoomHistoryNotFound) WithPayload(payload *GetUsersUsernameRoomHistoryNotFoundBody) *GetUsersUsernameRoomHistoryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users username room history not found response
func (o *GetUsersUsernameRoomHistoryNotFound) SetPayload(payload *GetUsersUsernameRoomHistoryNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersUsernameRoomHistoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersUsernameRoomHistoryInternalServerErrorCode is the HTTP code returned for type GetUsersUsernameRoomHistoryInternalServerError
const GetUsersUsernameRoomHistoryInternalServerErrorCode int = 500

/*GetUsersUsernameRoomHistoryInternalServerError Server error.

swagger:response getUsersUsernameRoomHistoryInternalServerError
*/
type GetUsersUsernameRoomHistoryInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetUsersUsernameRoomHistoryInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUsersUsernameRoomHistoryInternalServerError creates GetUsersUsernameRoomHistoryInternalServerError with default headers values
func NewGetUsersUsernameRoomHistoryInternalServerError() *GetUsersUsernameRoomHistoryInternalServerError {

	return &GetUsersUsernameRoomHistoryInternalServerError{}
}

// WithPayload adds the payload to the get users username room history internal server error response
func (o *GetUsersUsernameRoomHistoryInternalServerError) WithPayload(payload *GetUsersUsernameRoomHistoryInternalServerErrorBody) *GetUsersUsernameRoomHistoryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users username room history internal server error response
func (o *GetUsersUsernameRoomHistoryInternalServerError) SetPayload(payload *GetUsersUsernameRoomHistoryInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersUsernameRoomHistoryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
