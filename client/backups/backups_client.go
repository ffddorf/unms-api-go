// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new backups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for backups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteNmsBackupsBackupid(params *DeleteNmsBackupsBackupidParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNmsBackupsBackupidOK, error)

	GetNmsBackups(params *GetNmsBackupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNmsBackupsOK, error)

	GetNmsBackupsBackupid(params *GetNmsBackupsBackupidParams, authInfo runtime.ClientAuthInfoWriter) error

	PostNmsBackupsBackupidRestore(params *PostNmsBackupsBackupidRestoreParams, authInfo runtime.ClientAuthInfoWriter) (*PostNmsBackupsBackupidRestoreOK, error)

	PostNmsBackupsCreate(params *PostNmsBackupsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*PostNmsBackupsCreateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteNmsBackupsBackupid deletes u n m s backup
*/
func (a *Client) DeleteNmsBackupsBackupid(params *DeleteNmsBackupsBackupidParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNmsBackupsBackupidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNmsBackupsBackupidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNmsBackupsBackupid",
		Method:             "DELETE",
		PathPattern:        "/nms/backups/{backupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNmsBackupsBackupidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNmsBackupsBackupidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNmsBackupsBackupid: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNmsBackups gets u n m s backups
*/
func (a *Client) GetNmsBackups(params *GetNmsBackupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetNmsBackupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNmsBackupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNmsBackups",
		Method:             "GET",
		PathPattern:        "/nms/backups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNmsBackupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNmsBackupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNmsBackups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNmsBackupsBackupid gets u n m s backup file
*/
func (a *Client) GetNmsBackupsBackupid(params *GetNmsBackupsBackupidParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNmsBackupsBackupidParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNmsBackupsBackupid",
		Method:             "GET",
		PathPattern:        "/nms/backups/{backupId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNmsBackupsBackupidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  PostNmsBackupsBackupidRestore restores u n m s backup
*/
func (a *Client) PostNmsBackupsBackupidRestore(params *PostNmsBackupsBackupidRestoreParams, authInfo runtime.ClientAuthInfoWriter) (*PostNmsBackupsBackupidRestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNmsBackupsBackupidRestoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postNmsBackupsBackupidRestore",
		Method:             "POST",
		PathPattern:        "/nms/backups/{backupId}/restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostNmsBackupsBackupidRestoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNmsBackupsBackupidRestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postNmsBackupsBackupidRestore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostNmsBackupsCreate creates u n m s backup
*/
func (a *Client) PostNmsBackupsCreate(params *PostNmsBackupsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*PostNmsBackupsCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNmsBackupsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postNmsBackupsCreate",
		Method:             "POST",
		PathPattern:        "/nms/backups/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostNmsBackupsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNmsBackupsCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postNmsBackupsCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}