// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211001

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Deprecated version of SignalR_Spec. Use v1beta20211001.SignalR_Spec instead
type SignalR_SpecARM struct {
	Identity   *ManagedIdentityARM   `json:"identity,omitempty"`
	Kind       *SignalRSpecKind      `json:"kind,omitempty"`
	Location   *string               `json:"location,omitempty"`
	Name       string                `json:"name,omitempty"`
	Properties *SignalRPropertiesARM `json:"properties,omitempty"`
	Sku        *ResourceSkuARM       `json:"sku,omitempty"`
	Tags       map[string]string     `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &SignalR_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (signalR SignalR_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (signalR *SignalR_SpecARM) GetName() string {
	return signalR.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.SignalRService/signalR"
func (signalR *SignalR_SpecARM) GetType() string {
	return "Microsoft.SignalRService/signalR"
}

// Deprecated version of ManagedIdentity. Use v1beta20211001.ManagedIdentity instead
type ManagedIdentityARM struct {
	Type                   *ManagedIdentityType `json:"type,omitempty"`
	UserAssignedIdentities map[string]v1.JSON   `json:"userAssignedIdentities,omitempty"`
}

// Deprecated version of ResourceSku. Use v1beta20211001.ResourceSku instead
type ResourceSkuARM struct {
	Capacity *int             `json:"capacity,omitempty"`
	Name     *string          `json:"name,omitempty"`
	Tier     *ResourceSkuTier `json:"tier,omitempty"`
}

// Deprecated version of SignalRProperties. Use v1beta20211001.SignalRProperties instead
type SignalRPropertiesARM struct {
	Cors                     *SignalRCorsSettingsARM        `json:"cors,omitempty"`
	DisableAadAuth           *bool                          `json:"disableAadAuth,omitempty"`
	DisableLocalAuth         *bool                          `json:"disableLocalAuth,omitempty"`
	Features                 []SignalRFeatureARM            `json:"features,omitempty"`
	NetworkACLs              *SignalRNetworkACLsARM         `json:"networkACLs,omitempty"`
	PublicNetworkAccess      *string                        `json:"publicNetworkAccess,omitempty"`
	ResourceLogConfiguration *ResourceLogConfigurationARM   `json:"resourceLogConfiguration,omitempty"`
	Tls                      *SignalRTlsSettingsARM         `json:"tls,omitempty"`
	Upstream                 *ServerlessUpstreamSettingsARM `json:"upstream,omitempty"`
}

// Deprecated version of SignalRSpecKind. Use v1beta20211001.SignalRSpecKind instead
// +kubebuilder:validation:Enum={"RawWebSockets","SignalR"}
type SignalRSpecKind string

const (
	SignalRSpecKind_RawWebSockets = SignalRSpecKind("RawWebSockets")
	SignalRSpecKind_SignalR       = SignalRSpecKind("SignalR")
)

// Deprecated version of ManagedIdentityType. Use v1beta20211001.ManagedIdentityType instead
// +kubebuilder:validation:Enum={"None","SystemAssigned","UserAssigned"}
type ManagedIdentityType string

const (
	ManagedIdentityType_None           = ManagedIdentityType("None")
	ManagedIdentityType_SystemAssigned = ManagedIdentityType("SystemAssigned")
	ManagedIdentityType_UserAssigned   = ManagedIdentityType("UserAssigned")
)

// Deprecated version of ResourceLogConfiguration. Use v1beta20211001.ResourceLogConfiguration instead
type ResourceLogConfigurationARM struct {
	Categories []ResourceLogCategoryARM `json:"categories,omitempty"`
}

// Deprecated version of ResourceSkuTier. Use v1beta20211001.ResourceSkuTier instead
// +kubebuilder:validation:Enum={"Basic","Free","Premium","Standard"}
type ResourceSkuTier string

const (
	ResourceSkuTier_Basic    = ResourceSkuTier("Basic")
	ResourceSkuTier_Free     = ResourceSkuTier("Free")
	ResourceSkuTier_Premium  = ResourceSkuTier("Premium")
	ResourceSkuTier_Standard = ResourceSkuTier("Standard")
)

// Deprecated version of ServerlessUpstreamSettings. Use v1beta20211001.ServerlessUpstreamSettings instead
type ServerlessUpstreamSettingsARM struct {
	Templates []UpstreamTemplateARM `json:"templates,omitempty"`
}

// Deprecated version of SignalRCorsSettings. Use v1beta20211001.SignalRCorsSettings instead
type SignalRCorsSettingsARM struct {
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
}

// Deprecated version of SignalRFeature. Use v1beta20211001.SignalRFeature instead
type SignalRFeatureARM struct {
	Flag       *SignalRFeatureFlag `json:"flag,omitempty"`
	Properties map[string]string   `json:"properties,omitempty"`
	Value      *string             `json:"value,omitempty"`
}

// Deprecated version of SignalRNetworkACLs. Use v1beta20211001.SignalRNetworkACLs instead
type SignalRNetworkACLsARM struct {
	DefaultAction    *SignalRNetworkACLsDefaultAction `json:"defaultAction,omitempty"`
	PrivateEndpoints []PrivateEndpointACLARM          `json:"privateEndpoints,omitempty"`
	PublicNetwork    *NetworkACLARM                   `json:"publicNetwork,omitempty"`
}

// Deprecated version of SignalRTlsSettings. Use v1beta20211001.SignalRTlsSettings instead
type SignalRTlsSettingsARM struct {
	ClientCertEnabled *bool `json:"clientCertEnabled,omitempty"`
}

// Deprecated version of NetworkACL. Use v1beta20211001.NetworkACL instead
type NetworkACLARM struct {
	Allow []NetworkACLAllow `json:"allow,omitempty"`
	Deny  []NetworkACLDeny  `json:"deny,omitempty"`
}

// Deprecated version of PrivateEndpointACL. Use v1beta20211001.PrivateEndpointACL instead
type PrivateEndpointACLARM struct {
	Allow []PrivateEndpointACLAllow `json:"allow,omitempty"`
	Deny  []PrivateEndpointACLDeny  `json:"deny,omitempty"`
	Name  *string                   `json:"name,omitempty"`
}

// Deprecated version of ResourceLogCategory. Use v1beta20211001.ResourceLogCategory instead
type ResourceLogCategoryARM struct {
	Enabled *string `json:"enabled,omitempty"`
	Name    *string `json:"name,omitempty"`
}

// Deprecated version of UpstreamTemplate. Use v1beta20211001.UpstreamTemplate instead
type UpstreamTemplateARM struct {
	Auth            *UpstreamAuthSettingsARM `json:"auth,omitempty"`
	CategoryPattern *string                  `json:"categoryPattern,omitempty"`
	EventPattern    *string                  `json:"eventPattern,omitempty"`
	HubPattern      *string                  `json:"hubPattern,omitempty"`
	UrlTemplate     *string                  `json:"urlTemplate,omitempty"`
}

// Deprecated version of UpstreamAuthSettings. Use v1beta20211001.UpstreamAuthSettings instead
type UpstreamAuthSettingsARM struct {
	ManagedIdentity *ManagedIdentitySettingsARM `json:"managedIdentity,omitempty"`
	Type            *UpstreamAuthSettingsType   `json:"type,omitempty"`
}

// Deprecated version of ManagedIdentitySettings. Use v1beta20211001.ManagedIdentitySettings instead
type ManagedIdentitySettingsARM struct {
	Resource *string `json:"resource,omitempty"`
}
