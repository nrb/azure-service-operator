// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec. Use v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec instead
type DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM struct {
	Location   *string                                      `json:"location,omitempty"`
	Name       string                                       `json:"name,omitempty"`
	Properties *SqlStoredProcedureCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                            `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (procedures DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM) GetAPIVersion() string {
	return "2021-05-15"
}

// GetName returns the Name of the resource
func (procedures DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM) GetName() string {
	return procedures.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
func (procedures DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
}

//Deprecated version of SqlStoredProcedureCreateUpdateProperties. Use v1beta20210515.SqlStoredProcedureCreateUpdateProperties instead
type SqlStoredProcedureCreateUpdatePropertiesARM struct {
	Options  *CreateUpdateOptionsARM        `json:"options,omitempty"`
	Resource *SqlStoredProcedureResourceARM `json:"resource,omitempty"`
}

//Deprecated version of SqlStoredProcedureResource. Use v1beta20210515.SqlStoredProcedureResource instead
type SqlStoredProcedureResourceARM struct {
	Body *string `json:"body,omitempty"`
	Id   *string `json:"id,omitempty"`
}
