// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515"
	"github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type SqlDatabaseExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *SqlDatabaseExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&documentdb.SqlDatabase{},
		&v1alpha1api20210515storage.SqlDatabase{}}
}