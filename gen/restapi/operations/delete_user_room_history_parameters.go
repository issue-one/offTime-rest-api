// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeleteUserRoomHistoryParams creates a new DeleteUserRoomHistoryParams object
//
// There are no default values defined in the spec.
func NewDeleteUserRoomHistoryParams() DeleteUserRoomHistoryParams {

	return DeleteUserRoomHistoryParams{}
}

// DeleteUserRoomHistoryParams contains all the bound params for the delete user room history operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteUserRoomHistory
type DeleteUserRoomHistoryParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  Min Length: 5
	  In: path
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteUserRoomHistoryParams() beforehand.
func (o *DeleteUserRoomHistoryParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUsername, rhkUsername, _ := route.Params.GetOK("username")
	if err := o.bindUsername(rUsername, rhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUsername binds and validates parameter Username from path.
func (o *DeleteUserRoomHistoryParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Username = raw

	if err := o.validateUsername(formats); err != nil {
		return err
	}

	return nil
}

// validateUsername carries on validations for parameter Username
func (o *DeleteUserRoomHistoryParams) validateUsername(formats strfmt.Registry) error {

	if err := validate.MinLength("username", "path", o.Username, 5); err != nil {
		return err
	}

	return nil
}