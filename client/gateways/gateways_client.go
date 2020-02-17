// Code generated by go-swagger; DO NOT EDIT.

package gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new gateways API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for gateways API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteGatewaysID(params *DeleteGatewaysIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteGatewaysIDOK, error)

	GetGateways(params *GetGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*GetGatewaysOK, error)

	PostGateways(params *PostGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*PostGatewaysOK, error)

	PutGatewaysID(params *PutGatewaysIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutGatewaysIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteGatewaysID deletes a gateway
*/
func (a *Client) DeleteGatewaysID(params *DeleteGatewaysIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteGatewaysIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGatewaysIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGatewaysId",
		Method:             "DELETE",
		PathPattern:        "/gateways/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGatewaysIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteGatewaysIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteGatewaysId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGateways gets all gateways
*/
func (a *Client) GetGateways(params *GetGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*GetGatewaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGatewaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGateways",
		Method:             "GET",
		PathPattern:        "/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGatewaysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetGatewaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getGateways: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostGateways creates a new gateway
*/
func (a *Client) PostGateways(params *PostGatewaysParams, authInfo runtime.ClientAuthInfoWriter) (*PostGatewaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostGatewaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postGateways",
		Method:             "POST",
		PathPattern:        "/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostGatewaysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostGatewaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postGateways: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutGatewaysID updates a gateway
*/
func (a *Client) PutGatewaysID(params *PutGatewaysIDParams, authInfo runtime.ClientAuthInfoWriter) (*PutGatewaysIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutGatewaysIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putGatewaysId",
		Method:             "PUT",
		PathPattern:        "/gateways/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutGatewaysIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutGatewaysIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putGatewaysId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}