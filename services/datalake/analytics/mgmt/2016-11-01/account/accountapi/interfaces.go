// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-analytics/armdatalakeanalytics). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package accountapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/datalake/analytics/mgmt/2016-11-01/account"
	"github.com/Azure/go-autorest/autorest"
)

// AccountsClientAPI contains the set of methods on the AccountsClient type.
type AccountsClientAPI interface {
	CheckNameAvailability(ctx context.Context, location string, parameters account.CheckNameAvailabilityParameters) (result account.NameAvailabilityInformation, err error)
	Create(ctx context.Context, resourceGroupName string, accountName string, parameters account.CreateDataLakeAnalyticsAccountParameters) (result account.AccountsCreateFutureType, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string) (result account.AccountsDeleteFutureType, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string) (result account.DataLakeAnalyticsAccount, err error)
	List(ctx context.Context, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeAnalyticsAccountListResultPage, err error)
	ListComplete(ctx context.Context, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeAnalyticsAccountListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeAnalyticsAccountListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeAnalyticsAccountListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, parameters *account.UpdateDataLakeAnalyticsAccountParameters) (result account.AccountsUpdateFutureType, err error)
}

var _ AccountsClientAPI = (*account.AccountsClient)(nil)

// DataLakeStoreAccountsClientAPI contains the set of methods on the DataLakeStoreAccountsClient type.
type DataLakeStoreAccountsClientAPI interface {
	Add(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string, parameters *account.AddDataLakeStoreParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, dataLakeStoreAccountName string) (result account.DataLakeStoreAccountInformation, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeStoreAccountInformationListResultPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.DataLakeStoreAccountInformationListResultIterator, err error)
}

var _ DataLakeStoreAccountsClientAPI = (*account.DataLakeStoreAccountsClient)(nil)

// StorageAccountsClientAPI contains the set of methods on the StorageAccountsClient type.
type StorageAccountsClientAPI interface {
	Add(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters account.AddStorageAccountParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result account.StorageAccountInformation, err error)
	GetStorageContainer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result account.StorageContainer, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.StorageAccountInformationListResultPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string, filter string, top *int32, skip *int32, selectParameter string, orderby string, count *bool) (result account.StorageAccountInformationListResultIterator, err error)
	ListSasTokens(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result account.SasTokenInformationListResultPage, err error)
	ListSasTokensComplete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string) (result account.SasTokenInformationListResultIterator, err error)
	ListStorageContainers(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result account.StorageContainerListResultPage, err error)
	ListStorageContainersComplete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string) (result account.StorageContainerListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters *account.UpdateStorageAccountParameters) (result autorest.Response, err error)
}

var _ StorageAccountsClientAPI = (*account.StorageAccountsClient)(nil)

// ComputePoliciesClientAPI contains the set of methods on the ComputePoliciesClient type.
type ComputePoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, parameters account.CreateOrUpdateComputePolicyParameters) (result account.ComputePolicy, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string) (result account.ComputePolicy, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result account.ComputePolicyListResultPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result account.ComputePolicyListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, computePolicyName string, parameters *account.UpdateComputePolicyParameters) (result account.ComputePolicy, err error)
}

var _ ComputePoliciesClientAPI = (*account.ComputePoliciesClient)(nil)

// FirewallRulesClientAPI contains the set of methods on the FirewallRulesClient type.
type FirewallRulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string, parameters account.CreateOrUpdateFirewallRuleParameters) (result account.FirewallRule, err error)
	Delete(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string) (result account.FirewallRule, err error)
	ListByAccount(ctx context.Context, resourceGroupName string, accountName string) (result account.FirewallRuleListResultPage, err error)
	ListByAccountComplete(ctx context.Context, resourceGroupName string, accountName string) (result account.FirewallRuleListResultIterator, err error)
	Update(ctx context.Context, resourceGroupName string, accountName string, firewallRuleName string, parameters *account.UpdateFirewallRuleParameters) (result account.FirewallRule, err error)
}

var _ FirewallRulesClientAPI = (*account.FirewallRulesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result account.OperationListResult, err error)
}

var _ OperationsClientAPI = (*account.OperationsClient)(nil)

// LocationsClientAPI contains the set of methods on the LocationsClient type.
type LocationsClientAPI interface {
	GetCapability(ctx context.Context, location string) (result account.CapabilityInformation, err error)
}

var _ LocationsClientAPI = (*account.LocationsClient)(nil)