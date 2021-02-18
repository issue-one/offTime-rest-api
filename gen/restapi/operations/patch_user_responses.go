// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

// PatchUserOKCode is the HTTP code returned for type PatchUserOK
const PatchUserOKCode int = 200

/*PatchUserOK Success patching User.

swagger:response patchUserOK
*/
type PatchUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPatchUserOK creates PatchUserOK with default headers values
func NewPatchUserOK() *PatchUserOK {

	return &PatchUserOK{}
}

// WithPayload adds the payload to the patch user o k response
func (o *PatchUserOK) WithPayload(payload *models.User) *PatchUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch user o k response
func (o *PatchUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchUserBadRequestCode is the HTTP code returned for type PatchUserBadRequest
const PatchUserBadRequestCode int = 400

/*PatchUserBadRequest Illegal input for operation.

swagger:response patchUserBadRequest
*/
type PatchUserBadRequest struct {

	/*
	  In: Body
	*/
	Payload *PatchUserBadRequestBody `json:"body,omitempty"`
}

// NewPatchUserBadRequest creates PatchUserBadRequest with default headers values
func NewPatchUserBadRequest() *PatchUserBadRequest {

	return &PatchUserBadRequest{}
}

// WithPayload adds the payload to the patch user bad request response
func (o *PatchUserBadRequest) WithPayload(payload *PatchUserBadRequestBody) *PatchUserBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch user bad request response
func (o *PatchUserBadRequest) SetPayload(payload *PatchUserBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchUserUnauthorizedCode is the HTTP code returned for type PatchUserUnauthorized
const PatchUserUnauthorizedCode int = 401

/*PatchUserUnauthorized Unauthorized.

swagger:response patchUserUnauthorized
*/
type PatchUserUnauthorized struct {
}

// NewPatchUserUnauthorized creates PatchUserUnauthorized with default headers values
func NewPatchUserUnauthorized() *PatchUserUnauthorized {

	return &PatchUserUnauthorized{}
}

// WriteResponse to the client
func (o *PatchUserUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PatchUserForbiddenCode is the HTTP code returned for type PatchUserForbidden
const PatchUserForbiddenCode int = 403

/*PatchUserForbidden Forbidden.

swagger:response patchUserForbidden
*/
type PatchUserForbidden struct {
}

// NewPatchUserForbidden creates PatchUserForbidden with default headers values
func NewPatchUserForbidden() *PatchUserForbidden {

	return &PatchUserForbidden{}
}

// WriteResponse to the client
func (o *PatchUserForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PatchUserNotFoundCode is the HTTP code returned for type PatchUserNotFound
const PatchUserNotFoundCode int = 404

/*PatchUserNotFound Entity not found.

swagger:response patchUserNotFound
*/
type PatchUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *PatchUserNotFoundBody `json:"body,omitempty"`
}

// NewPatchUserNotFound creates PatchUserNotFound with default headers values
func NewPatchUserNotFound() *PatchUserNotFound {

	return &PatchUserNotFound{}
}

// WithPayload adds the payload to the patch user not found response
func (o *PatchUserNotFound) WithPayload(payload *PatchUserNotFoundBody) *PatchUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch user not found response
func (o *PatchUserNotFound) SetPayload(payload *PatchUserNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchUserConflictCode is the HTTP code returned for type PatchUserConflict
const PatchUserConflictCode int = 409

/*PatchUserConflict Field occupied.

swagger:response patchUserConflict
*/
type PatchUserConflict struct {

	/*
	  In: Body
	*/
	Payload *PatchUserConflictBody `json:"body,omitempty"`
}

// NewPatchUserConflict creates PatchUserConflict with default headers values
func NewPatchUserConflict() *PatchUserConflict {

	return &PatchUserConflict{}
}

// WithPayload adds the payload to the patch user conflict response
func (o *PatchUserConflict) WithPayload(payload *PatchUserConflictBody) *PatchUserConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch user conflict response
func (o *PatchUserConflict) SetPayload(payload *PatchUserConflictBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUserConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchUserInternalServerErrorCode is the HTTP code returned for type PatchUserInternalServerError
const PatchUserInternalServerErrorCode int = 500

/*PatchUserInternalServerError Server error.

swagger:response patchUserInternalServerError
*/
type PatchUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *PatchUserInternalServerErrorBody `json:"body,omitempty"`
}

// NewPatchUserInternalServerError creates PatchUserInternalServerError with default headers values
func NewPatchUserInternalServerError() *PatchUserInternalServerError {

	return &PatchUserInternalServerError{}
}

// WithPayload adds the payload to the patch user internal server error response
func (o *PatchUserInternalServerError) WithPayload(payload *PatchUserInternalServerErrorBody) *PatchUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch user internal server error response
func (o *PatchUserInternalServerError) SetPayload(payload *PatchUserInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}