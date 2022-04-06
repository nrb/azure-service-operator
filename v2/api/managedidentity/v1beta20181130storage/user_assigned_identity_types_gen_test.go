// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20181130storage

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

func Test_UserAssignedIdentity_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentity via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentity, UserAssignedIdentityGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentity runs a test to see if a specific instance of UserAssignedIdentity round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentity(subject UserAssignedIdentity) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentity
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

// Generator of UserAssignedIdentity instances for property testing - lazily instantiated by
//UserAssignedIdentityGenerator()
var userAssignedIdentityGenerator gopter.Gen

// UserAssignedIdentityGenerator returns a generator of UserAssignedIdentity instances for property testing.
func UserAssignedIdentityGenerator() gopter.Gen {
	if userAssignedIdentityGenerator != nil {
		return userAssignedIdentityGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForUserAssignedIdentity(generators)
	userAssignedIdentityGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentity{}), generators)

	return userAssignedIdentityGenerator
}

// AddRelatedPropertyGeneratorsForUserAssignedIdentity is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForUserAssignedIdentity(gens map[string]gopter.Gen) {
	gens["Spec"] = UserAssignedIdentitiesSpecGenerator()
	gens["Status"] = IdentityStatusGenerator()
}

func Test_Identity_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Identity_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForIdentityStatus, IdentityStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForIdentityStatus runs a test to see if a specific instance of Identity_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForIdentityStatus(subject Identity_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Identity_Status
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

// Generator of Identity_Status instances for property testing - lazily instantiated by IdentityStatusGenerator()
var identityStatusGenerator gopter.Gen

// IdentityStatusGenerator returns a generator of Identity_Status instances for property testing.
func IdentityStatusGenerator() gopter.Gen {
	if identityStatusGenerator != nil {
		return identityStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForIdentityStatus(generators)
	identityStatusGenerator = gen.Struct(reflect.TypeOf(Identity_Status{}), generators)

	return identityStatusGenerator
}

// AddIndependentPropertyGeneratorsForIdentityStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForIdentityStatus(gens map[string]gopter.Gen) {
	gens["ClientId"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_UserAssignedIdentities_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of UserAssignedIdentities_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForUserAssignedIdentitiesSpec, UserAssignedIdentitiesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForUserAssignedIdentitiesSpec runs a test to see if a specific instance of UserAssignedIdentities_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForUserAssignedIdentitiesSpec(subject UserAssignedIdentities_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual UserAssignedIdentities_Spec
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

// Generator of UserAssignedIdentities_Spec instances for property testing - lazily instantiated by
//UserAssignedIdentitiesSpecGenerator()
var userAssignedIdentitiesSpecGenerator gopter.Gen

// UserAssignedIdentitiesSpecGenerator returns a generator of UserAssignedIdentities_Spec instances for property testing.
func UserAssignedIdentitiesSpecGenerator() gopter.Gen {
	if userAssignedIdentitiesSpecGenerator != nil {
		return userAssignedIdentitiesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSpec(generators)
	userAssignedIdentitiesSpecGenerator = gen.Struct(reflect.TypeOf(UserAssignedIdentities_Spec{}), generators)

	return userAssignedIdentitiesSpecGenerator
}

// AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForUserAssignedIdentitiesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
