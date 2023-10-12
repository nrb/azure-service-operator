// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20220801 "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type ServiceExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *ServiceExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20220801.Service{},
		&v20220801s.Service{}}
}