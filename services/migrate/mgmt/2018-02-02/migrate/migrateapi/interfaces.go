// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package migrateapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/migrate/mgmt/2018-02-02/migrate"
	"github.com/Azure/go-autorest/autorest"
)

// LocationClientAPI contains the set of methods on the LocationClient type.
type LocationClientAPI interface {
	CheckNameAvailability(ctx context.Context, locationName string, parameters migrate.CheckNameAvailabilityParameters) (result migrate.CheckNameAvailabilityResult, err error)
}

var _ LocationClientAPI = (*migrate.LocationClient)(nil)

// AssessmentOptionsClientAPI contains the set of methods on the AssessmentOptionsClient type.
type AssessmentOptionsClientAPI interface {
	Get(ctx context.Context, locationName string) (result migrate.AssessmentOptionsResultList, err error)
}

var _ AssessmentOptionsClientAPI = (*migrate.AssessmentOptionsClient)(nil)

// ProjectsClientAPI contains the set of methods on the ProjectsClient type.
type ProjectsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, projectName string, project *migrate.Project) (result migrate.Project, err error)
	Delete(ctx context.Context, resourceGroupName string, projectName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, projectName string) (result migrate.Project, err error)
	GetKeys(ctx context.Context, resourceGroupName string, projectName string) (result migrate.ProjectKey, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result migrate.ProjectResultList, err error)
	ListBySubscription(ctx context.Context) (result migrate.ProjectResultList, err error)
	Update(ctx context.Context, resourceGroupName string, projectName string, project *migrate.Project) (result migrate.Project, err error)
}

var _ ProjectsClientAPI = (*migrate.ProjectsClient)(nil)

// MachinesClientAPI contains the set of methods on the MachinesClient type.
type MachinesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, projectName string, machineName string) (result migrate.Machine, err error)
	ListByProject(ctx context.Context, resourceGroupName string, projectName string) (result migrate.MachineResultList, err error)
}

var _ MachinesClientAPI = (*migrate.MachinesClient)(nil)

// GroupsClientAPI contains the set of methods on the GroupsClient type.
type GroupsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, projectName string, groupName string, group *migrate.Group) (result migrate.Group, err error)
	Delete(ctx context.Context, resourceGroupName string, projectName string, groupName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, projectName string, groupName string) (result migrate.Group, err error)
	ListByProject(ctx context.Context, resourceGroupName string, projectName string) (result migrate.GroupResultList, err error)
}

var _ GroupsClientAPI = (*migrate.GroupsClient)(nil)

// AssessmentsClientAPI contains the set of methods on the AssessmentsClient type.
type AssessmentsClientAPI interface {
	Create(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessment *migrate.Assessment) (result migrate.Assessment, err error)
	Delete(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (result migrate.Assessment, err error)
	GetReportDownloadURL(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (result migrate.DownloadURL, err error)
	ListByGroup(ctx context.Context, resourceGroupName string, projectName string, groupName string) (result migrate.AssessmentResultList, err error)
	ListByProject(ctx context.Context, resourceGroupName string, projectName string) (result migrate.AssessmentResultList, err error)
}

var _ AssessmentsClientAPI = (*migrate.AssessmentsClient)(nil)

// AssessedMachinesClientAPI contains the set of methods on the AssessedMachinesClient type.
type AssessedMachinesClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, assessedMachineName string) (result migrate.AssessedMachine, err error)
	ListByAssessment(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string) (result migrate.AssessedMachineResultList, err error)
}

var _ AssessedMachinesClientAPI = (*migrate.AssessedMachinesClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result migrate.OperationResultList, err error)
}

var _ OperationsClientAPI = (*migrate.OperationsClient)(nil)