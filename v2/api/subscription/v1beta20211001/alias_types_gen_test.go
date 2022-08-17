// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

import (
	"encoding/json"
	v20211001s "github.com/Azure/azure-service-operator/v2/api/subscription/v1beta20211001storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Alias_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Alias to hub returns original",
		prop.ForAll(RunResourceConversionTestForAlias, AliasGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForAlias tests if a specific instance of Alias round trips to the hub storage version and back losslessly
func RunResourceConversionTestForAlias(subject Alias) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211001s.Alias
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual Alias
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Alias_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Alias to Alias via AssignPropertiesToAlias & AssignPropertiesFromAlias returns original",
		prop.ForAll(RunPropertyAssignmentTestForAlias, AliasGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAlias tests if a specific instance of Alias can be assigned to v1beta20211001storage and back losslessly
func RunPropertyAssignmentTestForAlias(subject Alias) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211001s.Alias
	err := copied.AssignPropertiesToAlias(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Alias
	err = actual.AssignPropertiesFromAlias(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Alias_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Alias via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAlias, AliasGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAlias runs a test to see if a specific instance of Alias round trips to JSON and back losslessly
func RunJSONSerializationTestForAlias(subject Alias) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Alias
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Alias instances for property testing - lazily instantiated by AliasGenerator()
var aliasGenerator gopter.Gen

// AliasGenerator returns a generator of Alias instances for property testing.
func AliasGenerator() gopter.Gen {
	if aliasGenerator != nil {
		return aliasGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAlias(generators)
	aliasGenerator = gen.Struct(reflect.TypeOf(Alias{}), generators)

	return aliasGenerator
}

// AddRelatedPropertyGeneratorsForAlias is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAlias(gens map[string]gopter.Gen) {
	gens["Spec"] = AliasesSpecGenerator()
	gens["Status"] = SubscriptionAliasResponseStatusGenerator()
}

func Test_Aliases_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Aliases_Spec to Aliases_Spec via AssignPropertiesToAliasesSpec & AssignPropertiesFromAliasesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForAliasesSpec, AliasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAliasesSpec tests if a specific instance of Aliases_Spec can be assigned to v1beta20211001storage and back losslessly
func RunPropertyAssignmentTestForAliasesSpec(subject Aliases_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211001s.Aliases_Spec
	err := copied.AssignPropertiesToAliasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Aliases_Spec
	err = actual.AssignPropertiesFromAliasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Aliases_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Aliases_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAliasesSpec, AliasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAliasesSpec runs a test to see if a specific instance of Aliases_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForAliasesSpec(subject Aliases_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Aliases_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Aliases_Spec instances for property testing - lazily instantiated by AliasesSpecGenerator()
var aliasesSpecGenerator gopter.Gen

// AliasesSpecGenerator returns a generator of Aliases_Spec instances for property testing.
// We first initialize aliasesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AliasesSpecGenerator() gopter.Gen {
	if aliasesSpecGenerator != nil {
		return aliasesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAliasesSpec(generators)
	aliasesSpecGenerator = gen.Struct(reflect.TypeOf(Aliases_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAliasesSpec(generators)
	AddRelatedPropertyGeneratorsForAliasesSpec(generators)
	aliasesSpecGenerator = gen.Struct(reflect.TypeOf(Aliases_Spec{}), generators)

	return aliasesSpecGenerator
}

// AddIndependentPropertyGeneratorsForAliasesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAliasesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForAliasesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAliasesSpec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PutAliasRequestPropertiesGenerator())
}

func Test_SubscriptionAliasResponse_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SubscriptionAliasResponse_Status to SubscriptionAliasResponse_Status via AssignPropertiesToSubscriptionAliasResponseStatus & AssignPropertiesFromSubscriptionAliasResponseStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubscriptionAliasResponseStatus, SubscriptionAliasResponseStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubscriptionAliasResponseStatus tests if a specific instance of SubscriptionAliasResponse_Status can be assigned to v1beta20211001storage and back losslessly
func RunPropertyAssignmentTestForSubscriptionAliasResponseStatus(subject SubscriptionAliasResponse_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211001s.SubscriptionAliasResponse_Status
	err := copied.AssignPropertiesToSubscriptionAliasResponseStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SubscriptionAliasResponse_Status
	err = actual.AssignPropertiesFromSubscriptionAliasResponseStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SubscriptionAliasResponse_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionAliasResponse_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionAliasResponseStatus, SubscriptionAliasResponseStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionAliasResponseStatus runs a test to see if a specific instance of SubscriptionAliasResponse_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionAliasResponseStatus(subject SubscriptionAliasResponse_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionAliasResponse_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SubscriptionAliasResponse_Status instances for property testing - lazily instantiated by
// SubscriptionAliasResponseStatusGenerator()
var subscriptionAliasResponseStatusGenerator gopter.Gen

// SubscriptionAliasResponseStatusGenerator returns a generator of SubscriptionAliasResponse_Status instances for property testing.
// We first initialize subscriptionAliasResponseStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SubscriptionAliasResponseStatusGenerator() gopter.Gen {
	if subscriptionAliasResponseStatusGenerator != nil {
		return subscriptionAliasResponseStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponseStatus(generators)
	subscriptionAliasResponseStatusGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponse_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponseStatus(generators)
	AddRelatedPropertyGeneratorsForSubscriptionAliasResponseStatus(generators)
	subscriptionAliasResponseStatusGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponse_Status{}), generators)

	return subscriptionAliasResponseStatusGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionAliasResponseStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionAliasResponseStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSubscriptionAliasResponseStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSubscriptionAliasResponseStatus(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(SubscriptionAliasResponsePropertiesStatusGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusGenerator())
}

func Test_PutAliasRequestProperties_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PutAliasRequestProperties to PutAliasRequestProperties via AssignPropertiesToPutAliasRequestProperties & AssignPropertiesFromPutAliasRequestProperties returns original",
		prop.ForAll(RunPropertyAssignmentTestForPutAliasRequestProperties, PutAliasRequestPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPutAliasRequestProperties tests if a specific instance of PutAliasRequestProperties can be assigned to v1beta20211001storage and back losslessly
func RunPropertyAssignmentTestForPutAliasRequestProperties(subject PutAliasRequestProperties) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211001s.PutAliasRequestProperties
	err := copied.AssignPropertiesToPutAliasRequestProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PutAliasRequestProperties
	err = actual.AssignPropertiesFromPutAliasRequestProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_PutAliasRequestProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestProperties, PutAliasRequestPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestProperties runs a test to see if a specific instance of PutAliasRequestProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestProperties(subject PutAliasRequestProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestProperties
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of PutAliasRequestProperties instances for property testing - lazily instantiated by
// PutAliasRequestPropertiesGenerator()
var putAliasRequestPropertiesGenerator gopter.Gen

// PutAliasRequestPropertiesGenerator returns a generator of PutAliasRequestProperties instances for property testing.
// We first initialize putAliasRequestPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PutAliasRequestPropertiesGenerator() gopter.Gen {
	if putAliasRequestPropertiesGenerator != nil {
		return putAliasRequestPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestProperties(generators)
	putAliasRequestPropertiesGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestProperties(generators)
	AddRelatedPropertyGeneratorsForPutAliasRequestProperties(generators)
	putAliasRequestPropertiesGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestProperties{}), generators)

	return putAliasRequestPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestProperties(gens map[string]gopter.Gen) {
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(PutAliasRequestPropertiesWorkload_DevTest, PutAliasRequestPropertiesWorkload_Production))
}

// AddRelatedPropertyGeneratorsForPutAliasRequestProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPutAliasRequestProperties(gens map[string]gopter.Gen) {
	gens["AdditionalProperties"] = gen.PtrOf(PutAliasRequestAdditionalPropertiesGenerator())
}

func Test_SubscriptionAliasResponseProperties_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SubscriptionAliasResponseProperties_Status to SubscriptionAliasResponseProperties_Status via AssignPropertiesToSubscriptionAliasResponsePropertiesStatus & AssignPropertiesFromSubscriptionAliasResponsePropertiesStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSubscriptionAliasResponsePropertiesStatus, SubscriptionAliasResponsePropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSubscriptionAliasResponsePropertiesStatus tests if a specific instance of SubscriptionAliasResponseProperties_Status can be assigned to v1beta20211001storage and back losslessly
func RunPropertyAssignmentTestForSubscriptionAliasResponsePropertiesStatus(subject SubscriptionAliasResponseProperties_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211001s.SubscriptionAliasResponseProperties_Status
	err := copied.AssignPropertiesToSubscriptionAliasResponsePropertiesStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SubscriptionAliasResponseProperties_Status
	err = actual.AssignPropertiesFromSubscriptionAliasResponsePropertiesStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SubscriptionAliasResponseProperties_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubscriptionAliasResponseProperties_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubscriptionAliasResponsePropertiesStatus, SubscriptionAliasResponsePropertiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubscriptionAliasResponsePropertiesStatus runs a test to see if a specific instance of SubscriptionAliasResponseProperties_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSubscriptionAliasResponsePropertiesStatus(subject SubscriptionAliasResponseProperties_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubscriptionAliasResponseProperties_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SubscriptionAliasResponseProperties_Status instances for property testing - lazily instantiated by
// SubscriptionAliasResponsePropertiesStatusGenerator()
var subscriptionAliasResponsePropertiesStatusGenerator gopter.Gen

// SubscriptionAliasResponsePropertiesStatusGenerator returns a generator of SubscriptionAliasResponseProperties_Status instances for property testing.
func SubscriptionAliasResponsePropertiesStatusGenerator() gopter.Gen {
	if subscriptionAliasResponsePropertiesStatusGenerator != nil {
		return subscriptionAliasResponsePropertiesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubscriptionAliasResponsePropertiesStatus(generators)
	subscriptionAliasResponsePropertiesStatusGenerator = gen.Struct(reflect.TypeOf(SubscriptionAliasResponseProperties_Status{}), generators)

	return subscriptionAliasResponsePropertiesStatusGenerator
}

// AddIndependentPropertyGeneratorsForSubscriptionAliasResponsePropertiesStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubscriptionAliasResponsePropertiesStatus(gens map[string]gopter.Gen) {
	gens["AcceptOwnershipState"] = gen.PtrOf(gen.OneConstOf(AcceptOwnershipState_Status_Completed, AcceptOwnershipState_Status_Expired, AcceptOwnershipState_Status_Pending))
	gens["AcceptOwnershipUrl"] = gen.PtrOf(gen.AlphaString())
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedTime"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(SubscriptionAliasResponsePropertiesStatusProvisioningState_Accepted, SubscriptionAliasResponsePropertiesStatusProvisioningState_Failed, SubscriptionAliasResponsePropertiesStatusProvisioningState_Succeeded))
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(Workload_Status_DevTest, Workload_Status_Production))
}

func Test_SystemData_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SystemData_Status to SystemData_Status via AssignPropertiesToSystemDataStatus & AssignPropertiesFromSystemDataStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSystemDataStatus, SystemDataStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSystemDataStatus tests if a specific instance of SystemData_Status can be assigned to v1beta20211001storage and back losslessly
func RunPropertyAssignmentTestForSystemDataStatus(subject SystemData_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211001s.SystemData_Status
	err := copied.AssignPropertiesToSystemDataStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SystemData_Status
	err = actual.AssignPropertiesFromSystemDataStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SystemData_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SystemData_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSystemDataStatus, SystemDataStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSystemDataStatus runs a test to see if a specific instance of SystemData_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSystemDataStatus(subject SystemData_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SystemData_Status
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SystemData_Status instances for property testing - lazily instantiated by SystemDataStatusGenerator()
var systemDataStatusGenerator gopter.Gen

// SystemDataStatusGenerator returns a generator of SystemData_Status instances for property testing.
func SystemDataStatusGenerator() gopter.Gen {
	if systemDataStatusGenerator != nil {
		return systemDataStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSystemDataStatus(generators)
	systemDataStatusGenerator = gen.Struct(reflect.TypeOf(SystemData_Status{}), generators)

	return systemDataStatusGenerator
}

// AddIndependentPropertyGeneratorsForSystemDataStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSystemDataStatus(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataStatusCreatedByType_Application,
		SystemDataStatusCreatedByType_Key,
		SystemDataStatusCreatedByType_ManagedIdentity,
		SystemDataStatusCreatedByType_User))
	gens["LastModifiedAt"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedBy"] = gen.PtrOf(gen.AlphaString())
	gens["LastModifiedByType"] = gen.PtrOf(gen.OneConstOf(
		SystemDataStatusLastModifiedByType_Application,
		SystemDataStatusLastModifiedByType_Key,
		SystemDataStatusLastModifiedByType_ManagedIdentity,
		SystemDataStatusLastModifiedByType_User))
}

func Test_PutAliasRequestAdditionalProperties_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PutAliasRequestAdditionalProperties to PutAliasRequestAdditionalProperties via AssignPropertiesToPutAliasRequestAdditionalProperties & AssignPropertiesFromPutAliasRequestAdditionalProperties returns original",
		prop.ForAll(RunPropertyAssignmentTestForPutAliasRequestAdditionalProperties, PutAliasRequestAdditionalPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPutAliasRequestAdditionalProperties tests if a specific instance of PutAliasRequestAdditionalProperties can be assigned to v1beta20211001storage and back losslessly
func RunPropertyAssignmentTestForPutAliasRequestAdditionalProperties(subject PutAliasRequestAdditionalProperties) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211001s.PutAliasRequestAdditionalProperties
	err := copied.AssignPropertiesToPutAliasRequestAdditionalProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PutAliasRequestAdditionalProperties
	err = actual.AssignPropertiesFromPutAliasRequestAdditionalProperties(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_PutAliasRequestAdditionalProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestAdditionalProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestAdditionalProperties, PutAliasRequestAdditionalPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestAdditionalProperties runs a test to see if a specific instance of PutAliasRequestAdditionalProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestAdditionalProperties(subject PutAliasRequestAdditionalProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestAdditionalProperties
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of PutAliasRequestAdditionalProperties instances for property testing - lazily instantiated by
// PutAliasRequestAdditionalPropertiesGenerator()
var putAliasRequestAdditionalPropertiesGenerator gopter.Gen

// PutAliasRequestAdditionalPropertiesGenerator returns a generator of PutAliasRequestAdditionalProperties instances for property testing.
func PutAliasRequestAdditionalPropertiesGenerator() gopter.Gen {
	if putAliasRequestAdditionalPropertiesGenerator != nil {
		return putAliasRequestAdditionalPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties(generators)
	putAliasRequestAdditionalPropertiesGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestAdditionalProperties{}), generators)

	return putAliasRequestAdditionalPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties(gens map[string]gopter.Gen) {
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionTenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
