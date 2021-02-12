// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PutUsersUsernamePictureHandlerFunc turns a function with the right signature into a put users username picture handler
type PutUsersUsernamePictureHandlerFunc func(PutUsersUsernamePictureParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutUsersUsernamePictureHandlerFunc) Handle(params PutUsersUsernamePictureParams) middleware.Responder {
	return fn(params)
}

// PutUsersUsernamePictureHandler interface for that can handle valid put users username picture params
type PutUsersUsernamePictureHandler interface {
	Handle(PutUsersUsernamePictureParams) middleware.Responder
}

// NewPutUsersUsernamePicture creates a new http.Handler for the put users username picture operation
func NewPutUsersUsernamePicture(ctx *middleware.Context, handler PutUsersUsernamePictureHandler) *PutUsersUsernamePicture {
	return &PutUsersUsernamePicture{Context: ctx, Handler: handler}
}

/* PutUsersUsernamePicture swagger:route PUT /users/{username}/picture putUsersUsernamePicture

PutUsersUsernamePicture put users username picture API

*/
type PutUsersUsernamePicture struct {
	Context *middleware.Context
	Handler PutUsersUsernamePictureHandler
}

func (o *PutUsersUsernamePicture) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutUsersUsernamePictureParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutUsersUsernamePictureBadRequestBody put users username picture bad request body
//
// swagger:model PutUsersUsernamePictureBadRequestBody
type PutUsersUsernamePictureBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this put users username picture bad request body
func (o *PutUsersUsernamePictureBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put users username picture bad request body based on context it is used
func (o *PutUsersUsernamePictureBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutUsersUsernamePictureBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutUsersUsernamePictureBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PutUsersUsernamePictureBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutUsersUsernamePictureInternalServerErrorBody put users username picture internal server error body
//
// swagger:model PutUsersUsernamePictureInternalServerErrorBody
type PutUsersUsernamePictureInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this put users username picture internal server error body
func (o *PutUsersUsernamePictureInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put users username picture internal server error body based on context it is used
func (o *PutUsersUsernamePictureInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutUsersUsernamePictureInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutUsersUsernamePictureInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PutUsersUsernamePictureInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutUsersUsernamePictureNotFoundBody put users username picture not found body
//
// swagger:model PutUsersUsernamePictureNotFoundBody
type PutUsersUsernamePictureNotFoundBody struct {

	// May be null.
	Entity string `json:"entity,omitempty"`

	// May be null.
	Identifer string `json:"identifer,omitempty"`
}

// Validate validates this put users username picture not found body
func (o *PutUsersUsernamePictureNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this put users username picture not found body based on context it is used
func (o *PutUsersUsernamePictureNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PutUsersUsernamePictureNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutUsersUsernamePictureNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PutUsersUsernamePictureNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
