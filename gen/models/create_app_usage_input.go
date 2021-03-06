// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAppUsageInput create app usage input
//
// swagger:model CreateAppUsageInput
type CreateAppUsageInput struct {

	// app name
	// Example: Telegram
	// Required: true
	AppName *string `json:"appName"`

	// app package name
	// Example: com.example.telegram
	// Required: true
	AppPackageName *string `json:"appPackageName"`

	// date of use
	// Example: 24-04-1997
	// Required: true
	// Format: date
	DateOfUse *strfmt.Date `json:"dateOfUse"`

	// Usage time in seconds.
	// Required: true
	TimeDuration *int64 `json:"timeDuration"`
}

// Validate validates this create app usage input
func (m *CreateAppUsageInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAppPackageName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateOfUse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAppUsageInput) validateAppName(formats strfmt.Registry) error {

	if err := validate.Required("appName", "body", m.AppName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAppUsageInput) validateAppPackageName(formats strfmt.Registry) error {

	if err := validate.Required("appPackageName", "body", m.AppPackageName); err != nil {
		return err
	}

	return nil
}

func (m *CreateAppUsageInput) validateDateOfUse(formats strfmt.Registry) error {

	if err := validate.Required("dateOfUse", "body", m.DateOfUse); err != nil {
		return err
	}

	if err := validate.FormatOf("dateOfUse", "body", "date", m.DateOfUse.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateAppUsageInput) validateTimeDuration(formats strfmt.Registry) error {

	if err := validate.Required("timeDuration", "body", m.TimeDuration); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create app usage input based on context it is used
func (m *CreateAppUsageInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAppUsageInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAppUsageInput) UnmarshalBinary(b []byte) error {
	var res CreateAppUsageInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
