// Code generated by go-swagger; DO NOT EDIT.

package workflow_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new workflow secrets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for workflow secrets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateWorkflowSecret adds a new secret to the given workflow
*/
func (a *Client) CreateWorkflowSecret(params *CreateWorkflowSecretParams, authInfo runtime.ClientAuthInfoWriter) (*CreateWorkflowSecretCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateWorkflowSecretParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createWorkflowSecret",
		Method:             "POST",
		PathPattern:        "/api/workflows/{workflowName}/secrets",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateWorkflowSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateWorkflowSecretCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateWorkflowSecretDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteWorkflowSecret deletes the secret associated with the given workflow and secret key
*/
func (a *Client) DeleteWorkflowSecret(params *DeleteWorkflowSecretParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteWorkflowSecretOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkflowSecretParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWorkflowSecret",
		Method:             "DELETE",
		PathPattern:        "/api/workflows/{workflowName}/secrets/{workflowSecretKey}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWorkflowSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteWorkflowSecretOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteWorkflowSecretDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListWorkflowSecrets gets all secrets associated with the given workflow
*/
func (a *Client) ListWorkflowSecrets(params *ListWorkflowSecretsParams, authInfo runtime.ClientAuthInfoWriter) (*ListWorkflowSecretsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListWorkflowSecretsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listWorkflowSecrets",
		Method:             "GET",
		PathPattern:        "/api/workflows/{workflowName}/secrets",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListWorkflowSecretsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListWorkflowSecretsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListWorkflowSecretsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateWorkflowSecret updates the secret associated with the given workflow and secret key
*/
func (a *Client) UpdateWorkflowSecret(params *UpdateWorkflowSecretParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateWorkflowSecretOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateWorkflowSecretParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateWorkflowSecret",
		Method:             "PUT",
		PathPattern:        "/api/workflows/{workflowName}/secrets/{workflowSecretKey}",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateWorkflowSecretReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateWorkflowSecretOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateWorkflowSecretDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
