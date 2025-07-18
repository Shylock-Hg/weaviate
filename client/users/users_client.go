//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ActivateUser(params *ActivateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ActivateUserOK, error)

	CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserCreated, error)

	DeactivateUser(params *DeactivateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeactivateUserOK, error)

	DeleteUser(params *DeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserNoContent, error)

	GetOwnInfo(params *GetOwnInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOwnInfoOK, error)

	GetUserInfo(params *GetUserInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserInfoOK, error)

	ListAllUsers(params *ListAllUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAllUsersOK, error)

	RotateUserAPIKey(params *RotateUserAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RotateUserAPIKeyOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ActivateUser activates a deactivated user
*/
func (a *Client) ActivateUser(params *ActivateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ActivateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewActivateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "activateUser",
		Method:             "POST",
		PathPattern:        "/users/db/{user_id}/activate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ActivateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ActivateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for activateUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateUser creates new user
*/
func (a *Client) CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateUserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "POST",
		PathPattern:        "/users/db/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateUserCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeactivateUser deactivates a user
*/
func (a *Client) DeactivateUser(params *DeactivateUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeactivateUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeactivateUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deactivateUser",
		Method:             "POST",
		PathPattern:        "/users/db/{user_id}/deactivate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeactivateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeactivateUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deactivateUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteUser deletes user
*/
func (a *Client) DeleteUser(params *DeleteUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUser",
		Method:             "DELETE",
		PathPattern:        "/users/db/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUserNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetOwnInfo gets info relevant to own user e g username roles
*/
func (a *Client) GetOwnInfo(params *GetOwnInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOwnInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOwnInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getOwnInfo",
		Method:             "GET",
		PathPattern:        "/users/own-info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOwnInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOwnInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getOwnInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetUserInfo gets info relevant to user e g username roles
*/
func (a *Client) GetUserInfo(params *GetUserInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUserInfo",
		Method:             "GET",
		PathPattern:        "/users/db/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUserInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListAllUsers lists all db users
*/
func (a *Client) ListAllUsers(params *ListAllUsersParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListAllUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAllUsersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listAllUsers",
		Method:             "GET",
		PathPattern:        "/users/db",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAllUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAllUsersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAllUsers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RotateUserAPIKey rotates user api key
*/
func (a *Client) RotateUserAPIKey(params *RotateUserAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RotateUserAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRotateUserAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "rotateUserApiKey",
		Method:             "POST",
		PathPattern:        "/users/db/{user_id}/rotate-key",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RotateUserAPIKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RotateUserAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for rotateUserApiKey: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
