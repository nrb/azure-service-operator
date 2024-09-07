// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240302

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Snapshot_Spec_ARM struct {
	// ExtendedLocation: The extended location where the snapshot will be created. Extended location cannot be changed.
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Snapshot resource properties.
	Properties *SnapshotProperties_ARM `json:"properties,omitempty"`

	// Sku: The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for
	// incremental  snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
	Sku *SnapshotSku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Snapshot_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-03-02"
func (snapshot Snapshot_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (snapshot *Snapshot_Spec_ARM) GetName() string {
	return snapshot.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/snapshots"
func (snapshot *Snapshot_Spec_ARM) GetType() string {
	return "Microsoft.Compute/snapshots"
}

// Snapshot resource properties.
type SnapshotProperties_ARM struct {
	// CompletionPercent: Percentage complete for the background copy when a resource is created via the CopyStart operation.
	CompletionPercent *float64 `json:"completionPercent,omitempty"`

	// CopyCompletionError: Indicates the error details if the background copy of a resource created via the CopyStart
	// operation fails.
	CopyCompletionError *CopyCompletionError_ARM `json:"copyCompletionError,omitempty"`

	// CreationData: Disk source information. CreationData information cannot be changed after the disk has been created.
	CreationData *CreationData_ARM `json:"creationData,omitempty"`

	// DataAccessAuthMode: Additional authentication requirements when exporting or uploading to a disk or snapshot.
	DataAccessAuthMode *DataAccessAuthMode `json:"dataAccessAuthMode,omitempty"`
	DiskAccessId       *string             `json:"diskAccessId,omitempty"`

	// DiskSizeGB: If creationData.createOption is Empty, this field is mandatory and it indicates the size of the disk to
	// create. If this field is present for updates or creation with other options, it indicates a resize. Resizes are only
	// allowed if the disk is not attached to a running VM, and can only increase the disk's size.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// DiskState: The state of the snapshot.
	DiskState *DiskState `json:"diskState,omitempty"`

	// Encryption: Encryption property can be used to encrypt data at rest with customer managed keys or platform managed keys.
	Encryption *Encryption_ARM `json:"encryption,omitempty"`

	// EncryptionSettingsCollection: Encryption settings collection used be Azure Disk Encryption, can contain multiple
	// encryption settings per disk or snapshot.
	EncryptionSettingsCollection *EncryptionSettingsCollection_ARM `json:"encryptionSettingsCollection,omitempty"`

	// HyperVGeneration: The hypervisor generation of the Virtual Machine. Applicable to OS disks only.
	HyperVGeneration *SnapshotProperties_HyperVGeneration `json:"hyperVGeneration,omitempty"`

	// Incremental: Whether a snapshot is incremental. Incremental snapshots on the same disk occupy less space than full
	// snapshots and can be diffed.
	Incremental *bool `json:"incremental,omitempty"`

	// NetworkAccessPolicy: Policy for accessing the disk via network.
	NetworkAccessPolicy *NetworkAccessPolicy `json:"networkAccessPolicy,omitempty"`

	// OsType: The Operating System type.
	OsType *SnapshotProperties_OsType `json:"osType,omitempty"`

	// PublicNetworkAccess: Policy for controlling export on the disk.
	PublicNetworkAccess *PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`

	// PurchasePlan: Purchase plan information for the image from which the source disk for the snapshot was originally created.
	PurchasePlan *PurchasePlan_ARM `json:"purchasePlan,omitempty"`

	// SecurityProfile: Contains the security related information for the resource.
	SecurityProfile *DiskSecurityProfile_ARM `json:"securityProfile,omitempty"`

	// SupportedCapabilities: List of supported capabilities for the image from which the source disk from the snapshot was
	// originally created.
	SupportedCapabilities *SupportedCapabilities_ARM `json:"supportedCapabilities,omitempty"`

	// SupportsHibernation: Indicates the OS on a snapshot supports hibernation.
	SupportsHibernation *bool `json:"supportsHibernation,omitempty"`
}

// The snapshots sku name. Can be Standard_LRS, Premium_LRS, or Standard_ZRS. This is an optional parameter for incremental
// snapshot and the default behavior is the SKU will be set to the same sku as the previous snapshot
type SnapshotSku_ARM struct {
	// Name: The sku name.
	Name *SnapshotSku_Name `json:"name,omitempty"`
}

// Indicates the error details if the background copy of a resource created via the CopyStart operation fails.
type CopyCompletionError_ARM struct {
	// ErrorCode: Indicates the error code if the background copy of a resource created via the CopyStart operation fails.
	ErrorCode *CopyCompletionError_ErrorCode `json:"errorCode,omitempty"`

	// ErrorMessage: Indicates the error message if the background copy of a resource created via the CopyStart operation fails.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// +kubebuilder:validation:Enum={"Premium_LRS","Standard_LRS","Standard_ZRS"}
type SnapshotSku_Name string

const (
	SnapshotSku_Name_Premium_LRS  = SnapshotSku_Name("Premium_LRS")
	SnapshotSku_Name_Standard_LRS = SnapshotSku_Name("Standard_LRS")
	SnapshotSku_Name_Standard_ZRS = SnapshotSku_Name("Standard_ZRS")
)

// Mapping from string to SnapshotSku_Name
var snapshotSku_Name_Values = map[string]SnapshotSku_Name{
	"premium_lrs":  SnapshotSku_Name_Premium_LRS,
	"standard_lrs": SnapshotSku_Name_Standard_LRS,
	"standard_zrs": SnapshotSku_Name_Standard_ZRS,
}