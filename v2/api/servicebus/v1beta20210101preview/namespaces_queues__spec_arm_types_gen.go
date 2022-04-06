// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210101preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesQueues_SpecARM struct {
	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	//Name: Name of the resource
	Name string `json:"name,omitempty"`

	//Properties: The Queue Properties definition.
	Properties *SBQueuePropertiesARM `json:"properties,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesQueues_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (queues NamespacesQueues_SpecARM) GetAPIVersion() string {
	return "2021-01-01-preview"
}

// GetName returns the Name of the resource
func (queues NamespacesQueues_SpecARM) GetName() string {
	return queues.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/queues"
func (queues NamespacesQueues_SpecARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces/queues"
}

//Generated from: https://schema.management.azure.com/schemas/2021-01-01-preview/Microsoft.ServiceBus.json#/definitions/SBQueueProperties
type SBQueuePropertiesARM struct {
	//AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration
	//is 5 minutes.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	//DeadLetteringOnMessageExpiration: A value that indicates whether this queue has dead letter support when a message
	//expires.
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty"`

	//DefaultMessageTimeToLive: ISO 8601 default message timespan to live value. This is the duration after which the message
	//expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
	//set on a message itself.
	DefaultMessageTimeToLive *string `json:"defaultMessageTimeToLive,omitempty"`

	//DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
	//history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`

	//EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty"`

	//EnableExpress: A value that indicates whether Express Entities are enabled. An express queue holds a message in memory
	//temporarily before writing it to persistent storage.
	EnableExpress *bool `json:"enableExpress,omitempty"`

	//EnablePartitioning: A value that indicates whether the queue is to be partitioned across multiple message brokers.
	EnablePartitioning *bool `json:"enablePartitioning,omitempty"`

	//ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty"`

	//ForwardTo: Queue/Topic name to forward the messages
	ForwardTo *string `json:"forwardTo,omitempty"`

	//LockDuration: ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for
	//other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.
	LockDuration *string `json:"lockDuration,omitempty"`

	//MaxDeliveryCount: The maximum delivery count. A message is automatically deadlettered after this number of deliveries.
	//default value is 10.
	MaxDeliveryCount *int `json:"maxDeliveryCount,omitempty"`

	//MaxSizeInMegabytes: The maximum size of the queue in megabytes, which is the size of memory allocated for the queue.
	//Default is 1024.
	MaxSizeInMegabytes *int `json:"maxSizeInMegabytes,omitempty"`

	//RequiresDuplicateDetection: A value indicating if this queue requires duplicate detection.
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty"`

	//RequiresSession: A value that indicates whether the queue supports the concept of sessions.
	RequiresSession *bool `json:"requiresSession,omitempty"`
}
