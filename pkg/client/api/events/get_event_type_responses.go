// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula-cli/pkg/client/api/models"
)

// GetEventTypeReader is a Reader for the GetEventType structure.
type GetEventTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEventTypeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventTypeOK creates a GetEventTypeOK with default headers values
func NewGetEventTypeOK() *GetEventTypeOK {
	return &GetEventTypeOK{}
}

/*GetEventTypeOK handles this case with default header values.

A representation of an event type
*/
type GetEventTypeOK struct {
	Payload *GetEventTypeOKBody
}

func (o *GetEventTypeOK) Error() string {
	return fmt.Sprintf("[GET /api/event-types/{eventTypeId}][%d] getEventTypeOK  %+v", 200, o.Payload)
}

func (o *GetEventTypeOK) GetPayload() *GetEventTypeOKBody {
	return o.Payload
}

func (o *GetEventTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEventTypeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventTypeDefault creates a GetEventTypeDefault with default headers values
func NewGetEventTypeDefault(code int) *GetEventTypeDefault {
	return &GetEventTypeDefault{
		_statusCode: code,
	}
}

/*GetEventTypeDefault handles this case with default header values.

An error occurred
*/
type GetEventTypeDefault struct {
	_statusCode int

	Payload *GetEventTypeDefaultBody
}

// Code gets the status code for the get event type default response
func (o *GetEventTypeDefault) Code() int {
	return o._statusCode
}

func (o *GetEventTypeDefault) Error() string {
	return fmt.Sprintf("[GET /api/event-types/{eventTypeId}][%d] getEventType default  %+v", o._statusCode, o.Payload)
}

func (o *GetEventTypeDefault) GetPayload() *GetEventTypeDefaultBody {
	return o.Payload
}

func (o *GetEventTypeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEventTypeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetEventTypeDefaultBody Error response
swagger:model GetEventTypeDefaultBody
*/
type GetEventTypeDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this get event type default body
func (o *GetEventTypeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEventTypeDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEventType default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEventTypeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEventTypeDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetEventTypeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetEventTypeOKBody The response type for an event type
swagger:model GetEventTypeOKBody
*/
type GetEventTypeOKBody struct {

	// type
	Type *models.EventType `json:"type,omitempty"`
}

// Validate validates this get event type o k body
func (o *GetEventTypeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEventTypeOKBody) validateType(formats strfmt.Registry) error {

	if swag.IsZero(o.Type) { // not required
		return nil
	}

	if o.Type != nil {
		if err := o.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEventTypeOK" + "." + "type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEventTypeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEventTypeOKBody) UnmarshalBinary(b []byte) error {
	var res GetEventTypeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
