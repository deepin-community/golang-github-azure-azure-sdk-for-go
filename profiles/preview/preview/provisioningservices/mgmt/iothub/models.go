//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package iothub

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/provisioningservices/mgmt/2017-08-21-preview/iothub"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AccessRightsDescription = original.AccessRightsDescription

const (
	DeviceConnect           AccessRightsDescription = original.DeviceConnect
	EnrollmentRead          AccessRightsDescription = original.EnrollmentRead
	EnrollmentWrite         AccessRightsDescription = original.EnrollmentWrite
	RegistrationStatusRead  AccessRightsDescription = original.RegistrationStatusRead
	RegistrationStatusWrite AccessRightsDescription = original.RegistrationStatusWrite
	ServiceConfig           AccessRightsDescription = original.ServiceConfig
)

type AllocationPolicy = original.AllocationPolicy

const (
	GeoLatency AllocationPolicy = original.GeoLatency
	Hashed     AllocationPolicy = original.Hashed
	Static     AllocationPolicy = original.Static
)

type CertificatePurpose = original.CertificatePurpose

const (
	ClientAuthentication CertificatePurpose = original.ClientAuthentication
	ServerAuthentication CertificatePurpose = original.ServerAuthentication
)

type IotDpsSku = original.IotDpsSku

const (
	S1 IotDpsSku = original.S1
)

type NameUnavailabilityReason = original.NameUnavailabilityReason

const (
	AlreadyExists NameUnavailabilityReason = original.AlreadyExists
	Invalid       NameUnavailabilityReason = original.Invalid
)

type State = original.State

const (
	Activating       State = original.Activating
	ActivationFailed State = original.ActivationFailed
	Active           State = original.Active
	Deleted          State = original.Deleted
	Deleting         State = original.Deleting
	DeletionFailed   State = original.DeletionFailed
	FailingOver      State = original.FailingOver
	FailoverFailed   State = original.FailoverFailed
	Resuming         State = original.Resuming
	Suspended        State = original.Suspended
	Suspending       State = original.Suspending
	Transitioning    State = original.Transitioning
)

type AsyncOperationResult = original.AsyncOperationResult
type BaseClient = original.BaseClient
type CertificateBodyDescription = original.CertificateBodyDescription
type CertificateListDescription = original.CertificateListDescription
type CertificateProperties = original.CertificateProperties
type CertificateResponse = original.CertificateResponse
type DefinitionDescription = original.DefinitionDescription
type DpsCertificateClient = original.DpsCertificateClient
type DpsCertificatesClient = original.DpsCertificatesClient
type ErrorDetails = original.ErrorDetails
type ErrorMesssage = original.ErrorMesssage
type IotDpsPropertiesDescription = original.IotDpsPropertiesDescription
type IotDpsResourceClient = original.IotDpsResourceClient
type IotDpsResourceCreateOrUpdateFuture = original.IotDpsResourceCreateOrUpdateFuture
type IotDpsResourceDeleteFuture = original.IotDpsResourceDeleteFuture
type IotDpsSkuDefinition = original.IotDpsSkuDefinition
type IotDpsSkuDefinitionListResult = original.IotDpsSkuDefinitionListResult
type IotDpsSkuDefinitionListResultIterator = original.IotDpsSkuDefinitionListResultIterator
type IotDpsSkuDefinitionListResultPage = original.IotDpsSkuDefinitionListResultPage
type IotDpsSkuInfo = original.IotDpsSkuInfo
type NameAvailabilityInfo = original.NameAvailabilityInfo
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationInputs = original.OperationInputs
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ProvisioningServiceDescription = original.ProvisioningServiceDescription
type ProvisioningServiceDescriptionListResult = original.ProvisioningServiceDescriptionListResult
type ProvisioningServiceDescriptionListResultIterator = original.ProvisioningServiceDescriptionListResultIterator
type ProvisioningServiceDescriptionListResultPage = original.ProvisioningServiceDescriptionListResultPage
type Resource = original.Resource
type SharedAccessSignatureAuthorizationRuleAccessRightsDescription = original.SharedAccessSignatureAuthorizationRuleAccessRightsDescription
type SharedAccessSignatureAuthorizationRuleListResult = original.SharedAccessSignatureAuthorizationRuleListResult
type SharedAccessSignatureAuthorizationRuleListResultIterator = original.SharedAccessSignatureAuthorizationRuleListResultIterator
type SharedAccessSignatureAuthorizationRuleListResultPage = original.SharedAccessSignatureAuthorizationRuleListResultPage
type VerificationCodeRequest = original.VerificationCodeRequest
type VerificationCodeResponse = original.VerificationCodeResponse

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDpsCertificateClient(subscriptionID string) DpsCertificateClient {
	return original.NewDpsCertificateClient(subscriptionID)
}
func NewDpsCertificateClientWithBaseURI(baseURI string, subscriptionID string) DpsCertificateClient {
	return original.NewDpsCertificateClientWithBaseURI(baseURI, subscriptionID)
}
func NewDpsCertificatesClient(subscriptionID string) DpsCertificatesClient {
	return original.NewDpsCertificatesClient(subscriptionID)
}
func NewDpsCertificatesClientWithBaseURI(baseURI string, subscriptionID string) DpsCertificatesClient {
	return original.NewDpsCertificatesClientWithBaseURI(baseURI, subscriptionID)
}
func NewIotDpsResourceClient(subscriptionID string) IotDpsResourceClient {
	return original.NewIotDpsResourceClient(subscriptionID)
}
func NewIotDpsResourceClientWithBaseURI(baseURI string, subscriptionID string) IotDpsResourceClient {
	return original.NewIotDpsResourceClientWithBaseURI(baseURI, subscriptionID)
}
func NewIotDpsSkuDefinitionListResultIterator(page IotDpsSkuDefinitionListResultPage) IotDpsSkuDefinitionListResultIterator {
	return original.NewIotDpsSkuDefinitionListResultIterator(page)
}
func NewIotDpsSkuDefinitionListResultPage(cur IotDpsSkuDefinitionListResult, getNextPage func(context.Context, IotDpsSkuDefinitionListResult) (IotDpsSkuDefinitionListResult, error)) IotDpsSkuDefinitionListResultPage {
	return original.NewIotDpsSkuDefinitionListResultPage(cur, getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProvisioningServiceDescriptionListResultIterator(page ProvisioningServiceDescriptionListResultPage) ProvisioningServiceDescriptionListResultIterator {
	return original.NewProvisioningServiceDescriptionListResultIterator(page)
}
func NewProvisioningServiceDescriptionListResultPage(cur ProvisioningServiceDescriptionListResult, getNextPage func(context.Context, ProvisioningServiceDescriptionListResult) (ProvisioningServiceDescriptionListResult, error)) ProvisioningServiceDescriptionListResultPage {
	return original.NewProvisioningServiceDescriptionListResultPage(cur, getNextPage)
}
func NewSharedAccessSignatureAuthorizationRuleListResultIterator(page SharedAccessSignatureAuthorizationRuleListResultPage) SharedAccessSignatureAuthorizationRuleListResultIterator {
	return original.NewSharedAccessSignatureAuthorizationRuleListResultIterator(page)
}
func NewSharedAccessSignatureAuthorizationRuleListResultPage(cur SharedAccessSignatureAuthorizationRuleListResult, getNextPage func(context.Context, SharedAccessSignatureAuthorizationRuleListResult) (SharedAccessSignatureAuthorizationRuleListResult, error)) SharedAccessSignatureAuthorizationRuleListResultPage {
	return original.NewSharedAccessSignatureAuthorizationRuleListResultPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccessRightsDescriptionValues() []AccessRightsDescription {
	return original.PossibleAccessRightsDescriptionValues()
}
func PossibleAllocationPolicyValues() []AllocationPolicy {
	return original.PossibleAllocationPolicyValues()
}
func PossibleCertificatePurposeValues() []CertificatePurpose {
	return original.PossibleCertificatePurposeValues()
}
func PossibleIotDpsSkuValues() []IotDpsSku {
	return original.PossibleIotDpsSkuValues()
}
func PossibleNameUnavailabilityReasonValues() []NameUnavailabilityReason {
	return original.PossibleNameUnavailabilityReasonValues()
}
func PossibleStateValues() []State {
	return original.PossibleStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}