// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scheduler/armscheduler](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scheduler/armscheduler). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package schedulerapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/scheduler/mgmt/2016-03-01/scheduler"
	"github.com/Azure/go-autorest/autorest"
)

// JobCollectionsClientAPI contains the set of methods on the JobCollectionsClient type.
type JobCollectionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection scheduler.JobCollectionDefinition) (result scheduler.JobCollectionDefinition, err error)
	Delete(ctx context.Context, resourceGroupName string, jobCollectionName string) (result scheduler.JobCollectionsDeleteFuture, err error)
	Disable(ctx context.Context, resourceGroupName string, jobCollectionName string) (result scheduler.JobCollectionsDisableFuture, err error)
	Enable(ctx context.Context, resourceGroupName string, jobCollectionName string) (result scheduler.JobCollectionsEnableFuture, err error)
	Get(ctx context.Context, resourceGroupName string, jobCollectionName string) (result scheduler.JobCollectionDefinition, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result scheduler.JobCollectionListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result scheduler.JobCollectionListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result scheduler.JobCollectionListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result scheduler.JobCollectionListResultIterator, err error)
	Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobCollection scheduler.JobCollectionDefinition) (result scheduler.JobCollectionDefinition, err error)
}

var _ JobCollectionsClientAPI = (*scheduler.JobCollectionsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job scheduler.JobDefinition) (result scheduler.JobDefinition, err error)
	Delete(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (result scheduler.JobDefinition, err error)
	List(ctx context.Context, resourceGroupName string, jobCollectionName string, top *int32, skip *int32, filter string) (result scheduler.JobListResultPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, jobCollectionName string, top *int32, skip *int32, filter string) (result scheduler.JobListResultIterator, err error)
	ListJobHistory(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, top *int32, skip *int32, filter string) (result scheduler.JobHistoryListResultPage, err error)
	ListJobHistoryComplete(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, top *int32, skip *int32, filter string) (result scheduler.JobHistoryListResultIterator, err error)
	Patch(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string, job scheduler.JobDefinition) (result scheduler.JobDefinition, err error)
	Run(ctx context.Context, resourceGroupName string, jobCollectionName string, jobName string) (result autorest.Response, err error)
}

var _ JobsClientAPI = (*scheduler.JobsClient)(nil)