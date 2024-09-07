// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231115

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccounts_SqlRoleAssignment_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties to create and update an Azure Cosmos DB SQL Role Assignment.
	Properties *SqlRoleAssignmentResource_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SqlRoleAssignment_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-15"
func (assignment DatabaseAccounts_SqlRoleAssignment_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (assignment *DatabaseAccounts_SqlRoleAssignment_Spec_ARM) GetName() string {
	return assignment.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"
func (assignment *DatabaseAccounts_SqlRoleAssignment_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"
}

// Azure Cosmos DB SQL Role Assignment resource object.
type SqlRoleAssignmentResource_ARM struct {
	// PrincipalId: The unique identifier for the associated AAD principal in the AAD graph to which access is being granted
	// through this Role Assignment. Tenant ID for the principal is inferred using the tenant associated with the subscription.
	PrincipalId *string `json:"principalId,omitempty" optionalConfigMapPair:"PrincipalId"`

	// RoleDefinitionId: The unique identifier for the associated Role Definition.
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`

	// Scope: The data plane resource path for which access is being granted through this Role Assignment.
	Scope *string `json:"scope,omitempty"`
}