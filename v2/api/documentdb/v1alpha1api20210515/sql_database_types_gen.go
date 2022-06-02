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
// Deprecated version of SqlDatabase. Use v1beta20210515.SqlDatabase instead
type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabases_Spec `json:"spec,omitempty"`
	Status            SqlDatabaseGetResults_Status      `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabase{}

// GetConditions returns the conditions of the resource
func (database *SqlDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *SqlDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabase{}

// ConvertFrom populates our SqlDatabase from the provided hub SqlDatabase
func (database *SqlDatabase) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20210515s.SqlDatabase

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = database.AssignPropertiesFromSqlDatabase(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to database")
	}

	return nil
}

// ConvertTo populates the provided hub SqlDatabase from our SqlDatabase
func (database *SqlDatabase) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20210515s.SqlDatabase
	err := database.AssignPropertiesToSqlDatabase(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from database")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1alpha1api20210515-sqldatabase,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabases,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabases.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabase{}

// Default applies defaults to the SqlDatabase resource
func (database *SqlDatabase) Default() {
	database.defaultImpl()
	var temp interface{} = database
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (database *SqlDatabase) defaultAzureName() {
	if database.Spec.AzureName == "" {
		database.Spec.AzureName = database.Name
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabase resource
func (database *SqlDatabase) defaultImpl() { database.defaultAzureName() }

var _ genruntime.KubernetesResource = &SqlDatabase{}

// AzureName returns the Azure name of the resource
func (database *SqlDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database SqlDatabase) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (database *SqlDatabase) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (database *SqlDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *SqlDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
func (database *SqlDatabase) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *SqlDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabaseGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *SqlDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *SqlDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabaseGetResults_Status); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabaseGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1alpha1api20210515-sqldatabase,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabases,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabases.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabase{}

// ValidateCreate validates the creation of the resource
func (database *SqlDatabase) ValidateCreate() error {
	validations := database.createValidations()
	var temp interface{} = database
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
func (database *SqlDatabase) ValidateDelete() error {
	validations := database.deleteValidations()
	var temp interface{} = database
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
func (database *SqlDatabase) ValidateUpdate(old runtime.Object) error {
	validations := database.updateValidations()
	var temp interface{} = database
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
func (database *SqlDatabase) createValidations() []func() error {
	return []func() error{database.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (database *SqlDatabase) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (database *SqlDatabase) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return database.validateResourceReferences()
		},
		database.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (database *SqlDatabase) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&database.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (database *SqlDatabase) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*SqlDatabase)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, database)
}

// AssignPropertiesFromSqlDatabase populates our SqlDatabase from the provided source SqlDatabase
func (database *SqlDatabase) AssignPropertiesFromSqlDatabase(source *alpha20210515s.SqlDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabases_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status SqlDatabaseGetResults_Status
	err = status.AssignPropertiesFromSqlDatabaseGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseGetResultsStatus() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabase populates the provided destination SqlDatabase from our SqlDatabase
func (database *SqlDatabase) AssignPropertiesToSqlDatabase(destination *alpha20210515s.SqlDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20210515s.DatabaseAccountsSqlDatabases_Spec
	err := database.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20210515s.SqlDatabaseGetResults_Status
	err = database.Status.AssignPropertiesToSqlDatabaseGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseGetResultsStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *SqlDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion(),
		Kind:    "SqlDatabase",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of SqlDatabase. Use v1beta20210515.SqlDatabase instead
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabase `json:"items"`
}

type DatabaseAccountsSqlDatabases_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string               `json:"azureName,omitempty"`
	Location  *string              `json:"location,omitempty"`
	Options   *CreateUpdateOptions `json:"options,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/DatabaseAccount resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"DatabaseAccount"`

	// +kubebuilder:validation:Required
	Resource *SqlDatabaseResource `json:"resource,omitempty"`
	Tags     map[string]string    `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabases_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (databases *DatabaseAccountsSqlDatabases_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if databases == nil {
		return nil, nil
	}
	result := &DatabaseAccountsSqlDatabases_SpecARM{}

	// Set property ‘Location’:
	if databases.Location != nil {
		location := *databases.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if databases.Options != nil || databases.Resource != nil {
		result.Properties = &SqlDatabaseCreateUpdatePropertiesARM{}
	}
	if databases.Options != nil {
		optionsARM, err := (*databases.Options).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		options := *optionsARM.(*CreateUpdateOptionsARM)
		result.Properties.Options = &options
	}
	if databases.Resource != nil {
		resourceARM, err := (*databases.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resourceARM.(*SqlDatabaseResourceARM)
		result.Properties.Resource = &resource
	}

	// Set property ‘Tags’:
	if databases.Tags != nil {
		result.Tags = make(map[string]string, len(databases.Tags))
		for key, value := range databases.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (databases *DatabaseAccountsSqlDatabases_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabases_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (databases *DatabaseAccountsSqlDatabases_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabases_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabases_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	databases.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		databases.Location = &location
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Options != nil {
			var options1 CreateUpdateOptions
			err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
			if err != nil {
				return err
			}
			options := options1
			databases.Options = &options
		}
	}

	// Set property ‘Owner’:
	databases.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlDatabaseResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			databases.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		databases.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			databases.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabases_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabases_Spec from the provided source
func (databases *DatabaseAccountsSqlDatabases_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20210515s.DatabaseAccountsSqlDatabases_Spec)
	if ok {
		// Populate our instance from source
		return databases.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(src)
	}

	// Convert to an intermediate form
	src = &alpha20210515s.DatabaseAccountsSqlDatabases_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databases.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabases_Spec
func (databases *DatabaseAccountsSqlDatabases_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20210515s.DatabaseAccountsSqlDatabases_Spec)
	if ok {
		// Populate destination from our instance
		return databases.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210515s.DatabaseAccountsSqlDatabases_Spec{}
	err := databases.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec populates our DatabaseAccountsSqlDatabases_Spec from the provided source DatabaseAccountsSqlDatabases_Spec
func (databases *DatabaseAccountsSqlDatabases_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(source *alpha20210515s.DatabaseAccountsSqlDatabases_Spec) error {

	// AzureName
	databases.AzureName = source.AzureName

	// Location
	databases.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		databases.Options = &option
	} else {
		databases.Options = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		databases.Owner = &owner
	} else {
		databases.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseResource
		err := resource.AssignPropertiesFromSqlDatabaseResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseResource() to populate field Resource")
		}
		databases.Resource = &resource
	} else {
		databases.Resource = nil
	}

	// Tags
	databases.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesSpec populates the provided destination DatabaseAccountsSqlDatabases_Spec from our DatabaseAccountsSqlDatabases_Spec
func (databases *DatabaseAccountsSqlDatabases_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(destination *alpha20210515s.DatabaseAccountsSqlDatabases_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = databases.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(databases.Location)

	// Options
	if databases.Options != nil {
		var option alpha20210515s.CreateUpdateOptions
		err := databases.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = databases.OriginalVersion()

	// Owner
	if databases.Owner != nil {
		owner := databases.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if databases.Resource != nil {
		var resource alpha20210515s.SqlDatabaseResource
		err := databases.Resource.AssignPropertiesToSqlDatabaseResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(databases.Tags)

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
func (databases *DatabaseAccountsSqlDatabases_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (databases *DatabaseAccountsSqlDatabases_Spec) SetAzureName(azureName string) {
	databases.AzureName = azureName
}

// Deprecated version of SqlDatabaseGetResults_Status. Use v1beta20210515.SqlDatabaseGetResults_Status instead
type SqlDatabaseGetResults_Status struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition                    `json:"conditions,omitempty"`
	Id         *string                                   `json:"id,omitempty"`
	Location   *string                                   `json:"location,omitempty"`
	Name       *string                                   `json:"name,omitempty"`
	Options    *OptionsResource_Status                   `json:"options,omitempty"`
	Resource   *SqlDatabaseGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`
	Type       *string                                   `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabaseGetResults_Status{}

// ConvertStatusFrom populates our SqlDatabaseGetResults_Status from the provided source
func (results *SqlDatabaseGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20210515s.SqlDatabaseGetResults_Status)
	if ok {
		// Populate our instance from source
		return results.AssignPropertiesFromSqlDatabaseGetResultsStatus(src)
	}

	// Convert to an intermediate form
	src = &alpha20210515s.SqlDatabaseGetResults_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = results.AssignPropertiesFromSqlDatabaseGetResultsStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlDatabaseGetResults_Status
func (results *SqlDatabaseGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20210515s.SqlDatabaseGetResults_Status)
	if ok {
		// Populate destination from our instance
		return results.AssignPropertiesToSqlDatabaseGetResultsStatus(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210515s.SqlDatabaseGetResults_Status{}
	err := results.AssignPropertiesToSqlDatabaseGetResultsStatus(dst)
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

var _ genruntime.FromARMConverter = &SqlDatabaseGetResults_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (results *SqlDatabaseGetResults_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlDatabaseGetResults_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (results *SqlDatabaseGetResults_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlDatabaseGetResults_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlDatabaseGetResults_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		results.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		results.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		results.Name = &name
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Options != nil {
			var options1 OptionsResource_Status
			err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
			if err != nil {
				return err
			}
			options := options1
			results.Options = &options
		}
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlDatabaseGetProperties_Status_Resource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			results.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		results.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			results.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		results.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlDatabaseGetResultsStatus populates our SqlDatabaseGetResults_Status from the provided source SqlDatabaseGetResults_Status
func (results *SqlDatabaseGetResults_Status) AssignPropertiesFromSqlDatabaseGetResultsStatus(source *alpha20210515s.SqlDatabaseGetResults_Status) error {

	// Conditions
	results.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	results.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	results.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	results.Name = genruntime.ClonePointerToString(source.Name)

	// Options
	if source.Options != nil {
		var option OptionsResource_Status
		err := option.AssignPropertiesFromOptionsResourceStatus(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromOptionsResourceStatus() to populate field Options")
		}
		results.Options = &option
	} else {
		results.Options = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseGetProperties_Status_Resource
		err := resource.AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource() to populate field Resource")
		}
		results.Resource = &resource
	} else {
		results.Resource = nil
	}

	// Tags
	results.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	results.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseGetResultsStatus populates the provided destination SqlDatabaseGetResults_Status from our SqlDatabaseGetResults_Status
func (results *SqlDatabaseGetResults_Status) AssignPropertiesToSqlDatabaseGetResultsStatus(destination *alpha20210515s.SqlDatabaseGetResults_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(results.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(results.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(results.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(results.Name)

	// Options
	if results.Options != nil {
		var option alpha20210515s.OptionsResource_Status
		err := results.Options.AssignPropertiesToOptionsResourceStatus(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToOptionsResourceStatus() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// Resource
	if results.Resource != nil {
		var resource alpha20210515s.SqlDatabaseGetProperties_Status_Resource
		err := results.Resource.AssignPropertiesToSqlDatabaseGetPropertiesStatusResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseGetPropertiesStatusResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(results.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(results.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of SqlDatabaseGetProperties_Status_Resource. Use v1beta20210515.SqlDatabaseGetProperties_Status_Resource instead
type SqlDatabaseGetProperties_Status_Resource struct {
	Colls *string  `json:"_colls,omitempty"`
	Etag  *string  `json:"_etag,omitempty"`
	Id    *string  `json:"id,omitempty"`
	Rid   *string  `json:"_rid,omitempty"`
	Ts    *float64 `json:"_ts,omitempty"`
	Users *string  `json:"_users,omitempty"`
}

var _ genruntime.FromARMConverter = &SqlDatabaseGetProperties_Status_Resource{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlDatabaseGetProperties_Status_Resource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlDatabaseGetProperties_Status_ResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlDatabaseGetProperties_Status_Resource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlDatabaseGetProperties_Status_ResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlDatabaseGetProperties_Status_ResourceARM, got %T", armInput)
	}

	// Set property ‘Colls’:
	if typedInput.Colls != nil {
		colls := *typedInput.Colls
		resource.Colls = &colls
	}

	// Set property ‘Etag’:
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		resource.Etag = &etag
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		resource.Id = &id
	}

	// Set property ‘Rid’:
	if typedInput.Rid != nil {
		rid := *typedInput.Rid
		resource.Rid = &rid
	}

	// Set property ‘Ts’:
	if typedInput.Ts != nil {
		ts := *typedInput.Ts
		resource.Ts = &ts
	}

	// Set property ‘Users’:
	if typedInput.Users != nil {
		users := *typedInput.Users
		resource.Users = &users
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource populates our SqlDatabaseGetProperties_Status_Resource from the provided source SqlDatabaseGetProperties_Status_Resource
func (resource *SqlDatabaseGetProperties_Status_Resource) AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource(source *alpha20210515s.SqlDatabaseGetProperties_Status_Resource) error {

	// Colls
	resource.Colls = genruntime.ClonePointerToString(source.Colls)

	// Etag
	resource.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Rid
	resource.Rid = genruntime.ClonePointerToString(source.Rid)

	// Ts
	if source.Ts != nil {
		t := *source.Ts
		resource.Ts = &t
	} else {
		resource.Ts = nil
	}

	// Users
	resource.Users = genruntime.ClonePointerToString(source.Users)

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseGetPropertiesStatusResource populates the provided destination SqlDatabaseGetProperties_Status_Resource from our SqlDatabaseGetProperties_Status_Resource
func (resource *SqlDatabaseGetProperties_Status_Resource) AssignPropertiesToSqlDatabaseGetPropertiesStatusResource(destination *alpha20210515s.SqlDatabaseGetProperties_Status_Resource) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Colls
	destination.Colls = genruntime.ClonePointerToString(resource.Colls)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(resource.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Rid
	destination.Rid = genruntime.ClonePointerToString(resource.Rid)

	// Ts
	if resource.Ts != nil {
		t := *resource.Ts
		destination.Ts = &t
	} else {
		destination.Ts = nil
	}

	// Users
	destination.Users = genruntime.ClonePointerToString(resource.Users)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of SqlDatabaseResource. Use v1beta20210515.SqlDatabaseResource instead
type SqlDatabaseResource struct {
	// +kubebuilder:validation:Required
	Id *string `json:"id,omitempty"`
}

var _ genruntime.ARMTransformer = &SqlDatabaseResource{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (resource *SqlDatabaseResource) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if resource == nil {
		return nil, nil
	}
	result := &SqlDatabaseResourceARM{}

	// Set property ‘Id’:
	if resource.Id != nil {
		id := *resource.Id
		result.Id = &id
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlDatabaseResource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlDatabaseResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlDatabaseResource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlDatabaseResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlDatabaseResourceARM, got %T", armInput)
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		resource.Id = &id
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlDatabaseResource populates our SqlDatabaseResource from the provided source SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignPropertiesFromSqlDatabaseResource(source *alpha20210515s.SqlDatabaseResource) error {

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseResource populates the provided destination SqlDatabaseResource from our SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignPropertiesToSqlDatabaseResource(destination *alpha20210515s.SqlDatabaseResource) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

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
	SchemeBuilder.Register(&SqlDatabase{}, &SqlDatabaseList{})
}
