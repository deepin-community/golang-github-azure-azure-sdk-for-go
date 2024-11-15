// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package backupapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/recoveryservices/mgmt/2016-06-01/backup"
	"github.com/Azure/go-autorest/autorest"
)

// ItemLevelRecoveryConnectionsClientAPI contains the set of methods on the ItemLevelRecoveryConnectionsClient type.
type ItemLevelRecoveryConnectionsClientAPI interface {
	Provision(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, resourceILRRequest backup.ILRRequestResource) (result autorest.Response, err error)
	Revoke(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (result autorest.Response, err error)
}

var _ ItemLevelRecoveryConnectionsClientAPI = (*backup.ItemLevelRecoveryConnectionsClient)(nil)

// RestoresClientAPI contains the set of methods on the RestoresClient type.
type RestoresClientAPI interface {
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, resourceRestoreRequest backup.RestoreRequestResource) (result autorest.Response, err error)
}

var _ RestoresClientAPI = (*backup.RestoresClient)(nil)

// ProtectionPolicyOperationStatusesClientAPI contains the set of methods on the ProtectionPolicyOperationStatusesClient type.
type ProtectionPolicyOperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result backup.OperationStatus, err error)
}

var _ ProtectionPolicyOperationStatusesClientAPI = (*backup.ProtectionPolicyOperationStatusesClient)(nil)

// ProtectionPolicyOperationResultsClientAPI contains the set of methods on the ProtectionPolicyOperationResultsClient type.
type ProtectionPolicyOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string, operationID string) (result backup.ProtectionPolicyResource, err error)
}

var _ ProtectionPolicyOperationResultsClientAPI = (*backup.ProtectionPolicyOperationResultsClient)(nil)

// ProtectionPoliciesClientAPI contains the set of methods on the ProtectionPoliciesClient type.
type ProtectionPoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, policyName string, resourceProtectionPolicy backup.ProtectionPolicyResource) (result backup.ProtectionPolicyResource, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string, policyName string) (result autorest.Response, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, policyName string) (result backup.ProtectionPolicyResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result backup.ProtectionPolicyResourceList, err error)
}

var _ ProtectionPoliciesClientAPI = (*backup.ProtectionPoliciesClient)(nil)

// ProtectionContainerOperationResultsClientAPI contains the set of methods on the ProtectionContainerOperationResultsClient type.
type ProtectionContainerOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, operationID string) (result backup.ProtectionContainerResource, err error)
}

var _ ProtectionContainerOperationResultsClientAPI = (*backup.ProtectionContainerOperationResultsClient)(nil)

// ProtectionContainerRefreshOperationResultsClientAPI contains the set of methods on the ProtectionContainerRefreshOperationResultsClient type.
type ProtectionContainerRefreshOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, operationID string) (result autorest.Response, err error)
}

var _ ProtectionContainerRefreshOperationResultsClientAPI = (*backup.ProtectionContainerRefreshOperationResultsClient)(nil)

// ProtectionContainersClientAPI contains the set of methods on the ProtectionContainersClient type.
type ProtectionContainersClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string) (result backup.ProtectionContainerResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result backup.ProtectionContainerResourceList, err error)
	Refresh(ctx context.Context, vaultName string, resourceGroupName string, fabricName string) (result autorest.Response, err error)
	Unregister(ctx context.Context, resourceGroupName string, vaultName string, identityName string) (result autorest.Response, err error)
}

var _ ProtectionContainersClientAPI = (*backup.ProtectionContainersClient)(nil)

// RecoveryPointsClientAPI contains the set of methods on the RecoveryPointsClient type.
type RecoveryPointsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (result backup.RecoveryPointResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (result backup.RecoveryPointResourceList, err error)
}

var _ RecoveryPointsClientAPI = (*backup.RecoveryPointsClient)(nil)

// BackupsClientAPI contains the set of methods on the BackupsClient type.
type BackupsClientAPI interface {
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, resourceBackupRequest backup.RequestResource) (result autorest.Response, err error)
}

var _ BackupsClientAPI = (*backup.BackupsClient)(nil)

// ProtectedItemOperationStatusesClientAPI contains the set of methods on the ProtectedItemOperationStatusesClient type.
type ProtectedItemOperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result backup.OperationStatus, err error)
}

var _ ProtectedItemOperationStatusesClientAPI = (*backup.ProtectedItemOperationStatusesClient)(nil)

// ProtectedItemOperationResultsClientAPI contains the set of methods on the ProtectedItemOperationResultsClient type.
type ProtectedItemOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, operationID string) (result backup.ProtectedItemResource, err error)
}

var _ ProtectedItemOperationResultsClientAPI = (*backup.ProtectedItemOperationResultsClient)(nil)

// ProtectedItemsClientAPI contains the set of methods on the ProtectedItemsClient type.
type ProtectedItemsClientAPI interface {
	CreateOrUpdate(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, resourceProtectedItem backup.ProtectedItemResource) (result autorest.Response, err error)
	Delete(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string) (result autorest.Response, err error)
	Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (result backup.ProtectedItemResource, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.ProtectedItemResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.ProtectedItemResourceListIterator, err error)
}

var _ ProtectedItemsClientAPI = (*backup.ProtectedItemsClient)(nil)

// ProtectableItemsClientAPI contains the set of methods on the ProtectableItemsClient type.
type ProtectableItemsClientAPI interface {
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.WorkloadProtectableItemResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.WorkloadProtectableItemResourceListIterator, err error)
}

var _ ProtectableItemsClientAPI = (*backup.ProtectableItemsClient)(nil)

// ExportJobsOperationResultsClientAPI contains the set of methods on the ExportJobsOperationResultsClient type.
type ExportJobsOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result backup.OperationResultInfoBaseResource, err error)
}

var _ ExportJobsOperationResultsClientAPI = (*backup.ExportJobsOperationResultsClient)(nil)

// JobOperationResultsClientAPI contains the set of methods on the JobOperationResultsClient type.
type JobOperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, jobName string, operationID string) (result autorest.Response, err error)
}

var _ JobOperationResultsClientAPI = (*backup.JobOperationResultsClient)(nil)

// JobsClientAPI contains the set of methods on the JobsClient type.
type JobsClientAPI interface {
	Export(ctx context.Context, vaultName string, resourceGroupName string, filter string) (result autorest.Response, err error)
	List(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.JobResourceListPage, err error)
	ListComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.JobResourceListIterator, err error)
}

var _ JobsClientAPI = (*backup.JobsClient)(nil)

// JobCancellationsClientAPI contains the set of methods on the JobCancellationsClient type.
type JobCancellationsClientAPI interface {
	Trigger(ctx context.Context, vaultName string, resourceGroupName string, jobName string) (result autorest.Response, err error)
}

var _ JobCancellationsClientAPI = (*backup.JobCancellationsClient)(nil)

// JobDetailsClientAPI contains the set of methods on the JobDetailsClient type.
type JobDetailsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, jobName string) (result backup.JobResource, err error)
}

var _ JobDetailsClientAPI = (*backup.JobDetailsClient)(nil)

// OperationStatusesClientAPI contains the set of methods on the OperationStatusesClient type.
type OperationStatusesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result backup.OperationStatus, err error)
}

var _ OperationStatusesClientAPI = (*backup.OperationStatusesClient)(nil)

// OperationResultsClientAPI contains the set of methods on the OperationResultsClient type.
type OperationResultsClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, operationID string) (result autorest.Response, err error)
}

var _ OperationResultsClientAPI = (*backup.OperationResultsClient)(nil)

// EnginesClientAPI contains the set of methods on the EnginesClient type.
type EnginesClientAPI interface {
	Get(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.EngineBaseResourceListPage, err error)
	GetComplete(ctx context.Context, vaultName string, resourceGroupName string, filter string, skipToken string) (result backup.EngineBaseResourceListIterator, err error)
}

var _ EnginesClientAPI = (*backup.EnginesClient)(nil)
