// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
//
// Package compute implements the Azure ARM Compute service API version 2017-08-01-preview.
//
// These APIs allow end users to operate on Azure Machine Learning Compute resources. They support the following
// operations:<ul><li>Create or update a cluster</li><li>Get a cluster</li><li>Patch a cluster</li><li>Delete a
// cluster</li><li>Get keys for a cluster</li><li>Check if updates are available for system services in a
// cluster</li><li>Update system services in a cluster</li><li>Get all clusters in a resource group</li><li>Get all
// clusters in a subscription</li></ul>
package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Compute
	DefaultBaseURI = "https://management.azure.com"
)

// BaseClient is the base client for Compute.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}
