package vmwarecloudsimple

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AggregationType enumerates the values for aggregation type.
type AggregationType string

const (
	// Average ...
	Average AggregationType = "Average"
	// Total ...
	Total AggregationType = "Total"
)

// PossibleAggregationTypeValues returns an array of possible values for the AggregationType const type.
func PossibleAggregationTypeValues() []AggregationType {
	return []AggregationType{Average, Total}
}

// Allocation enumerates the values for allocation.
type Allocation string

const (
	// Dynamic ...
	Dynamic Allocation = "dynamic"
	// Static ...
	Static Allocation = "static"
)

// PossibleAllocationValues returns an array of possible values for the Allocation const type.
func PossibleAllocationValues() []Allocation {
	return []Allocation{Dynamic, Static}
}

// DiskIndependenceMode enumerates the values for disk independence mode.
type DiskIndependenceMode string

const (
	// IndependentNonpersistent ...
	IndependentNonpersistent DiskIndependenceMode = "independent_nonpersistent"
	// IndependentPersistent ...
	IndependentPersistent DiskIndependenceMode = "independent_persistent"
	// Persistent ...
	Persistent DiskIndependenceMode = "persistent"
)

// PossibleDiskIndependenceModeValues returns an array of possible values for the DiskIndependenceMode const type.
func PossibleDiskIndependenceModeValues() []DiskIndependenceMode {
	return []DiskIndependenceMode{IndependentNonpersistent, IndependentPersistent, Persistent}
}

// GuestOSType enumerates the values for guest os type.
type GuestOSType string

const (
	// Linux ...
	Linux GuestOSType = "linux"
	// Other ...
	Other GuestOSType = "other"
	// Windows ...
	Windows GuestOSType = "windows"
)

// PossibleGuestOSTypeValues returns an array of possible values for the GuestOSType const type.
func PossibleGuestOSTypeValues() []GuestOSType {
	return []GuestOSType{Linux, Other, Windows}
}

// NICType enumerates the values for nic type.
type NICType string

const (
	// E1000 ...
	E1000 NICType = "E1000"
	// E1000E ...
	E1000E NICType = "E1000E"
	// PCNET32 ...
	PCNET32 NICType = "PCNET32"
	// VMXNET ...
	VMXNET NICType = "VMXNET"
	// VMXNET2 ...
	VMXNET2 NICType = "VMXNET2"
	// VMXNET3 ...
	VMXNET3 NICType = "VMXNET3"
)

// PossibleNICTypeValues returns an array of possible values for the NICType const type.
func PossibleNICTypeValues() []NICType {
	return []NICType{E1000, E1000E, PCNET32, VMXNET, VMXNET2, VMXNET3}
}

// NodeStatus enumerates the values for node status.
type NodeStatus string

const (
	// Unused ...
	Unused NodeStatus = "unused"
	// Used ...
	Used NodeStatus = "used"
)

// PossibleNodeStatusValues returns an array of possible values for the NodeStatus const type.
func PossibleNodeStatusValues() []NodeStatus {
	return []NodeStatus{Unused, Used}
}

// OnboardingStatus enumerates the values for onboarding status.
type OnboardingStatus string

const (
	// NotOnBoarded ...
	NotOnBoarded OnboardingStatus = "notOnBoarded"
	// OnBoarded ...
	OnBoarded OnboardingStatus = "onBoarded"
	// OnBoarding ...
	OnBoarding OnboardingStatus = "onBoarding"
	// OnBoardingFailed ...
	OnBoardingFailed OnboardingStatus = "onBoardingFailed"
)

// PossibleOnboardingStatusValues returns an array of possible values for the OnboardingStatus const type.
func PossibleOnboardingStatusValues() []OnboardingStatus {
	return []OnboardingStatus{NotOnBoarded, OnBoarded, OnBoarding, OnBoardingFailed}
}

// OperationOrigin enumerates the values for operation origin.
type OperationOrigin string

const (
	// System ...
	System OperationOrigin = "system"
	// User ...
	User OperationOrigin = "user"
	// Usersystem ...
	Usersystem OperationOrigin = "user,system"
)

// PossibleOperationOriginValues returns an array of possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{System, User, Usersystem}
}

// PrivateCloudResourceType enumerates the values for private cloud resource type.
type PrivateCloudResourceType string

const (
	// MicrosoftVMwareCloudSimpleprivateClouds ...
	MicrosoftVMwareCloudSimpleprivateClouds PrivateCloudResourceType = "Microsoft.VMwareCloudSimple/privateClouds"
)

// PossiblePrivateCloudResourceTypeValues returns an array of possible values for the PrivateCloudResourceType const type.
func PossiblePrivateCloudResourceTypeValues() []PrivateCloudResourceType {
	return []PrivateCloudResourceType{MicrosoftVMwareCloudSimpleprivateClouds}
}

// StopMode enumerates the values for stop mode.
type StopMode string

const (
	// Poweroff ...
	Poweroff StopMode = "poweroff"
	// Reboot ...
	Reboot StopMode = "reboot"
	// Shutdown ...
	Shutdown StopMode = "shutdown"
	// Suspend ...
	Suspend StopMode = "suspend"
)

// PossibleStopModeValues returns an array of possible values for the StopMode const type.
func PossibleStopModeValues() []StopMode {
	return []StopMode{Poweroff, Reboot, Shutdown, Suspend}
}

// Type enumerates the values for type.
type Type string

const (
	// CUSTOMNAME ...
	CUSTOMNAME Type = "CUSTOM_NAME"
	// FIXED ...
	FIXED Type = "FIXED"
	// PREFIXBASED ...
	PREFIXBASED Type = "PREFIX_BASED"
	// USERDEFINED ...
	USERDEFINED Type = "USER_DEFINED"
	// VIRTUALMACHINENAME ...
	VIRTUALMACHINENAME Type = "VIRTUAL_MACHINE_NAME"
)

// PossibleTypeValues returns an array of possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{CUSTOMNAME, FIXED, PREFIXBASED, USERDEFINED, VIRTUALMACHINENAME}
}

// Type1 enumerates the values for type 1.
type Type1 string

const (
	// Type1CUSTOM ...
	Type1CUSTOM Type1 = "CUSTOM"
	// Type1DHCPIP ...
	Type1DHCPIP Type1 = "DHCP_IP"
	// Type1FIXEDIP ...
	Type1FIXEDIP Type1 = "FIXED_IP"
	// Type1USERDEFINED ...
	Type1USERDEFINED Type1 = "USER_DEFINED"
)

// PossibleType1Values returns an array of possible values for the Type1 const type.
func PossibleType1Values() []Type1 {
	return []Type1{Type1CUSTOM, Type1DHCPIP, Type1FIXEDIP, Type1USERDEFINED}
}

// Type2 enumerates the values for type 2.
type Type2 string

const (
	// LINUX ...
	LINUX Type2 = "LINUX"
	// WINDOWS ...
	WINDOWS Type2 = "WINDOWS"
	// WINDOWSTEXT ...
	WINDOWSTEXT Type2 = "WINDOWS_TEXT"
)

// PossibleType2Values returns an array of possible values for the Type2 const type.
func PossibleType2Values() []Type2 {
	return []Type2{LINUX, WINDOWS, WINDOWSTEXT}
}

// Type3 enumerates the values for type 3.
type Type3 string

const (
	// Type3LINUX ...
	Type3LINUX Type3 = "LINUX"
	// Type3WINDOWS ...
	Type3WINDOWS Type3 = "WINDOWS"
)

// PossibleType3Values returns an array of possible values for the Type3 const type.
func PossibleType3Values() []Type3 {
	return []Type3{Type3LINUX, Type3WINDOWS}
}

// UsageCount enumerates the values for usage count.
type UsageCount string

const (
	// Bytes ...
	Bytes UsageCount = "Bytes"
	// BytesPerSecond ...
	BytesPerSecond UsageCount = "BytesPerSecond"
	// Count ...
	Count UsageCount = "Count"
	// CountPerSecond ...
	CountPerSecond UsageCount = "CountPerSecond"
	// Percent ...
	Percent UsageCount = "Percent"
	// Seconds ...
	Seconds UsageCount = "Seconds"
)

// PossibleUsageCountValues returns an array of possible values for the UsageCount const type.
func PossibleUsageCountValues() []UsageCount {
	return []UsageCount{Bytes, BytesPerSecond, Count, CountPerSecond, Percent, Seconds}
}

// VirtualMachineStatus enumerates the values for virtual machine status.
type VirtualMachineStatus string

const (
	// Deallocating ...
	Deallocating VirtualMachineStatus = "deallocating"
	// Deleting ...
	Deleting VirtualMachineStatus = "deleting"
	// Poweredoff ...
	Poweredoff VirtualMachineStatus = "poweredoff"
	// Running ...
	Running VirtualMachineStatus = "running"
	// Suspended ...
	Suspended VirtualMachineStatus = "suspended"
	// Updating ...
	Updating VirtualMachineStatus = "updating"
)

// PossibleVirtualMachineStatusValues returns an array of possible values for the VirtualMachineStatus const type.
func PossibleVirtualMachineStatusValues() []VirtualMachineStatus {
	return []VirtualMachineStatus{Deallocating, Deleting, Poweredoff, Running, Suspended, Updating}
}