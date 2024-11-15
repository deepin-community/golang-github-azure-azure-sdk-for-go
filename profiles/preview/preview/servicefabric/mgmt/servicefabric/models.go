//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package servicefabric

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/servicefabric/mgmt/2017-07-01-preview/servicefabric"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ClusterState = original.ClusterState

const (
	AutoScale                 ClusterState = original.AutoScale
	BaselineUpgrade           ClusterState = original.BaselineUpgrade
	Deploying                 ClusterState = original.Deploying
	EnforcingClusterVersion   ClusterState = original.EnforcingClusterVersion
	Ready                     ClusterState = original.Ready
	UpdatingInfrastructure    ClusterState = original.UpdatingInfrastructure
	UpdatingUserCertificate   ClusterState = original.UpdatingUserCertificate
	UpdatingUserConfiguration ClusterState = original.UpdatingUserConfiguration
	UpgradeServiceUnreachable ClusterState = original.UpgradeServiceUnreachable
	WaitingForNodes           ClusterState = original.WaitingForNodes
)

type DefaultMoveCost = original.DefaultMoveCost

const (
	High   DefaultMoveCost = original.High
	Low    DefaultMoveCost = original.Low
	Medium DefaultMoveCost = original.Medium
	Zero   DefaultMoveCost = original.Zero
)

type DurabilityLevel = original.DurabilityLevel

const (
	Bronze DurabilityLevel = original.Bronze
	Gold   DurabilityLevel = original.Gold
	Silver DurabilityLevel = original.Silver
)

type Environment = original.Environment

const (
	Linux   Environment = original.Linux
	Windows Environment = original.Windows
)

type PartitionScheme = original.PartitionScheme

const (
	PartitionSchemeNamed                      PartitionScheme = original.PartitionSchemeNamed
	PartitionSchemePartitionSchemeDescription PartitionScheme = original.PartitionSchemePartitionSchemeDescription
	PartitionSchemeSingleton                  PartitionScheme = original.PartitionSchemeSingleton
	PartitionSchemeUniformInt64Range          PartitionScheme = original.PartitionSchemeUniformInt64Range
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type ReliabilityLevel = original.ReliabilityLevel

const (
	ReliabilityLevelBronze   ReliabilityLevel = original.ReliabilityLevelBronze
	ReliabilityLevelGold     ReliabilityLevel = original.ReliabilityLevelGold
	ReliabilityLevelNone     ReliabilityLevel = original.ReliabilityLevelNone
	ReliabilityLevelPlatinum ReliabilityLevel = original.ReliabilityLevelPlatinum
	ReliabilityLevelSilver   ReliabilityLevel = original.ReliabilityLevelSilver
)

type ReliabilityLevel1 = original.ReliabilityLevel1

const (
	ReliabilityLevel1Bronze ReliabilityLevel1 = original.ReliabilityLevel1Bronze
	ReliabilityLevel1Gold   ReliabilityLevel1 = original.ReliabilityLevel1Gold
	ReliabilityLevel1Silver ReliabilityLevel1 = original.ReliabilityLevel1Silver
)

type Scheme = original.Scheme

const (
	Affinity           Scheme = original.Affinity
	AlignedAffinity    Scheme = original.AlignedAffinity
	Invalid            Scheme = original.Invalid
	NonAlignedAffinity Scheme = original.NonAlignedAffinity
)

type ServiceKind = original.ServiceKind

const (
	ServiceKindServiceProperties ServiceKind = original.ServiceKindServiceProperties
	ServiceKindStateful          ServiceKind = original.ServiceKindStateful
	ServiceKindStateless         ServiceKind = original.ServiceKindStateless
)

type ServiceKindBasicServiceUpdateProperties = original.ServiceKindBasicServiceUpdateProperties

const (
	ServiceKindBasicServiceUpdatePropertiesServiceKindServiceUpdateProperties ServiceKindBasicServiceUpdateProperties = original.ServiceKindBasicServiceUpdatePropertiesServiceKindServiceUpdateProperties
	ServiceKindBasicServiceUpdatePropertiesServiceKindStateful                ServiceKindBasicServiceUpdateProperties = original.ServiceKindBasicServiceUpdatePropertiesServiceKindStateful
	ServiceKindBasicServiceUpdatePropertiesServiceKindStateless               ServiceKindBasicServiceUpdateProperties = original.ServiceKindBasicServiceUpdatePropertiesServiceKindStateless
)

type Type = original.Type

const (
	TypeServicePlacementPolicyDescription Type = original.TypeServicePlacementPolicyDescription
)

type UpgradeMode = original.UpgradeMode

const (
	Automatic UpgradeMode = original.Automatic
	Manual    UpgradeMode = original.Manual
)

type UpgradeMode1 = original.UpgradeMode1

const (
	UpgradeMode1Automatic UpgradeMode1 = original.UpgradeMode1Automatic
	UpgradeMode1Manual    UpgradeMode1 = original.UpgradeMode1Manual
)

type Weight = original.Weight

const (
	WeightHigh   Weight = original.WeightHigh
	WeightLow    Weight = original.WeightLow
	WeightMedium Weight = original.WeightMedium
	WeightZero   Weight = original.WeightZero
)

type X509StoreName = original.X509StoreName

const (
	AddressBook          X509StoreName = original.AddressBook
	AuthRoot             X509StoreName = original.AuthRoot
	CertificateAuthority X509StoreName = original.CertificateAuthority
	Disallowed           X509StoreName = original.Disallowed
	My                   X509StoreName = original.My
	Root                 X509StoreName = original.Root
	TrustedPeople        X509StoreName = original.TrustedPeople
	TrustedPublisher     X509StoreName = original.TrustedPublisher
)

type ApplicationHealthPolicy = original.ApplicationHealthPolicy
type ApplicationMetricDescription = original.ApplicationMetricDescription
type ApplicationParameter = original.ApplicationParameter
type ApplicationProperties = original.ApplicationProperties
type ApplicationResource = original.ApplicationResource
type ApplicationResourceList = original.ApplicationResourceList
type ApplicationResourceUpdate = original.ApplicationResourceUpdate
type ApplicationTypeProperties = original.ApplicationTypeProperties
type ApplicationTypeResource = original.ApplicationTypeResource
type ApplicationTypeResourceList = original.ApplicationTypeResourceList
type ApplicationUpdateProperties = original.ApplicationUpdateProperties
type ApplicationUpgradePolicy = original.ApplicationUpgradePolicy
type AvailableOperationDisplay = original.AvailableOperationDisplay
type AzureActiveDirectory = original.AzureActiveDirectory
type BaseClient = original.BaseClient
type BasicPartitionSchemeDescription = original.BasicPartitionSchemeDescription
type BasicServicePlacementPolicyDescription = original.BasicServicePlacementPolicyDescription
type BasicServiceProperties = original.BasicServiceProperties
type BasicServiceUpdateProperties = original.BasicServiceUpdateProperties
type CertificateDescription = original.CertificateDescription
type ClientCertificateCommonName = original.ClientCertificateCommonName
type ClientCertificateThumbprint = original.ClientCertificateThumbprint
type Cluster = original.Cluster
type ClusterCodeVersionsListResult = original.ClusterCodeVersionsListResult
type ClusterCodeVersionsResult = original.ClusterCodeVersionsResult
type ClusterHealthPolicy = original.ClusterHealthPolicy
type ClusterListResult = original.ClusterListResult
type ClusterProperties = original.ClusterProperties
type ClusterPropertiesUpdateParameters = original.ClusterPropertiesUpdateParameters
type ClusterUpdateParameters = original.ClusterUpdateParameters
type ClusterUpgradeDeltaHealthPolicy = original.ClusterUpgradeDeltaHealthPolicy
type ClusterUpgradePolicy = original.ClusterUpgradePolicy
type ClusterVersionDetails = original.ClusterVersionDetails
type ClusterVersionsClient = original.ClusterVersionsClient
type ClustersClient = original.ClustersClient
type ClustersCreateFuture = original.ClustersCreateFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type DiagnosticsStorageAccountConfig = original.DiagnosticsStorageAccountConfig
type EndpointRangeDescription = original.EndpointRangeDescription
type ErrorModel = original.ErrorModel
type NamedPartitionSchemeDescription = original.NamedPartitionSchemeDescription
type NodeTypeDescription = original.NodeTypeDescription
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type OperationsClient = original.OperationsClient
type PartitionSchemeDescription = original.PartitionSchemeDescription
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type RollingUpgradeMonitoringPolicy = original.RollingUpgradeMonitoringPolicy
type ServiceCorrelationDescription = original.ServiceCorrelationDescription
type ServiceLoadMetricDescription = original.ServiceLoadMetricDescription
type ServicePlacementPolicyDescription = original.ServicePlacementPolicyDescription
type ServiceProperties = original.ServiceProperties
type ServicePropertiesBase = original.ServicePropertiesBase
type ServiceResource = original.ServiceResource
type ServiceResourceList = original.ServiceResourceList
type ServiceResourceUpdate = original.ServiceResourceUpdate
type ServiceTypeDeltaHealthPolicy = original.ServiceTypeDeltaHealthPolicy
type ServiceTypeHealthPolicy = original.ServiceTypeHealthPolicy
type ServiceTypeHealthPolicyMapItem = original.ServiceTypeHealthPolicyMapItem
type ServiceUpdateProperties = original.ServiceUpdateProperties
type SettingsParameterDescription = original.SettingsParameterDescription
type SettingsSectionDescription = original.SettingsSectionDescription
type SingletonPartitionSchemeDescription = original.SingletonPartitionSchemeDescription
type StatefulServiceProperties = original.StatefulServiceProperties
type StatefulServiceUpdateProperties = original.StatefulServiceUpdateProperties
type StatelessServiceProperties = original.StatelessServiceProperties
type StatelessServiceUpdateProperties = original.StatelessServiceUpdateProperties
type UniformInt64RangePartitionSchemeDescription = original.UniformInt64RangePartitionSchemeDescription
type VersionProperties = original.VersionProperties
type VersionResource = original.VersionResource
type VersionResourceList = original.VersionResourceList

func New() BaseClient {
	return original.New()
}
func NewClusterVersionsClient() ClusterVersionsClient {
	return original.NewClusterVersionsClient()
}
func NewClusterVersionsClientWithBaseURI(baseURI string) ClusterVersionsClient {
	return original.NewClusterVersionsClientWithBaseURI(baseURI)
}
func NewClustersClient() ClustersClient {
	return original.NewClustersClient()
}
func NewClustersClientWithBaseURI(baseURI string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleClusterStateValues() []ClusterState {
	return original.PossibleClusterStateValues()
}
func PossibleDefaultMoveCostValues() []DefaultMoveCost {
	return original.PossibleDefaultMoveCostValues()
}
func PossibleDurabilityLevelValues() []DurabilityLevel {
	return original.PossibleDurabilityLevelValues()
}
func PossibleEnvironmentValues() []Environment {
	return original.PossibleEnvironmentValues()
}
func PossiblePartitionSchemeValues() []PartitionScheme {
	return original.PossiblePartitionSchemeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleReliabilityLevel1Values() []ReliabilityLevel1 {
	return original.PossibleReliabilityLevel1Values()
}
func PossibleReliabilityLevelValues() []ReliabilityLevel {
	return original.PossibleReliabilityLevelValues()
}
func PossibleSchemeValues() []Scheme {
	return original.PossibleSchemeValues()
}
func PossibleServiceKindBasicServiceUpdatePropertiesValues() []ServiceKindBasicServiceUpdateProperties {
	return original.PossibleServiceKindBasicServiceUpdatePropertiesValues()
}
func PossibleServiceKindValues() []ServiceKind {
	return original.PossibleServiceKindValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleUpgradeMode1Values() []UpgradeMode1 {
	return original.PossibleUpgradeMode1Values()
}
func PossibleUpgradeModeValues() []UpgradeMode {
	return original.PossibleUpgradeModeValues()
}
func PossibleWeightValues() []Weight {
	return original.PossibleWeightValues()
}
func PossibleX509StoreNameValues() []X509StoreName {
	return original.PossibleX509StoreNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
