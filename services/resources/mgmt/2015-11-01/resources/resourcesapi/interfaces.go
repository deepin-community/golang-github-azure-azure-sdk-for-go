// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package resourcesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2015-11-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

// DeploymentsClientAPI contains the set of methods on the DeploymentsClient type.
type DeploymentsClientAPI interface {
	CalculateTemplateHash(ctx context.Context, templateParameter interface{}) (result resources.TemplateHashResult, err error)
	Cancel(ctx context.Context, resourceGroupName string, deploymentName string) (result autorest.Response, err error)
	CheckExistence(ctx context.Context, resourceGroupName string, deploymentName string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, deploymentName string, parameters resources.Deployment) (result resources.DeploymentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, deploymentName string) (result resources.DeploymentsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, deploymentName string) (result resources.DeploymentExtended, err error)
	List(ctx context.Context, resourceGroupName string, filter string, top *int32) (result resources.DeploymentListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, filter string, top *int32) (result resources.DeploymentListResultIterator, err error)
	Validate(ctx context.Context, resourceGroupName string, deploymentName string, parameters resources.Deployment) (result resources.DeploymentValidateResult, err error)
}

var _ DeploymentsClientAPI = (*resources.DeploymentsClient)(nil)

// ProvidersClientAPI contains the set of methods on the ProvidersClient type.
type ProvidersClientAPI interface {
	Get(ctx context.Context, resourceProviderNamespace string) (result resources.Provider, err error)
	List(ctx context.Context, top *int32) (result resources.ProviderListResultPage, err error)
	ListComplete(ctx context.Context, top *int32) (result resources.ProviderListResultIterator, err error)
	Register(ctx context.Context, resourceProviderNamespace string) (result resources.Provider, err error)
	Unregister(ctx context.Context, resourceProviderNamespace string) (result resources.Provider, err error)
}

var _ ProvidersClientAPI = (*resources.ProvidersClient)(nil)

// GroupsClientAPI contains the set of methods on the GroupsClient type.
type GroupsClientAPI interface {
	CheckExistence(ctx context.Context, resourceGroupName string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, parameters resources.Group) (result resources.Group, err error)
	Delete(ctx context.Context, resourceGroupName string) (result resources.GroupsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string) (result resources.Group, err error)
	List(ctx context.Context, filter string, top *int32) (result resources.GroupListResultPage, err error)
	ListComplete(ctx context.Context, filter string, top *int32) (result resources.GroupListResultIterator, err error)
	ListResources(ctx context.Context, resourceGroupName string, filter string, expand string, top *int32) (result resources.ListResultPage, err error)
	ListResourcesComplete(ctx context.Context, resourceGroupName string, filter string, expand string, top *int32) (result resources.ListResultIterator, err error)
	Patch(ctx context.Context, resourceGroupName string, parameters resources.Group) (result resources.Group, err error)
}

var _ GroupsClientAPI = (*resources.GroupsClient)(nil)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	CheckExistence(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string) (result autorest.Response, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string, parameters resources.GenericResource) (result resources.GenericResource, err error)
	Delete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string) (result resources.GenericResource, err error)
	List(ctx context.Context, filter string, expand string, top *int32) (result resources.ListResultPage, err error)
	ListComplete(ctx context.Context, filter string, expand string, top *int32) (result resources.ListResultIterator, err error)
	MoveResources(ctx context.Context, sourceResourceGroupName string, parameters resources.MoveInfo) (result resources.MoveResourcesFuture, err error)
	Update(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, APIVersion string, parameters resources.GenericResource) (result resources.UpdateFuture, err error)
}

var _ ClientAPI = (*resources.Client)(nil)

// TagsClientAPI contains the set of methods on the TagsClient type.
type TagsClientAPI interface {
	CreateOrUpdate(ctx context.Context, tagName string) (result resources.TagDetails, err error)
	CreateOrUpdateValue(ctx context.Context, tagName string, tagValue string) (result resources.TagValue, err error)
	Delete(ctx context.Context, tagName string) (result autorest.Response, err error)
	DeleteValue(ctx context.Context, tagName string, tagValue string) (result autorest.Response, err error)
	List(ctx context.Context) (result resources.TagsListResultPage, err error)
	ListComplete(ctx context.Context) (result resources.TagsListResultIterator, err error)
}

var _ TagsClientAPI = (*resources.TagsClient)(nil)

// DeploymentOperationsClientAPI contains the set of methods on the DeploymentOperationsClient type.
type DeploymentOperationsClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, deploymentName string, operationID string) (result resources.DeploymentOperation, err error)
	List(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result resources.DeploymentOperationsListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, deploymentName string, top *int32) (result resources.DeploymentOperationsListResultIterator, err error)
}

var _ DeploymentOperationsClientAPI = (*resources.DeploymentOperationsClient)(nil)

// ProviderOperationDetailsClientAPI contains the set of methods on the ProviderOperationDetailsClient type.
type ProviderOperationDetailsClientAPI interface {
	List(ctx context.Context, resourceProviderNamespace string, APIVersion string) (result resources.ProviderOperationDetailListResultPage, err error)
	ListComplete(ctx context.Context, resourceProviderNamespace string, APIVersion string) (result resources.ProviderOperationDetailListResultIterator, err error)
}

var _ ProviderOperationDetailsClientAPI = (*resources.ProviderOperationDetailsClient)(nil)

// PolicyDefinitionsClientAPI contains the set of methods on the PolicyDefinitionsClient type.
type PolicyDefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters resources.PolicyDefinition) (result resources.PolicyDefinition, err error)
	Delete(ctx context.Context, policyDefinitionName string) (result autorest.Response, err error)
	Get(ctx context.Context, policyDefinitionName string) (result resources.PolicyDefinition, err error)
}

var _ PolicyDefinitionsClientAPI = (*resources.PolicyDefinitionsClient)(nil)

// PolicyAssignmentsClientAPI contains the set of methods on the PolicyAssignmentsClient type.
type PolicyAssignmentsClientAPI interface {
	Create(ctx context.Context, scope string, policyAssignmentName string, parameters resources.PolicyAssignment) (result resources.PolicyAssignment, err error)
	CreateByID(ctx context.Context, policyAssignmentID string, parameters resources.PolicyAssignment) (result resources.PolicyAssignment, err error)
	Delete(ctx context.Context, scope string, policyAssignmentName string) (result resources.PolicyAssignment, err error)
	DeleteByID(ctx context.Context, policyAssignmentID string) (result resources.PolicyAssignment, err error)
	Get(ctx context.Context, scope string, policyAssignmentName string) (result resources.PolicyAssignment, err error)
	GetByID(ctx context.Context, policyAssignmentID string) (result resources.PolicyAssignment, err error)
	List(ctx context.Context, filter string) (result resources.PolicyAssignmentListResultPage, err error)
	ListComplete(ctx context.Context, filter string) (result resources.PolicyAssignmentListResultIterator, err error)
	ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result resources.PolicyAssignmentListResultPage, err error)
	ListForResourceComplete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result resources.PolicyAssignmentListResultIterator, err error)
	ListForResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result resources.PolicyAssignmentListResultPage, err error)
	ListForResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result resources.PolicyAssignmentListResultIterator, err error)
	ListForScope(ctx context.Context, scope string, filter string) (result resources.PolicyAssignmentListResultPage, err error)
	ListForScopeComplete(ctx context.Context, scope string, filter string) (result resources.PolicyAssignmentListResultIterator, err error)
}

var _ PolicyAssignmentsClientAPI = (*resources.PolicyAssignmentsClient)(nil)