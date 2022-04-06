// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101storage"
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

func Test_VirtualNetwork_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetwork to hub returns original",
		prop.ForAll(RunResourceConversionTestForVirtualNetwork, VirtualNetworkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForVirtualNetwork tests if a specific instance of VirtualNetwork round trips to the hub storage version and back losslessly
func RunResourceConversionTestForVirtualNetwork(subject VirtualNetwork) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1beta20201101storage.VirtualNetwork
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual VirtualNetwork
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

func Test_VirtualNetwork_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetwork to VirtualNetwork via AssignPropertiesToVirtualNetwork & AssignPropertiesFromVirtualNetwork returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetwork, VirtualNetworkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetwork tests if a specific instance of VirtualNetwork can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetwork(subject VirtualNetwork) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201101storage.VirtualNetwork
	err := copied.AssignPropertiesToVirtualNetwork(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetwork
	err = actual.AssignPropertiesFromVirtualNetwork(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetwork_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetwork, VirtualNetworkGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetwork runs a test to see if a specific instance of VirtualNetwork round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetwork(subject VirtualNetwork) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork
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

// Generator of VirtualNetwork instances for property testing - lazily instantiated by VirtualNetworkGenerator()
var virtualNetworkGenerator gopter.Gen

// VirtualNetworkGenerator returns a generator of VirtualNetwork instances for property testing.
func VirtualNetworkGenerator() gopter.Gen {
	if virtualNetworkGenerator != nil {
		return virtualNetworkGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForVirtualNetwork(generators)
	virtualNetworkGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork{}), generators)

	return virtualNetworkGenerator
}

// AddRelatedPropertyGeneratorsForVirtualNetwork is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetwork(gens map[string]gopter.Gen) {
	gens["Spec"] = VirtualNetworksSpecGenerator()
	gens["Status"] = VirtualNetworkStatusGenerator()
}

func Test_VirtualNetwork_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetwork_Status to VirtualNetwork_Status via AssignPropertiesToVirtualNetworkStatus & AssignPropertiesFromVirtualNetworkStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworkStatus, VirtualNetworkStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworkStatus tests if a specific instance of VirtualNetwork_Status can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworkStatus(subject VirtualNetwork_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201101storage.VirtualNetwork_Status
	err := copied.AssignPropertiesToVirtualNetworkStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetwork_Status
	err = actual.AssignPropertiesFromVirtualNetworkStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetwork_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetwork_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkStatus, VirtualNetworkStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkStatus runs a test to see if a specific instance of VirtualNetwork_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkStatus(subject VirtualNetwork_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetwork_Status
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

// Generator of VirtualNetwork_Status instances for property testing - lazily instantiated by
//VirtualNetworkStatusGenerator()
var virtualNetworkStatusGenerator gopter.Gen

// VirtualNetworkStatusGenerator returns a generator of VirtualNetwork_Status instances for property testing.
// We first initialize virtualNetworkStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworkStatusGenerator() gopter.Gen {
	if virtualNetworkStatusGenerator != nil {
		return virtualNetworkStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkStatus(generators)
	virtualNetworkStatusGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkStatus(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworkStatus(generators)
	virtualNetworkStatusGenerator = gen.Struct(reflect.TypeOf(VirtualNetwork_Status{}), generators)

	return virtualNetworkStatusGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkStatus(gens map[string]gopter.Gen) {
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_StatusDeleting,
		ProvisioningState_StatusFailed,
		ProvisioningState_StatusSucceeded,
		ProvisioningState_StatusUpdating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworkStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworkStatus(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpaceStatusGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesStatusGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResourceStatusGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptionsStatusGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationStatusGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceStatusGenerator())
}

func Test_VirtualNetworks_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworks_Spec to VirtualNetworks_Spec via AssignPropertiesToVirtualNetworksSpec & AssignPropertiesFromVirtualNetworksSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworksSpec, VirtualNetworksSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworksSpec tests if a specific instance of VirtualNetworks_Spec can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworksSpec(subject VirtualNetworks_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201101storage.VirtualNetworks_Spec
	err := copied.AssignPropertiesToVirtualNetworksSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworks_Spec
	err = actual.AssignPropertiesFromVirtualNetworksSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworks_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworks_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworksSpec, VirtualNetworksSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworksSpec runs a test to see if a specific instance of VirtualNetworks_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworksSpec(subject VirtualNetworks_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworks_Spec
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

// Generator of VirtualNetworks_Spec instances for property testing - lazily instantiated by
//VirtualNetworksSpecGenerator()
var virtualNetworksSpecGenerator gopter.Gen

// VirtualNetworksSpecGenerator returns a generator of VirtualNetworks_Spec instances for property testing.
// We first initialize virtualNetworksSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func VirtualNetworksSpecGenerator() gopter.Gen {
	if virtualNetworksSpecGenerator != nil {
		return virtualNetworksSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpec(generators)
	virtualNetworksSpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworksSpec(generators)
	AddRelatedPropertyGeneratorsForVirtualNetworksSpec(generators)
	virtualNetworksSpecGenerator = gen.Struct(reflect.TypeOf(VirtualNetworks_Spec{}), generators)

	return virtualNetworksSpecGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworksSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworksSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EnableDdosProtection"] = gen.PtrOf(gen.Bool())
	gens["EnableVmProtection"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForVirtualNetworksSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForVirtualNetworksSpec(gens map[string]gopter.Gen) {
	gens["AddressSpace"] = gen.PtrOf(AddressSpaceGenerator())
	gens["BgpCommunities"] = gen.PtrOf(VirtualNetworkBgpCommunitiesGenerator())
	gens["DdosProtectionPlan"] = gen.PtrOf(SubResourceGenerator())
	gens["DhcpOptions"] = gen.PtrOf(DhcpOptionsGenerator())
	gens["ExtendedLocation"] = gen.PtrOf(ExtendedLocationGenerator())
	gens["IpAllocations"] = gen.SliceOf(SubResourceGenerator())
}

func Test_AddressSpace_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AddressSpace to AddressSpace via AssignPropertiesToAddressSpace & AssignPropertiesFromAddressSpace returns original",
		prop.ForAll(RunPropertyAssignmentTestForAddressSpace, AddressSpaceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAddressSpace tests if a specific instance of AddressSpace can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForAddressSpace(subject AddressSpace) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201101storage.AddressSpace
	err := copied.AssignPropertiesToAddressSpace(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AddressSpace
	err = actual.AssignPropertiesFromAddressSpace(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AddressSpace_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AddressSpace via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAddressSpace, AddressSpaceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAddressSpace runs a test to see if a specific instance of AddressSpace round trips to JSON and back losslessly
func RunJSONSerializationTestForAddressSpace(subject AddressSpace) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AddressSpace
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

// Generator of AddressSpace instances for property testing - lazily instantiated by AddressSpaceGenerator()
var addressSpaceGenerator gopter.Gen

// AddressSpaceGenerator returns a generator of AddressSpace instances for property testing.
func AddressSpaceGenerator() gopter.Gen {
	if addressSpaceGenerator != nil {
		return addressSpaceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAddressSpace(generators)
	addressSpaceGenerator = gen.Struct(reflect.TypeOf(AddressSpace{}), generators)

	return addressSpaceGenerator
}

// AddIndependentPropertyGeneratorsForAddressSpace is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAddressSpace(gens map[string]gopter.Gen) {
	gens["AddressPrefixes"] = gen.SliceOf(gen.AlphaString())
}

func Test_AddressSpace_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from AddressSpace_Status to AddressSpace_Status via AssignPropertiesToAddressSpaceStatus & AssignPropertiesFromAddressSpaceStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForAddressSpaceStatus, AddressSpaceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForAddressSpaceStatus tests if a specific instance of AddressSpace_Status can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForAddressSpaceStatus(subject AddressSpace_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201101storage.AddressSpace_Status
	err := copied.AssignPropertiesToAddressSpaceStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual AddressSpace_Status
	err = actual.AssignPropertiesFromAddressSpaceStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_AddressSpace_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AddressSpace_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAddressSpaceStatus, AddressSpaceStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAddressSpaceStatus runs a test to see if a specific instance of AddressSpace_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForAddressSpaceStatus(subject AddressSpace_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AddressSpace_Status
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

// Generator of AddressSpace_Status instances for property testing - lazily instantiated by AddressSpaceStatusGenerator()
var addressSpaceStatusGenerator gopter.Gen

// AddressSpaceStatusGenerator returns a generator of AddressSpace_Status instances for property testing.
func AddressSpaceStatusGenerator() gopter.Gen {
	if addressSpaceStatusGenerator != nil {
		return addressSpaceStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAddressSpaceStatus(generators)
	addressSpaceStatusGenerator = gen.Struct(reflect.TypeOf(AddressSpace_Status{}), generators)

	return addressSpaceStatusGenerator
}

// AddIndependentPropertyGeneratorsForAddressSpaceStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAddressSpaceStatus(gens map[string]gopter.Gen) {
	gens["AddressPrefixes"] = gen.SliceOf(gen.AlphaString())
}

func Test_DhcpOptions_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DhcpOptions to DhcpOptions via AssignPropertiesToDhcpOptions & AssignPropertiesFromDhcpOptions returns original",
		prop.ForAll(RunPropertyAssignmentTestForDhcpOptions, DhcpOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDhcpOptions tests if a specific instance of DhcpOptions can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForDhcpOptions(subject DhcpOptions) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201101storage.DhcpOptions
	err := copied.AssignPropertiesToDhcpOptions(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DhcpOptions
	err = actual.AssignPropertiesFromDhcpOptions(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DhcpOptions_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptions via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptions, DhcpOptionsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptions runs a test to see if a specific instance of DhcpOptions round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptions(subject DhcpOptions) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptions
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

// Generator of DhcpOptions instances for property testing - lazily instantiated by DhcpOptionsGenerator()
var dhcpOptionsGenerator gopter.Gen

// DhcpOptionsGenerator returns a generator of DhcpOptions instances for property testing.
func DhcpOptionsGenerator() gopter.Gen {
	if dhcpOptionsGenerator != nil {
		return dhcpOptionsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptions(generators)
	dhcpOptionsGenerator = gen.Struct(reflect.TypeOf(DhcpOptions{}), generators)

	return dhcpOptionsGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptions is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptions(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}

func Test_DhcpOptions_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DhcpOptions_Status to DhcpOptions_Status via AssignPropertiesToDhcpOptionsStatus & AssignPropertiesFromDhcpOptionsStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForDhcpOptionsStatus, DhcpOptionsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDhcpOptionsStatus tests if a specific instance of DhcpOptions_Status can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForDhcpOptionsStatus(subject DhcpOptions_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201101storage.DhcpOptions_Status
	err := copied.AssignPropertiesToDhcpOptionsStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DhcpOptions_Status
	err = actual.AssignPropertiesFromDhcpOptionsStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DhcpOptions_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DhcpOptions_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDhcpOptionsStatus, DhcpOptionsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDhcpOptionsStatus runs a test to see if a specific instance of DhcpOptions_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForDhcpOptionsStatus(subject DhcpOptions_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DhcpOptions_Status
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

// Generator of DhcpOptions_Status instances for property testing - lazily instantiated by DhcpOptionsStatusGenerator()
var dhcpOptionsStatusGenerator gopter.Gen

// DhcpOptionsStatusGenerator returns a generator of DhcpOptions_Status instances for property testing.
func DhcpOptionsStatusGenerator() gopter.Gen {
	if dhcpOptionsStatusGenerator != nil {
		return dhcpOptionsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDhcpOptionsStatus(generators)
	dhcpOptionsStatusGenerator = gen.Struct(reflect.TypeOf(DhcpOptions_Status{}), generators)

	return dhcpOptionsStatusGenerator
}

// AddIndependentPropertyGeneratorsForDhcpOptionsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDhcpOptionsStatus(gens map[string]gopter.Gen) {
	gens["DnsServers"] = gen.SliceOf(gen.AlphaString())
}

func Test_VirtualNetworkBgpCommunities_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworkBgpCommunities to VirtualNetworkBgpCommunities via AssignPropertiesToVirtualNetworkBgpCommunities & AssignPropertiesFromVirtualNetworkBgpCommunities returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworkBgpCommunities, VirtualNetworkBgpCommunitiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworkBgpCommunities tests if a specific instance of VirtualNetworkBgpCommunities can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworkBgpCommunities(subject VirtualNetworkBgpCommunities) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201101storage.VirtualNetworkBgpCommunities
	err := copied.AssignPropertiesToVirtualNetworkBgpCommunities(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworkBgpCommunities
	err = actual.AssignPropertiesFromVirtualNetworkBgpCommunities(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworkBgpCommunities_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunities, VirtualNetworkBgpCommunitiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunities runs a test to see if a specific instance of VirtualNetworkBgpCommunities round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunities(subject VirtualNetworkBgpCommunities) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities
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

// Generator of VirtualNetworkBgpCommunities instances for property testing - lazily instantiated by
//VirtualNetworkBgpCommunitiesGenerator()
var virtualNetworkBgpCommunitiesGenerator gopter.Gen

// VirtualNetworkBgpCommunitiesGenerator returns a generator of VirtualNetworkBgpCommunities instances for property testing.
func VirtualNetworkBgpCommunitiesGenerator() gopter.Gen {
	if virtualNetworkBgpCommunitiesGenerator != nil {
		return virtualNetworkBgpCommunitiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities(generators)
	virtualNetworkBgpCommunitiesGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities{}), generators)

	return virtualNetworkBgpCommunitiesGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunities(gens map[string]gopter.Gen) {
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}

func Test_VirtualNetworkBgpCommunities_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from VirtualNetworkBgpCommunities_Status to VirtualNetworkBgpCommunities_Status via AssignPropertiesToVirtualNetworkBgpCommunitiesStatus & AssignPropertiesFromVirtualNetworkBgpCommunitiesStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForVirtualNetworkBgpCommunitiesStatus, VirtualNetworkBgpCommunitiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForVirtualNetworkBgpCommunitiesStatus tests if a specific instance of VirtualNetworkBgpCommunities_Status can be assigned to v1beta20201101storage and back losslessly
func RunPropertyAssignmentTestForVirtualNetworkBgpCommunitiesStatus(subject VirtualNetworkBgpCommunities_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1beta20201101storage.VirtualNetworkBgpCommunities_Status
	err := copied.AssignPropertiesToVirtualNetworkBgpCommunitiesStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual VirtualNetworkBgpCommunities_Status
	err = actual.AssignPropertiesFromVirtualNetworkBgpCommunitiesStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_VirtualNetworkBgpCommunities_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of VirtualNetworkBgpCommunities_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForVirtualNetworkBgpCommunitiesStatus, VirtualNetworkBgpCommunitiesStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForVirtualNetworkBgpCommunitiesStatus runs a test to see if a specific instance of VirtualNetworkBgpCommunities_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForVirtualNetworkBgpCommunitiesStatus(subject VirtualNetworkBgpCommunities_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual VirtualNetworkBgpCommunities_Status
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

// Generator of VirtualNetworkBgpCommunities_Status instances for property testing - lazily instantiated by
//VirtualNetworkBgpCommunitiesStatusGenerator()
var virtualNetworkBgpCommunitiesStatusGenerator gopter.Gen

// VirtualNetworkBgpCommunitiesStatusGenerator returns a generator of VirtualNetworkBgpCommunities_Status instances for property testing.
func VirtualNetworkBgpCommunitiesStatusGenerator() gopter.Gen {
	if virtualNetworkBgpCommunitiesStatusGenerator != nil {
		return virtualNetworkBgpCommunitiesStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesStatus(generators)
	virtualNetworkBgpCommunitiesStatusGenerator = gen.Struct(reflect.TypeOf(VirtualNetworkBgpCommunities_Status{}), generators)

	return virtualNetworkBgpCommunitiesStatusGenerator
}

// AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForVirtualNetworkBgpCommunitiesStatus(gens map[string]gopter.Gen) {
	gens["RegionalCommunity"] = gen.PtrOf(gen.AlphaString())
	gens["VirtualNetworkCommunity"] = gen.PtrOf(gen.AlphaString())
}
