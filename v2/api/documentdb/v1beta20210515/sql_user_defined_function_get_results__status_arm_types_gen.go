// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

type SqlUserDefinedFunctionGetResults_StatusARM struct {
	//Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	//Properties: The properties of an Azure Cosmos DB userDefinedFunction
	Properties *SqlUserDefinedFunctionGetProperties_StatusARM `json:"properties,omitempty"`
	Tags       map[string]string                              `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

type SqlUserDefinedFunctionGetProperties_StatusARM struct {
	Resource *SqlUserDefinedFunctionGetProperties_Status_ResourceARM `json:"resource,omitempty"`
}

type SqlUserDefinedFunctionGetProperties_Status_ResourceARM struct {
	//Body: Body of the User Defined Function
	Body *string `json:"body,omitempty"`

	//Etag: A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	//Id: Name of the Cosmos DB SQL userDefinedFunction
	Id *string `json:"id,omitempty"`

	//Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	//Ts: A system generated property that denotes the last updated timestamp of the resource.
	Ts *float64 `json:"_ts,omitempty"`
}
