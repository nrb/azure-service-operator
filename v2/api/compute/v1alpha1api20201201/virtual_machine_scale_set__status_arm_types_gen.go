// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

// Deprecated version of VirtualMachineScaleSet_Status. Use v1beta20201201.VirtualMachineScaleSet_Status instead
type VirtualMachineScaleSet_StatusARM struct {
	ExtendedLocation *ExtendedLocation_StatusARM                 `json:"extendedLocation,omitempty"`
	Id               *string                                     `json:"id,omitempty"`
	Identity         *VirtualMachineScaleSetIdentity_StatusARM   `json:"identity,omitempty"`
	Location         *string                                     `json:"location,omitempty"`
	Name             *string                                     `json:"name,omitempty"`
	Plan             *Plan_StatusARM                             `json:"plan,omitempty"`
	Properties       *VirtualMachineScaleSetProperties_StatusARM `json:"properties,omitempty"`
	Sku              *Sku_StatusARM                              `json:"sku,omitempty"`
	Tags             map[string]string                           `json:"tags,omitempty"`
	Type             *string                                     `json:"type,omitempty"`
	Zones            []string                                    `json:"zones,omitempty"`
}

// Deprecated version of ExtendedLocation_Status. Use v1beta20201201.ExtendedLocation_Status instead
type ExtendedLocation_StatusARM struct {
	Name *string                      `json:"name,omitempty"`
	Type *ExtendedLocationType_Status `json:"type,omitempty"`
}

// Deprecated version of Plan_Status. Use v1beta20201201.Plan_Status instead
type Plan_StatusARM struct {
	Name          *string `json:"name,omitempty"`
	Product       *string `json:"product,omitempty"`
	PromotionCode *string `json:"promotionCode,omitempty"`
	Publisher     *string `json:"publisher,omitempty"`
}

// Deprecated version of Sku_Status. Use v1beta20201201.Sku_Status instead
type Sku_StatusARM struct {
	Capacity *int    `json:"capacity,omitempty"`
	Name     *string `json:"name,omitempty"`
	Tier     *string `json:"tier,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetIdentity_Status. Use v1beta20201201.VirtualMachineScaleSetIdentity_Status instead
type VirtualMachineScaleSetIdentity_StatusARM struct {
	PrincipalId *string                                   `json:"principalId,omitempty"`
	TenantId    *string                                   `json:"tenantId,omitempty"`
	Type        *VirtualMachineScaleSetIdentityStatusType `json:"type,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetProperties_Status. Use v1beta20201201.VirtualMachineScaleSetProperties_Status instead
type VirtualMachineScaleSetProperties_StatusARM struct {
	AdditionalCapabilities                 *AdditionalCapabilities_StatusARM          `json:"additionalCapabilities,omitempty"`
	AutomaticRepairsPolicy                 *AutomaticRepairsPolicy_StatusARM          `json:"automaticRepairsPolicy,omitempty"`
	DoNotRunExtensionsOnOverprovisionedVMs *bool                                      `json:"doNotRunExtensionsOnOverprovisionedVMs,omitempty"`
	HostGroup                              *SubResource_StatusARM                     `json:"hostGroup,omitempty"`
	OrchestrationMode                      *OrchestrationMode_Status                  `json:"orchestrationMode,omitempty"`
	Overprovision                          *bool                                      `json:"overprovision,omitempty"`
	PlatformFaultDomainCount               *int                                       `json:"platformFaultDomainCount,omitempty"`
	ProvisioningState                      *string                                    `json:"provisioningState,omitempty"`
	ProximityPlacementGroup                *SubResource_StatusARM                     `json:"proximityPlacementGroup,omitempty"`
	ScaleInPolicy                          *ScaleInPolicy_StatusARM                   `json:"scaleInPolicy,omitempty"`
	SinglePlacementGroup                   *bool                                      `json:"singlePlacementGroup,omitempty"`
	UniqueId                               *string                                    `json:"uniqueId,omitempty"`
	UpgradePolicy                          *UpgradePolicy_StatusARM                   `json:"upgradePolicy,omitempty"`
	VirtualMachineProfile                  *VirtualMachineScaleSetVMProfile_StatusARM `json:"virtualMachineProfile,omitempty"`
	ZoneBalance                            *bool                                      `json:"zoneBalance,omitempty"`
}

// Deprecated version of AdditionalCapabilities_Status. Use v1beta20201201.AdditionalCapabilities_Status instead
type AdditionalCapabilities_StatusARM struct {
	UltraSSDEnabled *bool `json:"ultraSSDEnabled,omitempty"`
}

// Deprecated version of AutomaticRepairsPolicy_Status. Use v1beta20201201.AutomaticRepairsPolicy_Status instead
type AutomaticRepairsPolicy_StatusARM struct {
	Enabled     *bool   `json:"enabled,omitempty"`
	GracePeriod *string `json:"gracePeriod,omitempty"`
}

// Deprecated version of ExtendedLocationType_Status. Use v1beta20201201.ExtendedLocationType_Status instead
type ExtendedLocationType_Status string

const ExtendedLocationType_Status_EdgeZone = ExtendedLocationType_Status("EdgeZone")

// Deprecated version of ScaleInPolicy_Status. Use v1beta20201201.ScaleInPolicy_Status instead
type ScaleInPolicy_StatusARM struct {
	Rules []ScaleInPolicyStatusRules `json:"rules,omitempty"`
}

// Deprecated version of SubResource_Status. Use v1beta20201201.SubResource_Status instead
type SubResource_StatusARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of UpgradePolicy_Status. Use v1beta20201201.UpgradePolicy_Status instead
type UpgradePolicy_StatusARM struct {
	AutomaticOSUpgradePolicy *AutomaticOSUpgradePolicy_StatusARM `json:"automaticOSUpgradePolicy,omitempty"`
	Mode                     *UpgradePolicyStatusMode            `json:"mode,omitempty"`
	RollingUpgradePolicy     *RollingUpgradePolicy_StatusARM     `json:"rollingUpgradePolicy,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetIdentityStatusType. Use
// v1beta20201201.VirtualMachineScaleSetIdentityStatusType instead
type VirtualMachineScaleSetIdentityStatusType string

const (
	VirtualMachineScaleSetIdentityStatusType_None                       = VirtualMachineScaleSetIdentityStatusType("None")
	VirtualMachineScaleSetIdentityStatusType_SystemAssigned             = VirtualMachineScaleSetIdentityStatusType("SystemAssigned")
	VirtualMachineScaleSetIdentityStatusType_SystemAssignedUserAssigned = VirtualMachineScaleSetIdentityStatusType("SystemAssigned, UserAssigned")
	VirtualMachineScaleSetIdentityStatusType_UserAssigned               = VirtualMachineScaleSetIdentityStatusType("UserAssigned")
)

// Deprecated version of VirtualMachineScaleSetVMProfile_Status. Use v1beta20201201.VirtualMachineScaleSetVMProfile_Status instead
type VirtualMachineScaleSetVMProfile_StatusARM struct {
	BillingProfile         *BillingProfile_StatusARM                         `json:"billingProfile,omitempty"`
	DiagnosticsProfile     *DiagnosticsProfile_StatusARM                     `json:"diagnosticsProfile,omitempty"`
	EvictionPolicy         *EvictionPolicy_Status                            `json:"evictionPolicy,omitempty"`
	ExtensionProfile       *VirtualMachineScaleSetExtensionProfile_StatusARM `json:"extensionProfile,omitempty"`
	LicenseType            *string                                           `json:"licenseType,omitempty"`
	NetworkProfile         *VirtualMachineScaleSetNetworkProfile_StatusARM   `json:"networkProfile,omitempty"`
	OsProfile              *VirtualMachineScaleSetOSProfile_StatusARM        `json:"osProfile,omitempty"`
	Priority               *Priority_Status                                  `json:"priority,omitempty"`
	ScheduledEventsProfile *ScheduledEventsProfile_StatusARM                 `json:"scheduledEventsProfile,omitempty"`
	SecurityProfile        *SecurityProfile_StatusARM                        `json:"securityProfile,omitempty"`
	StorageProfile         *VirtualMachineScaleSetStorageProfile_StatusARM   `json:"storageProfile,omitempty"`
}

// Deprecated version of AutomaticOSUpgradePolicy_Status. Use v1beta20201201.AutomaticOSUpgradePolicy_Status instead
type AutomaticOSUpgradePolicy_StatusARM struct {
	DisableAutomaticRollback *bool `json:"disableAutomaticRollback,omitempty"`
	EnableAutomaticOSUpgrade *bool `json:"enableAutomaticOSUpgrade,omitempty"`
}

// Deprecated version of RollingUpgradePolicy_Status. Use v1beta20201201.RollingUpgradePolicy_Status instead
type RollingUpgradePolicy_StatusARM struct {
	EnableCrossZoneUpgrade              *bool   `json:"enableCrossZoneUpgrade,omitempty"`
	MaxBatchInstancePercent             *int    `json:"maxBatchInstancePercent,omitempty"`
	MaxUnhealthyInstancePercent         *int    `json:"maxUnhealthyInstancePercent,omitempty"`
	MaxUnhealthyUpgradedInstancePercent *int    `json:"maxUnhealthyUpgradedInstancePercent,omitempty"`
	PauseTimeBetweenBatches             *string `json:"pauseTimeBetweenBatches,omitempty"`
	PrioritizeUnhealthyInstances        *bool   `json:"prioritizeUnhealthyInstances,omitempty"`
}

// Deprecated version of ScheduledEventsProfile_Status. Use v1beta20201201.ScheduledEventsProfile_Status instead
type ScheduledEventsProfile_StatusARM struct {
	TerminateNotificationProfile *TerminateNotificationProfile_StatusARM `json:"terminateNotificationProfile,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetExtensionProfile_Status. Use v1beta20201201.VirtualMachineScaleSetExtensionProfile_Status instead
type VirtualMachineScaleSetExtensionProfile_StatusARM struct {
	Extensions           []VirtualMachineScaleSetExtension_StatusARM `json:"extensions,omitempty"`
	ExtensionsTimeBudget *string                                     `json:"extensionsTimeBudget,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetNetworkProfile_Status. Use v1beta20201201.VirtualMachineScaleSetNetworkProfile_Status instead
type VirtualMachineScaleSetNetworkProfile_StatusARM struct {
	HealthProbe                    *ApiEntityReference_StatusARM                          `json:"healthProbe,omitempty"`
	NetworkInterfaceConfigurations []VirtualMachineScaleSetNetworkConfiguration_StatusARM `json:"networkInterfaceConfigurations,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetOSProfile_Status. Use v1beta20201201.VirtualMachineScaleSetOSProfile_Status instead
type VirtualMachineScaleSetOSProfile_StatusARM struct {
	AdminUsername        *string                         `json:"adminUsername,omitempty"`
	ComputerNamePrefix   *string                         `json:"computerNamePrefix,omitempty"`
	CustomData           *string                         `json:"customData,omitempty"`
	LinuxConfiguration   *LinuxConfiguration_StatusARM   `json:"linuxConfiguration,omitempty"`
	Secrets              []VaultSecretGroup_StatusARM    `json:"secrets,omitempty"`
	WindowsConfiguration *WindowsConfiguration_StatusARM `json:"windowsConfiguration,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetStorageProfile_Status. Use v1beta20201201.VirtualMachineScaleSetStorageProfile_Status instead
type VirtualMachineScaleSetStorageProfile_StatusARM struct {
	DataDisks      []VirtualMachineScaleSetDataDisk_StatusARM `json:"dataDisks,omitempty"`
	ImageReference *ImageReference_StatusARM                  `json:"imageReference,omitempty"`
	OsDisk         *VirtualMachineScaleSetOSDisk_StatusARM    `json:"osDisk,omitempty"`
}

// Deprecated version of ApiEntityReference_Status. Use v1beta20201201.ApiEntityReference_Status instead
type ApiEntityReference_StatusARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of TerminateNotificationProfile_Status. Use v1beta20201201.TerminateNotificationProfile_Status instead
type TerminateNotificationProfile_StatusARM struct {
	Enable           *bool   `json:"enable,omitempty"`
	NotBeforeTimeout *string `json:"notBeforeTimeout,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetDataDisk_Status. Use v1beta20201201.VirtualMachineScaleSetDataDisk_Status instead
type VirtualMachineScaleSetDataDisk_StatusARM struct {
	Caching                 *Caching_Status                                        `json:"caching,omitempty"`
	CreateOption            *CreateOption_Status                                   `json:"createOption,omitempty"`
	DiskIOPSReadWrite       *int                                                   `json:"diskIOPSReadWrite,omitempty"`
	DiskMBpsReadWrite       *int                                                   `json:"diskMBpsReadWrite,omitempty"`
	DiskSizeGB              *int                                                   `json:"diskSizeGB,omitempty"`
	Lun                     *int                                                   `json:"lun,omitempty"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParameters_StatusARM `json:"managedDisk,omitempty"`
	Name                    *string                                                `json:"name,omitempty"`
	WriteAcceleratorEnabled *bool                                                  `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetExtension_Status. Use v1beta20201201.VirtualMachineScaleSetExtension_Status instead
type VirtualMachineScaleSetExtension_StatusARM struct {
	Id         *string                                              `json:"id,omitempty"`
	Name       *string                                              `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetExtensionProperties_StatusARM `json:"properties,omitempty"`
	Type       *string                                              `json:"type,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetNetworkConfiguration_Status. Use v1beta20201201.VirtualMachineScaleSetNetworkConfiguration_Status instead
type VirtualMachineScaleSetNetworkConfiguration_StatusARM struct {
	Id         *string                                                         `json:"id,omitempty"`
	Name       *string                                                         `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetOSDisk_Status. Use v1beta20201201.VirtualMachineScaleSetOSDisk_Status instead
type VirtualMachineScaleSetOSDisk_StatusARM struct {
	Caching                 *Caching_Status                                        `json:"caching,omitempty"`
	CreateOption            *CreateOption_Status                                   `json:"createOption,omitempty"`
	DiffDiskSettings        *DiffDiskSettings_StatusARM                            `json:"diffDiskSettings,omitempty"`
	DiskSizeGB              *int                                                   `json:"diskSizeGB,omitempty"`
	Image                   *VirtualHardDisk_StatusARM                             `json:"image,omitempty"`
	ManagedDisk             *VirtualMachineScaleSetManagedDiskParameters_StatusARM `json:"managedDisk,omitempty"`
	Name                    *string                                                `json:"name,omitempty"`
	OsType                  *VirtualMachineScaleSetOSDiskStatusOsType              `json:"osType,omitempty"`
	VhdContainers           []string                                               `json:"vhdContainers,omitempty"`
	WriteAcceleratorEnabled *bool                                                  `json:"writeAcceleratorEnabled,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetExtensionProperties_Status. Use v1beta20201201.VirtualMachineScaleSetExtensionProperties_Status instead
type VirtualMachineScaleSetExtensionProperties_StatusARM struct {
	AutoUpgradeMinorVersion  *bool              `json:"autoUpgradeMinorVersion,omitempty"`
	EnableAutomaticUpgrade   *bool              `json:"enableAutomaticUpgrade,omitempty"`
	ForceUpdateTag           *string            `json:"forceUpdateTag,omitempty"`
	ProtectedSettings        map[string]v1.JSON `json:"protectedSettings,omitempty"`
	ProvisionAfterExtensions []string           `json:"provisionAfterExtensions,omitempty"`
	ProvisioningState        *string            `json:"provisioningState,omitempty"`
	Publisher                *string            `json:"publisher,omitempty"`
	Settings                 map[string]v1.JSON `json:"settings,omitempty"`
	Type                     *string            `json:"type,omitempty"`
	TypeHandlerVersion       *string            `json:"typeHandlerVersion,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetManagedDiskParameters_Status. Use v1beta20201201.VirtualMachineScaleSetManagedDiskParameters_Status instead
type VirtualMachineScaleSetManagedDiskParameters_StatusARM struct {
	DiskEncryptionSet  *SubResource_StatusARM     `json:"diskEncryptionSet,omitempty"`
	StorageAccountType *StorageAccountType_Status `json:"storageAccountType,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetNetworkConfigurationProperties_Status. Use v1beta20201201.VirtualMachineScaleSetNetworkConfigurationProperties_Status instead
type VirtualMachineScaleSetNetworkConfigurationProperties_StatusARM struct {
	DnsSettings                 *VirtualMachineScaleSetNetworkConfigurationDnsSettings_StatusARM `json:"dnsSettings,omitempty"`
	EnableAcceleratedNetworking *bool                                                            `json:"enableAcceleratedNetworking,omitempty"`
	EnableFpga                  *bool                                                            `json:"enableFpga,omitempty"`
	EnableIPForwarding          *bool                                                            `json:"enableIPForwarding,omitempty"`
	IpConfigurations            []VirtualMachineScaleSetIPConfiguration_StatusARM                `json:"ipConfigurations,omitempty"`
	NetworkSecurityGroup        *SubResource_StatusARM                                           `json:"networkSecurityGroup,omitempty"`
	Primary                     *bool                                                            `json:"primary,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetIPConfiguration_Status. Use v1beta20201201.VirtualMachineScaleSetIPConfiguration_Status instead
type VirtualMachineScaleSetIPConfiguration_StatusARM struct {
	Id         *string                                                    `json:"id,omitempty"`
	Name       *string                                                    `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetIPConfigurationProperties_StatusARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetNetworkConfigurationDnsSettings_Status. Use v1beta20201201.VirtualMachineScaleSetNetworkConfigurationDnsSettings_Status instead
type VirtualMachineScaleSetNetworkConfigurationDnsSettings_StatusARM struct {
	DnsServers []string `json:"dnsServers,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetIPConfigurationProperties_Status. Use v1beta20201201.VirtualMachineScaleSetIPConfigurationProperties_Status instead
type VirtualMachineScaleSetIPConfigurationProperties_StatusARM struct {
	ApplicationGatewayBackendAddressPools []SubResource_StatusARM                                                       `json:"applicationGatewayBackendAddressPools,omitempty"`
	ApplicationSecurityGroups             []SubResource_StatusARM                                                       `json:"applicationSecurityGroups,omitempty"`
	LoadBalancerBackendAddressPools       []SubResource_StatusARM                                                       `json:"loadBalancerBackendAddressPools,omitempty"`
	LoadBalancerInboundNatPools           []SubResource_StatusARM                                                       `json:"loadBalancerInboundNatPools,omitempty"`
	Primary                               *bool                                                                         `json:"primary,omitempty"`
	PrivateIPAddressVersion               *VirtualMachineScaleSetIPConfigurationPropertiesStatusPrivateIPAddressVersion `json:"privateIPAddressVersion,omitempty"`
	PublicIPAddressConfiguration          *VirtualMachineScaleSetPublicIPAddressConfiguration_StatusARM                 `json:"publicIPAddressConfiguration,omitempty"`
	Subnet                                *ApiEntityReference_StatusARM                                                 `json:"subnet,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetPublicIPAddressConfiguration_Status. Use v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfiguration_Status instead
type VirtualMachineScaleSetPublicIPAddressConfiguration_StatusARM struct {
	Name       *string                                                                 `json:"name,omitempty"`
	Properties *VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM `json:"properties,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetPublicIPAddressConfigurationProperties_Status. Use v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationProperties_Status instead
type VirtualMachineScaleSetPublicIPAddressConfigurationProperties_StatusARM struct {
	DnsSettings            *VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_StatusARM                  `json:"dnsSettings,omitempty"`
	IdleTimeoutInMinutes   *int                                                                                      `json:"idleTimeoutInMinutes,omitempty"`
	IpTags                 []VirtualMachineScaleSetIpTag_StatusARM                                                   `json:"ipTags,omitempty"`
	PublicIPAddressVersion *VirtualMachineScaleSetPublicIPAddressConfigurationPropertiesStatusPublicIPAddressVersion `json:"publicIPAddressVersion,omitempty"`
	PublicIPPrefix         *SubResource_StatusARM                                                                    `json:"publicIPPrefix,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetIpTag_Status. Use v1beta20201201.VirtualMachineScaleSetIpTag_Status instead
type VirtualMachineScaleSetIpTag_StatusARM struct {
	IpTagType *string `json:"ipTagType,omitempty"`
	Tag       *string `json:"tag,omitempty"`
}

// Deprecated version of VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_Status. Use v1beta20201201.VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_Status instead
type VirtualMachineScaleSetPublicIPAddressConfigurationDnsSettings_StatusARM struct {
	DomainNameLabel *string `json:"domainNameLabel,omitempty"`
}
