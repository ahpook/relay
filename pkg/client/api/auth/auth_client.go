// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new auth API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateSession requests an j w t access token from email and password
*/
func (a *Client) CreateSession(params *CreateSessionParams) (*CreateSessionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSession",
		Method:             "POST",
		PathPattern:        "/auth/sessions",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSessionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSessionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSessionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteSession invalidates attached j w t token used likely for logout
*/
func (a *Client) DeleteSession(params *DeleteSessionParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSession",
		Method:             "DELETE",
		PathPattern:        "/auth/sessions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSessionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSessionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ForgotPassword requests a password reset for a user identified by their email address
*/
func (a *Client) ForgotPassword(params *ForgotPasswordParams) (*ForgotPasswordAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewForgotPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "forgotPassword",
		Method:             "POST",
		PathPattern:        "/forgot-password",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ForgotPasswordReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ForgotPasswordAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ForgotPasswordDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProfile gets the currently authenticated user s profile
*/
func (a *Client) GetProfile(params *GetProfileParams, authInfo runtime.ClientAuthInfoWriter) (*GetProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProfile",
		Method:             "GET",
		PathPattern:        "/auth/profile",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetProfilePreferences gets the current user s public preferences
*/
func (a *Client) GetProfilePreferences(params *GetProfilePreferencesParams, authInfo runtime.ClientAuthInfoWriter) (*GetProfilePreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProfilePreferencesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProfilePreferences",
		Method:             "GET",
		PathPattern:        "/auth/profile/preferences",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProfilePreferencesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetProfilePreferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetProfilePreferencesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSession checks if attached j w t token is valid
*/
func (a *Client) GetSession(params *GetSessionParams, authInfo runtime.ClientAuthInfoWriter) (*GetSessionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSessionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSession",
		Method:             "GET",
		PathPattern:        "/auth/sessions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSessionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSessionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSessionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateProfile updates the currently authenticated user s profile
*/
func (a *Client) UpdateProfile(params *UpdateProfileParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProfile",
		Method:             "PUT",
		PathPattern:        "/auth/profile",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateProfilePassword updates your password
*/
func (a *Client) UpdateProfilePassword(params *UpdateProfilePasswordParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProfilePasswordOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProfilePasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProfilePassword",
		Method:             "PUT",
		PathPattern:        "/auth/profile/password",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProfilePasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProfilePasswordOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateProfilePasswordDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateProfilePreferences sets the current user s public preferences
*/
func (a *Client) UpdateProfilePreferences(params *UpdateProfilePreferencesParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProfilePreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProfilePreferencesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProfilePreferences",
		Method:             "PUT",
		PathPattern:        "/auth/profile/preferences",
		ProducesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		ConsumesMediaTypes: []string{"application/vnd.puppet.nebula.v20191223+json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProfilePreferencesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProfilePreferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateProfilePreferencesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
