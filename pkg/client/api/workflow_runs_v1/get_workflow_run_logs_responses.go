// Code generated by go-swagger; DO NOT EDIT.

package workflow_runs_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/puppetlabs/nebula/pkg/client/api/models"
)

// GetWorkflowRunLogsReader is a Reader for the GetWorkflowRunLogs structure.
type GetWorkflowRunLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkflowRunLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWorkflowRunLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWorkflowRunLogsOK creates a GetWorkflowRunLogsOK with default headers values
func NewGetWorkflowRunLogsOK() *GetWorkflowRunLogsOK {
	return &GetWorkflowRunLogsOK{}
}

/*GetWorkflowRunLogsOK handles this case with default header values.

A workflow log representation
*/
type GetWorkflowRunLogsOK struct {
	Payload *models.Log
}

func (o *GetWorkflowRunLogsOK) Error() string {
	return fmt.Sprintf("[GET /api/runs/{rid}/logs][%d] getWorkflowRunLogsOK  %+v", 200, o.Payload)
}

func (o *GetWorkflowRunLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Log)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}