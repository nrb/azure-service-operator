// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

import (
	"fmt"
	v20201101s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1alpha1api20201101.VirtualNetworksVirtualNetworkPeering
// Deprecated version of VirtualNetworksVirtualNetworkPeering. Use v1beta20201101.VirtualNetworksVirtualNetworkPeering instead
type VirtualNetworksVirtualNetworkPeering struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworksVirtualNetworkPeerings_Spec `json:"spec,omitempty"`
	Status            VirtualNetworkPeering_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &VirtualNetworksVirtualNetworkPeering{}

// GetConditions returns the conditions of the resource
func (peering *VirtualNetworksVirtualNetworkPeering) GetConditions() conditions.Conditions {
	return peering.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (peering *VirtualNetworksVirtualNetworkPeering) SetConditions(conditions conditions.Conditions) {
	peering.Status.Conditions = conditions
}

var _ conversion.Convertible = &VirtualNetworksVirtualNetworkPeering{}

// ConvertFrom populates our VirtualNetworksVirtualNetworkPeering from the provided hub VirtualNetworksVirtualNetworkPeering
func (peering *VirtualNetworksVirtualNetworkPeering) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20201101s.VirtualNetworksVirtualNetworkPeering)
	if !ok {
		return fmt.Errorf("expected network/v1beta20201101storage/VirtualNetworksVirtualNetworkPeering but received %T instead", hub)
	}

	return peering.AssignPropertiesFromVirtualNetworksVirtualNetworkPeering(source)
}

// ConvertTo populates the provided hub VirtualNetworksVirtualNetworkPeering from our VirtualNetworksVirtualNetworkPeering
func (peering *VirtualNetworksVirtualNetworkPeering) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20201101s.VirtualNetworksVirtualNetworkPeering)
	if !ok {
		return fmt.Errorf("expected network/v1beta20201101storage/VirtualNetworksVirtualNetworkPeering but received %T instead", hub)
	}

	return peering.AssignPropertiesToVirtualNetworksVirtualNetworkPeering(destination)
}

var _ genruntime.KubernetesResource = &VirtualNetworksVirtualNetworkPeering{}

// AzureName returns the Azure name of the resource
func (peering *VirtualNetworksVirtualNetworkPeering) AzureName() string {
	return peering.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (peering VirtualNetworksVirtualNetworkPeering) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (peering *VirtualNetworksVirtualNetworkPeering) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (peering *VirtualNetworksVirtualNetworkPeering) GetSpec() genruntime.ConvertibleSpec {
	return &peering.Spec
}

// GetStatus returns the status of this resource
func (peering *VirtualNetworksVirtualNetworkPeering) GetStatus() genruntime.ConvertibleStatus {
	return &peering.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
func (peering *VirtualNetworksVirtualNetworkPeering) GetType() string {
	return "Microsoft.Network/virtualNetworks/virtualNetworkPeerings"
}

// NewEmptyStatus returns a new empty (blank) status
func (peering *VirtualNetworksVirtualNetworkPeering) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &VirtualNetworkPeering_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (peering *VirtualNetworksVirtualNetworkPeering) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(peering.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  peering.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (peering *VirtualNetworksVirtualNetworkPeering) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*VirtualNetworkPeering_Status); ok {
		peering.Status = *st
		return nil
	}

	// Convert status to required version
	var st VirtualNetworkPeering_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	peering.Status = st
	return nil
}

// AssignPropertiesFromVirtualNetworksVirtualNetworkPeering populates our VirtualNetworksVirtualNetworkPeering from the provided source VirtualNetworksVirtualNetworkPeering
func (peering *VirtualNetworksVirtualNetworkPeering) AssignPropertiesFromVirtualNetworksVirtualNetworkPeering(source *v20201101s.VirtualNetworksVirtualNetworkPeering) error {

	// ObjectMeta
	peering.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec VirtualNetworksVirtualNetworkPeerings_Spec
	err := spec.AssignPropertiesFromVirtualNetworksVirtualNetworkPeeringsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromVirtualNetworksVirtualNetworkPeeringsSpec() to populate field Spec")
	}
	peering.Spec = spec

	// Status
	var status VirtualNetworkPeering_Status
	err = status.AssignPropertiesFromVirtualNetworkPeeringStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromVirtualNetworkPeeringStatus() to populate field Status")
	}
	peering.Status = status

	// No error
	return nil
}

// AssignPropertiesToVirtualNetworksVirtualNetworkPeering populates the provided destination VirtualNetworksVirtualNetworkPeering from our VirtualNetworksVirtualNetworkPeering
func (peering *VirtualNetworksVirtualNetworkPeering) AssignPropertiesToVirtualNetworksVirtualNetworkPeering(destination *v20201101s.VirtualNetworksVirtualNetworkPeering) error {

	// ObjectMeta
	destination.ObjectMeta = *peering.ObjectMeta.DeepCopy()

	// Spec
	var spec v20201101s.VirtualNetworksVirtualNetworkPeerings_Spec
	err := peering.Spec.AssignPropertiesToVirtualNetworksVirtualNetworkPeeringsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToVirtualNetworksVirtualNetworkPeeringsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20201101s.VirtualNetworkPeering_Status
	err = peering.Status.AssignPropertiesToVirtualNetworkPeeringStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToVirtualNetworkPeeringStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (peering *VirtualNetworksVirtualNetworkPeering) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: peering.Spec.OriginalVersion,
		Kind:    "VirtualNetworksVirtualNetworkPeering",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20201101.VirtualNetworksVirtualNetworkPeering
// Deprecated version of VirtualNetworksVirtualNetworkPeering. Use v1beta20201101.VirtualNetworksVirtualNetworkPeering instead
type VirtualNetworksVirtualNetworkPeeringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetworksVirtualNetworkPeering `json:"items"`
}

// Storage version of v1alpha1api20201101.VirtualNetworkPeering_Status
// Deprecated version of VirtualNetworkPeering_Status. Use v1beta20201101.VirtualNetworkPeering_Status instead
type VirtualNetworkPeering_Status struct {
	AllowForwardedTraffic     *bool                                `json:"allowForwardedTraffic,omitempty"`
	AllowGatewayTransit       *bool                                `json:"allowGatewayTransit,omitempty"`
	AllowVirtualNetworkAccess *bool                                `json:"allowVirtualNetworkAccess,omitempty"`
	Conditions                []conditions.Condition               `json:"conditions,omitempty"`
	DoNotVerifyRemoteGateways *bool                                `json:"doNotVerifyRemoteGateways,omitempty"`
	Etag                      *string                              `json:"etag,omitempty"`
	Id                        *string                              `json:"id,omitempty"`
	Name                      *string                              `json:"name,omitempty"`
	PeeringState              *string                              `json:"peeringState,omitempty"`
	PropertyBag               genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	ProvisioningState         *string                              `json:"provisioningState,omitempty"`
	RemoteAddressSpace        *AddressSpace_Status                 `json:"remoteAddressSpace,omitempty"`
	RemoteBgpCommunities      *VirtualNetworkBgpCommunities_Status `json:"remoteBgpCommunities,omitempty"`
	RemoteVirtualNetwork      *SubResource_Status                  `json:"remoteVirtualNetwork,omitempty"`
	ResourceGuid              *string                              `json:"resourceGuid,omitempty"`
	Type                      *string                              `json:"type,omitempty"`
	UseRemoteGateways         *bool                                `json:"useRemoteGateways,omitempty"`
}

var _ genruntime.ConvertibleStatus = &VirtualNetworkPeering_Status{}

// ConvertStatusFrom populates our VirtualNetworkPeering_Status from the provided source
func (peering *VirtualNetworkPeering_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20201101s.VirtualNetworkPeering_Status)
	if ok {
		// Populate our instance from source
		return peering.AssignPropertiesFromVirtualNetworkPeeringStatus(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.VirtualNetworkPeering_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = peering.AssignPropertiesFromVirtualNetworkPeeringStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our VirtualNetworkPeering_Status
func (peering *VirtualNetworkPeering_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20201101s.VirtualNetworkPeering_Status)
	if ok {
		// Populate destination from our instance
		return peering.AssignPropertiesToVirtualNetworkPeeringStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.VirtualNetworkPeering_Status{}
	err := peering.AssignPropertiesToVirtualNetworkPeeringStatus(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignPropertiesFromVirtualNetworkPeeringStatus populates our VirtualNetworkPeering_Status from the provided source VirtualNetworkPeering_Status
func (peering *VirtualNetworkPeering_Status) AssignPropertiesFromVirtualNetworkPeeringStatus(source *v20201101s.VirtualNetworkPeering_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AllowForwardedTraffic
	if source.AllowForwardedTraffic != nil {
		allowForwardedTraffic := *source.AllowForwardedTraffic
		peering.AllowForwardedTraffic = &allowForwardedTraffic
	} else {
		peering.AllowForwardedTraffic = nil
	}

	// AllowGatewayTransit
	if source.AllowGatewayTransit != nil {
		allowGatewayTransit := *source.AllowGatewayTransit
		peering.AllowGatewayTransit = &allowGatewayTransit
	} else {
		peering.AllowGatewayTransit = nil
	}

	// AllowVirtualNetworkAccess
	if source.AllowVirtualNetworkAccess != nil {
		allowVirtualNetworkAccess := *source.AllowVirtualNetworkAccess
		peering.AllowVirtualNetworkAccess = &allowVirtualNetworkAccess
	} else {
		peering.AllowVirtualNetworkAccess = nil
	}

	// Conditions
	peering.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DoNotVerifyRemoteGateways
	if source.DoNotVerifyRemoteGateways != nil {
		doNotVerifyRemoteGateway := *source.DoNotVerifyRemoteGateways
		peering.DoNotVerifyRemoteGateways = &doNotVerifyRemoteGateway
	} else {
		peering.DoNotVerifyRemoteGateways = nil
	}

	// Etag
	peering.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	peering.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	peering.Name = genruntime.ClonePointerToString(source.Name)

	// PeeringState
	peering.PeeringState = genruntime.ClonePointerToString(source.PeeringState)

	// ProvisioningState
	peering.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// RemoteAddressSpace
	if source.RemoteAddressSpace != nil {
		var remoteAddressSpace AddressSpace_Status
		err := remoteAddressSpace.AssignPropertiesFromAddressSpaceStatus(source.RemoteAddressSpace)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromAddressSpaceStatus() to populate field RemoteAddressSpace")
		}
		peering.RemoteAddressSpace = &remoteAddressSpace
	} else {
		peering.RemoteAddressSpace = nil
	}

	// RemoteBgpCommunities
	if source.RemoteBgpCommunities != nil {
		var remoteBgpCommunity VirtualNetworkBgpCommunities_Status
		err := remoteBgpCommunity.AssignPropertiesFromVirtualNetworkBgpCommunitiesStatus(source.RemoteBgpCommunities)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromVirtualNetworkBgpCommunitiesStatus() to populate field RemoteBgpCommunities")
		}
		peering.RemoteBgpCommunities = &remoteBgpCommunity
	} else {
		peering.RemoteBgpCommunities = nil
	}

	// RemoteVirtualNetwork
	if source.RemoteVirtualNetwork != nil {
		var remoteVirtualNetwork SubResource_Status
		err := remoteVirtualNetwork.AssignPropertiesFromSubResourceStatus(source.RemoteVirtualNetwork)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSubResourceStatus() to populate field RemoteVirtualNetwork")
		}
		peering.RemoteVirtualNetwork = &remoteVirtualNetwork
	} else {
		peering.RemoteVirtualNetwork = nil
	}

	// ResourceGuid
	peering.ResourceGuid = genruntime.ClonePointerToString(source.ResourceGuid)

	// Type
	peering.Type = genruntime.ClonePointerToString(source.Type)

	// UseRemoteGateways
	if source.UseRemoteGateways != nil {
		useRemoteGateway := *source.UseRemoteGateways
		peering.UseRemoteGateways = &useRemoteGateway
	} else {
		peering.UseRemoteGateways = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		peering.PropertyBag = propertyBag
	} else {
		peering.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToVirtualNetworkPeeringStatus populates the provided destination VirtualNetworkPeering_Status from our VirtualNetworkPeering_Status
func (peering *VirtualNetworkPeering_Status) AssignPropertiesToVirtualNetworkPeeringStatus(destination *v20201101s.VirtualNetworkPeering_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(peering.PropertyBag)

	// AllowForwardedTraffic
	if peering.AllowForwardedTraffic != nil {
		allowForwardedTraffic := *peering.AllowForwardedTraffic
		destination.AllowForwardedTraffic = &allowForwardedTraffic
	} else {
		destination.AllowForwardedTraffic = nil
	}

	// AllowGatewayTransit
	if peering.AllowGatewayTransit != nil {
		allowGatewayTransit := *peering.AllowGatewayTransit
		destination.AllowGatewayTransit = &allowGatewayTransit
	} else {
		destination.AllowGatewayTransit = nil
	}

	// AllowVirtualNetworkAccess
	if peering.AllowVirtualNetworkAccess != nil {
		allowVirtualNetworkAccess := *peering.AllowVirtualNetworkAccess
		destination.AllowVirtualNetworkAccess = &allowVirtualNetworkAccess
	} else {
		destination.AllowVirtualNetworkAccess = nil
	}

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(peering.Conditions)

	// DoNotVerifyRemoteGateways
	if peering.DoNotVerifyRemoteGateways != nil {
		doNotVerifyRemoteGateway := *peering.DoNotVerifyRemoteGateways
		destination.DoNotVerifyRemoteGateways = &doNotVerifyRemoteGateway
	} else {
		destination.DoNotVerifyRemoteGateways = nil
	}

	// Etag
	destination.Etag = genruntime.ClonePointerToString(peering.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(peering.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(peering.Name)

	// PeeringState
	destination.PeeringState = genruntime.ClonePointerToString(peering.PeeringState)

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(peering.ProvisioningState)

	// RemoteAddressSpace
	if peering.RemoteAddressSpace != nil {
		var remoteAddressSpace v20201101s.AddressSpace_Status
		err := peering.RemoteAddressSpace.AssignPropertiesToAddressSpaceStatus(&remoteAddressSpace)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToAddressSpaceStatus() to populate field RemoteAddressSpace")
		}
		destination.RemoteAddressSpace = &remoteAddressSpace
	} else {
		destination.RemoteAddressSpace = nil
	}

	// RemoteBgpCommunities
	if peering.RemoteBgpCommunities != nil {
		var remoteBgpCommunity v20201101s.VirtualNetworkBgpCommunities_Status
		err := peering.RemoteBgpCommunities.AssignPropertiesToVirtualNetworkBgpCommunitiesStatus(&remoteBgpCommunity)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToVirtualNetworkBgpCommunitiesStatus() to populate field RemoteBgpCommunities")
		}
		destination.RemoteBgpCommunities = &remoteBgpCommunity
	} else {
		destination.RemoteBgpCommunities = nil
	}

	// RemoteVirtualNetwork
	if peering.RemoteVirtualNetwork != nil {
		var remoteVirtualNetwork v20201101s.SubResource_Status
		err := peering.RemoteVirtualNetwork.AssignPropertiesToSubResourceStatus(&remoteVirtualNetwork)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSubResourceStatus() to populate field RemoteVirtualNetwork")
		}
		destination.RemoteVirtualNetwork = &remoteVirtualNetwork
	} else {
		destination.RemoteVirtualNetwork = nil
	}

	// ResourceGuid
	destination.ResourceGuid = genruntime.ClonePointerToString(peering.ResourceGuid)

	// Type
	destination.Type = genruntime.ClonePointerToString(peering.Type)

	// UseRemoteGateways
	if peering.UseRemoteGateways != nil {
		useRemoteGateway := *peering.UseRemoteGateways
		destination.UseRemoteGateways = &useRemoteGateway
	} else {
		destination.UseRemoteGateways = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20201101.VirtualNetworksVirtualNetworkPeerings_Spec
type VirtualNetworksVirtualNetworkPeerings_Spec struct {
	AllowForwardedTraffic     *bool `json:"allowForwardedTraffic,omitempty"`
	AllowGatewayTransit       *bool `json:"allowGatewayTransit,omitempty"`
	AllowVirtualNetworkAccess *bool `json:"allowVirtualNetworkAccess,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/VirtualNetwork resource
	Owner                *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"VirtualNetwork"`
	PeeringState         *string                            `json:"peeringState,omitempty"`
	PropertyBag          genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RemoteAddressSpace   *AddressSpace                      `json:"remoteAddressSpace,omitempty"`
	RemoteBgpCommunities *VirtualNetworkBgpCommunities      `json:"remoteBgpCommunities,omitempty"`
	RemoteVirtualNetwork *SubResource                       `json:"remoteVirtualNetwork,omitempty"`
	Tags                 map[string]string                  `json:"tags,omitempty"`
	UseRemoteGateways    *bool                              `json:"useRemoteGateways,omitempty"`
}

var _ genruntime.ConvertibleSpec = &VirtualNetworksVirtualNetworkPeerings_Spec{}

// ConvertSpecFrom populates our VirtualNetworksVirtualNetworkPeerings_Spec from the provided source
func (peerings *VirtualNetworksVirtualNetworkPeerings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20201101s.VirtualNetworksVirtualNetworkPeerings_Spec)
	if ok {
		// Populate our instance from source
		return peerings.AssignPropertiesFromVirtualNetworksVirtualNetworkPeeringsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20201101s.VirtualNetworksVirtualNetworkPeerings_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = peerings.AssignPropertiesFromVirtualNetworksVirtualNetworkPeeringsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our VirtualNetworksVirtualNetworkPeerings_Spec
func (peerings *VirtualNetworksVirtualNetworkPeerings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20201101s.VirtualNetworksVirtualNetworkPeerings_Spec)
	if ok {
		// Populate destination from our instance
		return peerings.AssignPropertiesToVirtualNetworksVirtualNetworkPeeringsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20201101s.VirtualNetworksVirtualNetworkPeerings_Spec{}
	err := peerings.AssignPropertiesToVirtualNetworksVirtualNetworkPeeringsSpec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignPropertiesFromVirtualNetworksVirtualNetworkPeeringsSpec populates our VirtualNetworksVirtualNetworkPeerings_Spec from the provided source VirtualNetworksVirtualNetworkPeerings_Spec
func (peerings *VirtualNetworksVirtualNetworkPeerings_Spec) AssignPropertiesFromVirtualNetworksVirtualNetworkPeeringsSpec(source *v20201101s.VirtualNetworksVirtualNetworkPeerings_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AllowForwardedTraffic
	if source.AllowForwardedTraffic != nil {
		allowForwardedTraffic := *source.AllowForwardedTraffic
		peerings.AllowForwardedTraffic = &allowForwardedTraffic
	} else {
		peerings.AllowForwardedTraffic = nil
	}

	// AllowGatewayTransit
	if source.AllowGatewayTransit != nil {
		allowGatewayTransit := *source.AllowGatewayTransit
		peerings.AllowGatewayTransit = &allowGatewayTransit
	} else {
		peerings.AllowGatewayTransit = nil
	}

	// AllowVirtualNetworkAccess
	if source.AllowVirtualNetworkAccess != nil {
		allowVirtualNetworkAccess := *source.AllowVirtualNetworkAccess
		peerings.AllowVirtualNetworkAccess = &allowVirtualNetworkAccess
	} else {
		peerings.AllowVirtualNetworkAccess = nil
	}

	// AzureName
	peerings.AzureName = source.AzureName

	// Location
	peerings.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	peerings.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		peerings.Owner = &owner
	} else {
		peerings.Owner = nil
	}

	// PeeringState
	peerings.PeeringState = genruntime.ClonePointerToString(source.PeeringState)

	// RemoteAddressSpace
	if source.RemoteAddressSpace != nil {
		var remoteAddressSpace AddressSpace
		err := remoteAddressSpace.AssignPropertiesFromAddressSpace(source.RemoteAddressSpace)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromAddressSpace() to populate field RemoteAddressSpace")
		}
		peerings.RemoteAddressSpace = &remoteAddressSpace
	} else {
		peerings.RemoteAddressSpace = nil
	}

	// RemoteBgpCommunities
	if source.RemoteBgpCommunities != nil {
		var remoteBgpCommunity VirtualNetworkBgpCommunities
		err := remoteBgpCommunity.AssignPropertiesFromVirtualNetworkBgpCommunities(source.RemoteBgpCommunities)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromVirtualNetworkBgpCommunities() to populate field RemoteBgpCommunities")
		}
		peerings.RemoteBgpCommunities = &remoteBgpCommunity
	} else {
		peerings.RemoteBgpCommunities = nil
	}

	// RemoteVirtualNetwork
	if source.RemoteVirtualNetwork != nil {
		var remoteVirtualNetwork SubResource
		err := remoteVirtualNetwork.AssignPropertiesFromSubResource(source.RemoteVirtualNetwork)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSubResource() to populate field RemoteVirtualNetwork")
		}
		peerings.RemoteVirtualNetwork = &remoteVirtualNetwork
	} else {
		peerings.RemoteVirtualNetwork = nil
	}

	// Tags
	peerings.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// UseRemoteGateways
	if source.UseRemoteGateways != nil {
		useRemoteGateway := *source.UseRemoteGateways
		peerings.UseRemoteGateways = &useRemoteGateway
	} else {
		peerings.UseRemoteGateways = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		peerings.PropertyBag = propertyBag
	} else {
		peerings.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToVirtualNetworksVirtualNetworkPeeringsSpec populates the provided destination VirtualNetworksVirtualNetworkPeerings_Spec from our VirtualNetworksVirtualNetworkPeerings_Spec
func (peerings *VirtualNetworksVirtualNetworkPeerings_Spec) AssignPropertiesToVirtualNetworksVirtualNetworkPeeringsSpec(destination *v20201101s.VirtualNetworksVirtualNetworkPeerings_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(peerings.PropertyBag)

	// AllowForwardedTraffic
	if peerings.AllowForwardedTraffic != nil {
		allowForwardedTraffic := *peerings.AllowForwardedTraffic
		destination.AllowForwardedTraffic = &allowForwardedTraffic
	} else {
		destination.AllowForwardedTraffic = nil
	}

	// AllowGatewayTransit
	if peerings.AllowGatewayTransit != nil {
		allowGatewayTransit := *peerings.AllowGatewayTransit
		destination.AllowGatewayTransit = &allowGatewayTransit
	} else {
		destination.AllowGatewayTransit = nil
	}

	// AllowVirtualNetworkAccess
	if peerings.AllowVirtualNetworkAccess != nil {
		allowVirtualNetworkAccess := *peerings.AllowVirtualNetworkAccess
		destination.AllowVirtualNetworkAccess = &allowVirtualNetworkAccess
	} else {
		destination.AllowVirtualNetworkAccess = nil
	}

	// AzureName
	destination.AzureName = peerings.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(peerings.Location)

	// OriginalVersion
	destination.OriginalVersion = peerings.OriginalVersion

	// Owner
	if peerings.Owner != nil {
		owner := peerings.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// PeeringState
	destination.PeeringState = genruntime.ClonePointerToString(peerings.PeeringState)

	// RemoteAddressSpace
	if peerings.RemoteAddressSpace != nil {
		var remoteAddressSpace v20201101s.AddressSpace
		err := peerings.RemoteAddressSpace.AssignPropertiesToAddressSpace(&remoteAddressSpace)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToAddressSpace() to populate field RemoteAddressSpace")
		}
		destination.RemoteAddressSpace = &remoteAddressSpace
	} else {
		destination.RemoteAddressSpace = nil
	}

	// RemoteBgpCommunities
	if peerings.RemoteBgpCommunities != nil {
		var remoteBgpCommunity v20201101s.VirtualNetworkBgpCommunities
		err := peerings.RemoteBgpCommunities.AssignPropertiesToVirtualNetworkBgpCommunities(&remoteBgpCommunity)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToVirtualNetworkBgpCommunities() to populate field RemoteBgpCommunities")
		}
		destination.RemoteBgpCommunities = &remoteBgpCommunity
	} else {
		destination.RemoteBgpCommunities = nil
	}

	// RemoteVirtualNetwork
	if peerings.RemoteVirtualNetwork != nil {
		var remoteVirtualNetwork v20201101s.SubResource
		err := peerings.RemoteVirtualNetwork.AssignPropertiesToSubResource(&remoteVirtualNetwork)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSubResource() to populate field RemoteVirtualNetwork")
		}
		destination.RemoteVirtualNetwork = &remoteVirtualNetwork
	} else {
		destination.RemoteVirtualNetwork = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(peerings.Tags)

	// UseRemoteGateways
	if peerings.UseRemoteGateways != nil {
		useRemoteGateway := *peerings.UseRemoteGateways
		destination.UseRemoteGateways = &useRemoteGateway
	} else {
		destination.UseRemoteGateways = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&VirtualNetworksVirtualNetworkPeering{}, &VirtualNetworksVirtualNetworkPeeringList{})
}
