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

// DeleteUsersUsernameUsageHistoryHandlerFunc turns a function with the right signature into a delete users username usage history handler
type DeleteUsersUsernameUsageHistoryHandlerFunc func(DeleteUsersUsernameUsageHistoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteUsersUsernameUsageHistoryHandlerFunc) Handle(params DeleteUsersUsernameUsageHistoryParams) middleware.Responder {
	return fn(params)
}

// DeleteUsersUsernameUsageHistoryHandler interface for that can handle valid delete users username usage history params
type DeleteUsersUsernameUsageHistoryHandler interface {
	Handle(DeleteUsersUsernameUsageHistoryParams) middleware.Responder
}

// NewDeleteUsersUsernameUsageHistory creates a new http.Handler for the delete users username usage history operation
func NewDeleteUsersUsernameUsageHistory(ctx *middleware.Context, handler DeleteUsersUsernameUsageHistoryHandler) *DeleteUsersUsernameUsageHistory {
	return &DeleteUsersUsernameUsageHistory{Context: ctx, Handler: handler}
}

/* DeleteUsersUsernameUsageHistory swagger:route DELETE /users/{username}/usageHistory deleteUsersUsernameUsageHistory

DeleteUsersUsernameUsageHistory delete users username usage history API

*/
type DeleteUsersUsernameUsageHistory struct {
	Context *middleware.Context
	Handler DeleteUsersUsernameUsageHistoryHandler
}

func (o *DeleteUsersUsernameUsageHistory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteUsersUsernameUsageHistoryParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteUsersUsernameUsageHistoryInternalServerErrorBody delete users username usage history internal server error body
//
// swagger:model DeleteUsersUsernameUsageHistoryInternalServerErrorBody
type DeleteUsersUsernameUsageHistoryInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete users username usage history internal server error body
func (o *DeleteUsersUsernameUsageHistoryInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete users username usage history internal server error body based on context it is used
func (o *DeleteUsersUsernameUsageHistoryInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteUsersUsernameUsageHistoryInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteUsersUsernameUsageHistoryInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res DeleteUsersUsernameUsageHistoryInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}