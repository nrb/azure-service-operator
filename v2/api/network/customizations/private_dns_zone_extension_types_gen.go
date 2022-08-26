// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20180901 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20180901"
	v20180901s "github.com/Azure/azure-service-operator/v2/api/network/v1beta20180901storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type PrivateDnsZoneExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *PrivateDnsZoneExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20180901.PrivateDnsZone{},
		&v20180901s.PrivateDnsZone{}}
}