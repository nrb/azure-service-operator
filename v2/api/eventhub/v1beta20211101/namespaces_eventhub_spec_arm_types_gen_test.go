// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

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

func Test_Namespaces_Eventhub_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhub_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhub_SpecARM, Namespaces_Eventhub_SpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhub_SpecARM runs a test to see if a specific instance of Namespaces_Eventhub_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhub_SpecARM(subject Namespaces_Eventhub_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhub_SpecARM
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

// Generator of Namespaces_Eventhub_SpecARM instances for property testing - lazily instantiated by
// Namespaces_Eventhub_SpecARMGenerator()
var namespaces_Eventhub_SpecARMGenerator gopter.Gen

// Namespaces_Eventhub_SpecARMGenerator returns a generator of Namespaces_Eventhub_SpecARM instances for property testing.
// We first initialize namespaces_Eventhub_SpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhub_SpecARMGenerator() gopter.Gen {
	if namespaces_Eventhub_SpecARMGenerator != nil {
		return namespaces_Eventhub_SpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_SpecARM(generators)
	namespaces_Eventhub_SpecARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_SpecARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhub_SpecARM(generators)
	namespaces_Eventhub_SpecARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_SpecARM{}), generators)

	return namespaces_Eventhub_SpecARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhub_SpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhub_SpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhub_SpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhub_SpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Namespaces_Eventhub_Spec_PropertiesARMGenerator())
}

func Test_Namespaces_Eventhub_Spec_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhub_Spec_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhub_Spec_PropertiesARM, Namespaces_Eventhub_Spec_PropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhub_Spec_PropertiesARM runs a test to see if a specific instance of Namespaces_Eventhub_Spec_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhub_Spec_PropertiesARM(subject Namespaces_Eventhub_Spec_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhub_Spec_PropertiesARM
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

// Generator of Namespaces_Eventhub_Spec_PropertiesARM instances for property testing - lazily instantiated by
// Namespaces_Eventhub_Spec_PropertiesARMGenerator()
var namespaces_Eventhub_Spec_PropertiesARMGenerator gopter.Gen

// Namespaces_Eventhub_Spec_PropertiesARMGenerator returns a generator of Namespaces_Eventhub_Spec_PropertiesARM instances for property testing.
// We first initialize namespaces_Eventhub_Spec_PropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhub_Spec_PropertiesARMGenerator() gopter.Gen {
	if namespaces_Eventhub_Spec_PropertiesARMGenerator != nil {
		return namespaces_Eventhub_Spec_PropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_PropertiesARM(generators)
	namespaces_Eventhub_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_PropertiesARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_PropertiesARM(generators)
	namespaces_Eventhub_Spec_PropertiesARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec_PropertiesARM{}), generators)

	return namespaces_Eventhub_Spec_PropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_PropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_PropertiesARM(gens map[string]gopter.Gen) {
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_PropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_PropertiesARM(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator())
}

func Test_Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM, Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM runs a test to see if a specific instance of Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM(subject Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARM
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

// Generator of Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARM instances for property testing - lazily
// instantiated by Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator()
var namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator gopter.Gen

// Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator returns a generator of Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARM instances for property testing.
// We first initialize namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator() gopter.Gen {
	if namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator != nil {
		return namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM(generators)
	namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM(generators)
	namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec_Properties_CaptureDescriptionARM{}), generators)

	return namespaces_Eventhub_Spec_Properties_CaptureDescriptionARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.OneConstOf(Namespaces_Eventhub_Spec_Properties_CaptureDescription_Encoding_Avro, Namespaces_Eventhub_Spec_Properties_CaptureDescription_Encoding_AvroDeflate))
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescriptionARM(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator())
}

func Test_Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM, Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM runs a test to see if a specific instance of Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM(subject Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM
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

// Generator of Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM instances for property testing -
// lazily instantiated by Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator()
var namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator gopter.Gen

// Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator returns a generator of Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM instances for property testing.
// We first initialize namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator() gopter.Gen {
	if namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator != nil {
		return namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM(generators)
	namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM(generators)
	AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM(generators)
	namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM{}), generators)

	return namespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespaces_Eventhub_Spec_Properties_CaptureDescription_DestinationARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DestinationPropertiesARMGenerator())
}

func Test_DestinationPropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DestinationPropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestinationPropertiesARM, DestinationPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestinationPropertiesARM runs a test to see if a specific instance of DestinationPropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDestinationPropertiesARM(subject DestinationPropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DestinationPropertiesARM
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

// Generator of DestinationPropertiesARM instances for property testing - lazily instantiated by
// DestinationPropertiesARMGenerator()
var destinationPropertiesARMGenerator gopter.Gen

// DestinationPropertiesARMGenerator returns a generator of DestinationPropertiesARM instances for property testing.
func DestinationPropertiesARMGenerator() gopter.Gen {
	if destinationPropertiesARMGenerator != nil {
		return destinationPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestinationPropertiesARM(generators)
	destinationPropertiesARMGenerator = gen.Struct(reflect.TypeOf(DestinationPropertiesARM{}), generators)

	return destinationPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForDestinationPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestinationPropertiesARM(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}