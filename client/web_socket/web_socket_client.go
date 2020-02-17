// Code generated by go-swagger; DO NOT EDIT.

package web_socket

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new web socket API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for web socket API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	PostWsapiConnection(params *PostWsapiConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*PostWsapiConnectionCreated, error)

	PostWsapiSubscription(params *PostWsapiSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PostWsapiSubscriptionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PostWsapiConnection creates websocket connection session
*/
func (a *Client) PostWsapiConnection(params *PostWsapiConnectionParams, authInfo runtime.ClientAuthInfoWriter) (*PostWsapiConnectionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWsapiConnectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postWsapiConnection",
		Method:             "POST",
		PathPattern:        "/ws-api/connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWsapiConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWsapiConnectionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postWsapiConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostWsapiSubscription updates websocket subscriptions
*/
func (a *Client) PostWsapiSubscription(params *PostWsapiSubscriptionParams, authInfo runtime.ClientAuthInfoWriter) (*PostWsapiSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWsapiSubscriptionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postWsapiSubscription",
		Method:             "POST",
		PathPattern:        "/ws-api/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWsapiSubscriptionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostWsapiSubscriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postWsapiSubscription: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
