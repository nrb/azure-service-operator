// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210101previewstorage

import (
	"fmt"
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1beta20210101previewstorage"
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
// Storage version of v1alpha1api20210101preview.NamespacesQueue
// Deprecated version of NamespacesQueue. Use v1beta20210101preview.NamespacesQueue instead
type NamespacesQueue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesQueues_Spec `json:"spec,omitempty"`
	Status            SBQueue_Status        `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesQueue{}

// GetConditions returns the conditions of the resource
func (queue *NamespacesQueue) GetConditions() conditions.Conditions {
	return queue.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (queue *NamespacesQueue) SetConditions(conditions conditions.Conditions) {
	queue.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesQueue{}

// ConvertFrom populates our NamespacesQueue from the provided hub NamespacesQueue
func (queue *NamespacesQueue) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210101ps.NamespacesQueue)
	if !ok {
		return fmt.Errorf("expected servicebus/v1beta20210101previewstorage/NamespacesQueue but received %T instead", hub)
	}

	return queue.AssignPropertiesFromNamespacesQueue(source)
}

// ConvertTo populates the provided hub NamespacesQueue from our NamespacesQueue
func (queue *NamespacesQueue) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210101ps.NamespacesQueue)
	if !ok {
		return fmt.Errorf("expected servicebus/v1beta20210101previewstorage/NamespacesQueue but received %T instead", hub)
	}

	return queue.AssignPropertiesToNamespacesQueue(destination)
}

var _ genruntime.KubernetesResource = &NamespacesQueue{}

// AzureName returns the Azure name of the resource
func (queue *NamespacesQueue) AzureName() string {
	return queue.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (queue NamespacesQueue) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (queue *NamespacesQueue) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (queue *NamespacesQueue) GetSpec() genruntime.ConvertibleSpec {
	return &queue.Spec
}

// GetStatus returns the status of this resource
func (queue *NamespacesQueue) GetStatus() genruntime.ConvertibleStatus {
	return &queue.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/queues"
func (queue *NamespacesQueue) GetType() string {
	return "Microsoft.ServiceBus/namespaces/queues"
}

// NewEmptyStatus returns a new empty (blank) status
func (queue *NamespacesQueue) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SBQueue_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (queue *NamespacesQueue) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(queue.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  queue.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (queue *NamespacesQueue) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SBQueue_Status); ok {
		queue.Status = *st
		return nil
	}

	// Convert status to required version
	var st SBQueue_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	queue.Status = st
	return nil
}

// AssignPropertiesFromNamespacesQueue populates our NamespacesQueue from the provided source NamespacesQueue
func (queue *NamespacesQueue) AssignPropertiesFromNamespacesQueue(source *v20210101ps.NamespacesQueue) error {

	// ObjectMeta
	queue.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesQueues_Spec
	err := spec.AssignPropertiesFromNamespacesQueuesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesQueuesSpec() to populate field Spec")
	}
	queue.Spec = spec

	// Status
	var status SBQueue_Status
	err = status.AssignPropertiesFromSBQueueStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSBQueueStatus() to populate field Status")
	}
	queue.Status = status

	// No error
	return nil
}

// AssignPropertiesToNamespacesQueue populates the provided destination NamespacesQueue from our NamespacesQueue
func (queue *NamespacesQueue) AssignPropertiesToNamespacesQueue(destination *v20210101ps.NamespacesQueue) error {

	// ObjectMeta
	destination.ObjectMeta = *queue.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210101ps.NamespacesQueues_Spec
	err := queue.Spec.AssignPropertiesToNamespacesQueuesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesQueuesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210101ps.SBQueue_Status
	err = queue.Status.AssignPropertiesToSBQueueStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSBQueueStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (queue *NamespacesQueue) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: queue.Spec.OriginalVersion,
		Kind:    "NamespacesQueue",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210101preview.NamespacesQueue
// Deprecated version of NamespacesQueue. Use v1beta20210101preview.NamespacesQueue instead
type NamespacesQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesQueue `json:"items"`
}

// Storage version of v1alpha1api20210101preview.NamespacesQueues_Spec
type NamespacesQueues_Spec struct {
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                           string  `json:"azureName,omitempty"`
	DeadLetteringOnMessageExpiration    *bool   `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool   `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool   `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool   `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string `json:"forwardTo,omitempty"`
	Location                            *string `json:"location,omitempty"`
	LockDuration                        *string `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int    `json:"maxDeliveryCount,omitempty"`
	MaxSizeInMegabytes                  *int    `json:"maxSizeInMegabytes,omitempty"`
	OriginalVersion                     string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/Namespace resource
	Owner                      *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag                genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection *bool                              `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession            *bool                              `json:"requiresSession,omitempty"`
	Tags                       map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesQueues_Spec{}

// ConvertSpecFrom populates our NamespacesQueues_Spec from the provided source
func (queues *NamespacesQueues_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210101ps.NamespacesQueues_Spec)
	if ok {
		// Populate our instance from source
		return queues.AssignPropertiesFromNamespacesQueuesSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210101ps.NamespacesQueues_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = queues.AssignPropertiesFromNamespacesQueuesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesQueues_Spec
func (queues *NamespacesQueues_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210101ps.NamespacesQueues_Spec)
	if ok {
		// Populate destination from our instance
		return queues.AssignPropertiesToNamespacesQueuesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210101ps.NamespacesQueues_Spec{}
	err := queues.AssignPropertiesToNamespacesQueuesSpec(dst)
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

// AssignPropertiesFromNamespacesQueuesSpec populates our NamespacesQueues_Spec from the provided source NamespacesQueues_Spec
func (queues *NamespacesQueues_Spec) AssignPropertiesFromNamespacesQueuesSpec(source *v20210101ps.NamespacesQueues_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AutoDeleteOnIdle
	queues.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// AzureName
	queues.AzureName = source.AzureName

	// DeadLetteringOnMessageExpiration
	if source.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *source.DeadLetteringOnMessageExpiration
		queues.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		queues.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	queues.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	queues.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		queues.EnableBatchedOperations = &enableBatchedOperation
	} else {
		queues.EnableBatchedOperations = nil
	}

	// EnableExpress
	if source.EnableExpress != nil {
		enableExpress := *source.EnableExpress
		queues.EnableExpress = &enableExpress
	} else {
		queues.EnableExpress = nil
	}

	// EnablePartitioning
	if source.EnablePartitioning != nil {
		enablePartitioning := *source.EnablePartitioning
		queues.EnablePartitioning = &enablePartitioning
	} else {
		queues.EnablePartitioning = nil
	}

	// ForwardDeadLetteredMessagesTo
	queues.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(source.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	queues.ForwardTo = genruntime.ClonePointerToString(source.ForwardTo)

	// Location
	queues.Location = genruntime.ClonePointerToString(source.Location)

	// LockDuration
	queues.LockDuration = genruntime.ClonePointerToString(source.LockDuration)

	// MaxDeliveryCount
	queues.MaxDeliveryCount = genruntime.ClonePointerToInt(source.MaxDeliveryCount)

	// MaxSizeInMegabytes
	queues.MaxSizeInMegabytes = genruntime.ClonePointerToInt(source.MaxSizeInMegabytes)

	// OriginalVersion
	queues.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		queues.Owner = &owner
	} else {
		queues.Owner = nil
	}

	// RequiresDuplicateDetection
	if source.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *source.RequiresDuplicateDetection
		queues.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		queues.RequiresDuplicateDetection = nil
	}

	// RequiresSession
	if source.RequiresSession != nil {
		requiresSession := *source.RequiresSession
		queues.RequiresSession = &requiresSession
	} else {
		queues.RequiresSession = nil
	}

	// Tags
	queues.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		queues.PropertyBag = propertyBag
	} else {
		queues.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNamespacesQueuesSpec populates the provided destination NamespacesQueues_Spec from our NamespacesQueues_Spec
func (queues *NamespacesQueues_Spec) AssignPropertiesToNamespacesQueuesSpec(destination *v20210101ps.NamespacesQueues_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(queues.PropertyBag)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(queues.AutoDeleteOnIdle)

	// AzureName
	destination.AzureName = queues.AzureName

	// DeadLetteringOnMessageExpiration
	if queues.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *queues.DeadLetteringOnMessageExpiration
		destination.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		destination.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(queues.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(queues.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if queues.EnableBatchedOperations != nil {
		enableBatchedOperation := *queues.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// EnableExpress
	if queues.EnableExpress != nil {
		enableExpress := *queues.EnableExpress
		destination.EnableExpress = &enableExpress
	} else {
		destination.EnableExpress = nil
	}

	// EnablePartitioning
	if queues.EnablePartitioning != nil {
		enablePartitioning := *queues.EnablePartitioning
		destination.EnablePartitioning = &enablePartitioning
	} else {
		destination.EnablePartitioning = nil
	}

	// ForwardDeadLetteredMessagesTo
	destination.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(queues.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	destination.ForwardTo = genruntime.ClonePointerToString(queues.ForwardTo)

	// Location
	destination.Location = genruntime.ClonePointerToString(queues.Location)

	// LockDuration
	destination.LockDuration = genruntime.ClonePointerToString(queues.LockDuration)

	// MaxDeliveryCount
	destination.MaxDeliveryCount = genruntime.ClonePointerToInt(queues.MaxDeliveryCount)

	// MaxSizeInMegabytes
	destination.MaxSizeInMegabytes = genruntime.ClonePointerToInt(queues.MaxSizeInMegabytes)

	// OriginalVersion
	destination.OriginalVersion = queues.OriginalVersion

	// Owner
	if queues.Owner != nil {
		owner := queues.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// RequiresDuplicateDetection
	if queues.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *queues.RequiresDuplicateDetection
		destination.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		destination.RequiresDuplicateDetection = nil
	}

	// RequiresSession
	if queues.RequiresSession != nil {
		requiresSession := *queues.RequiresSession
		destination.RequiresSession = &requiresSession
	} else {
		destination.RequiresSession = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(queues.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210101preview.SBQueue_Status
// Deprecated version of SBQueue_Status. Use v1beta20210101preview.SBQueue_Status instead
type SBQueue_Status struct {
	AccessedAt                          *string                     `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string                     `json:"autoDeleteOnIdle,omitempty"`
	Conditions                          []conditions.Condition      `json:"conditions,omitempty"`
	CountDetails                        *MessageCountDetails_Status `json:"countDetails,omitempty"`
	CreatedAt                           *string                     `json:"createdAt,omitempty"`
	DeadLetteringOnMessageExpiration    *bool                       `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string                     `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string                     `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                       `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                       `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                       `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string                     `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string                     `json:"forwardTo,omitempty"`
	Id                                  *string                     `json:"id,omitempty"`
	LockDuration                        *string                     `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int                        `json:"maxDeliveryCount,omitempty"`
	MaxSizeInMegabytes                  *int                        `json:"maxSizeInMegabytes,omitempty"`
	MessageCount                        *int                        `json:"messageCount,omitempty"`
	Name                                *string                     `json:"name,omitempty"`
	PropertyBag                         genruntime.PropertyBag      `json:"$propertyBag,omitempty"`
	RequiresDuplicateDetection          *bool                       `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession                     *bool                       `json:"requiresSession,omitempty"`
	SizeInBytes                         *int                        `json:"sizeInBytes,omitempty"`
	Status                              *string                     `json:"status,omitempty"`
	SystemData                          *SystemData_Status          `json:"systemData,omitempty"`
	Type                                *string                     `json:"type,omitempty"`
	UpdatedAt                           *string                     `json:"updatedAt,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SBQueue_Status{}

// ConvertStatusFrom populates our SBQueue_Status from the provided source
func (queue *SBQueue_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210101ps.SBQueue_Status)
	if ok {
		// Populate our instance from source
		return queue.AssignPropertiesFromSBQueueStatus(src)
	}

	// Convert to an intermediate form
	src = &v20210101ps.SBQueue_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = queue.AssignPropertiesFromSBQueueStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SBQueue_Status
func (queue *SBQueue_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210101ps.SBQueue_Status)
	if ok {
		// Populate destination from our instance
		return queue.AssignPropertiesToSBQueueStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20210101ps.SBQueue_Status{}
	err := queue.AssignPropertiesToSBQueueStatus(dst)
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

// AssignPropertiesFromSBQueueStatus populates our SBQueue_Status from the provided source SBQueue_Status
func (queue *SBQueue_Status) AssignPropertiesFromSBQueueStatus(source *v20210101ps.SBQueue_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AccessedAt
	queue.AccessedAt = genruntime.ClonePointerToString(source.AccessedAt)

	// AutoDeleteOnIdle
	queue.AutoDeleteOnIdle = genruntime.ClonePointerToString(source.AutoDeleteOnIdle)

	// Conditions
	queue.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CountDetails
	if source.CountDetails != nil {
		var countDetail MessageCountDetails_Status
		err := countDetail.AssignPropertiesFromMessageCountDetailsStatus(source.CountDetails)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromMessageCountDetailsStatus() to populate field CountDetails")
		}
		queue.CountDetails = &countDetail
	} else {
		queue.CountDetails = nil
	}

	// CreatedAt
	queue.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// DeadLetteringOnMessageExpiration
	if source.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *source.DeadLetteringOnMessageExpiration
		queue.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		queue.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	queue.DefaultMessageTimeToLive = genruntime.ClonePointerToString(source.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	queue.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(source.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if source.EnableBatchedOperations != nil {
		enableBatchedOperation := *source.EnableBatchedOperations
		queue.EnableBatchedOperations = &enableBatchedOperation
	} else {
		queue.EnableBatchedOperations = nil
	}

	// EnableExpress
	if source.EnableExpress != nil {
		enableExpress := *source.EnableExpress
		queue.EnableExpress = &enableExpress
	} else {
		queue.EnableExpress = nil
	}

	// EnablePartitioning
	if source.EnablePartitioning != nil {
		enablePartitioning := *source.EnablePartitioning
		queue.EnablePartitioning = &enablePartitioning
	} else {
		queue.EnablePartitioning = nil
	}

	// ForwardDeadLetteredMessagesTo
	queue.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(source.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	queue.ForwardTo = genruntime.ClonePointerToString(source.ForwardTo)

	// Id
	queue.Id = genruntime.ClonePointerToString(source.Id)

	// LockDuration
	queue.LockDuration = genruntime.ClonePointerToString(source.LockDuration)

	// MaxDeliveryCount
	queue.MaxDeliveryCount = genruntime.ClonePointerToInt(source.MaxDeliveryCount)

	// MaxSizeInMegabytes
	queue.MaxSizeInMegabytes = genruntime.ClonePointerToInt(source.MaxSizeInMegabytes)

	// MessageCount
	queue.MessageCount = genruntime.ClonePointerToInt(source.MessageCount)

	// Name
	queue.Name = genruntime.ClonePointerToString(source.Name)

	// RequiresDuplicateDetection
	if source.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *source.RequiresDuplicateDetection
		queue.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		queue.RequiresDuplicateDetection = nil
	}

	// RequiresSession
	if source.RequiresSession != nil {
		requiresSession := *source.RequiresSession
		queue.RequiresSession = &requiresSession
	} else {
		queue.RequiresSession = nil
	}

	// SizeInBytes
	queue.SizeInBytes = genruntime.ClonePointerToInt(source.SizeInBytes)

	// Status
	queue.Status = genruntime.ClonePointerToString(source.Status)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		queue.SystemData = &systemDatum
	} else {
		queue.SystemData = nil
	}

	// Type
	queue.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedAt
	queue.UpdatedAt = genruntime.ClonePointerToString(source.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		queue.PropertyBag = propertyBag
	} else {
		queue.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSBQueueStatus populates the provided destination SBQueue_Status from our SBQueue_Status
func (queue *SBQueue_Status) AssignPropertiesToSBQueueStatus(destination *v20210101ps.SBQueue_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(queue.PropertyBag)

	// AccessedAt
	destination.AccessedAt = genruntime.ClonePointerToString(queue.AccessedAt)

	// AutoDeleteOnIdle
	destination.AutoDeleteOnIdle = genruntime.ClonePointerToString(queue.AutoDeleteOnIdle)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(queue.Conditions)

	// CountDetails
	if queue.CountDetails != nil {
		var countDetail v20210101ps.MessageCountDetails_Status
		err := queue.CountDetails.AssignPropertiesToMessageCountDetailsStatus(&countDetail)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToMessageCountDetailsStatus() to populate field CountDetails")
		}
		destination.CountDetails = &countDetail
	} else {
		destination.CountDetails = nil
	}

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(queue.CreatedAt)

	// DeadLetteringOnMessageExpiration
	if queue.DeadLetteringOnMessageExpiration != nil {
		deadLetteringOnMessageExpiration := *queue.DeadLetteringOnMessageExpiration
		destination.DeadLetteringOnMessageExpiration = &deadLetteringOnMessageExpiration
	} else {
		destination.DeadLetteringOnMessageExpiration = nil
	}

	// DefaultMessageTimeToLive
	destination.DefaultMessageTimeToLive = genruntime.ClonePointerToString(queue.DefaultMessageTimeToLive)

	// DuplicateDetectionHistoryTimeWindow
	destination.DuplicateDetectionHistoryTimeWindow = genruntime.ClonePointerToString(queue.DuplicateDetectionHistoryTimeWindow)

	// EnableBatchedOperations
	if queue.EnableBatchedOperations != nil {
		enableBatchedOperation := *queue.EnableBatchedOperations
		destination.EnableBatchedOperations = &enableBatchedOperation
	} else {
		destination.EnableBatchedOperations = nil
	}

	// EnableExpress
	if queue.EnableExpress != nil {
		enableExpress := *queue.EnableExpress
		destination.EnableExpress = &enableExpress
	} else {
		destination.EnableExpress = nil
	}

	// EnablePartitioning
	if queue.EnablePartitioning != nil {
		enablePartitioning := *queue.EnablePartitioning
		destination.EnablePartitioning = &enablePartitioning
	} else {
		destination.EnablePartitioning = nil
	}

	// ForwardDeadLetteredMessagesTo
	destination.ForwardDeadLetteredMessagesTo = genruntime.ClonePointerToString(queue.ForwardDeadLetteredMessagesTo)

	// ForwardTo
	destination.ForwardTo = genruntime.ClonePointerToString(queue.ForwardTo)

	// Id
	destination.Id = genruntime.ClonePointerToString(queue.Id)

	// LockDuration
	destination.LockDuration = genruntime.ClonePointerToString(queue.LockDuration)

	// MaxDeliveryCount
	destination.MaxDeliveryCount = genruntime.ClonePointerToInt(queue.MaxDeliveryCount)

	// MaxSizeInMegabytes
	destination.MaxSizeInMegabytes = genruntime.ClonePointerToInt(queue.MaxSizeInMegabytes)

	// MessageCount
	destination.MessageCount = genruntime.ClonePointerToInt(queue.MessageCount)

	// Name
	destination.Name = genruntime.ClonePointerToString(queue.Name)

	// RequiresDuplicateDetection
	if queue.RequiresDuplicateDetection != nil {
		requiresDuplicateDetection := *queue.RequiresDuplicateDetection
		destination.RequiresDuplicateDetection = &requiresDuplicateDetection
	} else {
		destination.RequiresDuplicateDetection = nil
	}

	// RequiresSession
	if queue.RequiresSession != nil {
		requiresSession := *queue.RequiresSession
		destination.RequiresSession = &requiresSession
	} else {
		destination.RequiresSession = nil
	}

	// SizeInBytes
	destination.SizeInBytes = genruntime.ClonePointerToInt(queue.SizeInBytes)

	// Status
	destination.Status = genruntime.ClonePointerToString(queue.Status)

	// SystemData
	if queue.SystemData != nil {
		var systemDatum v20210101ps.SystemData_Status
		err := queue.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(queue.Type)

	// UpdatedAt
	destination.UpdatedAt = genruntime.ClonePointerToString(queue.UpdatedAt)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210101preview.MessageCountDetails_Status
// Deprecated version of MessageCountDetails_Status. Use v1beta20210101preview.MessageCountDetails_Status instead
type MessageCountDetails_Status struct {
	ActiveMessageCount             *int                   `json:"activeMessageCount,omitempty"`
	DeadLetterMessageCount         *int                   `json:"deadLetterMessageCount,omitempty"`
	PropertyBag                    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScheduledMessageCount          *int                   `json:"scheduledMessageCount,omitempty"`
	TransferDeadLetterMessageCount *int                   `json:"transferDeadLetterMessageCount,omitempty"`
	TransferMessageCount           *int                   `json:"transferMessageCount,omitempty"`
}

// AssignPropertiesFromMessageCountDetailsStatus populates our MessageCountDetails_Status from the provided source MessageCountDetails_Status
func (details *MessageCountDetails_Status) AssignPropertiesFromMessageCountDetailsStatus(source *v20210101ps.MessageCountDetails_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ActiveMessageCount
	details.ActiveMessageCount = genruntime.ClonePointerToInt(source.ActiveMessageCount)

	// DeadLetterMessageCount
	details.DeadLetterMessageCount = genruntime.ClonePointerToInt(source.DeadLetterMessageCount)

	// ScheduledMessageCount
	details.ScheduledMessageCount = genruntime.ClonePointerToInt(source.ScheduledMessageCount)

	// TransferDeadLetterMessageCount
	details.TransferDeadLetterMessageCount = genruntime.ClonePointerToInt(source.TransferDeadLetterMessageCount)

	// TransferMessageCount
	details.TransferMessageCount = genruntime.ClonePointerToInt(source.TransferMessageCount)

	// Update the property bag
	if len(propertyBag) > 0 {
		details.PropertyBag = propertyBag
	} else {
		details.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToMessageCountDetailsStatus populates the provided destination MessageCountDetails_Status from our MessageCountDetails_Status
func (details *MessageCountDetails_Status) AssignPropertiesToMessageCountDetailsStatus(destination *v20210101ps.MessageCountDetails_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(details.PropertyBag)

	// ActiveMessageCount
	destination.ActiveMessageCount = genruntime.ClonePointerToInt(details.ActiveMessageCount)

	// DeadLetterMessageCount
	destination.DeadLetterMessageCount = genruntime.ClonePointerToInt(details.DeadLetterMessageCount)

	// ScheduledMessageCount
	destination.ScheduledMessageCount = genruntime.ClonePointerToInt(details.ScheduledMessageCount)

	// TransferDeadLetterMessageCount
	destination.TransferDeadLetterMessageCount = genruntime.ClonePointerToInt(details.TransferDeadLetterMessageCount)

	// TransferMessageCount
	destination.TransferMessageCount = genruntime.ClonePointerToInt(details.TransferMessageCount)

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
	SchemeBuilder.Register(&NamespacesQueue{}, &NamespacesQueueList{})
}
