// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

// PostAppUsageOKCode is the HTTP code returned for type PostAppUsageOK
const PostAppUsageOKCode int = 200

/*PostAppUsageOK Success adding AppUsage entry.

swagger:response postAppUsageOK
*/
type PostAppUsageOK struct {

	/*
	  In: Body
	*/
	Payload *models.AppUsage `json:"body,omitempty"`
}

// NewPostAppUsageOK creates PostAppUsageOK with default headers values
func NewPostAppUsageOK() *PostAppUsageOK {

	return &PostAppUsageOK{}
}

// WithPayload adds the payload to the post app usage o k response
func (o *PostAppUsageOK) WithPayload(payload *models.AppUsage) *PostAppUsageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post app usage o k response
func (o *PostAppUsageOK) SetPayload(payload *models.AppUsage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAppUsageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAppUsageBadRequestCode is the HTTP code returned for type PostAppUsageBadRequest
const PostAppUsageBadRequestCode int = 400

/*PostAppUsageBadRequest Illegal input for operation.

swagger:response postAppUsageBadRequest
*/
type PostAppUsageBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PostAppUsageBadRequestBody `json:"body,omitempty"`
}

// NewPostAppUsageBadRequest creates PostAppUsageBadRequest with default headers values
func NewPostAppUsageBadRequest() *PostAppUsageBadRequest {

	return &PostAppUsageBadRequest{}
}

// WithPayload adds the payload to the post app usage bad request response
func (o *PostAppUsageBadRequest) WithPayload(payload *PostAppUsageBadRequestBody) *PostAppUsageBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post app usage bad request response
func (o *PostAppUsageBadRequest) SetPayload(payload *PostAppUsageBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAppUsageBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAppUsageUnauthorizedCode is the HTTP code returned for type PostAppUsageUnauthorized
const PostAppUsageUnauthorizedCode int = 401

/*PostAppUsageUnauthorized Unauthorized.

swagger:response postAppUsageUnauthorized
*/
type PostAppUsageUnauthorized struct {
}

// NewPostAppUsageUnauthorized creates PostAppUsageUnauthorized with default headers values
func NewPostAppUsageUnauthorized() *PostAppUsageUnauthorized {

	return &PostAppUsageUnauthorized{}
}

// WriteResponse to the client
func (o *PostAppUsageUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostAppUsageForbiddenCode is the HTTP code returned for type PostAppUsageForbidden
const PostAppUsageForbiddenCode int = 403

/*PostAppUsageForbidden Forbidden.

swagger:response postAppUsageForbidden
*/
type PostAppUsageForbidden struct {
}

// NewPostAppUsageForbidden creates PostAppUsageForbidden with default headers values
func NewPostAppUsageForbidden() *PostAppUsageForbidden {

	return &PostAppUsageForbidden{}
}

// WriteResponse to the client
func (o *PostAppUsageForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PostAppUsageNotFoundCode is the HTTP code returned for type PostAppUsageNotFound
const PostAppUsageNotFoundCode int = 404

/*PostAppUsageNotFound Entity not found.

swagger:response postAppUsageNotFound
*/
type PostAppUsageNotFound struct {

	/*
	  In: Body
	*/
	Payload *PostAppUsageNotFoundBody `json:"body,omitempty"`
}

// NewPostAppUsageNotFound creates PostAppUsageNotFound with default headers values
func NewPostAppUsageNotFound() *PostAppUsageNotFound {

	return &PostAppUsageNotFound{}
}

// WithPayload adds the payload to the post app usage not found response
func (o *PostAppUsageNotFound) WithPayload(payload *PostAppUsageNotFoundBody) *PostAppUsageNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post app usage not found response
func (o *PostAppUsageNotFound) SetPayload(payload *PostAppUsageNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAppUsageNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAppUsageConflictCode is the HTTP code returned for type PostAppUsageConflict
const PostAppUsageConflictCode int = 409

/*PostAppUsageConflict Field occupied.

swagger:response postAppUsageConflict
*/
type PostAppUsageConflict struct {

	/*
	  In: Body
	*/
	Payload *PostAppUsageConflictBody `json:"body,omitempty"`
}

// NewPostAppUsageConflict creates PostAppUsageConflict with default headers values
func NewPostAppUsageConflict() *PostAppUsageConflict {

	return &PostAppUsageConflict{}
}

// WithPayload adds the payload to the post app usage conflict response
func (o *PostAppUsageConflict) WithPayload(payload *PostAppUsageConflictBody) *PostAppUsageConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post app usage conflict response
func (o *PostAppUsageConflict) SetPayload(payload *PostAppUsageConflictBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAppUsageConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostAppUsageInternalServerErrorCode is the HTTP code returned for type PostAppUsageInternalServerError
const PostAppUsageInternalServerErrorCode int = 500

/*PostAppUsageInternalServerError Server error.

swagger:response postAppUsageInternalServerError
*/
type PostAppUsageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PostAppUsageInternalServerErrorBody `json:"body,omitempty"`
}

// NewPostAppUsageInternalServerError creates PostAppUsageInternalServerError with default headers values
func NewPostAppUsageInternalServerError() *PostAppUsageInternalServerError {

	return &PostAppUsageInternalServerError{}
}

// WithPayload adds the payload to the post app usage internal server error response
func (o *PostAppUsageInternalServerError) WithPayload(payload *PostAppUsageInternalServerErrorBody) *PostAppUsageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post app usage internal server error response
func (o *PostAppUsageInternalServerError) SetPayload(payload *PostAppUsageInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAppUsageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
