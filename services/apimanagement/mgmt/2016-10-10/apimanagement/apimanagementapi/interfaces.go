// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package apimanagementapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2016-10-10/apimanagement"
	"github.com/Azure/go-autorest/autorest"
	"io"
)

// PolicySnippetsClientAPI contains the set of methods on the PolicySnippetsClient type.
type PolicySnippetsClientAPI interface {
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, scope apimanagement.PolicyScopeContract) (result apimanagement.PolicySnippetsCollection, err error)
}

var _ PolicySnippetsClientAPI = (*apimanagement.PolicySnippetsClient)(nil)

// RegionsClientAPI contains the set of methods on the RegionsClient type.
type RegionsClientAPI interface {
	ListByService(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.RegionListResult, err error)
}

var _ RegionsClientAPI = (*apimanagement.RegionsClient)(nil)

// ApisClientAPI contains the set of methods on the ApisClient type.
type ApisClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, parameters apimanagement.APIContract, ifMatch string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string) (result apimanagement.APIContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, apiid string, parameters apimanagement.APIUpdateContract, ifMatch string) (result autorest.Response, err error)
}

var _ ApisClientAPI = (*apimanagement.ApisClient)(nil)

// APIOperationsClientAPI contains the set of methods on the APIOperationsClient type.
type APIOperationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string, parameters apimanagement.OperationContract) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string) (result apimanagement.OperationContract, err error)
	ListByApis(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.OperationCollectionPage, err error)
	ListByApisComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.OperationCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string, parameters apimanagement.OperationUpdateContract, ifMatch string) (result autorest.Response, err error)
}

var _ APIOperationsClientAPI = (*apimanagement.APIOperationsClient)(nil)

// APIOperationsPolicyClientAPI contains the set of methods on the APIOperationsPolicyClient type.
type APIOperationsPolicyClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string, parameters io.ReadCloser, ifMatch string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string) (result apimanagement.ReadCloser, err error)
}

var _ APIOperationsPolicyClientAPI = (*apimanagement.APIOperationsPolicyClient)(nil)

// APIProductsClientAPI contains the set of methods on the APIProductsClient type.
type APIProductsClientAPI interface {
	ListByApis(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.ProductCollectionPage, err error)
	ListByApisComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.ProductCollectionIterator, err error)
}

var _ APIProductsClientAPI = (*apimanagement.APIProductsClient)(nil)

// APIPolicyClientAPI contains the set of methods on the APIPolicyClient type.
type APIPolicyClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, parameters io.ReadCloser, ifMatch string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string) (result apimanagement.ReadCloser, err error)
}

var _ APIPolicyClientAPI = (*apimanagement.APIPolicyClient)(nil)

// AuthorizationServersClientAPI contains the set of methods on the AuthorizationServersClient type.
type AuthorizationServersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters apimanagement.OAuth2AuthorizationServerContract) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, authsid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, authsid string) (result apimanagement.OAuth2AuthorizationServerContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.AuthorizationServerCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.AuthorizationServerCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, authsid string, parameters apimanagement.OAuth2AuthorizationServerUpdateContract, ifMatch string) (result autorest.Response, err error)
}

var _ AuthorizationServersClientAPI = (*apimanagement.AuthorizationServersClient)(nil)

// BackendsClientAPI contains the set of methods on the BackendsClient type.
type BackendsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, backendid string, parameters apimanagement.BackendContract) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, backendid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, backendid string) (result apimanagement.BackendResponse, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.BackendCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.BackendCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, backendid string, parameters apimanagement.BackendUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ BackendsClientAPI = (*apimanagement.BackendsClient)(nil)

// CertificatesClientAPI contains the set of methods on the CertificatesClient type.
type CertificatesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, certificateID string, parameters apimanagement.CertificateCreateOrUpdateParameters, ifMatch string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, certificateID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, certificateID string) (result apimanagement.CertificateContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.CertificateCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.CertificateCollectionIterator, err error)
}

var _ CertificatesClientAPI = (*apimanagement.CertificatesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result apimanagement.OperationListResultPage, err error)
	ListComplete(ctx context.Context) (result apimanagement.OperationListResultIterator, err error)
}

var _ OperationsClientAPI = (*apimanagement.OperationsClient)(nil)

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	ApplyNetworkConfigurationUpdates(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.ServicesApplyNetworkConfigurationUpdatesFuture, err error)
	Backup(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceBackupRestoreParameters) (result apimanagement.ServicesBackupFuture, err error)
	CheckNameAvailability(ctx context.Context, parameters apimanagement.ServiceCheckNameAvailabilityParameters) (result apimanagement.ServiceNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceResource) (result apimanagement.ServiceResource, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.ServiceResource, err error)
	GetSsoToken(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.ServiceGetSsoTokenResult, err error)
	List(ctx context.Context) (result apimanagement.ServiceListResultPage, err error)
	ListComplete(ctx context.Context) (result apimanagement.ServiceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result apimanagement.ServiceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result apimanagement.ServiceListResultIterator, err error)
	ManageDeployments(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceManageDeploymentsParameters) (result apimanagement.ServicesManageDeploymentsFuture, err error)
	Restore(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceBackupRestoreParameters) (result apimanagement.ServicesRestoreFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceUpdateParameters) (result apimanagement.ServicesUpdateFuture, err error)
	UpdateHostname(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceUpdateHostnameParameters) (result apimanagement.ServicesUpdateHostnameFuture, err error)
	UploadCertificate(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceUploadCertificateParameters) (result apimanagement.CertificateInformation, err error)
}

var _ ServicesClientAPI = (*apimanagement.ServicesClient)(nil)

// GroupsClientAPI contains the set of methods on the GroupsClient type.
type GroupsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, groupID string, parameters apimanagement.GroupCreateParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, groupID string, ifMatch string) (result apimanagement.ErrorBodyContract, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, groupID string) (result apimanagement.GroupContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, groupID string, parameters apimanagement.GroupUpdateParameters, ifMatch string) (result apimanagement.ErrorBodyContract, err error)
}

var _ GroupsClientAPI = (*apimanagement.GroupsClient)(nil)

// GroupUsersClientAPI contains the set of methods on the GroupUsersClient type.
type GroupUsersClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, serviceName string, groupID string, UID string) (result apimanagement.ErrorBodyContract, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, groupID string, UID string) (result apimanagement.ErrorBodyContract, err error)
	ListByGroups(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionPage, err error)
	ListByGroupsComplete(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionIterator, err error)
}

var _ GroupUsersClientAPI = (*apimanagement.GroupUsersClient)(nil)

// IdentityProvidersClientAPI contains the set of methods on the IdentityProvidersClient type.
type IdentityProvidersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName apimanagement.IdentityProviderNameType, parameters apimanagement.IdentityProviderContract) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName apimanagement.IdentityProviderNameType, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName apimanagement.IdentityProviderNameType) (result apimanagement.IdentityProviderContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.IdentityProviderList, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName apimanagement.IdentityProviderNameType, parameters apimanagement.IdentityProviderUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ IdentityProvidersClientAPI = (*apimanagement.IdentityProvidersClient)(nil)

// LoggersClientAPI contains the set of methods on the LoggersClient type.
type LoggersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, loggerid string, parameters apimanagement.LoggerCreateParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, loggerid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, loggerid string) (result apimanagement.LoggerResponse, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.LoggerCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.LoggerCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, loggerid string, parameters apimanagement.LoggerUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ LoggersClientAPI = (*apimanagement.LoggersClient)(nil)

// NetworkStatusClientAPI contains the set of methods on the NetworkStatusClient type.
type NetworkStatusClientAPI interface {
	GetByService(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.NetworkStatusContract, err error)
}

var _ NetworkStatusClientAPI = (*apimanagement.NetworkStatusClient)(nil)

// OpenIDConnectProvidersClientAPI contains the set of methods on the OpenIDConnectProvidersClient type.
type OpenIDConnectProvidersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, opid string, parameters apimanagement.OpenidConnectProviderCreateContract) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, opid string, ifMatch string) (result apimanagement.ErrorBodyContract, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, opid string) (result apimanagement.OpenidConnectProviderContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.OpenIDConnectProviderCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.OpenIDConnectProviderCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, opid string, parameters apimanagement.OpenidConnectProviderUpdateContract, ifMatch string) (result autorest.Response, err error)
}

var _ OpenIDConnectProvidersClientAPI = (*apimanagement.OpenIDConnectProvidersClient)(nil)

// ProductsClientAPI contains the set of methods on the ProductsClient type.
type ProductsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, productID string, parameters apimanagement.ProductContract) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, ifMatch string, deleteSubscriptions *bool) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, productID string) (result apimanagement.ProductContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32, expandGroups *bool) (result apimanagement.ProductCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32, expandGroups *bool) (result apimanagement.ProductCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, productID string, parameters apimanagement.ProductUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ ProductsClientAPI = (*apimanagement.ProductsClient)(nil)

// ProductApisClientAPI contains the set of methods on the ProductApisClient type.
type ProductApisClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (result autorest.Response, err error)
	ListByProducts(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionPage, err error)
	ListByProductsComplete(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionIterator, err error)
}

var _ ProductApisClientAPI = (*apimanagement.ProductApisClient)(nil)

// ProductGroupsClientAPI contains the set of methods on the ProductGroupsClient type.
type ProductGroupsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string) (result apimanagement.ErrorBodyContract, err error)
	ListByProducts(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionPage, err error)
	ListByProductsComplete(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionIterator, err error)
}

var _ ProductGroupsClientAPI = (*apimanagement.ProductGroupsClient)(nil)

// ProductSubscriptionsClientAPI contains the set of methods on the ProductSubscriptionsClient type.
type ProductSubscriptionsClientAPI interface {
	ListByProducts(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionPage, err error)
	ListByProductsComplete(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionIterator, err error)
}

var _ ProductSubscriptionsClientAPI = (*apimanagement.ProductSubscriptionsClient)(nil)

// ProductPolicyClientAPI contains the set of methods on the ProductPolicyClient type.
type ProductPolicyClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, productID string, parameters io.ReadCloser, ifMatch string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, productID string) (result apimanagement.ReadCloser, err error)
}

var _ ProductPolicyClientAPI = (*apimanagement.ProductPolicyClient)(nil)

// PropertiesClientAPI contains the set of methods on the PropertiesClient type.
type PropertiesClientAPI interface {
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.PropertyCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.PropertyCollectionIterator, err error)
}

var _ PropertiesClientAPI = (*apimanagement.PropertiesClient)(nil)

// PropertyClientAPI contains the set of methods on the PropertyClient type.
type PropertyClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, propID string, parameters apimanagement.PropertyCreateParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, propID string, ifMatch string) (result apimanagement.ErrorBodyContract, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, propID string) (result apimanagement.PropertyContract, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, propID string, parameters apimanagement.PropertyUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ PropertyClientAPI = (*apimanagement.PropertyClient)(nil)

// QuotaByCounterKeysClientAPI contains the set of methods on the QuotaByCounterKeysClient type.
type QuotaByCounterKeysClientAPI interface {
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string) (result apimanagement.QuotaCounterCollection, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, parameters apimanagement.QuotaCounterValueContract) (result autorest.Response, err error)
}

var _ QuotaByCounterKeysClientAPI = (*apimanagement.QuotaByCounterKeysClient)(nil)

// QuotaByPeriodKeysClientAPI contains the set of methods on the QuotaByPeriodKeysClient type.
type QuotaByPeriodKeysClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string) (result apimanagement.QuotaCounterContract, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, quotaCounterKey string, quotaPeriodKey string, parameters apimanagement.QuotaCounterValueContract) (result autorest.Response, err error)
}

var _ QuotaByPeriodKeysClientAPI = (*apimanagement.QuotaByPeriodKeysClient)(nil)

// ReportsClientAPI contains the set of methods on the ReportsClient type.
type ReportsClientAPI interface {
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, aggregation apimanagement.ReportsAggregation, filter string, top *int32, skip *int32, interval *string) (result apimanagement.ReportCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, aggregation apimanagement.ReportsAggregation, filter string, top *int32, skip *int32, interval *string) (result apimanagement.ReportCollectionIterator, err error)
}

var _ ReportsClientAPI = (*apimanagement.ReportsClient)(nil)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters apimanagement.SubscriptionCreateParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, sid string) (result apimanagement.SubscriptionContract, err error)
	List(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionIterator, err error)
	RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string) (result autorest.Response, err error)
	RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters apimanagement.SubscriptionUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ SubscriptionsClientAPI = (*apimanagement.SubscriptionsClient)(nil)

// TenantAccessClientAPI contains the set of methods on the TenantAccessClient type.
type TenantAccessClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.AccessInformationContract, err error)
	RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
	RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.AccessInformationUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ TenantAccessClientAPI = (*apimanagement.TenantAccessClient)(nil)

// TenantAccessGitClientAPI contains the set of methods on the TenantAccessGitClient type.
type TenantAccessGitClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.AccessInformationContract, err error)
	RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
	RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
}

var _ TenantAccessGitClientAPI = (*apimanagement.TenantAccessGitClient)(nil)

// TenantConfigurationClientAPI contains the set of methods on the TenantConfigurationClient type.
type TenantConfigurationClientAPI interface {
	Deploy(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.DeployConfigurationParameters) (result apimanagement.TenantConfigurationDeployFuture, err error)
	Save(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.SaveConfigurationParameter) (result apimanagement.TenantConfigurationSaveFuture, err error)
	Validate(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.DeployConfigurationParameters) (result apimanagement.TenantConfigurationValidateFuture, err error)
}

var _ TenantConfigurationClientAPI = (*apimanagement.TenantConfigurationClient)(nil)

// TenantConfigurationSyncStateClientAPI contains the set of methods on the TenantConfigurationSyncStateClient type.
type TenantConfigurationSyncStateClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.TenantConfigurationSyncStateContract, err error)
}

var _ TenantConfigurationSyncStateClientAPI = (*apimanagement.TenantConfigurationSyncStateClient)(nil)

// TenantPolicyClientAPI contains the set of methods on the TenantPolicyClient type.
type TenantPolicyClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters io.ReadCloser, ifMatch string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.ReadCloser, err error)
}

var _ TenantPolicyClientAPI = (*apimanagement.TenantPolicyClient)(nil)

// UsersClientAPI contains the set of methods on the UsersClient type.
type UsersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters apimanagement.UserCreateParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, UID string, ifMatch string, deleteSubscriptions *bool) (result apimanagement.ErrorBodyContract, err error)
	GenerateSsoURL(ctx context.Context, resourceGroupName string, serviceName string, UID string) (result apimanagement.GenerateSsoURLResult, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, UID string) (result apimanagement.UserContract, err error)
	GetSharedAccessToken(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters apimanagement.UserTokenParameters) (result apimanagement.UserTokenResult, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters apimanagement.UserUpdateParameters, ifMatch string) (result apimanagement.ErrorBodyContract, err error)
}

var _ UsersClientAPI = (*apimanagement.UsersClient)(nil)

// UserGroupsClientAPI contains the set of methods on the UserGroupsClient type.
type UserGroupsClientAPI interface {
	ListByUsers(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionPage, err error)
	ListByUsersComplete(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionIterator, err error)
}

var _ UserGroupsClientAPI = (*apimanagement.UserGroupsClient)(nil)

// UserSubscriptionsClientAPI contains the set of methods on the UserSubscriptionsClient type.
type UserSubscriptionsClientAPI interface {
	ListByUsers(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionPage, err error)
	ListByUsersComplete(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionIterator, err error)
}

var _ UserSubscriptionsClientAPI = (*apimanagement.UserSubscriptionsClient)(nil)

// UserIdentitiesClientAPI contains the set of methods on the UserIdentitiesClient type.
type UserIdentitiesClientAPI interface {
	ListByUsers(ctx context.Context, resourceGroupName string, serviceName string, UID string) (result apimanagement.UserIdentityCollection, err error)
}

var _ UserIdentitiesClientAPI = (*apimanagement.UserIdentitiesClient)(nil)

// APIExportClientAPI contains the set of methods on the APIExportClient type.
type APIExportClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string) (result apimanagement.APIExportResult, err error)
}

var _ APIExportClientAPI = (*apimanagement.APIExportClient)(nil)
