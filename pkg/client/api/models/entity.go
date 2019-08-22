// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Entity A general response for a single entity
// swagger:model Entity
type Entity struct {

	// access
	Access *EntityAccess `json:"access,omitempty"`
}

// Validate validates this entity
func (m *Entity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Entity) validateAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.Access) { // not required
		return nil
	}

	if m.Access != nil {
		if err := m.Access.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("access")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Entity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Entity) UnmarshalBinary(b []byte) error {
	var res Entity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}