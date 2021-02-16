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

// PostRoomsRoomIDUserUsagesHandlerFunc turns a function with the right signature into a post rooms room ID user usages handler
type PostRoomsRoomIDUserUsagesHandlerFunc func(PostRoomsRoomIDUserUsagesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRoomsRoomIDUserUsagesHandlerFunc) Handle(params PostRoomsRoomIDUserUsagesParams) middleware.Responder {
	return fn(params)
}

// PostRoomsRoomIDUserUsagesHandler interface for that can handle valid post rooms room ID user usages params
type PostRoomsRoomIDUserUsagesHandler interface {
	Handle(PostRoomsRoomIDUserUsagesParams) middleware.Responder
}

// NewPostRoomsRoomIDUserUsages creates a new http.Handler for the post rooms room ID user usages operation
func NewPostRoomsRoomIDUserUsages(ctx *middleware.Context, handler PostRoomsRoomIDUserUsagesHandler) *PostRoomsRoomIDUserUsages {
	return &PostRoomsRoomIDUserUsages{Context: ctx, Handler: handler}
}

/* PostRoomsRoomIDUserUsages swagger:route POST /rooms/{roomID}/userUsages postRoomsRoomIdUserUsages

PostRoomsRoomIDUserUsages post rooms room ID user usages API

*/
type PostRoomsRoomIDUserUsages struct {
	Context *middleware.Context
	Handler PostRoomsRoomIDUserUsagesHandler
}

func (o *PostRoomsRoomIDUserUsages) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostRoomsRoomIDUserUsagesParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostRoomsRoomIDUserUsagesBadRequestBody post rooms room ID user usages bad request body
//
// swagger:model PostRoomsRoomIDUserUsagesBadRequestBody
type PostRoomsRoomIDUserUsagesBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post rooms room ID user usages bad request body
func (o *PostRoomsRoomIDUserUsagesBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post rooms room ID user usages bad request body based on context it is used
func (o *PostRoomsRoomIDUserUsagesBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRoomsRoomIDUserUsagesBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRoomsRoomIDUserUsagesBadRequestBody) UnmarshalBinary(b []byte) error {
	var res PostRoomsRoomIDUserUsagesBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRoomsRoomIDUserUsagesInternalServerErrorBody post rooms room ID user usages internal server error body
//
// swagger:model PostRoomsRoomIDUserUsagesInternalServerErrorBody
type PostRoomsRoomIDUserUsagesInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post rooms room ID user usages internal server error body
func (o *PostRoomsRoomIDUserUsagesInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post rooms room ID user usages internal server error body based on context it is used
func (o *PostRoomsRoomIDUserUsagesInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRoomsRoomIDUserUsagesInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRoomsRoomIDUserUsagesInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res PostRoomsRoomIDUserUsagesInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostRoomsRoomIDUserUsagesNotFoundBody post rooms room ID user usages not found body
//
// swagger:model PostRoomsRoomIDUserUsagesNotFoundBody
type PostRoomsRoomIDUserUsagesNotFoundBody struct {

	// May be null.
	Entity string `json:"entity,omitempty"`

	// May be null.
	Identifer string `json:"identifer,omitempty"`
}

// Validate validates this post rooms room ID user usages not found body
func (o *PostRoomsRoomIDUserUsagesNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post rooms room ID user usages not found body based on context it is used
func (o *PostRoomsRoomIDUserUsagesNotFoundBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostRoomsRoomIDUserUsagesNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostRoomsRoomIDUserUsagesNotFoundBody) UnmarshalBinary(b []byte) error {
	var res PostRoomsRoomIDUserUsagesNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}