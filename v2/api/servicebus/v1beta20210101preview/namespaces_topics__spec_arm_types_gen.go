// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesTopics_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Properties: The Topic Properties definition.
	Properties *SBTopicPropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesTopics_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (topics NamespacesTopics_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (topics *NamespacesTopics_SpecARM) GetName() string {
	return topics.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics"
func (topics *NamespacesTopics_SpecARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics"
}

// Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBTopicProperties
type SBTopicPropertiesARM struct {
	// AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
	// is 5 minutes.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
	// expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
	// set on a message itself.
	DefaultMessageTimeToLive *string `json:"defaultMessageTimeToLive,omitempty"`

	// DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
	// history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`

	// EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty"`

	// EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
	// temporarily before writing it to persistent storage.
	EnableExpress *bool `json:"enableExpress,omitempty"`

	// EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning *bool `json:"enablePartitioning,omitempty"`

	// MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
	// Default is 1024.
	MaxSizeInMegabytes *int `json:"maxSizeInMegabytes,omitempty"`

	// RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty"`

	// SupportOrdering: Value that indicates whether the topic supports ordering.
	SupportOrdering *bool `json:"supportOrdering,omitempty"`
}
