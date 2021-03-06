// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

// GetUserRoomHistoryOKCode is the HTTP code returned for type GetUserRoomHistoryOK
const GetUserRoomHistoryOKCode int = 200

/*GetUserRoomHistoryOK Success gettin all Room history of User.

swagger:response getUserRoomHistoryOK
*/
type GetUserRoomHistoryOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Room `json:"body,omitempty"`
}

// NewGetUserRoomHistoryOK creates GetUserRoomHistoryOK with default headers values
func NewGetUserRoomHistoryOK() *GetUserRoomHistoryOK {

	return &GetUserRoomHistoryOK{}
}

// WithPayload adds the payload to the get user room history o k response
func (o *GetUserRoomHistoryOK) WithPayload(payload []*models.Room) *GetUserRoomHistoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user room history o k response
func (o *GetUserRoomHistoryOK) SetPayload(payload []*models.Room) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserRoomHistoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// GetUserRoomHistoryUnauthorizedCode is the HTTP code returned for type GetUserRoomHistoryUnauthorized
const GetUserRoomHistoryUnauthorizedCode int = 401

/*GetUserRoomHistoryUnauthorized Unauthorized.

swagger:response getUserRoomHistoryUnauthorized
*/
type GetUserRoomHistoryUnauthorized struct {
}

// NewGetUserRoomHistoryUnauthorized creates GetUserRoomHistoryUnauthorized with default headers values
func NewGetUserRoomHistoryUnauthorized() *GetUserRoomHistoryUnauthorized {

	return &GetUserRoomHistoryUnauthorized{}
}

// WriteResponse to the client
func (o *GetUserRoomHistoryUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetUserRoomHistoryForbiddenCode is the HTTP code returned for type GetUserRoomHistoryForbidden
const GetUserRoomHistoryForbiddenCode int = 403

/*GetUserRoomHistoryForbidden Forbidden.

swagger:response getUserRoomHistoryForbidden
*/
type GetUserRoomHistoryForbidden struct {
}

// NewGetUserRoomHistoryForbidden creates GetUserRoomHistoryForbidden with default headers values
func NewGetUserRoomHistoryForbidden() *GetUserRoomHistoryForbidden {

	return &GetUserRoomHistoryForbidden{}
}

// WriteResponse to the client
func (o *GetUserRoomHistoryForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetUserRoomHistoryNotFoundCode is the HTTP code returned for type GetUserRoomHistoryNotFound
const GetUserRoomHistoryNotFoundCode int = 404

/*GetUserRoomHistoryNotFound Entity not found.

swagger:response getUserRoomHistoryNotFound
*/
type GetUserRoomHistoryNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetUserRoomHistoryNotFoundBody `json:"body,omitempty"`
}

// NewGetUserRoomHistoryNotFound creates GetUserRoomHistoryNotFound with default headers values
func NewGetUserRoomHistoryNotFound() *GetUserRoomHistoryNotFound {

	return &GetUserRoomHistoryNotFound{}
}

// WithPayload adds the payload to the get user room history not found response
func (o *GetUserRoomHistoryNotFound) WithPayload(payload *GetUserRoomHistoryNotFoundBody) *GetUserRoomHistoryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user room history not found response
func (o *GetUserRoomHistoryNotFound) SetPayload(payload *GetUserRoomHistoryNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserRoomHistoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserRoomHistoryInternalServerErrorCode is the HTTP code returned for type GetUserRoomHistoryInternalServerError
const GetUserRoomHistoryInternalServerErrorCode int = 500

/*GetUserRoomHistoryInternalServerError Server error.

swagger:response getUserRoomHistoryInternalServerError
*/
type GetUserRoomHistoryInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetUserRoomHistoryInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUserRoomHistoryInternalServerError creates GetUserRoomHistoryInternalServerError with default headers values
func NewGetUserRoomHistoryInternalServerError() *GetUserRoomHistoryInternalServerError {

	return &GetUserRoomHistoryInternalServerError{}
}

// WithPayload adds the payload to the get user room history internal server error response
func (o *GetUserRoomHistoryInternalServerError) WithPayload(payload *GetUserRoomHistoryInternalServerErrorBody) *GetUserRoomHistoryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user room history internal server error response
func (o *GetUserRoomHistoryInternalServerError) SetPayload(payload *GetUserRoomHistoryInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserRoomHistoryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
