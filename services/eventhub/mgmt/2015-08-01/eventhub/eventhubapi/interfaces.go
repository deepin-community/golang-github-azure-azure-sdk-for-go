// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package eventhubapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2015-08-01/eventhub"
	"github.com/Azure/go-autorest/autorest"
)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result eventhub.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result eventhub.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*eventhub.OperationsClient)(nil)

// NamespacesClientAPI contains the set of methods on the NamespacesClient type.
type NamespacesClientAPI interface {
	CheckNameAvailability(ctx context.Context, parameters eventhub.CheckNameAvailabilityParameter) (result eventhub.CheckNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, parameters eventhub.NamespaceCreateOrUpdateParameters) (result eventhub.NamespacesCreateOrUpdateFuture, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters eventhub.SharedAccessAuthorizationRuleCreateOrUpdateParameters) (result eventhub.SharedAccessAuthorizationRuleResource, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NamespacesDeleteFuture, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.NamespaceResource, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result eventhub.SharedAccessAuthorizationRuleResource, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.SharedAccessAuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.SharedAccessAuthorizationRuleListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result eventhub.NamespaceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result eventhub.NamespaceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result eventhub.NamespaceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result eventhub.NamespaceListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string) (result eventhub.ResourceListKeys, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters eventhub.RegenerateKeysParameters) (result eventhub.ResourceListKeys, err error)
	Update(ctx context.Context, resourceGroupName string, namespaceName string, parameters eventhub.NamespaceUpdateParameter) (result eventhub.NamespaceResource, err error)
}

var _ NamespacesClientAPI = (*eventhub.NamespacesClient)(nil)

// EventHubsClientAPI contains the set of methods on the EventHubsClient type.
type EventHubsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, parameters eventhub.CreateOrUpdateParameters) (result eventhub.ResourceType, err error)
	CreateOrUpdateAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.SharedAccessAuthorizationRuleCreateOrUpdateParameters) (result eventhub.SharedAccessAuthorizationRuleResource, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result autorest.Response, err error)
	DeleteAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result eventhub.ResourceType, err error)
	GetAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result eventhub.SharedAccessAuthorizationRuleResource, err error)
	ListAll(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.ListResultPage, err error)
	ListAllComplete(ctx context.Context, resourceGroupName string, namespaceName string) (result eventhub.ListResultIterator, err error)
	ListAuthorizationRules(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result eventhub.SharedAccessAuthorizationRuleListResultPage, err error)
	ListAuthorizationRulesComplete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result eventhub.SharedAccessAuthorizationRuleListResultIterator, err error)
	ListKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result eventhub.ResourceListKeys, err error)
	PostAuthorizationRule(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string) (result eventhub.SharedAccessAuthorizationRuleResource, err error)
	RegenerateKeys(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters eventhub.RegenerateKeysParameters) (result eventhub.ResourceListKeys, err error)
}

var _ EventHubsClientAPI = (*eventhub.EventHubsClient)(nil)

// ConsumerGroupsClientAPI contains the set of methods on the ConsumerGroupsClient type.
type ConsumerGroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, parameters eventhub.ConsumerGroupCreateOrUpdateParameters) (result eventhub.ConsumerGroupResource, err error)
	Delete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result eventhub.ConsumerGroupResource, err error)
	ListAll(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result eventhub.ConsumerGroupListResultPage, err error)
	ListAllComplete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string) (result eventhub.ConsumerGroupListResultIterator, err error)
}

var _ ConsumerGroupsClientAPI = (*eventhub.ConsumerGroupsClient)(nil)
