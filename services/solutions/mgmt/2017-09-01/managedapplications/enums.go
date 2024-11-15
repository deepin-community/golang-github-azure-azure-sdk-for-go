package managedapplications

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionType enumerates the values for action type.
type ActionType string

const (
	// Internal ...
	Internal ActionType = "Internal"
)

// PossibleActionTypeValues returns an array of possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{Internal}
}

// ApplicationArtifactType enumerates the values for application artifact type.
type ApplicationArtifactType string

const (
	// Custom ...
	Custom ApplicationArtifactType = "Custom"
	// Template ...
	Template ApplicationArtifactType = "Template"
)

// PossibleApplicationArtifactTypeValues returns an array of possible values for the ApplicationArtifactType const type.
func PossibleApplicationArtifactTypeValues() []ApplicationArtifactType {
	return []ApplicationArtifactType{Custom, Template}
}

// ApplicationLockLevel enumerates the values for application lock level.
type ApplicationLockLevel string

const (
	// CanNotDelete ...
	CanNotDelete ApplicationLockLevel = "CanNotDelete"
	// None ...
	None ApplicationLockLevel = "None"
	// ReadOnly ...
	ReadOnly ApplicationLockLevel = "ReadOnly"
)

// PossibleApplicationLockLevelValues returns an array of possible values for the ApplicationLockLevel const type.
func PossibleApplicationLockLevelValues() []ApplicationLockLevel {
	return []ApplicationLockLevel{CanNotDelete, None, ReadOnly}
}

// Origin enumerates the values for origin.
type Origin string

const (
	// System ...
	System Origin = "system"
	// User ...
	User Origin = "user"
	// Usersystem ...
	Usersystem Origin = "user,system"
)

// PossibleOriginValues returns an array of possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{System, User, Usersystem}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Accepted ...
	Accepted ProvisioningState = "Accepted"
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Created ...
	Created ProvisioningState = "Created"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Deleted ...
	Deleted ProvisioningState = "Deleted"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Ready ...
	Ready ProvisioningState = "Ready"
	// Running ...
	Running ProvisioningState = "Running"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Accepted, Canceled, Created, Creating, Deleted, Deleting, Failed, Ready, Running, Succeeded, Updating}
}

// ResourceIdentityType enumerates the values for resource identity type.
type ResourceIdentityType string

const (
	// SystemAssigned ...
	SystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns an array of possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{SystemAssigned}
}
