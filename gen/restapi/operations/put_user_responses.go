// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

// PutUserOKCode is the HTTP code returned for type PutUserOK
const PutUserOKCode int = 200

/*PutUserOK Success.

swagger:response putUserOK
*/
type PutUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPutUserOK creates PutUserOK with default headers values
func NewPutUserOK() *PutUserOK {

	return &PutUserOK{}
}

// WithPayload adds the payload to the put user o k response
func (o *PutUserOK) WithPayload(payload *models.User) *PutUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put user o k response
func (o *PutUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutUserBadRequestCode is the HTTP code returned for type PutUserBadRequest
const PutUserBadRequestCode int = 400

/*PutUserBadRequest Illegal input for operation.

swagger:response putUserBadRequest
*/
type PutUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PutUserBadRequestBody `json:"body,omitempty"`
}

// NewPutUserBadRequest creates PutUserBadRequest with default headers values
func NewPutUserBadRequest() *PutUserBadRequest {

	return &PutUserBadRequest{}
}

// WithPayload adds the payload to the put user bad request response
func (o *PutUserBadRequest) WithPayload(payload *PutUserBadRequestBody) *PutUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put user bad request response
func (o *PutUserBadRequest) SetPayload(payload *PutUserBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutUserUnauthorizedCode is the HTTP code returned for type PutUserUnauthorized
const PutUserUnauthorizedCode int = 401

/*PutUserUnauthorized Unauthorized.

swagger:response putUserUnauthorized
*/
type PutUserUnauthorized struct {
}

// NewPutUserUnauthorized creates PutUserUnauthorized with default headers values
func NewPutUserUnauthorized() *PutUserUnauthorized {

	return &PutUserUnauthorized{}
}

// WriteResponse to the client
func (o *PutUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PutUserConflictCode is the HTTP code returned for type PutUserConflict
const PutUserConflictCode int = 409

/*PutUserConflict Field occupied.

swagger:response putUserConflict
*/
type PutUserConflict struct {

	/*
	  In: Body
	*/
	Payload *PutUserConflictBody `json:"body,omitempty"`
}

// NewPutUserConflict creates PutUserConflict with default headers values
func NewPutUserConflict() *PutUserConflict {

	return &PutUserConflict{}
}

// WithPayload adds the payload to the put user conflict response
func (o *PutUserConflict) WithPayload(payload *PutUserConflictBody) *PutUserConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put user conflict response
func (o *PutUserConflict) SetPayload(payload *PutUserConflictBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutUserInternalServerErrorCode is the HTTP code returned for type PutUserInternalServerError
const PutUserInternalServerErrorCode int = 500

/*PutUserInternalServerError Server error.

swagger:response putUserInternalServerError
*/
type PutUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PutUserInternalServerErrorBody `json:"body,omitempty"`
}

// NewPutUserInternalServerError creates PutUserInternalServerError with default headers values
func NewPutUserInternalServerError() *PutUserInternalServerError {

	return &PutUserInternalServerError{}
}

// WithPayload adds the payload to the put user internal server error response
func (o *PutUserInternalServerError) WithPayload(payload *PutUserInternalServerErrorBody) *PutUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put user internal server error response
func (o *PutUserInternalServerError) SetPayload(payload *PutUserInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
