// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

type DomainTopic_StatusARM struct {
	// Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	// Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the Domain Topic.
	Properties *DomainTopicProperties_StatusARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to Domain Topic resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	// Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type DomainTopicProperties_StatusARM struct {
	// ProvisioningState: Provisioning state of the domain topic.
	ProvisioningState *DomainTopicPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`
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

type DomainTopicPropertiesStatusProvisioningState string

const (
	DomainTopicPropertiesStatusProvisioningState_Canceled  = DomainTopicPropertiesStatusProvisioningState("Canceled")
	DomainTopicPropertiesStatusProvisioningState_Creating  = DomainTopicPropertiesStatusProvisioningState("Creating")
	DomainTopicPropertiesStatusProvisioningState_Deleting  = DomainTopicPropertiesStatusProvisioningState("Deleting")
	DomainTopicPropertiesStatusProvisioningState_Failed    = DomainTopicPropertiesStatusProvisioningState("Failed")
	DomainTopicPropertiesStatusProvisioningState_Succeeded = DomainTopicPropertiesStatusProvisioningState("Succeeded")
	DomainTopicPropertiesStatusProvisioningState_Updating  = DomainTopicPropertiesStatusProvisioningState("Updating")
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
