// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

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

func Test_AgentPool_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPool_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolStatusARM, AgentPoolStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolStatusARM runs a test to see if a specific instance of AgentPool_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolStatusARM(subject AgentPool_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPool_StatusARM
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

// Generator of AgentPool_StatusARM instances for property testing - lazily instantiated by AgentPoolStatusARMGenerator()
var agentPoolStatusARMGenerator gopter.Gen

// AgentPoolStatusARMGenerator returns a generator of AgentPool_StatusARM instances for property testing.
// We first initialize agentPoolStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AgentPoolStatusARMGenerator() gopter.Gen {
	if agentPoolStatusARMGenerator != nil {
		return agentPoolStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolStatusARM(generators)
	agentPoolStatusARMGenerator = gen.Struct(reflect.TypeOf(AgentPool_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolStatusARM(generators)
	AddRelatedPropertyGeneratorsForAgentPoolStatusARM(generators)
	agentPoolStatusARMGenerator = gen.Struct(reflect.TypeOf(AgentPool_StatusARM{}), generators)

	return agentPoolStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForAgentPoolStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAgentPoolStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ManagedClusterAgentPoolProfilePropertiesStatusARMGenerator())
}

func Test_ManagedClusterAgentPoolProfileProperties_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ManagedClusterAgentPoolProfileProperties_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForManagedClusterAgentPoolProfilePropertiesStatusARM, ManagedClusterAgentPoolProfilePropertiesStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForManagedClusterAgentPoolProfilePropertiesStatusARM runs a test to see if a specific instance of ManagedClusterAgentPoolProfileProperties_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForManagedClusterAgentPoolProfilePropertiesStatusARM(subject ManagedClusterAgentPoolProfileProperties_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ManagedClusterAgentPoolProfileProperties_StatusARM
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

// Generator of ManagedClusterAgentPoolProfileProperties_StatusARM instances for property testing - lazily instantiated
// by ManagedClusterAgentPoolProfilePropertiesStatusARMGenerator()
var managedClusterAgentPoolProfilePropertiesStatusARMGenerator gopter.Gen

// ManagedClusterAgentPoolProfilePropertiesStatusARMGenerator returns a generator of ManagedClusterAgentPoolProfileProperties_StatusARM instances for property testing.
// We first initialize managedClusterAgentPoolProfilePropertiesStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ManagedClusterAgentPoolProfilePropertiesStatusARMGenerator() gopter.Gen {
	if managedClusterAgentPoolProfilePropertiesStatusARMGenerator != nil {
		return managedClusterAgentPoolProfilePropertiesStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesStatusARM(generators)
	managedClusterAgentPoolProfilePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesStatusARM(generators)
	AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesStatusARM(generators)
	managedClusterAgentPoolProfilePropertiesStatusARMGenerator = gen.Struct(reflect.TypeOf(ManagedClusterAgentPoolProfileProperties_StatusARM{}), generators)

	return managedClusterAgentPoolProfilePropertiesStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["AvailabilityZones"] = gen.SliceOf(gen.AlphaString())
	gens["Count"] = gen.PtrOf(gen.Int())
	gens["EnableAutoScaling"] = gen.PtrOf(gen.Bool())
	gens["EnableEncryptionAtHost"] = gen.PtrOf(gen.Bool())
	gens["EnableFIPS"] = gen.PtrOf(gen.Bool())
	gens["EnableNodePublicIP"] = gen.PtrOf(gen.Bool())
	gens["EnableUltraSSD"] = gen.PtrOf(gen.Bool())
	gens["GpuInstanceProfile"] = gen.PtrOf(gen.OneConstOf(
		GPUInstanceProfile_Status_MIG1G,
		GPUInstanceProfile_Status_MIG2G,
		GPUInstanceProfile_Status_MIG3G,
		GPUInstanceProfile_Status_MIG4G,
		GPUInstanceProfile_Status_MIG7G))
	gens["KubeletDiskType"] = gen.PtrOf(gen.OneConstOf(KubeletDiskType_Status_OS, KubeletDiskType_Status_Temporary))
	gens["MaxCount"] = gen.PtrOf(gen.Int())
	gens["MaxPods"] = gen.PtrOf(gen.Int())
	gens["MinCount"] = gen.PtrOf(gen.Int())
	gens["Mode"] = gen.PtrOf(gen.OneConstOf(AgentPoolMode_Status_System, AgentPoolMode_Status_User))
	gens["NodeImageVersion"] = gen.PtrOf(gen.AlphaString())
	gens["NodeLabels"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["NodePublicIPPrefixID"] = gen.PtrOf(gen.AlphaString())
	gens["NodeTaints"] = gen.SliceOf(gen.AlphaString())
	gens["OrchestratorVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OsDiskSizeGB"] = gen.PtrOf(gen.Int())
	gens["OsDiskType"] = gen.PtrOf(gen.OneConstOf(OSDiskType_Status_Ephemeral, OSDiskType_Status_Managed))
	gens["OsSKU"] = gen.PtrOf(gen.OneConstOf(OSSKU_Status_CBLMariner, OSSKU_Status_Ubuntu))
	gens["OsType"] = gen.PtrOf(gen.OneConstOf(OSType_Status_Linux, OSType_Status_Windows))
	gens["PodSubnetID"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["ProximityPlacementGroupID"] = gen.PtrOf(gen.AlphaString())
	gens["ScaleSetEvictionPolicy"] = gen.PtrOf(gen.OneConstOf(ScaleSetEvictionPolicy_Status_Deallocate, ScaleSetEvictionPolicy_Status_Delete))
	gens["ScaleSetPriority"] = gen.PtrOf(gen.OneConstOf(ScaleSetPriority_Status_Regular, ScaleSetPriority_Status_Spot))
	gens["SpotMaxPrice"] = gen.PtrOf(gen.Float64())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.OneConstOf(AgentPoolType_Status_AvailabilitySet, AgentPoolType_Status_VirtualMachineScaleSets))
	gens["VmSize"] = gen.PtrOf(gen.AlphaString())
	gens["VnetSubnetID"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForManagedClusterAgentPoolProfilePropertiesStatusARM(gens map[string]gopter.Gen) {
	gens["KubeletConfig"] = gen.PtrOf(KubeletConfigStatusARMGenerator())
	gens["LinuxOSConfig"] = gen.PtrOf(LinuxOSConfigStatusARMGenerator())
	gens["PowerState"] = gen.PtrOf(PowerStateStatusARMGenerator())
	gens["UpgradeSettings"] = gen.PtrOf(AgentPoolUpgradeSettingsStatusARMGenerator())
}

func Test_AgentPoolUpgradeSettings_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AgentPoolUpgradeSettings_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAgentPoolUpgradeSettingsStatusARM, AgentPoolUpgradeSettingsStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAgentPoolUpgradeSettingsStatusARM runs a test to see if a specific instance of AgentPoolUpgradeSettings_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAgentPoolUpgradeSettingsStatusARM(subject AgentPoolUpgradeSettings_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AgentPoolUpgradeSettings_StatusARM
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

// Generator of AgentPoolUpgradeSettings_StatusARM instances for property testing - lazily instantiated by
// AgentPoolUpgradeSettingsStatusARMGenerator()
var agentPoolUpgradeSettingsStatusARMGenerator gopter.Gen

// AgentPoolUpgradeSettingsStatusARMGenerator returns a generator of AgentPoolUpgradeSettings_StatusARM instances for property testing.
func AgentPoolUpgradeSettingsStatusARMGenerator() gopter.Gen {
	if agentPoolUpgradeSettingsStatusARMGenerator != nil {
		return agentPoolUpgradeSettingsStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettingsStatusARM(generators)
	agentPoolUpgradeSettingsStatusARMGenerator = gen.Struct(reflect.TypeOf(AgentPoolUpgradeSettings_StatusARM{}), generators)

	return agentPoolUpgradeSettingsStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettingsStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAgentPoolUpgradeSettingsStatusARM(gens map[string]gopter.Gen) {
	gens["MaxSurge"] = gen.PtrOf(gen.AlphaString())
}

func Test_KubeletConfig_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of KubeletConfig_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForKubeletConfigStatusARM, KubeletConfigStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForKubeletConfigStatusARM runs a test to see if a specific instance of KubeletConfig_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForKubeletConfigStatusARM(subject KubeletConfig_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual KubeletConfig_StatusARM
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

// Generator of KubeletConfig_StatusARM instances for property testing - lazily instantiated by
// KubeletConfigStatusARMGenerator()
var kubeletConfigStatusARMGenerator gopter.Gen

// KubeletConfigStatusARMGenerator returns a generator of KubeletConfig_StatusARM instances for property testing.
func KubeletConfigStatusARMGenerator() gopter.Gen {
	if kubeletConfigStatusARMGenerator != nil {
		return kubeletConfigStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForKubeletConfigStatusARM(generators)
	kubeletConfigStatusARMGenerator = gen.Struct(reflect.TypeOf(KubeletConfig_StatusARM{}), generators)

	return kubeletConfigStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForKubeletConfigStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForKubeletConfigStatusARM(gens map[string]gopter.Gen) {
	gens["AllowedUnsafeSysctls"] = gen.SliceOf(gen.AlphaString())
	gens["ContainerLogMaxFiles"] = gen.PtrOf(gen.Int())
	gens["ContainerLogMaxSizeMB"] = gen.PtrOf(gen.Int())
	gens["CpuCfsQuota"] = gen.PtrOf(gen.Bool())
	gens["CpuCfsQuotaPeriod"] = gen.PtrOf(gen.AlphaString())
	gens["CpuManagerPolicy"] = gen.PtrOf(gen.AlphaString())
	gens["FailSwapOn"] = gen.PtrOf(gen.Bool())
	gens["ImageGcHighThreshold"] = gen.PtrOf(gen.Int())
	gens["ImageGcLowThreshold"] = gen.PtrOf(gen.Int())
	gens["PodMaxPids"] = gen.PtrOf(gen.Int())
	gens["TopologyManagerPolicy"] = gen.PtrOf(gen.AlphaString())
}

func Test_LinuxOSConfig_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinuxOSConfig_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinuxOSConfigStatusARM, LinuxOSConfigStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinuxOSConfigStatusARM runs a test to see if a specific instance of LinuxOSConfig_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLinuxOSConfigStatusARM(subject LinuxOSConfig_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinuxOSConfig_StatusARM
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

// Generator of LinuxOSConfig_StatusARM instances for property testing - lazily instantiated by
// LinuxOSConfigStatusARMGenerator()
var linuxOSConfigStatusARMGenerator gopter.Gen

// LinuxOSConfigStatusARMGenerator returns a generator of LinuxOSConfig_StatusARM instances for property testing.
// We first initialize linuxOSConfigStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LinuxOSConfigStatusARMGenerator() gopter.Gen {
	if linuxOSConfigStatusARMGenerator != nil {
		return linuxOSConfigStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfigStatusARM(generators)
	linuxOSConfigStatusARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinuxOSConfigStatusARM(generators)
	AddRelatedPropertyGeneratorsForLinuxOSConfigStatusARM(generators)
	linuxOSConfigStatusARMGenerator = gen.Struct(reflect.TypeOf(LinuxOSConfig_StatusARM{}), generators)

	return linuxOSConfigStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForLinuxOSConfigStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinuxOSConfigStatusARM(gens map[string]gopter.Gen) {
	gens["SwapFileSizeMB"] = gen.PtrOf(gen.Int())
	gens["TransparentHugePageDefrag"] = gen.PtrOf(gen.AlphaString())
	gens["TransparentHugePageEnabled"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForLinuxOSConfigStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLinuxOSConfigStatusARM(gens map[string]gopter.Gen) {
	gens["Sysctls"] = gen.PtrOf(SysctlConfigStatusARMGenerator())
}

func Test_PowerState_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PowerState_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPowerStateStatusARM, PowerStateStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPowerStateStatusARM runs a test to see if a specific instance of PowerState_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPowerStateStatusARM(subject PowerState_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PowerState_StatusARM
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

// Generator of PowerState_StatusARM instances for property testing - lazily instantiated by
// PowerStateStatusARMGenerator()
var powerStateStatusARMGenerator gopter.Gen

// PowerStateStatusARMGenerator returns a generator of PowerState_StatusARM instances for property testing.
func PowerStateStatusARMGenerator() gopter.Gen {
	if powerStateStatusARMGenerator != nil {
		return powerStateStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPowerStateStatusARM(generators)
	powerStateStatusARMGenerator = gen.Struct(reflect.TypeOf(PowerState_StatusARM{}), generators)

	return powerStateStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForPowerStateStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPowerStateStatusARM(gens map[string]gopter.Gen) {
	gens["Code"] = gen.PtrOf(gen.OneConstOf(PowerStateStatusCode_Running, PowerStateStatusCode_Stopped))
}

func Test_SysctlConfig_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SysctlConfig_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSysctlConfigStatusARM, SysctlConfigStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSysctlConfigStatusARM runs a test to see if a specific instance of SysctlConfig_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSysctlConfigStatusARM(subject SysctlConfig_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SysctlConfig_StatusARM
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

// Generator of SysctlConfig_StatusARM instances for property testing - lazily instantiated by
// SysctlConfigStatusARMGenerator()
var sysctlConfigStatusARMGenerator gopter.Gen

// SysctlConfigStatusARMGenerator returns a generator of SysctlConfig_StatusARM instances for property testing.
func SysctlConfigStatusARMGenerator() gopter.Gen {
	if sysctlConfigStatusARMGenerator != nil {
		return sysctlConfigStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSysctlConfigStatusARM(generators)
	sysctlConfigStatusARMGenerator = gen.Struct(reflect.TypeOf(SysctlConfig_StatusARM{}), generators)

	return sysctlConfigStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForSysctlConfigStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSysctlConfigStatusARM(gens map[string]gopter.Gen) {
	gens["FsAioMaxNr"] = gen.PtrOf(gen.Int())
	gens["FsFileMax"] = gen.PtrOf(gen.Int())
	gens["FsInotifyMaxUserWatches"] = gen.PtrOf(gen.Int())
	gens["FsNrOpen"] = gen.PtrOf(gen.Int())
	gens["KernelThreadsMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreNetdevMaxBacklog"] = gen.PtrOf(gen.Int())
	gens["NetCoreOptmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreRmemMax"] = gen.PtrOf(gen.Int())
	gens["NetCoreSomaxconn"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemDefault"] = gen.PtrOf(gen.Int())
	gens["NetCoreWmemMax"] = gen.PtrOf(gen.Int())
	gens["NetIpv4IpLocalPortRange"] = gen.PtrOf(gen.AlphaString())
	gens["NetIpv4NeighDefaultGcThresh1"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh2"] = gen.PtrOf(gen.Int())
	gens["NetIpv4NeighDefaultGcThresh3"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpFinTimeout"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveProbes"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpKeepaliveTime"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxSynBacklog"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpMaxTwBuckets"] = gen.PtrOf(gen.Int())
	gens["NetIpv4TcpTwReuse"] = gen.PtrOf(gen.Bool())
	gens["NetIpv4TcpkeepaliveIntvl"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackBuckets"] = gen.PtrOf(gen.Int())
	gens["NetNetfilterNfConntrackMax"] = gen.PtrOf(gen.Int())
	gens["VmMaxMapCount"] = gen.PtrOf(gen.Int())
	gens["VmSwappiness"] = gen.PtrOf(gen.Int())
	gens["VmVfsCachePressure"] = gen.PtrOf(gen.Int())
}
