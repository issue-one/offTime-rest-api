package models

import (
	"github.com/go-openapi/strfmt"
)

type UpdateRoomInput struct {
	Name      string          `json:"name,omitempty"`
	EndTime   strfmt.DateTime `json:"endTime,omitempty"`
	StartTime strfmt.DateTime `json:"startTime,omitempty"`
}
