// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package containerinstanceapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/containerinstance/mgmt/2017-10-01-preview/containerinstance"
)

// ContainerGroupsClientAPI contains the set of methods on the ContainerGroupsClient type.
type ContainerGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, containerGroupName string, containerGroup containerinstance.ContainerGroup) (result containerinstance.ContainerGroup, err error)
	Delete(ctx context.Context, resourceGroupName string, containerGroupName string) (result containerinstance.ContainerGroup, err error)
	Get(ctx context.Context, resourceGroupName string, containerGroupName string) (result containerinstance.ContainerGroup, err error)
	List(ctx context.Context) (result containerinstance.ContainerGroupListResultPage, err error)
	ListComplete(ctx context.Context) (result containerinstance.ContainerGroupListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerinstance.ContainerGroupListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerinstance.ContainerGroupListResultIterator, err error)
}

var _ ContainerGroupsClientAPI = (*containerinstance.ContainerGroupsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result containerinstance.OperationListResult, err error)
}

var _ OperationsClientAPI = (*containerinstance.OperationsClient)(nil)

// ContainerLogsClientAPI contains the set of methods on the ContainerLogsClient type.
type ContainerLogsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, tail *int32) (result containerinstance.Logs, err error)
}

var _ ContainerLogsClientAPI = (*containerinstance.ContainerLogsClient)(nil)