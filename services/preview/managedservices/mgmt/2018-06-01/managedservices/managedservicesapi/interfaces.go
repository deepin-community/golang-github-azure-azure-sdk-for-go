// Deprecated: Please note, this package has been deprecated. A replacement package is available [github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices](https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices). We strongly encourage you to upgrade to continue receiving updates. See [Migration Guide](https://aka.ms/azsdk/golang/t2/migration) for guidance on upgrading. Refer to our [deprecation policy](https://azure.github.io/azure-sdk/policies_support.html) for more details.
package managedservicesapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/managedservices/mgmt/2018-06-01/managedservices"
	"github.com/Azure/go-autorest/autorest"
)

// RegistrationDefinitionsClientAPI contains the set of methods on the RegistrationDefinitionsClient type.
type RegistrationDefinitionsClientAPI interface {
	CreateOrUpdate(ctx context.Context, registrationDefinitionID string, APIVersion string, scope string, requestBody managedservices.RegistrationDefinition) (result managedservices.RegistrationDefinition, err error)
	Delete(ctx context.Context, registrationDefinitionID string, APIVersion string, scope string) (result autorest.Response, err error)
	Get(ctx context.Context, scope string, registrationDefinitionID string, APIVersion string) (result managedservices.RegistrationDefinition, err error)
	List(ctx context.Context, scope string, APIVersion string) (result managedservices.RegistrationDefinitionListPage, err error)
	ListComplete(ctx context.Context, scope string, APIVersion string) (result managedservices.RegistrationDefinitionListIterator, err error)
}

var _ RegistrationDefinitionsClientAPI = (*managedservices.RegistrationDefinitionsClient)(nil)

// RegistrationAssignmentsClientAPI contains the set of methods on the RegistrationAssignmentsClient type.
type RegistrationAssignmentsClientAPI interface {
	CreateOrUpdate(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string, requestBody managedservices.RegistrationAssignment) (result managedservices.RegistrationAssignment, err error)
	Delete(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string) (result autorest.Response, err error)
	Get(ctx context.Context, scope string, registrationAssignmentID string, APIVersion string, expandRegistrationDefinition *bool) (result managedservices.RegistrationAssignment, err error)
	List(ctx context.Context, scope string, APIVersion string, expandRegistrationDefinition *bool) (result managedservices.RegistrationAssignmentListPage, err error)
	ListComplete(ctx context.Context, scope string, APIVersion string, expandRegistrationDefinition *bool) (result managedservices.RegistrationAssignmentListIterator, err error)
}

var _ RegistrationAssignmentsClientAPI = (*managedservices.RegistrationAssignmentsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context, APIVersion string) (result managedservices.OperationList, err error)
}

var _ OperationsClientAPI = (*managedservices.OperationsClient)(nil)
