// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=signalrservice.azure.com,resources=signalrs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=signalrservice.azure.com,resources={signalrs/status,signalrs/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20211001.SignalR
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/resourceDefinitions/signalR
type SignalR struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SignalR_Spec           `json:"spec,omitempty"`
	Status            SignalRResource_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SignalR{}

// GetConditions returns the conditions of the resource
func (signalR *SignalR) GetConditions() conditions.Conditions {
	return signalR.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (signalR *SignalR) SetConditions(conditions conditions.Conditions) {
	signalR.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SignalR{}

// AzureName returns the Azure name of the resource
func (signalR *SignalR) AzureName() string {
	return signalR.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (signalR SignalR) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (signalR *SignalR) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (signalR *SignalR) GetSpec() genruntime.ConvertibleSpec {
	return &signalR.Spec
}

// GetStatus returns the status of this resource
func (signalR *SignalR) GetStatus() genruntime.ConvertibleStatus {
	return &signalR.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.SignalRService/signalR"
func (signalR *SignalR) GetType() string {
	return "Microsoft.SignalRService/signalR"
}

// NewEmptyStatus returns a new empty (blank) status
func (signalR *SignalR) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SignalRResource_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (signalR *SignalR) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(signalR.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  signalR.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (signalR *SignalR) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SignalRResource_Status); ok {
		signalR.Status = *st
		return nil
	}

	// Convert status to required version
	var st SignalRResource_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	signalR.Status = st
	return nil
}

// Hub marks that this SignalR is the hub type for conversion
func (signalR *SignalR) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (signalR *SignalR) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: signalR.Spec.OriginalVersion,
		Kind:    "SignalR",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20211001.SignalR
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/resourceDefinitions/signalR
type SignalRList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SignalR `json:"items"`
}

// Storage version of v1beta20211001.APIVersion
// +kubebuilder:validation:Enum={"2021-10-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2021-10-01")

// Storage version of v1beta20211001.SignalRResource_Status
type SignalRResource_Status struct {
	Conditions                 []conditions.Condition                                         `json:"conditions,omitempty"`
	Cors                       *SignalRCorsSettings_Status                                    `json:"cors,omitempty"`
	DisableAadAuth             *bool                                                          `json:"disableAadAuth,omitempty"`
	DisableLocalAuth           *bool                                                          `json:"disableLocalAuth,omitempty"`
	ExternalIP                 *string                                                        `json:"externalIP,omitempty"`
	Features                   []SignalRFeature_Status                                        `json:"features,omitempty"`
	HostName                   *string                                                        `json:"hostName,omitempty"`
	HostNamePrefix             *string                                                        `json:"hostNamePrefix,omitempty"`
	Id                         *string                                                        `json:"id,omitempty"`
	Identity                   *ManagedIdentity_Status                                        `json:"identity,omitempty"`
	Kind                       *string                                                        `json:"kind,omitempty"`
	Location                   *string                                                        `json:"location,omitempty"`
	Name                       *string                                                        `json:"name,omitempty"`
	NetworkACLs                *SignalRNetworkACLs_Status                                     `json:"networkACLs,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded `json:"privateEndpointConnections,omitempty"`
	PropertyBag                genruntime.PropertyBag                                         `json:"$propertyBag,omitempty"`
	ProvisioningState          *string                                                        `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *string                                                        `json:"publicNetworkAccess,omitempty"`
	PublicPort                 *int                                                           `json:"publicPort,omitempty"`
	ResourceLogConfiguration   *ResourceLogConfiguration_Status                               `json:"resourceLogConfiguration,omitempty"`
	ServerPort                 *int                                                           `json:"serverPort,omitempty"`
	SharedPrivateLinkResources []SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded `json:"sharedPrivateLinkResources,omitempty"`
	Sku                        *ResourceSku_Status                                            `json:"sku,omitempty"`
	SystemData                 *SystemData_Status                                             `json:"systemData,omitempty"`
	Tags                       map[string]string                                              `json:"tags,omitempty"`
	Tls                        *SignalRTlsSettings_Status                                     `json:"tls,omitempty"`
	Type                       *string                                                        `json:"type,omitempty"`
	Upstream                   *ServerlessUpstreamSettings_Status                             `json:"upstream,omitempty"`
	Version                    *string                                                        `json:"version,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SignalRResource_Status{}

// ConvertStatusFrom populates our SignalRResource_Status from the provided source
func (resource *SignalRResource_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == resource {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(resource)
}

// ConvertStatusTo populates the provided destination from our SignalRResource_Status
func (resource *SignalRResource_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == resource {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(resource)
}

// Storage version of v1beta20211001.SignalR_Spec
type SignalR_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string               `json:"azureName,omitempty"`
	Cors             *SignalRCorsSettings `json:"cors,omitempty"`
	DisableAadAuth   *bool                `json:"disableAadAuth,omitempty"`
	DisableLocalAuth *bool                `json:"disableLocalAuth,omitempty"`
	Features         []SignalRFeature     `json:"features,omitempty"`
	Identity         *ManagedIdentity     `json:"identity,omitempty"`
	Kind             *string              `json:"kind,omitempty"`
	Location         *string              `json:"location,omitempty"`
	NetworkACLs      *SignalRNetworkACLs  `json:"networkACLs,omitempty"`
	OriginalVersion  string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                    *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag              genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	PublicNetworkAccess      *string                            `json:"publicNetworkAccess,omitempty"`
	ResourceLogConfiguration *ResourceLogConfiguration          `json:"resourceLogConfiguration,omitempty"`
	Sku                      *ResourceSku                       `json:"sku,omitempty"`
	Tags                     map[string]string                  `json:"tags,omitempty"`
	Tls                      *SignalRTlsSettings                `json:"tls,omitempty"`
	Upstream                 *ServerlessUpstreamSettings        `json:"upstream,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SignalR_Spec{}

// ConvertSpecFrom populates our SignalR_Spec from the provided source
func (signalR *SignalR_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == signalR {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(signalR)
}

// ConvertSpecTo populates the provided destination from our SignalR_Spec
func (signalR *SignalR_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == signalR {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(signalR)
}

// Storage version of v1beta20211001.ManagedIdentity
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentity
type ManagedIdentity struct {
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type                   *string                `json:"type,omitempty"`
	UserAssignedIdentities map[string]v1.JSON     `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20211001.ManagedIdentity_Status
type ManagedIdentity_Status struct {
	PrincipalId            *string                                        `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                         `json:"$propertyBag,omitempty"`
	TenantId               *string                                        `json:"tenantId,omitempty"`
	Type                   *string                                        `json:"type,omitempty"`
	UserAssignedIdentities map[string]UserAssignedIdentityProperty_Status `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1beta20211001.PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded
type PrivateEndpointConnection_Status_SignalR_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_Status     `json:"systemData,omitempty"`
}

// Storage version of v1beta20211001.ResourceLogConfiguration
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogConfiguration
type ResourceLogConfiguration struct {
	Categories  []ResourceLogCategory  `json:"categories,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ResourceLogConfiguration_Status
type ResourceLogConfiguration_Status struct {
	Categories  []ResourceLogCategory_Status `json:"categories,omitempty"`
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ResourceSku
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceSku
type ResourceSku struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20211001.ResourceSku_Status
type ResourceSku_Status struct {
	Capacity    *int                   `json:"capacity,omitempty"`
	Family      *string                `json:"family,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Size        *string                `json:"size,omitempty"`
	Tier        *string                `json:"tier,omitempty"`
}

// Storage version of v1beta20211001.ServerlessUpstreamSettings
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ServerlessUpstreamSettings
type ServerlessUpstreamSettings struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Templates   []UpstreamTemplate     `json:"templates,omitempty"`
}

// Storage version of v1beta20211001.ServerlessUpstreamSettings_Status
type ServerlessUpstreamSettings_Status struct {
	PropertyBag genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Templates   []UpstreamTemplate_Status `json:"templates,omitempty"`
}

// Storage version of v1beta20211001.SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded
type SharedPrivateLinkResource_Status_SignalR_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_Status     `json:"systemData,omitempty"`
}

// Storage version of v1beta20211001.SignalRCorsSettings
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRCorsSettings
type SignalRCorsSettings struct {
	AllowedOrigins []string               `json:"allowedOrigins,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.SignalRCorsSettings_Status
type SignalRCorsSettings_Status struct {
	AllowedOrigins []string               `json:"allowedOrigins,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.SignalRFeature
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRFeature
type SignalRFeature struct {
	Flag        *string                `json:"flag,omitempty"`
	Properties  map[string]string      `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20211001.SignalRFeature_Status
type SignalRFeature_Status struct {
	Flag        *string                `json:"flag,omitempty"`
	Properties  map[string]string      `json:"properties,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

// Storage version of v1beta20211001.SignalRNetworkACLs
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRNetworkACLs
type SignalRNetworkACLs struct {
	DefaultAction    *string                `json:"defaultAction,omitempty"`
	PrivateEndpoints []PrivateEndpointACL   `json:"privateEndpoints,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicNetwork    *NetworkACL            `json:"publicNetwork,omitempty"`
}

// Storage version of v1beta20211001.SignalRNetworkACLs_Status
type SignalRNetworkACLs_Status struct {
	DefaultAction    *string                     `json:"defaultAction,omitempty"`
	PrivateEndpoints []PrivateEndpointACL_Status `json:"privateEndpoints,omitempty"`
	PropertyBag      genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	PublicNetwork    *NetworkACL_Status          `json:"publicNetwork,omitempty"`
}

// Storage version of v1beta20211001.SignalRTlsSettings
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/SignalRTlsSettings
type SignalRTlsSettings struct {
	ClientCertEnabled *bool                  `json:"clientCertEnabled,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.SignalRTlsSettings_Status
type SignalRTlsSettings_Status struct {
	ClientCertEnabled *bool                  `json:"clientCertEnabled,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.SystemData_Status
type SystemData_Status struct {
	CreatedAt          *string                `json:"createdAt,omitempty"`
	CreatedBy          *string                `json:"createdBy,omitempty"`
	CreatedByType      *string                `json:"createdByType,omitempty"`
	LastModifiedAt     *string                `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *string                `json:"lastModifiedByType,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.NetworkACL
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/NetworkACL
type NetworkACL struct {
	Allow       []string               `json:"allow,omitempty"`
	Deny        []string               `json:"deny,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.NetworkACL_Status
type NetworkACL_Status struct {
	Allow       []string               `json:"allow,omitempty"`
	Deny        []string               `json:"deny,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.PrivateEndpointACL
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/PrivateEndpointACL
type PrivateEndpointACL struct {
	Allow       []string               `json:"allow,omitempty"`
	Deny        []string               `json:"deny,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.PrivateEndpointACL_Status
type PrivateEndpointACL_Status struct {
	Allow       []string               `json:"allow,omitempty"`
	Deny        []string               `json:"deny,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ResourceLogCategory
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ResourceLogCategory
type ResourceLogCategory struct {
	Enabled     *string                `json:"enabled,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.ResourceLogCategory_Status
type ResourceLogCategory_Status struct {
	Enabled     *string                `json:"enabled,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.UpstreamTemplate
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamTemplate
type UpstreamTemplate struct {
	Auth            *UpstreamAuthSettings  `json:"auth,omitempty"`
	CategoryPattern *string                `json:"categoryPattern,omitempty"`
	EventPattern    *string                `json:"eventPattern,omitempty"`
	HubPattern      *string                `json:"hubPattern,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	UrlTemplate     *string                `json:"urlTemplate,omitempty"`
}

// Storage version of v1beta20211001.UpstreamTemplate_Status
type UpstreamTemplate_Status struct {
	Auth            *UpstreamAuthSettings_Status `json:"auth,omitempty"`
	CategoryPattern *string                      `json:"categoryPattern,omitempty"`
	EventPattern    *string                      `json:"eventPattern,omitempty"`
	HubPattern      *string                      `json:"hubPattern,omitempty"`
	PropertyBag     genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	UrlTemplate     *string                      `json:"urlTemplate,omitempty"`
}

// Storage version of v1beta20211001.UserAssignedIdentityProperty_Status
type UserAssignedIdentityProperty_Status struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20211001.UpstreamAuthSettings
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/UpstreamAuthSettings
type UpstreamAuthSettings struct {
	ManagedIdentity *ManagedIdentitySettings `json:"managedIdentity,omitempty"`
	PropertyBag     genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
	Type            *string                  `json:"type,omitempty"`
}

// Storage version of v1beta20211001.UpstreamAuthSettings_Status
type UpstreamAuthSettings_Status struct {
	ManagedIdentity *ManagedIdentitySettings_Status `json:"managedIdentity,omitempty"`
	PropertyBag     genruntime.PropertyBag          `json:"$propertyBag,omitempty"`
	Type            *string                         `json:"type,omitempty"`
}

// Storage version of v1beta20211001.ManagedIdentitySettings
// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.SignalRService.json#/definitions/ManagedIdentitySettings
type ManagedIdentitySettings struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Resource    *string                `json:"resource,omitempty"`
}

// Storage version of v1beta20211001.ManagedIdentitySettings_Status
type ManagedIdentitySettings_Status struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Resource    *string                `json:"resource,omitempty"`
}

func init() {
	SchemeBuilder.Register(&SignalR{}, &SignalRList{})
}
