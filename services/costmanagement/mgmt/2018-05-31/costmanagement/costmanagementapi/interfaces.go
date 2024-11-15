// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package costmanagementapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/costmanagement/mgmt/2018-05-31/costmanagement"
	"github.com/Azure/go-autorest/autorest"
)

// BaseClientAPI contains the set of methods on the BaseClient type.
type BaseClientAPI interface {
	QueryBillingAccount(ctx context.Context, billingAccountID string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	QueryResourceGroup(ctx context.Context, resourceGroupName string, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
	QuerySubscription(ctx context.Context, parameters costmanagement.ReportConfigDefinition) (result costmanagement.QueryResult, err error)
}

var _ BaseClientAPI = (*costmanagement.BaseClient)(nil)

// ReportConfigClientAPI contains the set of methods on the ReportConfigClient type.
type ReportConfigClientAPI interface {
	CreateOrUpdate(ctx context.Context, reportConfigName string, parameters costmanagement.ReportConfig) (result costmanagement.ReportConfig, err error)
	CreateOrUpdateByResourceGroupName(ctx context.Context, resourceGroupName string, reportConfigName string, parameters costmanagement.ReportConfig) (result costmanagement.ReportConfig, err error)
	Delete(ctx context.Context, reportConfigName string) (result autorest.Response, err error)
	DeleteByResourceGroupName(ctx context.Context, resourceGroupName string, reportConfigName string) (result autorest.Response, err error)
	Get(ctx context.Context, reportConfigName string) (result costmanagement.ReportConfig, err error)
	GetByResourceGroupName(ctx context.Context, resourceGroupName string, reportConfigName string) (result costmanagement.ReportConfig, err error)
	List(ctx context.Context) (result costmanagement.ReportConfigListResult, err error)
	ListByResourceGroupName(ctx context.Context, resourceGroupName string) (result costmanagement.ReportConfigListResult, err error)
}

var _ ReportConfigClientAPI = (*costmanagement.ReportConfigClient)(nil)

// BillingAccountDimensionsClientAPI contains the set of methods on the BillingAccountDimensionsClient type.
type BillingAccountDimensionsClientAPI interface {
	List(ctx context.Context, billingAccountID string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
}

var _ BillingAccountDimensionsClientAPI = (*costmanagement.BillingAccountDimensionsClient)(nil)

// SubscriptionDimensionsClientAPI contains the set of methods on the SubscriptionDimensionsClient type.
type SubscriptionDimensionsClientAPI interface {
	List(ctx context.Context, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
}

var _ SubscriptionDimensionsClientAPI = (*costmanagement.SubscriptionDimensionsClient)(nil)

// ResourceGroupDimensionsClientAPI contains the set of methods on the ResourceGroupDimensionsClient type.
type ResourceGroupDimensionsClientAPI interface {
	List(ctx context.Context, resourceGroupName string, filter string, expand string, skiptoken string, top *int32) (result costmanagement.DimensionsListResult, err error)
}

var _ ResourceGroupDimensionsClientAPI = (*costmanagement.ResourceGroupDimensionsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result costmanagement.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result costmanagement.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*costmanagement.OperationsClient)(nil)
