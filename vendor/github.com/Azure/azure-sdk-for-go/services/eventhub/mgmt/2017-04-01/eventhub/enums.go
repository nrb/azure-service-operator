package eventhub

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccessRights enumerates the values for access rights.
type AccessRights string

const (
	// Listen ...
	Listen AccessRights = "Listen"
	// Manage ...
	Manage AccessRights = "Manage"
	// SendEnumValue ...
	SendEnumValue AccessRights = "Send"
)

// PossibleAccessRightsValues returns an array of possible values for the AccessRights const type.
func PossibleAccessRightsValues() []AccessRights {
	return []AccessRights{Listen, Manage, SendEnumValue}
}

// DefaultAction enumerates the values for default action.
type DefaultAction string

const (
	// Allow ...
	Allow DefaultAction = "Allow"
	// Deny ...
	Deny DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns an array of possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{Allow, Deny}
}

// EncodingCaptureDescription enumerates the values for encoding capture description.
type EncodingCaptureDescription string

const (
	// Avro ...
	Avro EncodingCaptureDescription = "Avro"
	// AvroDeflate ...
	AvroDeflate EncodingCaptureDescription = "AvroDeflate"
)

// PossibleEncodingCaptureDescriptionValues returns an array of possible values for the EncodingCaptureDescription const type.
func PossibleEncodingCaptureDescriptionValues() []EncodingCaptureDescription {
	return []EncodingCaptureDescription{Avro, AvroDeflate}
}

// EntityStatus enumerates the values for entity status.
type EntityStatus string

const (
	// Active ...
	Active EntityStatus = "Active"
	// Creating ...
	Creating EntityStatus = "Creating"
	// Deleting ...
	Deleting EntityStatus = "Deleting"
	// Disabled ...
	Disabled EntityStatus = "Disabled"
	// ReceiveDisabled ...
	ReceiveDisabled EntityStatus = "ReceiveDisabled"
	// Renaming ...
	Renaming EntityStatus = "Renaming"
	// Restoring ...
	Restoring EntityStatus = "Restoring"
	// SendDisabled ...
	SendDisabled EntityStatus = "SendDisabled"
	// Unknown ...
	Unknown EntityStatus = "Unknown"
)

// PossibleEntityStatusValues returns an array of possible values for the EntityStatus const type.
func PossibleEntityStatusValues() []EntityStatus {
	return []EntityStatus{Active, Creating, Deleting, Disabled, ReceiveDisabled, Renaming, Restoring, SendDisabled, Unknown}
}

// KeyType enumerates the values for key type.
type KeyType string

const (
	// PrimaryKey ...
	PrimaryKey KeyType = "PrimaryKey"
	// SecondaryKey ...
	SecondaryKey KeyType = "SecondaryKey"
)

// PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{PrimaryKey, SecondaryKey}
}

// NetworkRuleIPAction enumerates the values for network rule ip action.
type NetworkRuleIPAction string

const (
	// NetworkRuleIPActionAllow ...
	NetworkRuleIPActionAllow NetworkRuleIPAction = "Allow"
)

// PossibleNetworkRuleIPActionValues returns an array of possible values for the NetworkRuleIPAction const type.
func PossibleNetworkRuleIPActionValues() []NetworkRuleIPAction {
	return []NetworkRuleIPAction{NetworkRuleIPActionAllow}
}

// ProvisioningStateDR enumerates the values for provisioning state dr.
type ProvisioningStateDR string

const (
	// Accepted ...
	Accepted ProvisioningStateDR = "Accepted"
	// Failed ...
	Failed ProvisioningStateDR = "Failed"
	// Succeeded ...
	Succeeded ProvisioningStateDR = "Succeeded"
)

// PossibleProvisioningStateDRValues returns an array of possible values for the ProvisioningStateDR const type.
func PossibleProvisioningStateDRValues() []ProvisioningStateDR {
	return []ProvisioningStateDR{Accepted, Failed, Succeeded}
}

// RoleDisasterRecovery enumerates the values for role disaster recovery.
type RoleDisasterRecovery string

const (
	// Primary ...
	Primary RoleDisasterRecovery = "Primary"
	// PrimaryNotReplicating ...
	PrimaryNotReplicating RoleDisasterRecovery = "PrimaryNotReplicating"
	// Secondary ...
	Secondary RoleDisasterRecovery = "Secondary"
)

// PossibleRoleDisasterRecoveryValues returns an array of possible values for the RoleDisasterRecovery const type.
func PossibleRoleDisasterRecoveryValues() []RoleDisasterRecovery {
	return []RoleDisasterRecovery{Primary, PrimaryNotReplicating, Secondary}
}

// SkuName enumerates the values for sku name.
type SkuName string

const (
	// Basic ...
	Basic SkuName = "Basic"
	// Standard ...
	Standard SkuName = "Standard"
)

// PossibleSkuNameValues returns an array of possible values for the SkuName const type.
func PossibleSkuNameValues() []SkuName {
	return []SkuName{Basic, Standard}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// SkuTierBasic ...
	SkuTierBasic SkuTier = "Basic"
	// SkuTierStandard ...
	SkuTierStandard SkuTier = "Standard"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{SkuTierBasic, SkuTierStandard}
}

// UnavailableReason enumerates the values for unavailable reason.
type UnavailableReason string

const (
	// InvalidName ...
	InvalidName UnavailableReason = "InvalidName"
	// NameInLockdown ...
	NameInLockdown UnavailableReason = "NameInLockdown"
	// NameInUse ...
	NameInUse UnavailableReason = "NameInUse"
	// None ...
	None UnavailableReason = "None"
	// SubscriptionIsDisabled ...
	SubscriptionIsDisabled UnavailableReason = "SubscriptionIsDisabled"
	// TooManyNamespaceInCurrentSubscription ...
	TooManyNamespaceInCurrentSubscription UnavailableReason = "TooManyNamespaceInCurrentSubscription"
)

// PossibleUnavailableReasonValues returns an array of possible values for the UnavailableReason const type.
func PossibleUnavailableReasonValues() []UnavailableReason {
	return []UnavailableReason{InvalidName, NameInLockdown, NameInUse, None, SubscriptionIsDisabled, TooManyNamespaceInCurrentSubscription}
}
