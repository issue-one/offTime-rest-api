// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostTokenAuthOKCode is the HTTP code returned for type PostTokenAuthOK
const PostTokenAuthOKCode int = 200

/*PostTokenAuthOK Successful POSTing rooms

swagger:response postTokenAuthOK
*/
type PostTokenAuthOK struct {

	/*
	  In: Body
	*/
	Payload *PostTokenAuthOKBody `json:"body,omitempty"`
}

// NewPostTokenAuthOK creates PostTokenAuthOK with default headers values
func NewPostTokenAuthOK() *PostTokenAuthOK {

	return &PostTokenAuthOK{}
}

// WithPayload adds the payload to the post token auth o k response
func (o *PostTokenAuthOK) WithPayload(payload *PostTokenAuthOKBody) *PostTokenAuthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post token auth o k response
func (o *PostTokenAuthOK) SetPayload(payload *PostTokenAuthOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTokenAuthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTokenAuthBadRequestCode is the HTTP code returned for type PostTokenAuthBadRequest
const PostTokenAuthBadRequestCode int = 400

/*PostTokenAuthBadRequest Illegal input for operation.

swagger:response postTokenAuthBadRequest
*/
type PostTokenAuthBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostTokenAuthBadRequestBody `json:"body,omitempty"`
}

// NewPostTokenAuthBadRequest creates PostTokenAuthBadRequest with default headers values
func NewPostTokenAuthBadRequest() *PostTokenAuthBadRequest {

	return &PostTokenAuthBadRequest{}
}

// WithPayload adds the payload to the post token auth bad request response
func (o *PostTokenAuthBadRequest) WithPayload(payload *PostTokenAuthBadRequestBody) *PostTokenAuthBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post token auth bad request response
func (o *PostTokenAuthBadRequest) SetPayload(payload *PostTokenAuthBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTokenAuthBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostTokenAuthForbiddenCode is the HTTP code returned for type PostTokenAuthForbidden
const PostTokenAuthForbiddenCode int = 403

/*PostTokenAuthForbidden Forbidden.

swagger:response postTokenAuthForbidden
*/
type PostTokenAuthForbidden struct {
}

// NewPostTokenAuthForbidden creates PostTokenAuthForbidden with default headers values
func NewPostTokenAuthForbidden() *PostTokenAuthForbidden {

	return &PostTokenAuthForbidden{}
}

// WriteResponse to the client
func (o *PostTokenAuthForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PostTokenAuthInternalServerErrorCode is the HTTP code returned for type PostTokenAuthInternalServerError
const PostTokenAuthInternalServerErrorCode int = 500

/*PostTokenAuthInternalServerError Server error.

swagger:response postTokenAuthInternalServerError
*/
type PostTokenAuthInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostTokenAuthInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostTokenAuthInternalServerError creates PostTokenAuthInternalServerError with default headers values
func NewPostTokenAuthInternalServerError() *PostTokenAuthInternalServerError {

	return &PostTokenAuthInternalServerError{}
}

// WithPayload adds the payload to the post token auth internal server error response
func (o *PostTokenAuthInternalServerError) WithPayload(payload *PostTokenAuthInternalServerErrorBody) *PostTokenAuthInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post token auth internal server error response
func (o *PostTokenAuthInternalServerError) SetPayload(payload *PostTokenAuthInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostTokenAuthInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
