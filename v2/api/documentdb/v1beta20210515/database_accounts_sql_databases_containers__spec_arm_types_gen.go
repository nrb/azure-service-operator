// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsSqlDatabasesContainers_SpecARM struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: Cosmos DB container name.
	Name string `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB container.
	Properties *SqlContainerCreateUpdatePropertiesARM `json:"properties,omitempty"`

	// Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
	// greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
	// type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph",
	// "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesContainers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (containers DatabaseAccountsSqlDatabasesContainers_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (containers *DatabaseAccountsSqlDatabasesContainers_SpecARM) GetName() string {
	return containers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
func (containers *DatabaseAccountsSqlDatabasesContainers_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers"
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerCreateUpdateProperties
type SqlContainerCreateUpdatePropertiesARM struct {
	// Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are "If-Match",
	// "If-None-Match", "Session-Token" and "Throughput"
	Options *CreateUpdateOptionsARM `json:"options,omitempty"`

	// Resource: Cosmos DB SQL container resource object
	Resource *SqlContainerResourceARM `json:"resource,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlContainerResource
type SqlContainerResourceARM struct {
	// AnalyticalStorageTtl: Analytical TTL.
	AnalyticalStorageTtl *int `json:"analyticalStorageTtl,omitempty"`

	// ConflictResolutionPolicy: The conflict resolution policy for the container.
	ConflictResolutionPolicy *ConflictResolutionPolicyARM `json:"conflictResolutionPolicy,omitempty"`

	// DefaultTtl: Default time to live
	DefaultTtl *int `json:"defaultTtl,omitempty"`

	// Id: Name of the Cosmos DB SQL container
	Id *string `json:"id,omitempty"`

	// IndexingPolicy: Cosmos DB indexing policy
	IndexingPolicy *IndexingPolicyARM `json:"indexingPolicy,omitempty"`

	// PartitionKey: The configuration of the partition key to be used for partitioning data into multiple partitions
	PartitionKey *ContainerPartitionKeyARM `json:"partitionKey,omitempty"`

	// UniqueKeyPolicy: The unique key policy configuration for specifying uniqueness constraints on documents in the
	// collection in the Azure Cosmos DB service.
	UniqueKeyPolicy *UniqueKeyPolicyARM `json:"uniqueKeyPolicy,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ConflictResolutionPolicy
type ConflictResolutionPolicyARM struct {
	// ConflictResolutionPath: The conflict resolution path in the case of LastWriterWins mode.
	ConflictResolutionPath *string `json:"conflictResolutionPath,omitempty"`

	// ConflictResolutionProcedure: The procedure to resolve conflicts in the case of custom mode.
	ConflictResolutionProcedure *string `json:"conflictResolutionProcedure,omitempty"`

	// Mode: Indicates the conflict resolution mode.
	Mode *ConflictResolutionPolicyMode `json:"mode,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ContainerPartitionKey
type ContainerPartitionKeyARM struct {
	// Kind: Indicates the kind of algorithm used for partitioning. For MultiHash, multiple partition keys (upto three maximum)
	// are supported for container create.
	Kind *ContainerPartitionKeyKind `json:"kind,omitempty"`

	// Paths: List of paths using which data within the container can be partitioned
	Paths []string `json:"paths,omitempty"`

	// Version: Indicates the version of the partition key definition
	Version *int `json:"version,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IndexingPolicy
type IndexingPolicyARM struct {
	// Automatic: Indicates if the indexing policy is automatic
	Automatic *bool `json:"automatic,omitempty"`

	// CompositeIndexes: List of composite path list
	CompositeIndexes [][]CompositePathARM `json:"compositeIndexes,omitempty"`

	// ExcludedPaths: List of paths to exclude from indexing
	ExcludedPaths []ExcludedPathARM `json:"excludedPaths,omitempty"`

	// IncludedPaths: List of paths to include in the indexing
	IncludedPaths []IncludedPathARM `json:"includedPaths,omitempty"`

	// IndexingMode: Indicates the indexing mode.
	IndexingMode *IndexingPolicyIndexingMode `json:"indexingMode,omitempty"`

	// SpatialIndexes: List of spatial specifics
	SpatialIndexes []SpatialSpecARM `json:"spatialIndexes,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKeyPolicy
type UniqueKeyPolicyARM struct {
	// UniqueKeys: List of unique keys on that enforces uniqueness constraint on documents in the collection in the Azure
	// Cosmos DB service.
	UniqueKeys []UniqueKeyARM `json:"uniqueKeys,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/CompositePath
type CompositePathARM struct {
	// Order: Sort order for composite paths.
	Order *CompositePathOrder `json:"order,omitempty"`

	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// +kubebuilder:validation:Enum={"Custom","LastWriterWins"}
type ConflictResolutionPolicyMode string

const (
	ConflictResolutionPolicyMode_Custom         = ConflictResolutionPolicyMode("Custom")
	ConflictResolutionPolicyMode_LastWriterWins = ConflictResolutionPolicyMode("LastWriterWins")
)

// +kubebuilder:validation:Enum={"Hash","MultiHash","Range"}
type ContainerPartitionKeyKind string

const (
	ContainerPartitionKeyKind_Hash      = ContainerPartitionKeyKind("Hash")
	ContainerPartitionKeyKind_MultiHash = ContainerPartitionKeyKind("MultiHash")
	ContainerPartitionKeyKind_Range     = ContainerPartitionKeyKind("Range")
)

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ExcludedPath
type ExcludedPathARM struct {
	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/IncludedPath
type IncludedPathARM struct {
	// Indexes: List of indexes for this path
	Indexes []IndexesARM `json:"indexes,omitempty"`

	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`
}

// +kubebuilder:validation:Enum={"consistent","lazy","none"}
type IndexingPolicyIndexingMode string

const (
	IndexingPolicyIndexingMode_Consistent = IndexingPolicyIndexingMode("consistent")
	IndexingPolicyIndexingMode_Lazy       = IndexingPolicyIndexingMode("lazy")
	IndexingPolicyIndexingMode_None       = IndexingPolicyIndexingMode("none")
)

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SpatialSpec
type SpatialSpecARM struct {
	// Path: The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard
	// (/path/*)
	Path *string `json:"path,omitempty"`

	// Types: List of path's spatial type
	Types []SpatialSpecTypes `json:"types,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/UniqueKey
type UniqueKeyARM struct {
	// Paths: List of paths must be unique for each document in the Azure Cosmos DB service
	Paths []string `json:"paths,omitempty"`
}

// +kubebuilder:validation:Enum={"ascending","descending"}
type CompositePathOrder string

const (
	CompositePathOrder_Ascending  = CompositePathOrder("ascending")
	CompositePathOrder_Descending = CompositePathOrder("descending")
)

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/Indexes
type IndexesARM struct {
	// DataType: The datatype for which the indexing behavior is applied to.
	DataType *IndexesDataType `json:"dataType,omitempty"`

	// Kind: Indicates the type of index.
	Kind *IndexesKind `json:"kind,omitempty"`

	// Precision: The precision of the index. -1 is maximum precision.
	Precision *int `json:"precision,omitempty"`
}

// +kubebuilder:validation:Enum={"LineString","MultiPolygon","Point","Polygon"}
type SpatialSpecTypes string

const (
	SpatialSpecTypes_LineString   = SpatialSpecTypes("LineString")
	SpatialSpecTypes_MultiPolygon = SpatialSpecTypes("MultiPolygon")
	SpatialSpecTypes_Point        = SpatialSpecTypes("Point")
	SpatialSpecTypes_Polygon      = SpatialSpecTypes("Polygon")
)

// +kubebuilder:validation:Enum={"LineString","MultiPolygon","Number","Point","Polygon","String"}
type IndexesDataType string

const (
	IndexesDataType_LineString   = IndexesDataType("LineString")
	IndexesDataType_MultiPolygon = IndexesDataType("MultiPolygon")
	IndexesDataType_Number       = IndexesDataType("Number")
	IndexesDataType_Point        = IndexesDataType("Point")
	IndexesDataType_Polygon      = IndexesDataType("Polygon")
	IndexesDataType_String       = IndexesDataType("String")
)

// +kubebuilder:validation:Enum={"Hash","Range","Spatial"}
type IndexesKind string

const (
	IndexesKind_Hash    = IndexesKind("Hash")
	IndexesKind_Range   = IndexesKind("Range")
	IndexesKind_Spatial = IndexesKind("Spatial")
)
