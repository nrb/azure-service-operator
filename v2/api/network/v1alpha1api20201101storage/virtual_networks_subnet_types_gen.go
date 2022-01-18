// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=virtualnetworkssubnets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={virtualnetworkssubnets/status,virtualnetworkssubnets/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.VirtualNetworksSubnet
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks_subnets
type VirtualNetworksSubnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworksSubnets_Spec                             `json:"spec,omitempty"`
	Status            Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetworksSubnet{}

// GetConditions returns the conditions of the resource
func (subnet *VirtualNetworksSubnet) GetConditions() conditions.Conditions {
	return subnet.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (subnet *VirtualNetworksSubnet) SetConditions(conditions conditions.Conditions) {
	subnet.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &VirtualNetworksSubnet{}

// AzureName returns the Azure name of the resource
func (subnet *VirtualNetworksSubnet) AzureName() string {
	return subnet.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (subnet VirtualNetworksSubnet) GetAPIVersion() string {
	return "2020-11-01"
}

// GetResourceKind returns the kind of the resource
func (subnet *VirtualNetworksSubnet) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (subnet *VirtualNetworksSubnet) GetSpec() genruntime.ConvertibleSpec {
	return &subnet.Spec
}

// GetStatus returns the status of this resource
func (subnet *VirtualNetworksSubnet) GetStatus() genruntime.ConvertibleStatus {
	return &subnet.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/subnets"
func (subnet *VirtualNetworksSubnet) GetType() string {
	return "Microsoft.Network/virtualNetworks/subnets"
}

// NewEmptyStatus returns a new empty (blank) status
func (subnet *VirtualNetworksSubnet) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (subnet *VirtualNetworksSubnet) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(subnet.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  subnet.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (subnet *VirtualNetworksSubnet) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded); ok {
		subnet.Status = *st
		return nil
	}

	// Convert status to required version
	var st Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	subnet.Status = st
	return nil
}

// Hub marks that this VirtualNetworksSubnet is the hub type for conversion
func (subnet *VirtualNetworksSubnet) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (subnet *VirtualNetworksSubnet) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: subnet.Spec.OriginalVersion,
		Kind:    "VirtualNetworksSubnet",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.VirtualNetworksSubnet
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/virtualNetworks_subnets
type VirtualNetworksSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworksSubnet `json:"items"`
}

//Storage version of v1alpha1api20201101.Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded
type Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded struct {
	AddressPrefix                      *string                                                                   `json:"addressPrefix,omitempty"`
	AddressPrefixes                    []string                                                                  `json:"addressPrefixes,omitempty"`
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfiguration_Status                                `json:"applicationGatewayIpConfigurations,omitempty"`
	Conditions                         []conditions.Condition                                                    `json:"conditions,omitempty"`
	Delegations                        []Delegation_Status                                                       `json:"delegations,omitempty"`
	Etag                               *string                                                                   `json:"etag,omitempty"`
	Id                                 *string                                                                   `json:"id,omitempty"`
	IpAllocations                      []SubResource_Status                                                      `json:"ipAllocations,omitempty"`
	IpConfigurationProfiles            []IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded `json:"ipConfigurationProfiles,omitempty"`
	IpConfigurations                   []IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded        `json:"ipConfigurations,omitempty"`
	Name                               *string                                                                   `json:"name,omitempty"`
	NatGateway                         *SubResource_Status                                                       `json:"natGateway,omitempty"`
	NetworkSecurityGroup               *NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbedded    `json:"networkSecurityGroup,omitempty"`
	PrivateEndpointNetworkPolicies     *string                                                                   `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateEndpoints                   []PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded        `json:"privateEndpoints,omitempty"`
	PrivateLinkServiceNetworkPolicies  *string                                                                   `json:"privateLinkServiceNetworkPolicies,omitempty"`
	PropertyBag                        genruntime.PropertyBag                                                    `json:"$propertyBag,omitempty"`
	ProvisioningState                  *string                                                                   `json:"provisioningState,omitempty"`
	Purpose                            *string                                                                   `json:"purpose,omitempty"`
	ResourceNavigationLinks            []ResourceNavigationLink_Status                                           `json:"resourceNavigationLinks,omitempty"`
	RouteTable                         *RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbedded              `json:"routeTable,omitempty"`
	ServiceAssociationLinks            []ServiceAssociationLink_Status                                           `json:"serviceAssociationLinks,omitempty"`
	ServiceEndpointPolicies            []ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbedded  `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                   []ServiceEndpointPropertiesFormat_Status                                  `json:"serviceEndpoints,omitempty"`
	Type                               *string                                                                   `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded{}

// ConvertStatusFrom populates our Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded from the provided source
func (embedded *Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(embedded)
}

// ConvertStatusTo populates the provided destination from our Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded
func (embedded *Subnet_Status_VirtualNetworksSubnet_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == embedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(embedded)
}

//Storage version of v1alpha1api20201101.VirtualNetworksSubnets_Spec
type VirtualNetworksSubnets_Spec struct {
	AddressPrefix   *string  `json:"addressPrefix,omitempty"`
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName            string                                               `json:"azureName"`
	Delegations          []VirtualNetworksSubnets_Spec_Properties_Delegations `json:"delegations,omitempty"`
	IpAllocations        []SubResource                                        `json:"ipAllocations,omitempty"`
	Location             *string                                              `json:"location,omitempty"`
	NatGateway           *SubResource                                         `json:"natGateway,omitempty"`
	NetworkSecurityGroup *SubResource                                         `json:"networkSecurityGroup,omitempty"`
	OriginalVersion      string                                               `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner                             genruntime.KnownResourceReference `group:"network.azure.com" json:"owner" kind:"VirtualNetwork"`
	PrivateEndpointNetworkPolicies    *string                           `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateLinkServiceNetworkPolicies *string                           `json:"privateLinkServiceNetworkPolicies,omitempty"`
	PropertyBag                       genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RouteTable                        *SubResource                      `json:"routeTable,omitempty"`
	ServiceEndpointPolicies           []SubResource                     `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                  []ServiceEndpointPropertiesFormat `json:"serviceEndpoints,omitempty"`
	Tags                              map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetworksSubnets_Spec{}

// ConvertSpecFrom populates our VirtualNetworksSubnets_Spec from the provided source
func (subnets *VirtualNetworksSubnets_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == subnets {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(subnets)
}

// ConvertSpecTo populates the provided destination from our VirtualNetworksSubnets_Spec
func (subnets *VirtualNetworksSubnets_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == subnets {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(subnets)
}

//Storage version of v1alpha1api20201101.ApplicationGatewayIPConfiguration_Status
type ApplicationGatewayIPConfiguration_Status struct {
	Etag              *string                `json:"etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Subnet            *SubResource_Status    `json:"subnet,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.Delegation_Status
type Delegation_Status struct {
	Actions           []string               `json:"actions,omitempty"`
	Etag              *string                `json:"etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	ServiceName       *string                `json:"serviceName,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded
type IPConfigurationProfile_Status_VirtualNetworksSubnet_SubResourceEmbedded struct {
	Etag              *string                `json:"etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Type              *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded
type IPConfiguration_Status_VirtualNetworksSubnet_SubResourceEmbedded struct {
	Etag                      *string                                                           `json:"etag,omitempty"`
	Id                        *string                                                           `json:"id,omitempty"`
	Name                      *string                                                           `json:"name,omitempty"`
	PrivateIPAddress          *string                                                           `json:"privateIPAddress,omitempty"`
	PrivateIPAllocationMethod *string                                                           `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag               genruntime.PropertyBag                                            `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                                                           `json:"provisioningState,omitempty"`
	PublicIPAddress           *PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded `json:"publicIPAddress,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbedded
type NetworkSecurityGroup_Status_VirtualNetworksSubnet_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded
type PrivateEndpoint_Status_VirtualNetworksSubnet_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status `json:"extendedLocation,omitempty"`
	Id               *string                  `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.ResourceNavigationLink_Status
type ResourceNavigationLink_Status struct {
	Etag               *string                `json:"etag,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	Link               *string                `json:"link,omitempty"`
	LinkedResourceType *string                `json:"linkedResourceType,omitempty"`
	Name               *string                `json:"name,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState  *string                `json:"provisioningState,omitempty"`
	Type               *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbedded
type RouteTable_Status_VirtualNetworksSubnet_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.ServiceAssociationLink_Status
type ServiceAssociationLink_Status struct {
	AllowDelete        *bool                  `json:"allowDelete,omitempty"`
	Etag               *string                `json:"etag,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	Link               *string                `json:"link,omitempty"`
	LinkedResourceType *string                `json:"linkedResourceType,omitempty"`
	Locations          []string               `json:"locations,omitempty"`
	Name               *string                `json:"name,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState  *string                `json:"provisioningState,omitempty"`
	Type               *string                `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbedded
type ServiceEndpointPolicy_Status_VirtualNetworksSubnet_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	Kind        *string                `json:"kind,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.ServiceEndpointPropertiesFormat
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/ServiceEndpointPropertiesFormat
type ServiceEndpointPropertiesFormat struct {
	Locations   []string               `json:"locations,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Service     *string                `json:"service,omitempty"`
}

//Storage version of v1alpha1api20201101.ServiceEndpointPropertiesFormat_Status
type ServiceEndpointPropertiesFormat_Status struct {
	Locations         []string               `json:"locations,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Service           *string                `json:"service,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworksSubnets_Spec_Properties_Delegations
type VirtualNetworksSubnets_Spec_Properties_Delegations struct {
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ServiceName *string                `json:"serviceName,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded
type PublicIPAddress_Status_VirtualNetworksSubnet_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status   `json:"extendedLocation,omitempty"`
	Id               *string                    `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Sku              *PublicIPAddressSku_Status `json:"sku,omitempty"`
	Zones            []string                   `json:"zones,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VirtualNetworksSubnet{}, &VirtualNetworksSubnetList{})
}