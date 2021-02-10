// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

// GetUsersUsernameOKCode is the HTTP code returned for type GetUsersUsernameOK
const GetUsersUsernameOKCode int = 200

/*GetUsersUsernameOK Success GETting User under given username.

swagger:response getUsersUsernameOK
*/
type GetUsersUsernameOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUsersUsernameOK creates GetUsersUsernameOK with default headers values
func NewGetUsersUsernameOK() *GetUsersUsernameOK {

	return &GetUsersUsernameOK{}
}

// WithPayload adds the payload to the get users username o k response
func (o *GetUsersUsernameOK) WithPayload(payload *models.User) *GetUsersUsernameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users username o k response
func (o *GetUsersUsernameOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersUsernameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUsersUsernameNotFoundCode is the HTTP code returned for type GetUsersUsernameNotFound
const GetUsersUsernameNotFoundCode int = 404

/*GetUsersUsernameNotFound Entity not found.

swagger:response getUsersUsernameNotFound
*/
type GetUsersUsernameNotFound struct {
}

// NewGetUsersUsernameNotFound creates GetUsersUsernameNotFound with default headers values
func NewGetUsersUsernameNotFound() *GetUsersUsernameNotFound {

	return &GetUsersUsernameNotFound{}
}

// WriteResponse to the client
func (o *GetUsersUsernameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetUsersUsernameInternalServerErrorCode is the HTTP code returned for type GetUsersUsernameInternalServerError
const GetUsersUsernameInternalServerErrorCode int = 500

/*GetUsersUsernameInternalServerError Server error.

swagger:response getUsersUsernameInternalServerError
*/
type GetUsersUsernameInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetUsersUsernameInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUsersUsernameInternalServerError creates GetUsersUsernameInternalServerError with default headers values
func NewGetUsersUsernameInternalServerError() *GetUsersUsernameInternalServerError {

	return &GetUsersUsernameInternalServerError{}
}

// WithPayload adds the payload to the get users username internal server error response
func (o *GetUsersUsernameInternalServerError) WithPayload(payload *GetUsersUsernameInternalServerErrorBody) *GetUsersUsernameInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users username internal server error response
func (o *GetUsersUsernameInternalServerError) SetPayload(payload *GetUsersUsernameInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersUsernameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
