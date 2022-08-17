// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	alpha20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Deprecated version of SqlDatabaseContainerThroughputSetting. Use v1beta20210515.SqlDatabaseContainerThroughputSetting instead
type SqlDatabaseContainerThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_Status                           `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *SqlDatabaseContainerThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerThroughputSetting{}

// ConvertFrom populates our SqlDatabaseContainerThroughputSetting from the provided hub SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20210515s.SqlDatabaseContainerThroughputSetting

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = setting.AssignPropertiesFromSqlDatabaseContainerThroughputSetting(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to setting")
	}

	return nil
}

// ConvertTo populates the provided hub SqlDatabaseContainerThroughputSetting from our SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20210515s.SqlDatabaseContainerThroughputSetting
	err := setting.AssignPropertiesToSqlDatabaseContainerThroughputSetting(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from setting")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainerthroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerthroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabasecontainerthroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &SqlDatabaseContainerThroughputSetting{}

// Default applies defaults to the SqlDatabaseContainerThroughputSetting resource
func (setting *SqlDatabaseContainerThroughputSetting) Default() {
	setting.defaultImpl()
	var temp interface{} = setting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseContainerThroughputSetting resource
func (setting *SqlDatabaseContainerThroughputSetting) defaultImpl() {}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *SqlDatabaseContainerThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting SqlDatabaseContainerThroughputSetting) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
func (setting *SqlDatabaseContainerThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *SqlDatabaseContainerThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (setting *SqlDatabaseContainerThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_Status); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainerthroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerthroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabasecontainerthroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &SqlDatabaseContainerThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (setting *SqlDatabaseContainerThroughputSetting) ValidateCreate() error {
	validations := setting.createValidations()
	var temp interface{} = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (setting *SqlDatabaseContainerThroughputSetting) ValidateDelete() error {
	validations := setting.deleteValidations()
	var temp interface{} = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (setting *SqlDatabaseContainerThroughputSetting) ValidateUpdate(old runtime.Object) error {
	validations := setting.updateValidations()
	var temp interface{} = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (setting *SqlDatabaseContainerThroughputSetting) createValidations() []func() error {
	return []func() error{setting.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (setting *SqlDatabaseContainerThroughputSetting) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (setting *SqlDatabaseContainerThroughputSetting) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return setting.validateResourceReferences()
		},
		setting.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (setting *SqlDatabaseContainerThroughputSetting) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&setting.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *SqlDatabaseContainerThroughputSetting) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*SqlDatabaseContainerThroughputSetting)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, setting)
}

// AssignPropertiesFromSqlDatabaseContainerThroughputSetting populates our SqlDatabaseContainerThroughputSetting from the provided source SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) AssignPropertiesFromSqlDatabaseContainerThroughputSetting(source *alpha20210515s.SqlDatabaseContainerThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status ThroughputSettingsGetResults_Status
	err = status.AssignPropertiesFromThroughputSettingsGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromThroughputSettingsGetResultsStatus() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerThroughputSetting populates the provided destination SqlDatabaseContainerThroughputSetting from our SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) AssignPropertiesToSqlDatabaseContainerThroughputSetting(destination *alpha20210515s.SqlDatabaseContainerThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20210515s.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
	err := setting.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20210515s.ThroughputSettingsGetResults_Status
	err = setting.Status.AssignPropertiesToThroughputSettingsGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToThroughputSettingsGetResultsStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *SqlDatabaseContainerThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseContainerThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of SqlDatabaseContainerThroughputSetting. Use v1beta20210515.SqlDatabaseContainerThroughputSetting instead
type SqlDatabaseContainerThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerThroughputSetting `json:"items"`
}

type DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec struct {
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`

	// +kubebuilder:validation:Required
	Resource *ThroughputSettingsResource `json:"resource,omitempty"`
	Tags     map[string]string           `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if settings == nil {
		return nil, nil
	}
	result := &DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM{}

	// Set property ‘Location’:
	if settings.Location != nil {
		location := *settings.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if settings.Resource != nil {
		result.Properties = &ThroughputSettingsUpdatePropertiesARM{}
	}
	if settings.Resource != nil {
		resourceARM, err := (*settings.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resourceARM.(*ThroughputSettingsResourceARM)
		result.Properties.Resource = &resource
	}

	// Set property ‘Tags’:
	if settings.Tags != nil {
		result.Tags = make(map[string]string, len(settings.Tags))
		for key, value := range settings.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM, got %T", armInput)
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		settings.Location = &location
	}

	// Set property ‘Owner’:
	settings.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			settings.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		settings.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			settings.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec from the provided source
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20210515s.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec)
	if ok {
		// Populate our instance from source
		return settings.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(src)
	}

	// Convert to an intermediate form
	src = &alpha20210515s.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = settings.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20210515s.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec)
	if ok {
		// Populate destination from our instance
		return settings.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210515s.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}
	err := settings.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec populates our DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec from the provided source DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(source *alpha20210515s.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) error {

	// Location
	settings.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		settings.Owner = &owner
	} else {
		settings.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignPropertiesFromThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromThroughputSettingsResource() to populate field Resource")
		}
		settings.Resource = &resource
	} else {
		settings.Resource = nil
	}

	// Tags
	settings.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec populates the provided destination DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec from our DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(destination *alpha20210515s.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(settings.Location)

	// OriginalVersion
	destination.OriginalVersion = settings.OriginalVersion()

	// Owner
	if settings.Owner != nil {
		owner := settings.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if settings.Resource != nil {
		var resource alpha20210515s.ThroughputSettingsResource
		err := settings.Resource.AssignPropertiesToThroughputSettingsResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(settings.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerThroughputSetting{}, &SqlDatabaseContainerThroughputSettingList{})
}
