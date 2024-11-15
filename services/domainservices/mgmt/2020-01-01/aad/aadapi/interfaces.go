// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package aadapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/domainservices/mgmt/2020-01-01/aad"
)

// DomainServiceOperationsClientAPI contains the set of methods on the DomainServiceOperationsClient type.
type DomainServiceOperationsClientAPI interface {
	List(ctx context.Context) (result aad.OperationEntityListResultPage, err error)
	ListComplete(ctx context.Context) (result aad.OperationEntityListResultIterator, err error)
}

var _ DomainServiceOperationsClientAPI = (*aad.DomainServiceOperationsClient)(nil)

// DomainServicesClientAPI contains the set of methods on the DomainServicesClient type.
type DomainServicesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, domainServiceName string, domainService aad.DomainService) (result aad.DomainServicesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, domainServiceName string) (result aad.DomainServicesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, domainServiceName string) (result aad.DomainService, err error)
	List(ctx context.Context) (result aad.DomainServiceListResultPage, err error)
	ListComplete(ctx context.Context) (result aad.DomainServiceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result aad.DomainServiceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result aad.DomainServiceListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, domainServiceName string, domainService aad.DomainService) (result aad.DomainServicesUpdateFuture, err error)
}

var _ DomainServicesClientAPI = (*aad.DomainServicesClient)(nil)

// OuContainerOperationsClientAPI contains the set of methods on the OuContainerOperationsClient type.
type OuContainerOperationsClientAPI interface {
	List(ctx context.Context) (result aad.OperationEntityListResultPage, err error)
	ListComplete(ctx context.Context) (result aad.OperationEntityListResultIterator, err error)
}

var _ OuContainerOperationsClientAPI = (*aad.OuContainerOperationsClient)(nil)

// OuContainerClientAPI contains the set of methods on the OuContainerClient type.
type OuContainerClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount aad.ContainerAccount) (result aad.OuContainerCreateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string) (result aad.OuContainerDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string) (result aad.OuContainer, err error)
	List(ctx context.Context, resourceGroupName string, domainServiceName string) (result aad.OuContainerListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, domainServiceName string) (result aad.OuContainerListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, domainServiceName string, ouContainerName string, containerAccount aad.ContainerAccount) (result aad.OuContainerUpdateFuture, err error)
}

var _ OuContainerClientAPI = (*aad.OuContainerClient)(nil)
