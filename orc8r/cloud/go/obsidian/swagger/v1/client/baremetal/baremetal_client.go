// Code generated by go-swagger; DO NOT EDIT.

package baremetal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new baremetal API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for baremetal API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteCiNodesNodeID deletes a c i node
*/
func (a *Client) DeleteCiNodesNodeID(params *DeleteCiNodesNodeIDParams) (*DeleteCiNodesNodeIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCiNodesNodeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteCiNodesNodeID",
		Method:             "DELETE",
		PathPattern:        "/ci/nodes/{node_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCiNodesNodeIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCiNodesNodeIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteCiNodesNodeIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCiNodes lists c i worker nodes
*/
func (a *Client) GetCiNodes(params *GetCiNodesParams) (*GetCiNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCiNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCiNodes",
		Method:             "GET",
		PathPattern:        "/ci/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCiNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCiNodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCiNodesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetCiNodesNodeID gets a specific c i node
*/
func (a *Client) GetCiNodesNodeID(params *GetCiNodesNodeIDParams) (*GetCiNodesNodeIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCiNodesNodeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCiNodesNodeID",
		Method:             "GET",
		PathPattern:        "/ci/nodes/{node_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCiNodesNodeIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCiNodesNodeIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetCiNodesNodeIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostCiNodes creates new available c i worker node
*/
func (a *Client) PostCiNodes(params *PostCiNodesParams) (*PostCiNodesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCiNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCiNodes",
		Method:             "POST",
		PathPattern:        "/ci/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCiNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCiNodesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCiNodesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostCiNodesNodeIDRelease releases a manually reserved c i node
*/
func (a *Client) PostCiNodesNodeIDRelease(params *PostCiNodesNodeIDReleaseParams) (*PostCiNodesNodeIDReleaseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCiNodesNodeIDReleaseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCiNodesNodeIDRelease",
		Method:             "POST",
		PathPattern:        "/ci/nodes/{node_id}/release",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCiNodesNodeIDReleaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCiNodesNodeIDReleaseNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCiNodesNodeIDReleaseDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostCiNodesNodeIDReleaseLeaseID releases a c i worker node
*/
func (a *Client) PostCiNodesNodeIDReleaseLeaseID(params *PostCiNodesNodeIDReleaseLeaseIDParams) (*PostCiNodesNodeIDReleaseLeaseIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCiNodesNodeIDReleaseLeaseIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCiNodesNodeIDReleaseLeaseID",
		Method:             "POST",
		PathPattern:        "/ci/nodes/{node_id}/release/{lease_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCiNodesNodeIDReleaseLeaseIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCiNodesNodeIDReleaseLeaseIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCiNodesNodeIDReleaseLeaseIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostCiNodesNodeIDReserve manuallies reserve a specific c i node
*/
func (a *Client) PostCiNodesNodeIDReserve(params *PostCiNodesNodeIDReserveParams) (*PostCiNodesNodeIDReserveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCiNodesNodeIDReserveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCiNodesNodeIDReserve",
		Method:             "POST",
		PathPattern:        "/ci/nodes/{node_id}/reserve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCiNodesNodeIDReserveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCiNodesNodeIDReserveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCiNodesNodeIDReserveDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostCiReserve reserves a c i worker node if available
*/
func (a *Client) PostCiReserve(params *PostCiReserveParams) (*PostCiReserveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostCiReserveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostCiReserve",
		Method:             "POST",
		PathPattern:        "/ci/reserve",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostCiReserveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostCiReserveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostCiReserveDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PutCiNodesNodeID updates a c i node
*/
func (a *Client) PutCiNodesNodeID(params *PutCiNodesNodeIDParams) (*PutCiNodesNodeIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutCiNodesNodeIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutCiNodesNodeID",
		Method:             "PUT",
		PathPattern:        "/ci/nodes/{node_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutCiNodesNodeIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutCiNodesNodeIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PutCiNodesNodeIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}