// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
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

func Test_ThroughputSettingsGetResults_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsGetResults_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsGetResultsStatusARM, ThroughputSettingsGetResultsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsGetResultsStatusARM runs a test to see if a specific instance of ThroughputSettingsGetResults_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsGetResultsStatusARM(subject ThroughputSettingsGetResults_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsGetResults_StatusARM
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

// Generator of ThroughputSettingsGetResults_StatusARM instances for property testing - lazily instantiated by
//ThroughputSettingsGetResultsStatusARMGenerator()
var throughputSettingsGetResultsStatusARMGenerator gopter.Gen

// ThroughputSettingsGetResultsStatusARMGenerator returns a generator of ThroughputSettingsGetResults_StatusARM instances for property testing.
// We first initialize throughputSettingsGetResultsStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ThroughputSettingsGetResultsStatusARMGenerator() gopter.Gen {
	if throughputSettingsGetResultsStatusARMGenerator != nil {
		return throughputSettingsGetResultsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsGetResultsStatusARM(generators)
	throughputSettingsGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetResults_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsGetResultsStatusARM(generators)
	AddRelatedPropertyGeneratorsForThroughputSettingsGetResultsStatusARM(generators)
	throughputSettingsGetResultsStatusARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetResults_StatusARM{}), generators)

	return throughputSettingsGetResultsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputSettingsGetResultsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputSettingsGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForThroughputSettingsGetResultsStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsGetResultsStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsGetPropertiesStatusARMGenerator())
}

func Test_ThroughputSettingsGetProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsGetProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsGetPropertiesStatusARM, ThroughputSettingsGetPropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsGetPropertiesStatusARM runs a test to see if a specific instance of ThroughputSettingsGetProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsGetPropertiesStatusARM(subject ThroughputSettingsGetProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsGetProperties_StatusARM
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

// Generator of ThroughputSettingsGetProperties_StatusARM instances for property testing - lazily instantiated by
//ThroughputSettingsGetPropertiesStatusARMGenerator()
var throughputSettingsGetPropertiesStatusARMGenerator gopter.Gen

// ThroughputSettingsGetPropertiesStatusARMGenerator returns a generator of ThroughputSettingsGetProperties_StatusARM instances for property testing.
func ThroughputSettingsGetPropertiesStatusARMGenerator() gopter.Gen {
	if throughputSettingsGetPropertiesStatusARMGenerator != nil {
		return throughputSettingsGetPropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForThroughputSettingsGetPropertiesStatusARM(generators)
	throughputSettingsGetPropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetProperties_StatusARM{}), generators)

	return throughputSettingsGetPropertiesStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForThroughputSettingsGetPropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsGetPropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsGetPropertiesStatusResourceARMGenerator())
}

func Test_ThroughputSettingsGetProperties_Status_ResourceARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsGetProperties_Status_ResourceARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsGetPropertiesStatusResourceARM, ThroughputSettingsGetPropertiesStatusResourceARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsGetPropertiesStatusResourceARM runs a test to see if a specific instance of ThroughputSettingsGetProperties_Status_ResourceARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsGetPropertiesStatusResourceARM(subject ThroughputSettingsGetProperties_Status_ResourceARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsGetProperties_Status_ResourceARM
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

// Generator of ThroughputSettingsGetProperties_Status_ResourceARM instances for property testing - lazily instantiated
//by ThroughputSettingsGetPropertiesStatusResourceARMGenerator()
var throughputSettingsGetPropertiesStatusResourceARMGenerator gopter.Gen

// ThroughputSettingsGetPropertiesStatusResourceARMGenerator returns a generator of ThroughputSettingsGetProperties_Status_ResourceARM instances for property testing.
// We first initialize throughputSettingsGetPropertiesStatusResourceARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ThroughputSettingsGetPropertiesStatusResourceARMGenerator() gopter.Gen {
	if throughputSettingsGetPropertiesStatusResourceARMGenerator != nil {
		return throughputSettingsGetPropertiesStatusResourceARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsGetPropertiesStatusResourceARM(generators)
	throughputSettingsGetPropertiesStatusResourceARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetProperties_Status_ResourceARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsGetPropertiesStatusResourceARM(generators)
	AddRelatedPropertyGeneratorsForThroughputSettingsGetPropertiesStatusResourceARM(generators)
	throughputSettingsGetPropertiesStatusResourceARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsGetProperties_Status_ResourceARM{}), generators)

	return throughputSettingsGetPropertiesStatusResourceARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputSettingsGetPropertiesStatusResourceARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputSettingsGetPropertiesStatusResourceARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumThroughput"] = gen.PtrOf(gen.AlphaString())
	gens["OfferReplacePending"] = gen.PtrOf(gen.AlphaString())
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Throughput"] = gen.PtrOf(gen.Int())
	gens["Ts"] = gen.PtrOf(gen.Float64())
}

// AddRelatedPropertyGeneratorsForThroughputSettingsGetPropertiesStatusResourceARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsGetPropertiesStatusResourceARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsResourceStatusARMGenerator())
}

func Test_AutoscaleSettingsResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsResourceStatusARM, AutoscaleSettingsResourceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsResourceStatusARM runs a test to see if a specific instance of AutoscaleSettingsResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsResourceStatusARM(subject AutoscaleSettingsResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsResource_StatusARM
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

// Generator of AutoscaleSettingsResource_StatusARM instances for property testing - lazily instantiated by
//AutoscaleSettingsResourceStatusARMGenerator()
var autoscaleSettingsResourceStatusARMGenerator gopter.Gen

// AutoscaleSettingsResourceStatusARMGenerator returns a generator of AutoscaleSettingsResource_StatusARM instances for property testing.
// We first initialize autoscaleSettingsResourceStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AutoscaleSettingsResourceStatusARMGenerator() gopter.Gen {
	if autoscaleSettingsResourceStatusARMGenerator != nil {
		return autoscaleSettingsResourceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceStatusARM(generators)
	autoscaleSettingsResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceStatusARM(generators)
	AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceStatusARM(generators)
	autoscaleSettingsResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_StatusARM{}), generators)

	return autoscaleSettingsResourceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsResourceStatusARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
	gens["TargetMaxThroughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoscaleSettingsResourceStatusARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradePolicy"] = gen.PtrOf(AutoUpgradePolicyResourceStatusARMGenerator())
}

func Test_AutoUpgradePolicyResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoUpgradePolicyResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoUpgradePolicyResourceStatusARM, AutoUpgradePolicyResourceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoUpgradePolicyResourceStatusARM runs a test to see if a specific instance of AutoUpgradePolicyResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoUpgradePolicyResourceStatusARM(subject AutoUpgradePolicyResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoUpgradePolicyResource_StatusARM
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

// Generator of AutoUpgradePolicyResource_StatusARM instances for property testing - lazily instantiated by
//AutoUpgradePolicyResourceStatusARMGenerator()
var autoUpgradePolicyResourceStatusARMGenerator gopter.Gen

// AutoUpgradePolicyResourceStatusARMGenerator returns a generator of AutoUpgradePolicyResource_StatusARM instances for property testing.
func AutoUpgradePolicyResourceStatusARMGenerator() gopter.Gen {
	if autoUpgradePolicyResourceStatusARMGenerator != nil {
		return autoUpgradePolicyResourceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceStatusARM(generators)
	autoUpgradePolicyResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(AutoUpgradePolicyResource_StatusARM{}), generators)

	return autoUpgradePolicyResourceStatusARMGenerator
}

// AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoUpgradePolicyResourceStatusARM(gens map[string]gopter.Gen) {
	gens["ThroughputPolicy"] = gen.PtrOf(ThroughputPolicyResourceStatusARMGenerator())
}

func Test_ThroughputPolicyResource_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputPolicyResource_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputPolicyResourceStatusARM, ThroughputPolicyResourceStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputPolicyResourceStatusARM runs a test to see if a specific instance of ThroughputPolicyResource_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputPolicyResourceStatusARM(subject ThroughputPolicyResource_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputPolicyResource_StatusARM
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

// Generator of ThroughputPolicyResource_StatusARM instances for property testing - lazily instantiated by
//ThroughputPolicyResourceStatusARMGenerator()
var throughputPolicyResourceStatusARMGenerator gopter.Gen

// ThroughputPolicyResourceStatusARMGenerator returns a generator of ThroughputPolicyResource_StatusARM instances for property testing.
func ThroughputPolicyResourceStatusARMGenerator() gopter.Gen {
	if throughputPolicyResourceStatusARMGenerator != nil {
		return throughputPolicyResourceStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputPolicyResourceStatusARM(generators)
	throughputPolicyResourceStatusARMGenerator = gen.Struct(reflect.TypeOf(ThroughputPolicyResource_StatusARM{}), generators)

	return throughputPolicyResourceStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputPolicyResourceStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputPolicyResourceStatusARM(gens map[string]gopter.Gen) {
	gens["IncrementPercent"] = gen.PtrOf(gen.Int())
	gens["IsEnabled"] = gen.PtrOf(gen.Bool())
}
