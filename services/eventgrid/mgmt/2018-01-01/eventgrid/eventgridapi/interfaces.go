// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventgrid/armeventgrid). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package eventgridapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/eventgrid/mgmt/2018-01-01/eventgrid"
)

// EventSubscriptionsClientAPI contains the set of methods on the EventSubscriptionsClient type.
type EventSubscriptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, eventSubscriptionName string, eventSubscriptionInfo eventgrid.EventSubscription) (result eventgrid.EventSubscriptionsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.EventSubscriptionsDeleteFuture, err error)
	Get(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.EventSubscription, err error)
	GetFullURL(ctx context.Context, scope string, eventSubscriptionName string) (result eventgrid.EventSubscriptionFullURL, err error)
	ListByResource(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string) (result eventgrid.EventSubscriptionsListResult, err error)
	ListGlobalByResourceGroup(ctx context.Context, resourceGroupName string) (result eventgrid.EventSubscriptionsListResult, err error)
	ListGlobalByResourceGroupForTopicType(ctx context.Context, resourceGroupName string, topicTypeName string) (result eventgrid.EventSubscriptionsListResult, err error)
	ListGlobalBySubscription(ctx context.Context) (result eventgrid.EventSubscriptionsListResult, err error)
	ListGlobalBySubscriptionForTopicType(ctx context.Context, topicTypeName string) (result eventgrid.EventSubscriptionsListResult, err error)
	ListRegionalByResourceGroup(ctx context.Context, resourceGroupName string, location string) (result eventgrid.EventSubscriptionsListResult, err error)
	ListRegionalByResourceGroupForTopicType(ctx context.Context, resourceGroupName string, location string, topicTypeName string) (result eventgrid.EventSubscriptionsListResult, err error)
	ListRegionalBySubscription(ctx context.Context, location string) (result eventgrid.EventSubscriptionsListResult, err error)
	ListRegionalBySubscriptionForTopicType(ctx context.Context, location string, topicTypeName string) (result eventgrid.EventSubscriptionsListResult, err error)
	Update(ctx context.Context, scope string, eventSubscriptionName string, eventSubscriptionUpdateParameters eventgrid.EventSubscriptionUpdateParameters) (result eventgrid.EventSubscriptionsUpdateFuture, err error)
}

var _ EventSubscriptionsClientAPI = (*eventgrid.EventSubscriptionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result eventgrid.OperationsListResult, err error)
}

var _ OperationsClientAPI = (*eventgrid.OperationsClient)(nil)

// TopicsClientAPI contains the set of methods on the TopicsClient type.
type TopicsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, topicName string, topicInfo eventgrid.Topic) (result eventgrid.TopicsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, topicName string) (result eventgrid.TopicsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, topicName string) (result eventgrid.Topic, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result eventgrid.TopicsListResult, err error)
	ListBySubscription(ctx context.Context) (result eventgrid.TopicsListResult, err error)
	ListEventTypes(ctx context.Context, resourceGroupName string, providerNamespace string, resourceTypeName string, resourceName string) (result eventgrid.EventTypesListResult, err error)
	ListSharedAccessKeys(ctx context.Context, resourceGroupName string, topicName string) (result eventgrid.TopicSharedAccessKeys, err error)
	RegenerateKey(ctx context.Context, resourceGroupName string, topicName string, regenerateKeyRequest eventgrid.TopicRegenerateKeyRequest) (result eventgrid.TopicSharedAccessKeys, err error)
	Update(ctx context.Context, resourceGroupName string, topicName string, topicUpdateParameters eventgrid.TopicUpdateParameters) (result eventgrid.TopicsUpdateFuture, err error)
}

var _ TopicsClientAPI = (*eventgrid.TopicsClient)(nil)

// TopicTypesClientAPI contains the set of methods on the TopicTypesClient type.
type TopicTypesClientAPI interface {
	Get(ctx context.Context, topicTypeName string) (result eventgrid.TopicTypeInfo, err error)
	List(ctx context.Context) (result eventgrid.TopicTypesListResult, err error)
	ListEventTypes(ctx context.Context, topicTypeName string) (result eventgrid.EventTypesListResult, err error)
}

var _ TopicTypesClientAPI = (*eventgrid.TopicTypesClient)(nil)
