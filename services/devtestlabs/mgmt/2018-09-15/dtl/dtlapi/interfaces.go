// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package dtlapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/devtestlabs/mgmt/2018-09-15/dtl"
	"github.com/Azure/go-autorest/autorest"
)

// ProviderOperationsClientAPI contains the set of methods on the ProviderOperationsClient type.
type ProviderOperationsClientAPI interface {
	List(ctx context.Context) (result dtl.ProviderOperationResultPage, err error)
	ListComplete(ctx context.Context) (result dtl.ProviderOperationResultIterator, err error)
}

var _ ProviderOperationsClientAPI = (*dtl.ProviderOperationsClient)(nil)

// LabsClientAPI contains the set of methods on the LabsClient type.
type LabsClientAPI interface {
	ClaimAnyVM(ctx context.Context, resourceGroupName string, name string) (result dtl.LabsClaimAnyVMFuture, err error)
	CreateEnvironment(ctx context.Context, resourceGroupName string, name string, labVirtualMachineCreationParameter dtl.LabVirtualMachineCreationParameter) (result dtl.LabsCreateEnvironmentFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, lab dtl.Lab) (result dtl.LabsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, name string) (result dtl.LabsDeleteFuture, err error)
	ExportResourceUsage(ctx context.Context, resourceGroupName string, name string, exportResourceUsageParameters dtl.ExportResourceUsageParameters) (result dtl.LabsExportResourceUsageFuture, err error)
	GenerateUploadURI(ctx context.Context, resourceGroupName string, name string, generateUploadURIParameter dtl.GenerateUploadURIParameter) (result dtl.GenerateUploadURIResponse, err error)
	Get(ctx context.Context, resourceGroupName string, name string, expand string) (result dtl.Lab, err error)
	ImportVirtualMachine(ctx context.Context, resourceGroupName string, name string, importLabVirtualMachineRequest dtl.ImportLabVirtualMachineRequest) (result dtl.LabsImportVirtualMachineFuture, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, expand string, filter string, top *int32, orderby string) (result dtl.LabListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, expand string, filter string, top *int32, orderby string) (result dtl.LabListIterator, err error)
	ListBySubscription(ctx context.Context, expand string, filter string, top *int32, orderby string) (result dtl.LabListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, expand string, filter string, top *int32, orderby string) (result dtl.LabListIterator, err error)
	ListVhds(ctx context.Context, resourceGroupName string, name string) (result dtl.LabVhdListPage, err error)
	ListVhdsComplete(ctx context.Context, resourceGroupName string, name string) (result dtl.LabVhdListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, name string, lab dtl.LabFragment) (result dtl.Lab, err error)
}

var _ LabsClientAPI = (*dtl.LabsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	Get(ctx context.Context, locationName string, name string) (result dtl.OperationResult, err error)
}

var _ OperationsClientAPI = (*dtl.OperationsClient)(nil)

// GlobalSchedulesClientAPI contains the set of methods on the GlobalSchedulesClient type.
type GlobalSchedulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, schedule dtl.Schedule) (result dtl.Schedule, err error)
	Delete(ctx context.Context, resourceGroupName string, name string) (result autorest.Response, err error)
	Execute(ctx context.Context, resourceGroupName string, name string) (result dtl.GlobalSchedulesExecuteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, name string, expand string) (result dtl.Schedule, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListIterator, err error)
	ListBySubscription(ctx context.Context, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListPage, err error)
	ListBySubscriptionComplete(ctx context.Context, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListIterator, err error)
	Retarget(ctx context.Context, resourceGroupName string, name string, retargetScheduleProperties dtl.RetargetScheduleProperties) (result dtl.GlobalSchedulesRetargetFuture, err error)
	Update(ctx context.Context, resourceGroupName string, name string, schedule dtl.ScheduleFragment) (result dtl.Schedule, err error)
}

var _ GlobalSchedulesClientAPI = (*dtl.GlobalSchedulesClient)(nil)

// ArtifactSourcesClientAPI contains the set of methods on the ArtifactSourcesClient type.
type ArtifactSourcesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, artifactSource dtl.ArtifactSource) (result dtl.ArtifactSource, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string, expand string) (result dtl.ArtifactSource, err error)
	List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.ArtifactSourceListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.ArtifactSourceListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, name string, artifactSource dtl.ArtifactSourceFragment) (result dtl.ArtifactSource, err error)
}

var _ ArtifactSourcesClientAPI = (*dtl.ArtifactSourcesClient)(nil)

// ArmTemplatesClientAPI contains the set of methods on the ArmTemplatesClient type.
type ArmTemplatesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, expand string) (result dtl.ArmTemplate, err error)
	List(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (result dtl.ArmTemplateListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (result dtl.ArmTemplateListIterator, err error)
}

var _ ArmTemplatesClientAPI = (*dtl.ArmTemplatesClient)(nil)

// ArtifactsClientAPI contains the set of methods on the ArtifactsClient type.
type ArtifactsClientAPI interface {
	GenerateArmTemplate(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, generateArmTemplateRequest dtl.GenerateArmTemplateRequest) (result dtl.ArmTemplateInfo, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, name string, expand string) (result dtl.Artifact, err error)
	List(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (result dtl.ArtifactListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, artifactSourceName string, expand string, filter string, top *int32, orderby string) (result dtl.ArtifactListIterator, err error)
}

var _ ArtifactsClientAPI = (*dtl.ArtifactsClient)(nil)

// CostsClientAPI contains the set of methods on the CostsClient type.
type CostsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, labCost dtl.LabCost) (result dtl.LabCost, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string, expand string) (result dtl.LabCost, err error)
}

var _ CostsClientAPI = (*dtl.CostsClient)(nil)

// CustomImagesClientAPI contains the set of methods on the CustomImagesClient type.
type CustomImagesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, customImage dtl.CustomImage) (result dtl.CustomImagesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.CustomImagesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string, expand string) (result dtl.CustomImage, err error)
	List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.CustomImageListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.CustomImageListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, name string, customImage dtl.CustomImageFragment) (result dtl.CustomImage, err error)
}

var _ CustomImagesClientAPI = (*dtl.CustomImagesClient)(nil)

// FormulasClientAPI contains the set of methods on the FormulasClient type.
type FormulasClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, formula dtl.Formula) (result dtl.FormulasCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string, expand string) (result dtl.Formula, err error)
	List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.FormulaListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.FormulaListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, name string, formula dtl.FormulaFragment) (result dtl.Formula, err error)
}

var _ FormulasClientAPI = (*dtl.FormulasClient)(nil)

// GalleryImagesClientAPI contains the set of methods on the GalleryImagesClient type.
type GalleryImagesClientAPI interface {
	List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.GalleryImageListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.GalleryImageListIterator, err error)
}

var _ GalleryImagesClientAPI = (*dtl.GalleryImagesClient)(nil)

// NotificationChannelsClientAPI contains the set of methods on the NotificationChannelsClient type.
type NotificationChannelsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, notificationChannel dtl.NotificationChannel) (result dtl.NotificationChannel, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string, expand string) (result dtl.NotificationChannel, err error)
	List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.NotificationChannelListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.NotificationChannelListIterator, err error)
	Notify(ctx context.Context, resourceGroupName string, labName string, name string, notifyParameters dtl.NotifyParameters) (result autorest.Response, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, name string, notificationChannel dtl.NotificationChannelFragment) (result dtl.NotificationChannel, err error)
}

var _ NotificationChannelsClientAPI = (*dtl.NotificationChannelsClient)(nil)

// PolicySetsClientAPI contains the set of methods on the PolicySetsClient type.
type PolicySetsClientAPI interface {
	EvaluatePolicies(ctx context.Context, resourceGroupName string, labName string, name string, evaluatePoliciesRequest dtl.EvaluatePoliciesRequest) (result dtl.EvaluatePoliciesResponse, err error)
}

var _ PolicySetsClientAPI = (*dtl.PolicySetsClient)(nil)

// PoliciesClientAPI contains the set of methods on the PoliciesClient type.
type PoliciesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, policySetName string, name string, policy dtl.Policy) (result dtl.Policy, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, policySetName string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, policySetName string, name string, expand string) (result dtl.Policy, err error)
	List(ctx context.Context, resourceGroupName string, labName string, policySetName string, expand string, filter string, top *int32, orderby string) (result dtl.PolicyListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, policySetName string, expand string, filter string, top *int32, orderby string) (result dtl.PolicyListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, policySetName string, name string, policy dtl.PolicyFragment) (result dtl.Policy, err error)
}

var _ PoliciesClientAPI = (*dtl.PoliciesClient)(nil)

// SchedulesClientAPI contains the set of methods on the SchedulesClient type.
type SchedulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, schedule dtl.Schedule) (result dtl.Schedule, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, name string) (result autorest.Response, err error)
	Execute(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.SchedulesExecuteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string, expand string) (result dtl.Schedule, err error)
	List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListIterator, err error)
	ListApplicable(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.ScheduleListPage, err error)
	ListApplicableComplete(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.ScheduleListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, name string, schedule dtl.ScheduleFragment) (result dtl.Schedule, err error)
}

var _ SchedulesClientAPI = (*dtl.SchedulesClient)(nil)

// ServiceRunnersClientAPI contains the set of methods on the ServiceRunnersClient type.
type ServiceRunnersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, serviceRunner dtl.ServiceRunner) (result dtl.ServiceRunner, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.ServiceRunner, err error)
}

var _ ServiceRunnersClientAPI = (*dtl.ServiceRunnersClient)(nil)

// UsersClientAPI contains the set of methods on the UsersClient type.
type UsersClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, userParameter dtl.User) (result dtl.UsersCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.UsersDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string, expand string) (result dtl.User, err error)
	List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.UserListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.UserListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, name string, userParameter dtl.UserFragment) (result dtl.User, err error)
}

var _ UsersClientAPI = (*dtl.UsersClient)(nil)

// DisksClientAPI contains the set of methods on the DisksClient type.
type DisksClientAPI interface {
	Attach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties dtl.AttachDiskProperties) (result dtl.DisksAttachFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk dtl.Disk) (result dtl.DisksCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result dtl.DisksDeleteFuture, err error)
	Detach(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties dtl.DetachDiskProperties) (result dtl.DisksDetachFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, expand string) (result dtl.Disk, err error)
	List(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result dtl.DiskListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result dtl.DiskListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk dtl.DiskFragment) (result dtl.Disk, err error)
}

var _ DisksClientAPI = (*dtl.DisksClient)(nil)

// EnvironmentsClientAPI contains the set of methods on the EnvironmentsClient type.
type EnvironmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, dtlEnvironment dtl.Environment) (result dtl.EnvironmentsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result dtl.EnvironmentsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, expand string) (result dtl.Environment, err error)
	List(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result dtl.EnvironmentListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result dtl.EnvironmentListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, dtlEnvironment dtl.EnvironmentFragment) (result dtl.Environment, err error)
}

var _ EnvironmentsClientAPI = (*dtl.EnvironmentsClient)(nil)

// SecretsClientAPI contains the set of methods on the SecretsClient type.
type SecretsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, secret dtl.Secret) (result dtl.SecretsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, expand string) (result dtl.Secret, err error)
	List(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result dtl.SecretListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result dtl.SecretListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, secret dtl.SecretFragment) (result dtl.Secret, err error)
}

var _ SecretsClientAPI = (*dtl.SecretsClient)(nil)

// ServiceFabricsClientAPI contains the set of methods on the ServiceFabricsClient type.
type ServiceFabricsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric dtl.ServiceFabric) (result dtl.ServiceFabricsCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result dtl.ServiceFabricsDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, userName string, name string, expand string) (result dtl.ServiceFabric, err error)
	List(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result dtl.ServiceFabricListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, userName string, expand string, filter string, top *int32, orderby string) (result dtl.ServiceFabricListIterator, err error)
	ListApplicableSchedules(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result dtl.ApplicableSchedule, err error)
	Start(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result dtl.ServiceFabricsStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, labName string, userName string, name string) (result dtl.ServiceFabricsStopFuture, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, userName string, name string, serviceFabric dtl.ServiceFabricFragment) (result dtl.ServiceFabric, err error)
}

var _ ServiceFabricsClientAPI = (*dtl.ServiceFabricsClient)(nil)

// ServiceFabricSchedulesClientAPI contains the set of methods on the ServiceFabricSchedulesClient type.
type ServiceFabricSchedulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule dtl.Schedule) (result dtl.Schedule, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string) (result autorest.Response, err error)
	Execute(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string) (result dtl.ServiceFabricSchedulesExecuteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, expand string) (result dtl.Schedule, err error)
	List(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, userName string, serviceFabricName string, name string, schedule dtl.ScheduleFragment) (result dtl.Schedule, err error)
}

var _ ServiceFabricSchedulesClientAPI = (*dtl.ServiceFabricSchedulesClient)(nil)

// VirtualMachinesClientAPI contains the set of methods on the VirtualMachinesClient type.
type VirtualMachinesClientAPI interface {
	AddDataDisk(ctx context.Context, resourceGroupName string, labName string, name string, dataDiskProperties dtl.DataDiskProperties) (result dtl.VirtualMachinesAddDataDiskFuture, err error)
	ApplyArtifacts(ctx context.Context, resourceGroupName string, labName string, name string, applyArtifactsRequest dtl.ApplyArtifactsRequest) (result dtl.VirtualMachinesApplyArtifactsFuture, err error)
	Claim(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachinesClaimFuture, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, labVirtualMachine dtl.LabVirtualMachine) (result dtl.VirtualMachinesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachinesDeleteFuture, err error)
	DetachDataDisk(ctx context.Context, resourceGroupName string, labName string, name string, detachDataDiskProperties dtl.DetachDataDiskProperties) (result dtl.VirtualMachinesDetachDataDiskFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string, expand string) (result dtl.LabVirtualMachine, err error)
	GetRdpFileContents(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.RdpConnection, err error)
	List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.LabVirtualMachineListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.LabVirtualMachineListIterator, err error)
	ListApplicableSchedules(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.ApplicableSchedule, err error)
	Redeploy(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachinesRedeployFuture, err error)
	Resize(ctx context.Context, resourceGroupName string, labName string, name string, resizeLabVirtualMachineProperties dtl.ResizeLabVirtualMachineProperties) (result dtl.VirtualMachinesResizeFuture, err error)
	Restart(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachinesRestartFuture, err error)
	Start(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachinesStartFuture, err error)
	Stop(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachinesStopFuture, err error)
	TransferDisks(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachinesTransferDisksFuture, err error)
	UnClaim(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualMachinesUnClaimFuture, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, name string, labVirtualMachine dtl.LabVirtualMachineFragment) (result dtl.LabVirtualMachine, err error)
}

var _ VirtualMachinesClientAPI = (*dtl.VirtualMachinesClient)(nil)

// VirtualMachineSchedulesClientAPI contains the set of methods on the VirtualMachineSchedulesClient type.
type VirtualMachineSchedulesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, schedule dtl.Schedule) (result dtl.Schedule, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string) (result autorest.Response, err error)
	Execute(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string) (result dtl.VirtualMachineSchedulesExecuteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, expand string) (result dtl.Schedule, err error)
	List(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, expand string, filter string, top *int32, orderby string) (result dtl.ScheduleListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, schedule dtl.ScheduleFragment) (result dtl.Schedule, err error)
}

var _ VirtualMachineSchedulesClientAPI = (*dtl.VirtualMachineSchedulesClient)(nil)

// VirtualNetworksClientAPI contains the set of methods on the VirtualNetworksClient type.
type VirtualNetworksClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, virtualNetwork dtl.VirtualNetwork) (result dtl.VirtualNetworksCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, labName string, name string) (result dtl.VirtualNetworksDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, labName string, name string, expand string) (result dtl.VirtualNetwork, err error)
	List(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.VirtualNetworkListPage, err error)
	ListComplete(ctx context.Context, resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result dtl.VirtualNetworkListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, labName string, name string, virtualNetwork dtl.VirtualNetworkFragment) (result dtl.VirtualNetwork, err error)
}

var _ VirtualNetworksClientAPI = (*dtl.VirtualNetworksClient)(nil)