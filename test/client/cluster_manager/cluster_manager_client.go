// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new cluster manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddClusterNodes adds cluster nodes
*/
func (a *Client) AddClusterNodes(params *AddClusterNodesParams) (*AddClusterNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddClusterNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddClusterNodes",
		Method:             "POST",
		PathPattern:        "/v1/clusters/add_nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddClusterNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddClusterNodesOK), nil

}

/*
AttachKeyPairs attaches key pairs
*/
func (a *Client) AttachKeyPairs(params *AttachKeyPairsParams) (*AttachKeyPairsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachKeyPairsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AttachKeyPairs",
		Method:             "POST",
		PathPattern:        "/v1/clusters/key_pair/attach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AttachKeyPairsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AttachKeyPairsOK), nil

}

/*
CeaseClusters ceases clusters
*/
func (a *Client) CeaseClusters(params *CeaseClustersParams) (*CeaseClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCeaseClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CeaseClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/cease",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CeaseClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CeaseClustersOK), nil

}

/*
CreateCluster creates cluster
*/
func (a *Client) CreateCluster(params *CreateClusterParams) (*CreateClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateCluster",
		Method:             "POST",
		PathPattern:        "/v1/clusters/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterOK), nil

}

/*
CreateKeyPair creates key pair
*/
func (a *Client) CreateKeyPair(params *CreateKeyPairParams) (*CreateKeyPairOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateKeyPairParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateKeyPair",
		Method:             "POST",
		PathPattern:        "/v1/clusters/key_pairs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateKeyPairReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateKeyPairOK), nil

}

/*
DeleteClusterNodes deletes cluster nodes
*/
func (a *Client) DeleteClusterNodes(params *DeleteClusterNodesParams) (*DeleteClusterNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusterNodes",
		Method:             "POST",
		PathPattern:        "/v1/clusters/delete_nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClusterNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterNodesOK), nil

}

/*
DeleteClusters deletes clusters
*/
func (a *Client) DeleteClusters(params *DeleteClustersParams) (*DeleteClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/delete",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClustersOK), nil

}

/*
DeleteKeyPairs deletes key pairs
*/
func (a *Client) DeleteKeyPairs(params *DeleteKeyPairsParams) (*DeleteKeyPairsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKeyPairsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteKeyPairs",
		Method:             "DELETE",
		PathPattern:        "/v1/clusters/key_pairs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteKeyPairsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteKeyPairsOK), nil

}

/*
DescribeClusterNodes describes cluster nodes
*/
func (a *Client) DescribeClusterNodes(params *DescribeClusterNodesParams) (*DescribeClusterNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeClusterNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeClusterNodes",
		Method:             "GET",
		PathPattern:        "/v1/clusters/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeClusterNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeClusterNodesOK), nil

}

/*
DescribeClusters describes clusters
*/
func (a *Client) DescribeClusters(params *DescribeClustersParams) (*DescribeClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeClusters",
		Method:             "GET",
		PathPattern:        "/v1/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeClustersOK), nil

}

/*
DescribeKeyPairs describes key pairs
*/
func (a *Client) DescribeKeyPairs(params *DescribeKeyPairsParams) (*DescribeKeyPairsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeKeyPairsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeKeyPairs",
		Method:             "GET",
		PathPattern:        "/v1/clusters/key_pairs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeKeyPairsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeKeyPairsOK), nil

}

/*
DescribeSubnets describes subnets
*/
func (a *Client) DescribeSubnets(params *DescribeSubnetsParams) (*DescribeSubnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeSubnetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeSubnets",
		Method:             "GET",
		PathPattern:        "/v1/clusters/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeSubnetsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeSubnetsOK), nil

}

/*
DetachKeyPairs detaches key pairs
*/
func (a *Client) DetachKeyPairs(params *DetachKeyPairsParams) (*DetachKeyPairsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachKeyPairsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DetachKeyPairs",
		Method:             "POST",
		PathPattern:        "/v1/clusters/key_pair/detach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DetachKeyPairsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetachKeyPairsOK), nil

}

/*
GetClusterStatistics gets cluster statistics
*/
func (a *Client) GetClusterStatistics(params *GetClusterStatisticsParams) (*GetClusterStatisticsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterStatisticsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetClusterStatistics",
		Method:             "GET",
		PathPattern:        "/v1/clusters/statistics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterStatisticsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterStatisticsOK), nil

}

/*
ModifyClusterAttributes modifies cluster attributes
*/
func (a *Client) ModifyClusterAttributes(params *ModifyClusterAttributesParams) (*ModifyClusterAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyClusterAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyClusterAttributes",
		Method:             "POST",
		PathPattern:        "/v1/clusters/modify",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyClusterAttributesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyClusterAttributesOK), nil

}

/*
ModifyClusterNodeAttributes modifies cluster node attributes
*/
func (a *Client) ModifyClusterNodeAttributes(params *ModifyClusterNodeAttributesParams) (*ModifyClusterNodeAttributesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyClusterNodeAttributesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyClusterNodeAttributes",
		Method:             "POST",
		PathPattern:        "/v1/clusters/modify_nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyClusterNodeAttributesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyClusterNodeAttributesOK), nil

}

/*
RecoverClusters recovers clusters
*/
func (a *Client) RecoverClusters(params *RecoverClustersParams) (*RecoverClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRecoverClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RecoverClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/recover",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RecoverClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RecoverClustersOK), nil

}

/*
ResizeCluster resizes cluster
*/
func (a *Client) ResizeCluster(params *ResizeClusterParams) (*ResizeClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResizeClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ResizeCluster",
		Method:             "POST",
		PathPattern:        "/v1/clusters/resize",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ResizeClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResizeClusterOK), nil

}

/*
RollbackCluster rollbacks cluster
*/
func (a *Client) RollbackCluster(params *RollbackClusterParams) (*RollbackClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRollbackClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RollbackCluster",
		Method:             "POST",
		PathPattern:        "/v1/clusters/rollback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RollbackClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RollbackClusterOK), nil

}

/*
StartClusters starts clusters
*/
func (a *Client) StartClusters(params *StartClustersParams) (*StartClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StartClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StartClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartClustersOK), nil

}

/*
StopClusters stops clusters
*/
func (a *Client) StopClusters(params *StopClustersParams) (*StopClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "StopClusters",
		Method:             "POST",
		PathPattern:        "/v1/clusters/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StopClustersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopClustersOK), nil

}

/*
UpdateClusterEnv updates cluster env
*/
func (a *Client) UpdateClusterEnv(params *UpdateClusterEnvParams) (*UpdateClusterEnvOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateClusterEnvParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateClusterEnv",
		Method:             "PATCH",
		PathPattern:        "/v1/clusters/update_env",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateClusterEnvReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateClusterEnvOK), nil

}

/*
UpgradeCluster upgrades cluster
*/
func (a *Client) UpgradeCluster(params *UpgradeClusterParams) (*UpgradeClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpgradeCluster",
		Method:             "POST",
		PathPattern:        "/v1/clusters/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpgradeClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeClusterOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
