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

// GetEventSourceReader is a Reader for the GetEventSource structure.
type GetEventSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEventSourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetEventSourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetEventSourceOK creates a GetEventSourceOK with default headers values
func NewGetEventSourceOK() *GetEventSourceOK {
	return &GetEventSourceOK{}
}

/*GetEventSourceOK handles this case with default header values.

A representation of an event source
*/
type GetEventSourceOK struct {
	Payload *GetEventSourceOKBody
}

func (o *GetEventSourceOK) Error() string {
	return fmt.Sprintf("[GET /api/event-sources/{eventSourceId}][%d] getEventSourceOK  %+v", 200, o.Payload)
}

func (o *GetEventSourceOK) GetPayload() *GetEventSourceOKBody {
	return o.Payload
}

func (o *GetEventSourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEventSourceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventSourceDefault creates a GetEventSourceDefault with default headers values
func NewGetEventSourceDefault(code int) *GetEventSourceDefault {
	return &GetEventSourceDefault{
		_statusCode: code,
	}
}

/*GetEventSourceDefault handles this case with default header values.

An error occurred
*/
type GetEventSourceDefault struct {
	_statusCode int

	Payload *GetEventSourceDefaultBody
}

// Code gets the status code for the get event source default response
func (o *GetEventSourceDefault) Code() int {
	return o._statusCode
}

func (o *GetEventSourceDefault) Error() string {
	return fmt.Sprintf("[GET /api/event-sources/{eventSourceId}][%d] getEventSource default  %+v", o._statusCode, o.Payload)
}

func (o *GetEventSourceDefault) GetPayload() *GetEventSourceDefaultBody {
	return o.Payload
}

func (o *GetEventSourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetEventSourceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetEventSourceDefaultBody Error response
swagger:model GetEventSourceDefaultBody
*/
type GetEventSourceDefaultBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`
}

// Validate validates this get event source default body
func (o *GetEventSourceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEventSourceDefaultBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEventSource default" + "." + "error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEventSourceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEventSourceDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetEventSourceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetEventSourceOKBody get event source o k body
swagger:model GetEventSourceOKBody
*/
type GetEventSourceOKBody struct {
	models.EventSourceEntity

	// source
	Source *models.EventSource `json:"source,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetEventSourceOKBody) UnmarshalJSON(raw []byte) error {
	// GetEventSourceOKBodyAO0
	var getEventSourceOKBodyAO0 models.EventSourceEntity
	if err := swag.ReadJSON(raw, &getEventSourceOKBodyAO0); err != nil {
		return err
	}
	o.EventSourceEntity = getEventSourceOKBodyAO0

	// GetEventSourceOKBodyAO1
	var dataGetEventSourceOKBodyAO1 struct {
		Source *models.EventSource `json:"source,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataGetEventSourceOKBodyAO1); err != nil {
		return err
	}

	o.Source = dataGetEventSourceOKBodyAO1.Source

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetEventSourceOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getEventSourceOKBodyAO0, err := swag.WriteJSON(o.EventSourceEntity)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getEventSourceOKBodyAO0)

	var dataGetEventSourceOKBodyAO1 struct {
		Source *models.EventSource `json:"source,omitempty"`
	}

	dataGetEventSourceOKBodyAO1.Source = o.Source

	jsonDataGetEventSourceOKBodyAO1, errGetEventSourceOKBodyAO1 := swag.WriteJSON(dataGetEventSourceOKBodyAO1)
	if errGetEventSourceOKBodyAO1 != nil {
		return nil, errGetEventSourceOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetEventSourceOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get event source o k body
func (o *GetEventSourceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.EventSourceEntity
	if err := o.EventSourceEntity.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetEventSourceOKBody) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(o.Source) { // not required
		return nil
	}

	if o.Source != nil {
		if err := o.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getEventSourceOK" + "." + "source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetEventSourceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetEventSourceOKBody) UnmarshalBinary(b []byte) error {
	var res GetEventSourceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
