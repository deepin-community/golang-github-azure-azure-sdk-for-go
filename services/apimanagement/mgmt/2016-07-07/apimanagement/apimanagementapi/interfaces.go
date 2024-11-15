// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package apimanagementapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2016-07-07/apimanagement"
	"github.com/Azure/go-autorest/autorest"
)

// ApisClientAPI contains the set of methods on the ApisClient type.
type ApisClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, parameters apimanagement.APIContract, ifMatch string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, ifMatch string) (result autorest.Response, err error)
	Export(ctx context.Context, resourceGroupName string, serviceName string, apiid string) (result apimanagement.APIExportResult, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string) (result apimanagement.APIContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, apiid string, parameters apimanagement.PatchParameters, ifMatch string) (result autorest.Response, err error)
}

var _ ApisClientAPI = (*apimanagement.ApisClient)(nil)

// APIOperationsClientAPI contains the set of methods on the APIOperationsClient type.
type APIOperationsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string, parameters apimanagement.OperationContract) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string) (result apimanagement.OperationContract, err error)
	ListByAPI(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.OperationCollectionPage, err error)
	ListByAPIComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.OperationCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, apiid string, operationID string, parameters apimanagement.PatchParameters, ifMatch string) (result autorest.Response, err error)
}

var _ APIOperationsClientAPI = (*apimanagement.APIOperationsClient)(nil)

// APIProductsClientAPI contains the set of methods on the APIProductsClient type.
type APIProductsClientAPI interface {
	ListByAPI(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.ProductCollectionPage, err error)
	ListByAPIComplete(ctx context.Context, resourceGroupName string, serviceName string, apiid string, filter string, top *int32, skip *int32) (result apimanagement.ProductCollectionIterator, err error)
}

var _ APIProductsClientAPI = (*apimanagement.APIProductsClient)(nil)

// SubscriptionsClientAPI contains the set of methods on the SubscriptionsClient type.
type SubscriptionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters apimanagement.SubscriptionCreateParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, sid string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, sid string) (result apimanagement.SubscriptionContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionIterator, err error)
	RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string) (result autorest.Response, err error)
	RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, sid string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, sid string, parameters apimanagement.SubscriptionUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ SubscriptionsClientAPI = (*apimanagement.SubscriptionsClient)(nil)

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
	Add(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (result autorest.Response, err error)
	ListByProduct(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionPage, err error)
	ListByProductComplete(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.APICollectionIterator, err error)
	Remove(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiid string) (result autorest.Response, err error)
}

var _ ProductApisClientAPI = (*apimanagement.ProductApisClient)(nil)

// ProductGroupsClientAPI contains the set of methods on the ProductGroupsClient type.
type ProductGroupsClientAPI interface {
	Add(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string) (result autorest.Response, err error)
	ListByProduct(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionPage, err error)
	ListByProductComplete(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionIterator, err error)
	Remove(ctx context.Context, resourceGroupName string, serviceName string, productID string, groupID string) (result apimanagement.ErrorBodyContract, err error)
}

var _ ProductGroupsClientAPI = (*apimanagement.ProductGroupsClient)(nil)

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
	Add(ctx context.Context, resourceGroupName string, serviceName string, groupID string, UID string) (result apimanagement.ErrorBodyContract, err error)
	ListByGroup(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionPage, err error)
	ListByGroupComplete(ctx context.Context, resourceGroupName string, serviceName string, groupID string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionIterator, err error)
	Remove(ctx context.Context, resourceGroupName string, serviceName string, groupID string, UID string) (result apimanagement.ErrorBodyContract, err error)
}

var _ GroupUsersClientAPI = (*apimanagement.GroupUsersClient)(nil)

// CertificatesClientAPI contains the set of methods on the CertificatesClient type.
type CertificatesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, certificateID string, parameters apimanagement.CertificateCreateOrUpdateParameters, ifMatch string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, certificateID string, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, certificateID string) (result apimanagement.CertificateContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.CertificateCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.CertificateCollectionIterator, err error)
}

var _ CertificatesClientAPI = (*apimanagement.CertificatesClient)(nil)

// PolicySnippetsClientAPI contains the set of methods on the PolicySnippetsClient type.
type PolicySnippetsClientAPI interface {
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, scope apimanagement.PolicyScopeContract) (result apimanagement.ListPolicySnippetContract, err error)
}

var _ PolicySnippetsClientAPI = (*apimanagement.PolicySnippetsClient)(nil)

// ProductSubscriptionsClientAPI contains the set of methods on the ProductSubscriptionsClient type.
type ProductSubscriptionsClientAPI interface {
	ListByProduct(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionPage, err error)
	ListByProductComplete(ctx context.Context, resourceGroupName string, serviceName string, productID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionIterator, err error)
}

var _ ProductSubscriptionsClientAPI = (*apimanagement.ProductSubscriptionsClient)(nil)

// UsersClientAPI contains the set of methods on the UsersClient type.
type UsersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters apimanagement.UserCreateParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, UID string, ifMatch string, deleteSubscriptions *bool) (result apimanagement.ErrorBodyContract, err error)
	GenerateSsoURL(ctx context.Context, resourceGroupName string, serviceName string, UID string) (result apimanagement.GenerateSsoURLResult, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, UID string) (result apimanagement.UserContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.UserCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, UID string, parameters apimanagement.UserUpdateParameters, ifMatch string) (result apimanagement.ErrorBodyContract, err error)
}

var _ UsersClientAPI = (*apimanagement.UsersClient)(nil)

// UserGroupsClientAPI contains the set of methods on the UserGroupsClient type.
type UserGroupsClientAPI interface {
	ListByUser(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionPage, err error)
	ListByUserComplete(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result apimanagement.GroupCollectionIterator, err error)
}

var _ UserGroupsClientAPI = (*apimanagement.UserGroupsClient)(nil)

// UserSubscriptionsClientAPI contains the set of methods on the UserSubscriptionsClient type.
type UserSubscriptionsClientAPI interface {
	ListByUser(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionPage, err error)
	ListByUserComplete(ctx context.Context, resourceGroupName string, serviceName string, UID string, filter string, top *int32, skip *int32) (result apimanagement.SubscriptionCollectionIterator, err error)
}

var _ UserSubscriptionsClientAPI = (*apimanagement.UserSubscriptionsClient)(nil)

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

// RegionsClientAPI contains the set of methods on the RegionsClient type.
type RegionsClientAPI interface {
	ListByService(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.RegionListResult, err error)
}

var _ RegionsClientAPI = (*apimanagement.RegionsClient)(nil)

// UserIdentitiesClientAPI contains the set of methods on the UserIdentitiesClient type.
type UserIdentitiesClientAPI interface {
	ListByUser(ctx context.Context, resourceGroupName string, serviceName string, UID string) (result apimanagement.ListUserIdentityContract, err error)
}

var _ UserIdentitiesClientAPI = (*apimanagement.UserIdentitiesClient)(nil)

// ReportsClientAPI contains the set of methods on the ReportsClient type.
type ReportsClientAPI interface {
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, aggregation apimanagement.ReportsAggregation, filter string, top *int32, skip *int32, interval string) (result apimanagement.ReportCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, aggregation apimanagement.ReportsAggregation, filter string, top *int32, skip *int32, interval string) (result apimanagement.ReportCollectionIterator, err error)
}

var _ ReportsClientAPI = (*apimanagement.ReportsClient)(nil)

// TenantAccessClientAPI contains the set of methods on the TenantAccessClient type.
type TenantAccessClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.AccessInformationContract, err error)
	RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
	RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.AccessInformationUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ TenantAccessClientAPI = (*apimanagement.TenantAccessClient)(nil)

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

// PropertyClientAPI contains the set of methods on the PropertyClient type.
type PropertyClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, propID string, parameters apimanagement.PropertyCreateParameters) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, propID string, ifMatch string) (result apimanagement.ErrorBodyContract, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, propID string) (result apimanagement.PropertyContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.PropertyCollectionPage, err error)
	ListByServiceComplete(ctx context.Context, resourceGroupName string, serviceName string, filter string, top *int32, skip *int32) (result apimanagement.PropertyCollectionIterator, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, propID string, parameters apimanagement.PropertyUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ PropertyClientAPI = (*apimanagement.PropertyClient)(nil)

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

// IdentityProvidersClientAPI contains the set of methods on the IdentityProvidersClient type.
type IdentityProvidersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName apimanagement.IdentityProviderNameType, parameters apimanagement.IdentityProviderContract) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName apimanagement.IdentityProviderNameType, ifMatch string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName apimanagement.IdentityProviderNameType) (result apimanagement.IdentityProviderContract, err error)
	ListByService(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.IdentityProviderList, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, identityProviderName apimanagement.IdentityProviderNameType, parameters apimanagement.IdentityProviderUpdateParameters, ifMatch string) (result autorest.Response, err error)
}

var _ IdentityProvidersClientAPI = (*apimanagement.IdentityProvidersClient)(nil)

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

// ServicesClientAPI contains the set of methods on the ServicesClient type.
type ServicesClientAPI interface {
	Backup(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceBackupRestoreParameters) (result apimanagement.ServicesBackupFuture, err error)
	CheckNameAvailability(ctx context.Context, parameters apimanagement.ServiceCheckNameAvailabilityParameters) (result apimanagement.ServiceNameAvailabilityResult, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceResource) (result apimanagement.ServiceResource, err error)
	Delete(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.ErrorResponse, err error)
	Get(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.SetObject, err error)
	GetSsoToken(ctx context.Context, resourceGroupName string, serviceName string) (result apimanagement.ServiceGetSsoTokenResult, err error)
	List(ctx context.Context) (result apimanagement.ServiceListResultPage, err error)
	ListComplete(ctx context.Context) (result apimanagement.ServiceListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result apimanagement.ServiceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result apimanagement.ServiceListResultIterator, err error)
	ManageDeployments(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceManageDeploymentsParameters) (result apimanagement.ServicesManageDeploymentsFuture, err error)
	Restore(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceBackupRestoreParameters) (result apimanagement.ServicesRestoreFuture, err error)
	Update(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceBaseParameters) (result apimanagement.ServicesUpdateFuture, err error)
	UpdateHostname(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceUpdateHostnameParameters) (result apimanagement.ServicesUpdateHostnameFuture, err error)
	UploadCertificate(ctx context.Context, resourceGroupName string, serviceName string, parameters apimanagement.ServiceUploadCertificateParameters) (result apimanagement.CertificateInformation, err error)
}

var _ ServicesClientAPI = (*apimanagement.ServicesClient)(nil)
