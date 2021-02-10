// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/issue-one/offTime-rest-api/gen/models"
)

// GetRoomsHandlerFunc turns a function with the right signature into a get rooms handler
type GetRoomsHandlerFunc func(GetRoomsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRoomsHandlerFunc) Handle(params GetRoomsParams) middleware.Responder {
	return fn(params)
}

// GetRoomsHandler interface for that can handle valid get rooms params
type GetRoomsHandler interface {
	Handle(GetRoomsParams) middleware.Responder
}

// NewGetRooms creates a new http.Handler for the get rooms operation
func NewGetRooms(ctx *middleware.Context, handler GetRoomsHandler) *GetRooms {
	return &GetRooms{Context: ctx, Handler: handler}
}

/* GetRooms swagger:route GET /rooms getRooms

GetRooms get rooms API

*/
type GetRooms struct {
	Context *middleware.Context
	Handler GetRoomsHandler
}

func (o *GetRooms) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRoomsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetRoomsBadRequestBody get rooms bad request body
//
// swagger:model GetRoomsBadRequestBody
type GetRoomsBadRequestBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get rooms bad request body
func (o *GetRoomsBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get rooms bad request body based on context it is used
func (o *GetRoomsBadRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoomsBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoomsBadRequestBody) UnmarshalBinary(b []byte) error {
	var res GetRoomsBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetRoomsInternalServerErrorBody get rooms internal server error body
//
// swagger:model GetRoomsInternalServerErrorBody
type GetRoomsInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get rooms internal server error body
func (o *GetRoomsInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get rooms internal server error body based on context it is used
func (o *GetRoomsInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRoomsInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoomsInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetRoomsInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetRoomsOKBody get rooms o k body
//
// swagger:model GetRoomsOKBody
type GetRoomsOKBody struct {

	// items
	Items []*models.Room `json:"items"`

	// total count
	// Example: 98738772
	TotalCount int64 `json:"totalCount,omitempty"`
}

// Validate validates this get rooms o k body
func (o *GetRoomsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoomsOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRoomsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get rooms o k body based on the context it is used
func (o *GetRoomsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRoomsOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRoomsOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRoomsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRoomsOKBody) UnmarshalBinary(b []byte) error {
	var res GetRoomsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
