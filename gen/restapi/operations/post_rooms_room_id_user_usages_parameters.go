// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewPostRoomsRoomIDUserUsagesParams creates a new PostRoomsRoomIDUserUsagesParams object
//
// There are no default values defined in the spec.
func NewPostRoomsRoomIDUserUsagesParams() PostRoomsRoomIDUserUsagesParams {

	return PostRoomsRoomIDUserUsagesParams{}
}

// PostRoomsRoomIDUserUsagesParams contains all the bound params for the post rooms room ID user usages operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostRoomsRoomIDUserUsages
type PostRoomsRoomIDUserUsagesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	RoomID strfmt.UUID
	/*
	  Required: true
	  In: query
	*/
	Seconds int64
	/*
	  Required: true
	  Min Length: 5
	  In: query
	*/
	Username string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostRoomsRoomIDUserUsagesParams() beforehand.
func (o *PostRoomsRoomIDUserUsagesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rRoomID, rhkRoomID, _ := route.Params.GetOK("roomID")
	if err := o.bindRoomID(rRoomID, rhkRoomID, route.Formats); err != nil {
		res = append(res, err)
	}

	qSeconds, qhkSeconds, _ := qs.GetOK("seconds")
	if err := o.bindSeconds(qSeconds, qhkSeconds, route.Formats); err != nil {
		res = append(res, err)
	}

	qUsername, qhkUsername, _ := qs.GetOK("username")
	if err := o.bindUsername(qUsername, qhkUsername, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindRoomID binds and validates parameter RoomID from path.
func (o *PostRoomsRoomIDUserUsagesParams) bindRoomID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("roomID", "path", "strfmt.UUID", raw)
	}
	o.RoomID = *(value.(*strfmt.UUID))

	if err := o.validateRoomID(formats); err != nil {
		return err
	}

	return nil
}

// validateRoomID carries on validations for parameter RoomID
func (o *PostRoomsRoomIDUserUsagesParams) validateRoomID(formats strfmt.Registry) error {

	if err := validate.FormatOf("roomID", "path", "uuid", o.RoomID.String(), formats); err != nil {
		return err
	}
	return nil
}

// bindSeconds binds and validates parameter Seconds from query.
func (o *PostRoomsRoomIDUserUsagesParams) bindSeconds(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("seconds", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("seconds", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("seconds", "query", "int64", raw)
	}
	o.Seconds = value

	return nil
}

// bindUsername binds and validates parameter Username from query.
func (o *PostRoomsRoomIDUserUsagesParams) bindUsername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("username", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("username", "query", raw); err != nil {
		return err
	}
	o.Username = raw

	if err := o.validateUsername(formats); err != nil {
		return err
	}

	return nil
}

// validateUsername carries on validations for parameter Username
func (o *PostRoomsRoomIDUserUsagesParams) validateUsername(formats strfmt.Registry) error {

	if err := validate.MinLength("username", "query", o.Username, 5); err != nil {
		return err
	}

	return nil
}