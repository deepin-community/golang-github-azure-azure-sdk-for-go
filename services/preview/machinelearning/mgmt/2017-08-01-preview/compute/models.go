package compute

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2017-08-01-preview/compute"

// AcsClusterProperties information about the container service backing the cluster
type AcsClusterProperties struct {
	// ClusterFqdn - READ-ONLY; The FQDN of the cluster.
	ClusterFqdn *string `json:"clusterFqdn,omitempty"`
	// OrchestratorType - Type of orchestrator. It cannot be changed once the cluster is created. Possible values include: 'Kubernetes', 'None'
	OrchestratorType OrchestratorType `json:"orchestratorType,omitempty"`
	// OrchestratorProperties - Orchestrator specific properties
	OrchestratorProperties *KubernetesClusterProperties `json:"orchestratorProperties,omitempty"`
	// SystemServices - The system services deployed to the cluster
	SystemServices *[]SystemService `json:"systemServices,omitempty"`
	// MasterCount - The number of master nodes in the container service.
	MasterCount *int32 `json:"masterCount,omitempty"`
	// AgentCount - The number of agent nodes in the Container Service. This can be changed to scale the cluster.
	AgentCount *int32 `json:"agentCount,omitempty"`
	// AgentVMSize - The Azure VM size of the agent VM nodes. This cannot be changed once the cluster is created. This list is non exhaustive; refer to https://docs.microsoft.com/en-us/azure/virtual-machines/windows/sizes for the possible VM sizes. Possible values include: 'StandardA0', 'StandardA1', 'StandardA2', 'StandardA3', 'StandardA4', 'StandardA5', 'StandardA6', 'StandardA7', 'StandardA8', 'StandardA9', 'StandardA10', 'StandardA11', 'StandardD1', 'StandardD2', 'StandardD3', 'StandardD4', 'StandardD11', 'StandardD12', 'StandardD13', 'StandardD14', 'StandardD1V2', 'StandardD2V2', 'StandardD3V2', 'StandardD4V2', 'StandardD5V2', 'StandardD11V2', 'StandardD12V2', 'StandardD13V2', 'StandardD14V2', 'StandardG1', 'StandardG2', 'StandardG3', 'StandardG4', 'StandardG5', 'StandardDS1', 'StandardDS2', 'StandardDS3', 'StandardDS4', 'StandardDS11', 'StandardDS12', 'StandardDS13', 'StandardDS14', 'StandardGS1', 'StandardGS2', 'StandardGS3', 'StandardGS4', 'StandardGS5'
	AgentVMSize AgentVMSizeTypes `json:"agentVmSize,omitempty"`
}

// MarshalJSON is the custom marshaler for AcsClusterProperties.
func (acp AcsClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if acp.OrchestratorType != "" {
		objectMap["orchestratorType"] = acp.OrchestratorType
	}
	if acp.OrchestratorProperties != nil {
		objectMap["orchestratorProperties"] = acp.OrchestratorProperties
	}
	if acp.SystemServices != nil {
		objectMap["systemServices"] = acp.SystemServices
	}
	if acp.MasterCount != nil {
		objectMap["masterCount"] = acp.MasterCount
	}
	if acp.AgentCount != nil {
		objectMap["agentCount"] = acp.AgentCount
	}
	if acp.AgentVMSize != "" {
		objectMap["agentVmSize"] = acp.AgentVMSize
	}
	return json.Marshal(objectMap)
}

// AppInsightsCredentials appInsights credentials.
type AppInsightsCredentials struct {
	// AppID - The AppInsights application ID.
	AppID *string `json:"appId,omitempty"`
	// InstrumentationKey - The AppInsights instrumentation key. This is not returned in response of GET/PUT on the resource. To see this please call listKeys API.
	InstrumentationKey *string `json:"instrumentationKey,omitempty"`
}

// AppInsightsProperties properties of App Insights.
type AppInsightsProperties struct {
	// ResourceID - ARM resource ID of the App Insights.
	ResourceID *string `json:"resourceId,omitempty"`
}

// AutoScaleConfiguration autoScale configuration properties.
type AutoScaleConfiguration struct {
	// Status - If auto-scale is enabled for all services. Each service can turn it off individually. Possible values include: 'Enabled', 'Disabled'
	Status Status `json:"status,omitempty"`
	// MinReplicas - The minimum number of replicas for each service.
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// MaxReplicas - The maximum number of replicas for each service.
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
	// TargetUtilization - The target utilization.
	TargetUtilization *float64 `json:"targetUtilization,omitempty"`
	// RefreshPeriodInSeconds - Refresh period in seconds.
	RefreshPeriodInSeconds *int32 `json:"refreshPeriodInSeconds,omitempty"`
}

// AvailableOperations available operation list.
type AvailableOperations struct {
	autorest.Response `json:"-"`
	// Value - An array of available operations.
	Value *[]ResourceOperation `json:"value,omitempty"`
}

// CheckSystemServicesUpdatesAvailableResponse information about updates available for system services in a
// cluster.
type CheckSystemServicesUpdatesAvailableResponse struct {
	autorest.Response `json:"-"`
	// UpdatesAvailable - READ-ONLY; Yes if updates are available for the system services, No if not. Possible values include: 'Yes', 'No'
	UpdatesAvailable UpdatesAvailable `json:"updatesAvailable,omitempty"`
}

// MarshalJSON is the custom marshaler for CheckSystemServicesUpdatesAvailableResponse.
func (cssuar CheckSystemServicesUpdatesAvailableResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ContainerRegistryCredentials information about the Azure Container Registry which contains the images
// deployed to the cluster.
type ContainerRegistryCredentials struct {
	// LoginServer - READ-ONLY; The ACR login server name. User name is the first part of the FQDN.
	LoginServer *string `json:"loginServer,omitempty"`
	// Password - READ-ONLY; The ACR primary password.
	Password *string `json:"password,omitempty"`
	// Password2 - READ-ONLY; The ACR secondary password.
	Password2 *string `json:"password2,omitempty"`
	// Username - READ-ONLY; The ACR login username.
	Username *string `json:"username,omitempty"`
}

// MarshalJSON is the custom marshaler for ContainerRegistryCredentials.
func (crc ContainerRegistryCredentials) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ContainerRegistryProperties properties of Azure Container Registry.
type ContainerRegistryProperties struct {
	// ResourceID - ARM resource ID of the Azure Container Registry used to store Docker images for web services in the cluster. If not provided one will be created. This cannot be changed once the cluster is created.
	ResourceID *string `json:"resourceId,omitempty"`
}

// ContainerServiceCredentials information about the Azure Container Registry which contains the images
// deployed to the cluster.
type ContainerServiceCredentials struct {
	// AcsKubeConfig - READ-ONLY; The ACS kube config file.
	AcsKubeConfig *string `json:"acsKubeConfig,omitempty"`
	// ServicePrincipalConfiguration - READ-ONLY; Service principal configuration used by Kubernetes.
	ServicePrincipalConfiguration *ServicePrincipalProperties `json:"servicePrincipalConfiguration,omitempty"`
	// ImagePullSecretName - READ-ONLY; The ACR image pull secret name which was created in Kubernetes.
	ImagePullSecretName *string `json:"imagePullSecretName,omitempty"`
}

// MarshalJSON is the custom marshaler for ContainerServiceCredentials.
func (csc ContainerServiceCredentials) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// ErrorDetail error detail information.
type ErrorDetail struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message.
	Message *string `json:"message,omitempty"`
}

// ErrorResponse error response information.
type ErrorResponse struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message.
	Message *string `json:"message,omitempty"`
	// Details - An array of error detail objects.
	Details *[]ErrorDetail `json:"details,omitempty"`
}

// ErrorResponseWrapper wrapper for error response to follow ARM guidelines.
type ErrorResponseWrapper struct {
	// Error - The error response.
	Error *ErrorResponse `json:"error,omitempty"`
}

// GlobalServiceConfiguration global configuration for services in the cluster.
type GlobalServiceConfiguration struct {
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	// Etag - The configuration ETag for updates.
	Etag *string `json:"etag,omitempty"`
	// Ssl - The SSL configuration properties
	Ssl *SslConfiguration `json:"ssl,omitempty"`
	// ServiceAuth - Optional global authorization keys for all user services deployed in cluster. These are used if the service does not have auth keys.
	ServiceAuth *ServiceAuthConfiguration `json:"serviceAuth,omitempty"`
	// AutoScale - The auto-scale configuration
	AutoScale *AutoScaleConfiguration `json:"autoScale,omitempty"`
}

// MarshalJSON is the custom marshaler for GlobalServiceConfiguration.
func (gsc GlobalServiceConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if gsc.Etag != nil {
		objectMap["etag"] = gsc.Etag
	}
	if gsc.Ssl != nil {
		objectMap["ssl"] = gsc.Ssl
	}
	if gsc.ServiceAuth != nil {
		objectMap["serviceAuth"] = gsc.ServiceAuth
	}
	if gsc.AutoScale != nil {
		objectMap["autoScale"] = gsc.AutoScale
	}
	for k, v := range gsc.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for GlobalServiceConfiguration struct.
func (gsc *GlobalServiceConfiguration) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if gsc.AdditionalProperties == nil {
					gsc.AdditionalProperties = make(map[string]interface{})
				}
				gsc.AdditionalProperties[k] = additionalProperties
			}
		case "etag":
			if v != nil {
				var etag string
				err = json.Unmarshal(*v, &etag)
				if err != nil {
					return err
				}
				gsc.Etag = &etag
			}
		case "ssl":
			if v != nil {
				var ssl SslConfiguration
				err = json.Unmarshal(*v, &ssl)
				if err != nil {
					return err
				}
				gsc.Ssl = &ssl
			}
		case "serviceAuth":
			if v != nil {
				var serviceAuth ServiceAuthConfiguration
				err = json.Unmarshal(*v, &serviceAuth)
				if err != nil {
					return err
				}
				gsc.ServiceAuth = &serviceAuth
			}
		case "autoScale":
			if v != nil {
				var autoScale AutoScaleConfiguration
				err = json.Unmarshal(*v, &autoScale)
				if err != nil {
					return err
				}
				gsc.AutoScale = &autoScale
			}
		}
	}

	return nil
}

// KubernetesClusterProperties kubernetes cluster specific properties
type KubernetesClusterProperties struct {
	// ServicePrincipal - The Azure Service Principal used by Kubernetes
	ServicePrincipal *ServicePrincipalProperties `json:"servicePrincipal,omitempty"`
}

// OperationalizationCluster instance of an Azure ML Operationalization Cluster resource.
type OperationalizationCluster struct {
	autorest.Response `json:"-"`
	// OperationalizationClusterProperties - Properties of an operationalization cluster.
	*OperationalizationClusterProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Specifies the resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Specifies the name of the resource.
	Name *string `json:"name,omitempty"`
	// Location - Specifies the location of the resource.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; Specifies the type of the resource.
	Type *string `json:"type,omitempty"`
	// Tags - Contains resource tags defined as key/value pairs.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for OperationalizationCluster.
func (oc OperationalizationCluster) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if oc.OperationalizationClusterProperties != nil {
		objectMap["properties"] = oc.OperationalizationClusterProperties
	}
	if oc.Location != nil {
		objectMap["location"] = oc.Location
	}
	if oc.Tags != nil {
		objectMap["tags"] = oc.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for OperationalizationCluster struct.
func (oc *OperationalizationCluster) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var operationalizationClusterProperties OperationalizationClusterProperties
				err = json.Unmarshal(*v, &operationalizationClusterProperties)
				if err != nil {
					return err
				}
				oc.OperationalizationClusterProperties = &operationalizationClusterProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				oc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				oc.Name = &name
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				oc.Location = &location
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				oc.Type = &typeVar
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				oc.Tags = tags
			}
		}
	}

	return nil
}

// OperationalizationClusterCredentials credentials to resources in the cluster.
type OperationalizationClusterCredentials struct {
	autorest.Response `json:"-"`
	// StorageAccount - Credentials for the Storage Account.
	StorageAccount *StorageAccountCredentials `json:"storageAccount,omitempty"`
	// ContainerRegistry - Credentials for Azure Container Registry.
	ContainerRegistry *ContainerRegistryCredentials `json:"containerRegistry,omitempty"`
	// ContainerService - Credentials for Azure Container Service.
	ContainerService *ContainerServiceCredentials `json:"containerService,omitempty"`
	// AppInsights - Credentials for Azure AppInsights.
	AppInsights *AppInsightsCredentials `json:"appInsights,omitempty"`
	// ServiceAuthConfiguration - Global authorization keys for all user services deployed in cluster. These are used if the service does not have auth keys.
	ServiceAuthConfiguration *ServiceAuthConfiguration `json:"serviceAuthConfiguration,omitempty"`
	// SslConfiguration - The SSL configuration for the services.
	SslConfiguration *SslConfiguration `json:"sslConfiguration,omitempty"`
}

// OperationalizationClusterProperties properties of an operationalization cluster
type OperationalizationClusterProperties struct {
	// Description - The description of the cluster.
	Description *string `json:"description,omitempty"`
	// CreatedOn - READ-ONLY; The date and time when the cluster was created.
	CreatedOn *date.Time `json:"createdOn,omitempty"`
	// ModifiedOn - READ-ONLY; The date and time when the cluster was last modified.
	ModifiedOn *date.Time `json:"modifiedOn,omitempty"`
	// ProvisioningState - READ-ONLY; The provision state of the cluster. Valid values are Unknown, Updating, Provisioning, Succeeded, and Failed. Possible values include: 'Unknown', 'Updating', 'Creating', 'Deleting', 'Succeeded', 'Failed', 'Canceled'
	ProvisioningState OperationStatus `json:"provisioningState,omitempty"`
	// ProvisioningErrors - READ-ONLY; List of provisioning errors reported by the resource provider.
	ProvisioningErrors *[]ErrorResponseWrapper `json:"provisioningErrors,omitempty"`
	// ClusterType - The cluster type. Possible values include: 'ACS', 'Local'
	ClusterType ClusterType `json:"clusterType,omitempty"`
	// StorageAccount - Storage Account properties.
	StorageAccount *StorageAccountProperties `json:"storageAccount,omitempty"`
	// ContainerRegistry - Container Registry properties.
	ContainerRegistry *ContainerRegistryProperties `json:"containerRegistry,omitempty"`
	// ContainerService - Parameters for the Azure Container Service cluster.
	ContainerService *AcsClusterProperties `json:"containerService,omitempty"`
	// AppInsights - AppInsights configuration.
	AppInsights *AppInsightsProperties `json:"appInsights,omitempty"`
	// GlobalServiceConfiguration - Contains global configuration for the web services in the cluster.
	GlobalServiceConfiguration *GlobalServiceConfiguration `json:"globalServiceConfiguration,omitempty"`
}

// MarshalJSON is the custom marshaler for OperationalizationClusterProperties.
func (ocp OperationalizationClusterProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ocp.Description != nil {
		objectMap["description"] = ocp.Description
	}
	if ocp.ClusterType != "" {
		objectMap["clusterType"] = ocp.ClusterType
	}
	if ocp.StorageAccount != nil {
		objectMap["storageAccount"] = ocp.StorageAccount
	}
	if ocp.ContainerRegistry != nil {
		objectMap["containerRegistry"] = ocp.ContainerRegistry
	}
	if ocp.ContainerService != nil {
		objectMap["containerService"] = ocp.ContainerService
	}
	if ocp.AppInsights != nil {
		objectMap["appInsights"] = ocp.AppInsights
	}
	if ocp.GlobalServiceConfiguration != nil {
		objectMap["globalServiceConfiguration"] = ocp.GlobalServiceConfiguration
	}
	return json.Marshal(objectMap)
}

// OperationalizationClustersCreateOrUpdateFuture an abstraction for monitoring and retrieving the results
// of a long-running operation.
type OperationalizationClustersCreateOrUpdateFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(OperationalizationClustersClient) (OperationalizationCluster, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *OperationalizationClustersCreateOrUpdateFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for OperationalizationClustersCreateOrUpdateFuture.Result.
func (future *OperationalizationClustersCreateOrUpdateFuture) result(client OperationalizationClustersClient) (oc OperationalizationCluster, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.OperationalizationClustersCreateOrUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		oc.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("compute.OperationalizationClustersCreateOrUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if oc.Response.Response, err = future.GetResult(sender); err == nil && oc.Response.Response.StatusCode != http.StatusNoContent {
		oc, err = client.CreateOrUpdateResponder(oc.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.OperationalizationClustersCreateOrUpdateFuture", "Result", oc.Response.Response, "Failure responding to request")
		}
	}
	return
}

// OperationalizationClustersDeleteFuture an abstraction for monitoring and retrieving the results of a
// long-running operation.
type OperationalizationClustersDeleteFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(OperationalizationClustersClient) (autorest.Response, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *OperationalizationClustersDeleteFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for OperationalizationClustersDeleteFuture.Result.
func (future *OperationalizationClustersDeleteFuture) result(client OperationalizationClustersClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.OperationalizationClustersDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		ar.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("compute.OperationalizationClustersDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// OperationalizationClustersUpdateSystemServicesFuture an abstraction for monitoring and retrieving the
// results of a long-running operation.
type OperationalizationClustersUpdateSystemServicesFuture struct {
	azure.FutureAPI
	// Result returns the result of the asynchronous operation.
	// If the operation has not completed it will return an error.
	Result func(OperationalizationClustersClient) (UpdateSystemServicesResponse, error)
}

// UnmarshalJSON is the custom unmarshaller for CreateFuture.
func (future *OperationalizationClustersUpdateSystemServicesFuture) UnmarshalJSON(body []byte) error {
	var azFuture azure.Future
	if err := json.Unmarshal(body, &azFuture); err != nil {
		return err
	}
	future.FutureAPI = &azFuture
	future.Result = future.result
	return nil
}

// result is the default implementation for OperationalizationClustersUpdateSystemServicesFuture.Result.
func (future *OperationalizationClustersUpdateSystemServicesFuture) result(client OperationalizationClustersClient) (ussr UpdateSystemServicesResponse, err error) {
	var done bool
	done, err = future.DoneWithContext(context.Background(), client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "compute.OperationalizationClustersUpdateSystemServicesFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		ussr.Response.Response = future.Response()
		err = azure.NewAsyncOpIncompleteError("compute.OperationalizationClustersUpdateSystemServicesFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if ussr.Response.Response, err = future.GetResult(sender); err == nil && ussr.Response.Response.StatusCode != http.StatusNoContent {
		ussr, err = client.UpdateSystemServicesResponder(ussr.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "compute.OperationalizationClustersUpdateSystemServicesFuture", "Result", ussr.Response.Response, "Failure responding to request")
		}
	}
	return
}

// OperationalizationClusterUpdateParameters parameters for PATCH operation on an operationalization
// cluster
type OperationalizationClusterUpdateParameters struct {
	// Tags - Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater in length than 128 characters and a value no greater in length than 256 characters.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for OperationalizationClusterUpdateParameters.
func (ocup OperationalizationClusterUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ocup.Tags != nil {
		objectMap["tags"] = ocup.Tags
	}
	return json.Marshal(objectMap)
}

// PaginatedOperationalizationClustersList paginated list of operationalization clusters.
type PaginatedOperationalizationClustersList struct {
	autorest.Response `json:"-"`
	// Value - An array of cluster objects.
	Value *[]OperationalizationCluster `json:"value,omitempty"`
	// NextLink - A continuation link (absolute URI) to the next page of results in the list.
	NextLink *string `json:"nextLink,omitempty"`
}

// PaginatedOperationalizationClustersListIterator provides access to a complete listing of
// OperationalizationCluster values.
type PaginatedOperationalizationClustersListIterator struct {
	i    int
	page PaginatedOperationalizationClustersListPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *PaginatedOperationalizationClustersListIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PaginatedOperationalizationClustersListIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *PaginatedOperationalizationClustersListIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter PaginatedOperationalizationClustersListIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter PaginatedOperationalizationClustersListIterator) Response() PaginatedOperationalizationClustersList {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter PaginatedOperationalizationClustersListIterator) Value() OperationalizationCluster {
	if !iter.page.NotDone() {
		return OperationalizationCluster{}
	}
	return iter.page.Values()[iter.i]
}

// Creates a new instance of the PaginatedOperationalizationClustersListIterator type.
func NewPaginatedOperationalizationClustersListIterator(page PaginatedOperationalizationClustersListPage) PaginatedOperationalizationClustersListIterator {
	return PaginatedOperationalizationClustersListIterator{page: page}
}

// IsEmpty returns true if the ListResult contains no values.
func (pocl PaginatedOperationalizationClustersList) IsEmpty() bool {
	return pocl.Value == nil || len(*pocl.Value) == 0
}

// hasNextLink returns true if the NextLink is not empty.
func (pocl PaginatedOperationalizationClustersList) hasNextLink() bool {
	return pocl.NextLink != nil && len(*pocl.NextLink) != 0
}

// paginatedOperationalizationClustersListPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (pocl PaginatedOperationalizationClustersList) paginatedOperationalizationClustersListPreparer(ctx context.Context) (*http.Request, error) {
	if !pocl.hasNextLink() {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(pocl.NextLink)))
}

// PaginatedOperationalizationClustersListPage contains a page of OperationalizationCluster values.
type PaginatedOperationalizationClustersListPage struct {
	fn   func(context.Context, PaginatedOperationalizationClustersList) (PaginatedOperationalizationClustersList, error)
	pocl PaginatedOperationalizationClustersList
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *PaginatedOperationalizationClustersListPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PaginatedOperationalizationClustersListPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	for {
		next, err := page.fn(ctx, page.pocl)
		if err != nil {
			return err
		}
		page.pocl = next
		if !next.hasNextLink() || !next.IsEmpty() {
			break
		}
	}
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *PaginatedOperationalizationClustersListPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page PaginatedOperationalizationClustersListPage) NotDone() bool {
	return !page.pocl.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page PaginatedOperationalizationClustersListPage) Response() PaginatedOperationalizationClustersList {
	return page.pocl
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page PaginatedOperationalizationClustersListPage) Values() []OperationalizationCluster {
	if page.pocl.IsEmpty() {
		return nil
	}
	return *page.pocl.Value
}

// Creates a new instance of the PaginatedOperationalizationClustersListPage type.
func NewPaginatedOperationalizationClustersListPage(cur PaginatedOperationalizationClustersList, getNextPage func(context.Context, PaginatedOperationalizationClustersList) (PaginatedOperationalizationClustersList, error)) PaginatedOperationalizationClustersListPage {
	return PaginatedOperationalizationClustersListPage{
		fn:   getNextPage,
		pocl: cur,
	}
}

// Resource azure resource
type Resource struct {
	// ID - READ-ONLY; Specifies the resource ID.
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Specifies the name of the resource.
	Name *string `json:"name,omitempty"`
	// Location - Specifies the location of the resource.
	Location *string `json:"location,omitempty"`
	// Type - READ-ONLY; Specifies the type of the resource.
	Type *string `json:"type,omitempty"`
	// Tags - Contains resource tags defined as key/value pairs.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceOperation resource operation.
type ResourceOperation struct {
	// Name - Name of this operation.
	Name *string `json:"name,omitempty"`
	// Display - Display of the operation.
	Display *ResourceOperationDisplay `json:"display,omitempty"`
	// Origin - The operation origin.
	Origin *string `json:"origin,omitempty"`
}

// ResourceOperationDisplay display of the operation.
type ResourceOperationDisplay struct {
	// Provider - The resource provider name.
	Provider *string `json:"provider,omitempty"`
	// Resource - The resource name.
	Resource *string `json:"resource,omitempty"`
	// Operation - The operation.
	Operation *string `json:"operation,omitempty"`
	// Description - The description of the operation.
	Description *string `json:"description,omitempty"`
}

// ServiceAuthConfiguration global service auth configuration properties. These are the data-plane
// authorization keys and are used if a service doesn't define it's own.
type ServiceAuthConfiguration struct {
	// PrimaryAuthKeyHash - The primary auth key hash. This is not returned in response of GET/PUT on the resource.. To see this please call listKeys API.
	PrimaryAuthKeyHash *string `json:"primaryAuthKeyHash,omitempty"`
	// SecondaryAuthKeyHash - The secondary auth key hash. This is not returned in response of GET/PUT on the resource.. To see this please call listKeys API.
	SecondaryAuthKeyHash *string `json:"secondaryAuthKeyHash,omitempty"`
}

// ServicePrincipalProperties the Azure service principal used by Kubernetes for configuring load balancers
type ServicePrincipalProperties struct {
	// ClientID - The service principal client ID
	ClientID *string `json:"clientId,omitempty"`
	// Secret - The service principal secret. This is not returned in response of GET/PUT on the resource. To see this please call listKeys.
	Secret *string `json:"secret,omitempty"`
}

// SslConfiguration SSL configuration. If configured data-plane calls to user services will be exposed over
// SSL only.
type SslConfiguration struct {
	// Status - SSL status. Allowed values are Enabled and Disabled. Possible values include: 'Enabled', 'Disabled'
	Status Status `json:"status,omitempty"`
	// Cert - The SSL cert data in PEM format.
	Cert *string `json:"cert,omitempty"`
	// Key - The SSL key data in PEM format. This is not returned in response of GET/PUT on the resource. To see this please call listKeys API.
	Key *string `json:"key,omitempty"`
	// Cname - The CName of the certificate.
	Cname *string `json:"cname,omitempty"`
}

// StorageAccountCredentials access information for the storage account.
type StorageAccountCredentials struct {
	// ResourceID - READ-ONLY; The ARM resource ID of the storage account.
	ResourceID *string `json:"resourceId,omitempty"`
	// PrimaryKey - READ-ONLY; The primary key of the storage account.
	PrimaryKey *string `json:"primaryKey,omitempty"`
	// SecondaryKey - READ-ONLY; The secondary key of the storage account.
	SecondaryKey *string `json:"secondaryKey,omitempty"`
}

// MarshalJSON is the custom marshaler for StorageAccountCredentials.
func (sac StorageAccountCredentials) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// StorageAccountProperties properties of Storage Account.
type StorageAccountProperties struct {
	// ResourceID - ARM resource ID of the Azure Storage Account to store CLI specific files. If not provided one will be created. This cannot be changed once the cluster is created.
	ResourceID *string `json:"resourceId,omitempty"`
}

// SystemService information about a system service deployed in the cluster
type SystemService struct {
	// SystemServiceType - The system service type. Possible values include: 'SystemServiceTypeNone', 'SystemServiceTypeScoringFrontEnd', 'SystemServiceTypeBatchFrontEnd'
	SystemServiceType SystemServiceType `json:"systemServiceType,omitempty"`
	// PublicIPAddress - READ-ONLY; The public IP address of the system service
	PublicIPAddress *string `json:"publicIpAddress,omitempty"`
	// Version - READ-ONLY; The state of the system service
	Version *string `json:"version,omitempty"`
}

// MarshalJSON is the custom marshaler for SystemService.
func (ss SystemService) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ss.SystemServiceType != "" {
		objectMap["systemServiceType"] = ss.SystemServiceType
	}
	return json.Marshal(objectMap)
}

// UpdateSystemServicesResponse response of the update system services API
type UpdateSystemServicesResponse struct {
	autorest.Response `json:"-"`
	// UpdateStatus - READ-ONLY; Update status. Possible values include: 'Unknown', 'Updating', 'Creating', 'Deleting', 'Succeeded', 'Failed', 'Canceled'
	UpdateStatus OperationStatus `json:"updateStatus,omitempty"`
	// UpdateStartedOn - READ-ONLY; The date and time when the last system services update was started.
	UpdateStartedOn *date.Time `json:"updateStartedOn,omitempty"`
	// UpdateCompletedOn - READ-ONLY; The date and time when the last system services update completed.
	UpdateCompletedOn *date.Time `json:"updateCompletedOn,omitempty"`
}

// MarshalJSON is the custom marshaler for UpdateSystemServicesResponse.
func (ussr UpdateSystemServicesResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}