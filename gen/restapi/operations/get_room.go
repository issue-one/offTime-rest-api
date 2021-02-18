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

	"github.com/issue-one/offTime-rest-api/gen/models"
)

// GetRoomHandlerFunc turns a function with the right signature into a get room handler
type GetRoomHandlerFunc func(GetRoomParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRoomHandlerFunc) Handle(params GetRoomParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetRoomHandler interface for that can handle valid get room params
type GetRoomHandler interface {
	Handle(GetRoomParams, *models.User) middleware.Responder
}

// NewGetRoom creates a new http.Handler for the get room operation
func NewGetRoom(ctx *middleware.Context, handler GetRoomHandler) *GetRoom {
	return &GetRoom{Context: ctx, Handler: handler}
}

/* GetRoom swagger:route GET /rooms/{roomID} getRoom

GetRoom get room API

*/
type GetRoom struct {
	Context *middleware.Context
	Handler GetRoomHandler
}

func (o *GetRoom) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRoomParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetRoomInternalServerErrorBody get room internal server error body
//
// swagger:model GetRoomInternalServerErrorBody
type GetRoomInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get room internal server error body
func (o *GetRoomInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get room internal server error body based on context it is used
func (o *GetRoomInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoomInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoomInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetRoomInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetRoomNotFoundBody get room not found body
//
// swagger:model GetRoomNotFoundBody
type GetRoomNotFoundBody struct {

	// May be null.
	Entity string `json:"entity,omitempty"`

	// May be null.
	Identifer string `json:"identifer,omitempty"`
}

// Validate validates this get room not found body
func (o *GetRoomNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get room not found body based on context it is used
func (o *GetRoomNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoomNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoomNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetRoomNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}