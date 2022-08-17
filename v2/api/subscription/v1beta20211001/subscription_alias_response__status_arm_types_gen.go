// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

type SubscriptionAliasResponse_StatusARM struct {
	// Id: Fully qualified ID for the alias resource.
	Id *string `json:"id,omitempty"`

	// Name: Alias ID.
	Name *string `json:"name,omitempty"`

	// Properties: Subscription Alias response properties.
	Properties *SubscriptionAliasResponseProperties_StatusARM `json:"properties,omitempty"`
	SystemData *SystemData_StatusARM                          `json:"systemData,omitempty"`

	// Type: Resource type, Microsoft.Subscription/aliases.
	Type *string `json:"type,omitempty"`
}

type SubscriptionAliasResponseProperties_StatusARM struct {
	AcceptOwnershipState *AcceptOwnershipState_Status `json:"acceptOwnershipState,omitempty"`

	// AcceptOwnershipUrl: Url to accept ownership of the subscription.
	AcceptOwnershipUrl *string `json:"acceptOwnershipUrl,omitempty"`
	BillingScope       *string `json:"billingScope,omitempty"`

	// CreatedTime: Created Time
	CreatedTime *string `json:"createdTime,omitempty"`

	// DisplayName: The display name of the subscription.
	DisplayName *string `json:"displayName,omitempty"`

	// ManagementGroupId: The Management Group Id.
	ManagementGroupId *string `json:"managementGroupId,omitempty"`

	// ProvisioningState: The provisioning state of the resource.
	ProvisioningState *SubscriptionAliasResponsePropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	// ResellerId: Reseller Id
	ResellerId *string `json:"resellerId,omitempty"`

	// SubscriptionId: Newly created subscription Id.
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	// SubscriptionOwnerId: Owner Id of the subscription
	SubscriptionOwnerId *string `json:"subscriptionOwnerId,omitempty"`

	// Tags: Tags for the subscription
	Tags     map[string]string `json:"tags,omitempty"`
	Workload *Workload_Status  `json:"workload,omitempty"`
}

type SystemData_StatusARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemDataStatusCreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type AcceptOwnershipState_Status string

const (
	AcceptOwnershipState_Status_Completed = AcceptOwnershipState_Status("Completed")
	AcceptOwnershipState_Status_Expired   = AcceptOwnershipState_Status("Expired")
	AcceptOwnershipState_Status_Pending   = AcceptOwnershipState_Status("Pending")
)

type SubscriptionAliasResponsePropertiesStatusProvisioningState string

const (
	SubscriptionAliasResponsePropertiesStatusProvisioningState_Accepted  = SubscriptionAliasResponsePropertiesStatusProvisioningState("Accepted")
	SubscriptionAliasResponsePropertiesStatusProvisioningState_Failed    = SubscriptionAliasResponsePropertiesStatusProvisioningState("Failed")
	SubscriptionAliasResponsePropertiesStatusProvisioningState_Succeeded = SubscriptionAliasResponsePropertiesStatusProvisioningState("Succeeded")
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

type Workload_Status string

const (
	Workload_Status_DevTest    = Workload_Status("DevTest")
	Workload_Status_Production = Workload_Status("Production")
)
