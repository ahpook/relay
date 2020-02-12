// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowTriggerSummary workflow trigger summary
// swagger:model WorkflowTriggerSummary
type WorkflowTriggerSummary struct {
	WorkflowTriggerIdentifier

	// event source
	// Required: true
	EventSource *EventSourceIdentifier `json:"event_source"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WorkflowTriggerSummary) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WorkflowTriggerIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WorkflowTriggerIdentifier = aO0

	// AO1
	var dataAO1 struct {
		EventSource *EventSourceIdentifier `json:"event_source"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EventSource = dataAO1.EventSource

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WorkflowTriggerSummary) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WorkflowTriggerIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		EventSource *EventSourceIdentifier `json:"event_source"`
	}

	dataAO1.EventSource = m.EventSource

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this workflow trigger summary
func (m *WorkflowTriggerSummary) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WorkflowTriggerIdentifier
	if err := m.WorkflowTriggerIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowTriggerSummary) validateEventSource(formats strfmt.Registry) error {

	if err := validate.Required("event_source", "body", m.EventSource); err != nil {
		return err
	}

	if m.EventSource != nil {
		if err := m.EventSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowTriggerSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowTriggerSummary) UnmarshalBinary(b []byte) error {
	var res WorkflowTriggerSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
