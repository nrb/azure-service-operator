// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

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

func Test_RedisLinkedServers_SpecARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServers_SpecARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServersSpecARM, RedisLinkedServersSpecARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServersSpecARM runs a test to see if a specific instance of RedisLinkedServers_SpecARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServersSpecARM(subject RedisLinkedServers_SpecARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServers_SpecARM
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

// Generator of RedisLinkedServers_SpecARM instances for property testing - lazily instantiated by
// RedisLinkedServersSpecARMGenerator()
var redisLinkedServersSpecARMGenerator gopter.Gen

// RedisLinkedServersSpecARMGenerator returns a generator of RedisLinkedServers_SpecARM instances for property testing.
// We first initialize redisLinkedServersSpecARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisLinkedServersSpecARMGenerator() gopter.Gen {
	if redisLinkedServersSpecARMGenerator != nil {
		return redisLinkedServersSpecARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServersSpecARM(generators)
	redisLinkedServersSpecARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServers_SpecARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServersSpecARM(generators)
	AddRelatedPropertyGeneratorsForRedisLinkedServersSpecARM(generators)
	redisLinkedServersSpecARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServers_SpecARM{}), generators)

	return redisLinkedServersSpecARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServersSpecARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServersSpecARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisLinkedServersSpecARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisLinkedServersSpecARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RedisLinkedServerCreatePropertiesARMGenerator())
}

func Test_RedisLinkedServerCreatePropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServerCreatePropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerCreatePropertiesARM, RedisLinkedServerCreatePropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerCreatePropertiesARM runs a test to see if a specific instance of RedisLinkedServerCreatePropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerCreatePropertiesARM(subject RedisLinkedServerCreatePropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServerCreatePropertiesARM
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

// Generator of RedisLinkedServerCreatePropertiesARM instances for property testing - lazily instantiated by
// RedisLinkedServerCreatePropertiesARMGenerator()
var redisLinkedServerCreatePropertiesARMGenerator gopter.Gen

// RedisLinkedServerCreatePropertiesARMGenerator returns a generator of RedisLinkedServerCreatePropertiesARM instances for property testing.
func RedisLinkedServerCreatePropertiesARMGenerator() gopter.Gen {
	if redisLinkedServerCreatePropertiesARMGenerator != nil {
		return redisLinkedServerCreatePropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerCreatePropertiesARM(generators)
	redisLinkedServerCreatePropertiesARMGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServerCreatePropertiesARM{}), generators)

	return redisLinkedServerCreatePropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerCreatePropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerCreatePropertiesARM(gens map[string]gopter.Gen) {
	gens["LinkedRedisCacheId"] = gen.PtrOf(gen.AlphaString())
	gens["LinkedRedisCacheLocation"] = gen.PtrOf(gen.AlphaString())
	gens["ServerRole"] = gen.PtrOf(gen.OneConstOf(RedisLinkedServerCreatePropertiesServerRole_Primary, RedisLinkedServerCreatePropertiesServerRole_Secondary))
}
