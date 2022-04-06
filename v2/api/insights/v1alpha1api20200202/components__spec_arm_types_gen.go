// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200202

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of Components_Spec. Use v1beta20200202.Components_Spec instead
type Components_SpecARM struct {
	Etag       *string                                    `json:"etag,omitempty"`
	Kind       *string                                    `json:"kind,omitempty"`
	Location   *string                                    `json:"location,omitempty"`
	Name       string                                     `json:"name,omitempty"`
	Properties *ApplicationInsightsComponentPropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                          `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Components_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-02-02"
func (components Components_SpecARM) GetAPIVersion() string {
	return "2020-02-02"
}

// GetName returns the Name of the resource
func (components Components_SpecARM) GetName() string {
	return components.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/components"
func (components Components_SpecARM) GetType() string {
	return "Microsoft.Insights/components"
}

//Deprecated version of ApplicationInsightsComponentProperties. Use v1beta20200202.ApplicationInsightsComponentProperties instead
type ApplicationInsightsComponentPropertiesARM struct {
	ApplicationType                 *ApplicationInsightsComponentPropertiesApplicationType                 `json:"Application_Type,omitempty"`
	DisableIpMasking                *bool                                                                  `json:"DisableIpMasking,omitempty"`
	DisableLocalAuth                *bool                                                                  `json:"DisableLocalAuth,omitempty"`
	FlowType                        *ApplicationInsightsComponentPropertiesFlowType                        `json:"Flow_Type,omitempty"`
	ForceCustomerStorageForProfiler *bool                                                                  `json:"ForceCustomerStorageForProfiler,omitempty"`
	HockeyAppId                     *string                                                                `json:"HockeyAppId,omitempty"`
	ImmediatePurgeDataOn30Days      *bool                                                                  `json:"ImmediatePurgeDataOn30Days,omitempty"`
	IngestionMode                   *ApplicationInsightsComponentPropertiesIngestionMode                   `json:"IngestionMode,omitempty"`
	PublicNetworkAccessForIngestion *ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery     `json:"publicNetworkAccessForQuery,omitempty"`
	RequestSource                   *ApplicationInsightsComponentPropertiesRequestSource                   `json:"Request_Source,omitempty"`
	RetentionInDays                 *int                                                                   `json:"RetentionInDays,omitempty"`
	SamplingPercentage              *float64                                                               `json:"SamplingPercentage,omitempty"`
	WorkspaceResourceId             *string                                                                `json:"workspaceResourceId,omitempty"`
}
