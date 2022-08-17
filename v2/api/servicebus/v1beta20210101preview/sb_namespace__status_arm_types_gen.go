// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

type SBNamespace_StatusARM struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Identity: Properties of BYOK Identity description
	Identity *Identity_StatusARM `json:"identity,omitempty"`

	// Location: The Geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the namespace.
	Properties *SBNamespaceProperties_StatusARM `json:"properties,omitempty"`

	// Sku: Properties of SKU
	Sku *SBSku_StatusARM `json:"sku,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

type Identity_StatusARM struct {
	// PrincipalId: ObjectId from the KeyVault
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: TenantId from the KeyVault
	TenantId *string `json:"tenantId,omitempty"`

	// Type: Type of managed service identity.
	Type *IdentityStatusType `json:"type,omitempty"`

	// UserAssignedIdentities: Properties for User Assigned Identities
	UserAssignedIdentities map[string]DictionaryValue_StatusARM `json:"userAssignedIdentities,omitempty"`
}

type SBNamespaceProperties_StatusARM struct {
	// CreatedAt: The time the namespace was created
	CreatedAt *string `json:"createdAt,omitempty"`

	// Encryption: Properties of BYOK Encryption description
	Encryption *Encryption_StatusARM `json:"encryption,omitempty"`

	// MetricId: Identifier for Azure Insights metrics
	MetricId *string `json:"metricId,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections.
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`

	// ProvisioningState: Provisioning state of the namespace.
	ProvisioningState *string `json:"provisioningState,omitempty"`

	// ServiceBusEndpoint: Endpoint you can use to perform Service Bus operations.
	ServiceBusEndpoint *string `json:"serviceBusEndpoint,omitempty"`

	// Status: Status of the namespace.
	Status *string `json:"status,omitempty"`

	// UpdatedAt: The time the namespace was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`

	// ZoneRedundant: Enabling this property creates a Premium Service Bus Namespace in regions supported availability zones.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty"`
}

type SBSku_StatusARM struct {
	// Capacity: The specified messaging units for the tier. For Premium tier, capacity are 1,2 and 4.
	Capacity *int `json:"capacity,omitempty"`

	// Name: Name of this SKU.
	Name *SBSkuStatusName `json:"name,omitempty"`

	// Tier: The billing tier of this particular SKU.
	Tier *SBSkuStatusTier `json:"tier,omitempty"`
}

type SystemData_StatusARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemDataStatusCreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The type of identity that last modified the resource.
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type DictionaryValue_StatusARM struct {
	// ClientId: Client Id of user assigned identity
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: Principal Id of user assigned identity
	PrincipalId *string `json:"principalId,omitempty"`
}

type Encryption_StatusARM struct {
	// KeySource: Enumerates the possible value of keySource for Encryption
	KeySource *EncryptionStatusKeySource `json:"keySource,omitempty"`

	// KeyVaultProperties: Properties of KeyVault
	KeyVaultProperties []KeyVaultProperties_StatusARM `json:"keyVaultProperties,omitempty"`

	// RequireInfrastructureEncryption: Enable Infrastructure Encryption (Double Encryption)
	RequireInfrastructureEncryption *bool `json:"requireInfrastructureEncryption,omitempty"`
}

type IdentityStatusType string

const (
	IdentityStatusType_None                       = IdentityStatusType("None")
	IdentityStatusType_SystemAssigned             = IdentityStatusType("SystemAssigned")
	IdentityStatusType_SystemAssignedUserAssigned = IdentityStatusType("SystemAssigned, UserAssigned")
	IdentityStatusType_UserAssigned               = IdentityStatusType("UserAssigned")
)

type PrivateEndpointConnection_Status_SubResourceEmbeddedARM struct {
	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`
}

type SBSkuStatusName string

const (
	SBSkuStatusName_Basic    = SBSkuStatusName("Basic")
	SBSkuStatusName_Premium  = SBSkuStatusName("Premium")
	SBSkuStatusName_Standard = SBSkuStatusName("Standard")
)

type SBSkuStatusTier string

const (
	SBSkuStatusTier_Basic    = SBSkuStatusTier("Basic")
	SBSkuStatusTier_Premium  = SBSkuStatusTier("Premium")
	SBSkuStatusTier_Standard = SBSkuStatusTier("Standard")
)

type SystemDataStatusCreatedByType string

const (
	SystemDataStatusCreatedByType_Application     = SystemDataStatusCreatedByType("Application")
	SystemDataStatusCreatedByType_Key             = SystemDataStatusCreatedByType("Key")
	SystemDataStatusCreatedByType_ManagedIdentity = SystemDataStatusCreatedByType("ManagedIdentity")
	SystemDataStatusCreatedByType_User            = SystemDataStatusCreatedByType("User")
)

type SystemDataStatusLastModifiedByType string

const (
	SystemDataStatusLastModifiedByType_Application     = SystemDataStatusLastModifiedByType("Application")
	SystemDataStatusLastModifiedByType_Key             = SystemDataStatusLastModifiedByType("Key")
	SystemDataStatusLastModifiedByType_ManagedIdentity = SystemDataStatusLastModifiedByType("ManagedIdentity")
	SystemDataStatusLastModifiedByType_User            = SystemDataStatusLastModifiedByType("User")
)

type KeyVaultProperties_StatusARM struct {
	Identity *UserAssignedIdentityProperties_StatusARM `json:"identity,omitempty"`

	// KeyName: Name of the Key from KeyVault
	KeyName *string `json:"keyName,omitempty"`

	// KeyVaultUri: Uri of KeyVault
	KeyVaultUri *string `json:"keyVaultUri,omitempty"`

	// KeyVersion: Version of KeyVault
	KeyVersion *string `json:"keyVersion,omitempty"`
}

type UserAssignedIdentityProperties_StatusARM struct {
	// UserAssignedIdentity: ARM ID of user Identity selected for encryption
	UserAssignedIdentity *string `json:"userAssignedIdentity,omitempty"`
}
