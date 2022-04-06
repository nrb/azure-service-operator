// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200202

import "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

//Deprecated version of ApplicationInsightsComponent_Status. Use v1beta20200202.ApplicationInsightsComponent_Status instead
type ApplicationInsightsComponent_StatusARM struct {
	Etag       *string                                           `json:"etag,omitempty"`
	Id         *string                                           `json:"id,omitempty"`
	Kind       *string                                           `json:"kind,omitempty"`
	Location   *string                                           `json:"location,omitempty"`
	Name       *string                                           `json:"name,omitempty"`
	Properties *ApplicationInsightsComponentProperties_StatusARM `json:"properties,omitempty"`
	Tags       *v1.JSON                                          `json:"tags,omitempty"`
	Type       *string                                           `json:"type,omitempty"`
}

//Deprecated version of ApplicationInsightsComponentProperties_Status. Use v1beta20200202.ApplicationInsightsComponentProperties_Status instead
type ApplicationInsightsComponentProperties_StatusARM struct {
	AppId                           *string                                                      `json:"AppId,omitempty"`
	ApplicationId                   *string                                                      `json:"ApplicationId,omitempty"`
	ApplicationType                 *ApplicationInsightsComponentPropertiesStatusApplicationType `json:"Application_Type,omitempty"`
	ConnectionString                *string                                                      `json:"ConnectionString,omitempty"`
	CreationDate                    *string                                                      `json:"CreationDate,omitempty"`
	DisableIpMasking                *bool                                                        `json:"DisableIpMasking,omitempty"`
	DisableLocalAuth                *bool                                                        `json:"DisableLocalAuth,omitempty"`
	FlowType                        *ApplicationInsightsComponentPropertiesStatusFlowType        `json:"Flow_Type,omitempty"`
	ForceCustomerStorageForProfiler *bool                                                        `json:"ForceCustomerStorageForProfiler,omitempty"`
	HockeyAppId                     *string                                                      `json:"HockeyAppId,omitempty"`
	HockeyAppToken                  *string                                                      `json:"HockeyAppToken,omitempty"`
	ImmediatePurgeDataOn30Days      *bool                                                        `json:"ImmediatePurgeDataOn30Days,omitempty"`
	IngestionMode                   *ApplicationInsightsComponentPropertiesStatusIngestionMode   `json:"IngestionMode,omitempty"`
	InstrumentationKey              *string                                                      `json:"InstrumentationKey,omitempty"`
	LaMigrationDate                 *string                                                      `json:"LaMigrationDate,omitempty"`
	Name                            *string                                                      `json:"Name,omitempty"`
	PrivateLinkScopedResources      []PrivateLinkScopedResource_StatusARM                        `json:"PrivateLinkScopedResources,omitempty"`
	ProvisioningState               *string                                                      `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *PublicNetworkAccessType_Status                              `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *PublicNetworkAccessType_Status                              `json:"publicNetworkAccessForQuery,omitempty"`
	RequestSource                   *ApplicationInsightsComponentPropertiesStatusRequestSource   `json:"Request_Source,omitempty"`
	RetentionInDays                 *int                                                         `json:"RetentionInDays,omitempty"`
	SamplingPercentage              *float64                                                     `json:"SamplingPercentage,omitempty"`
	TenantId                        *string                                                      `json:"TenantId,omitempty"`
	WorkspaceResourceId             *string                                                      `json:"WorkspaceResourceId,omitempty"`
}

//Deprecated version of ApplicationInsightsComponentPropertiesStatusApplicationType. Use
//v1beta20200202.ApplicationInsightsComponentPropertiesStatusApplicationType instead
type ApplicationInsightsComponentPropertiesStatusApplicationType string

const (
	ApplicationInsightsComponentPropertiesStatusApplicationTypeOther = ApplicationInsightsComponentPropertiesStatusApplicationType("other")
	ApplicationInsightsComponentPropertiesStatusApplicationTypeWeb   = ApplicationInsightsComponentPropertiesStatusApplicationType("web")
)

//Deprecated version of ApplicationInsightsComponentPropertiesStatusFlowType. Use
//v1beta20200202.ApplicationInsightsComponentPropertiesStatusFlowType instead
type ApplicationInsightsComponentPropertiesStatusFlowType string

const ApplicationInsightsComponentPropertiesStatusFlowTypeBluefield = ApplicationInsightsComponentPropertiesStatusFlowType("Bluefield")

//Deprecated version of ApplicationInsightsComponentPropertiesStatusIngestionMode. Use
//v1beta20200202.ApplicationInsightsComponentPropertiesStatusIngestionMode instead
type ApplicationInsightsComponentPropertiesStatusIngestionMode string

const (
	ApplicationInsightsComponentPropertiesStatusIngestionModeApplicationInsights                       = ApplicationInsightsComponentPropertiesStatusIngestionMode("ApplicationInsights")
	ApplicationInsightsComponentPropertiesStatusIngestionModeApplicationInsightsWithDiagnosticSettings = ApplicationInsightsComponentPropertiesStatusIngestionMode("ApplicationInsightsWithDiagnosticSettings")
	ApplicationInsightsComponentPropertiesStatusIngestionModeLogAnalytics                              = ApplicationInsightsComponentPropertiesStatusIngestionMode("LogAnalytics")
)

//Deprecated version of ApplicationInsightsComponentPropertiesStatusRequestSource. Use
//v1beta20200202.ApplicationInsightsComponentPropertiesStatusRequestSource instead
type ApplicationInsightsComponentPropertiesStatusRequestSource string

const ApplicationInsightsComponentPropertiesStatusRequestSourceRest = ApplicationInsightsComponentPropertiesStatusRequestSource("rest")

//Deprecated version of PrivateLinkScopedResource_Status. Use v1beta20200202.PrivateLinkScopedResource_Status instead
type PrivateLinkScopedResource_StatusARM struct {
	ResourceId *string `json:"ResourceId,omitempty"`
	ScopeId    *string `json:"ScopeId,omitempty"`
}

//Deprecated version of PublicNetworkAccessType_Status. Use v1beta20200202.PublicNetworkAccessType_Status instead
type PublicNetworkAccessType_Status string

const (
	PublicNetworkAccessType_StatusDisabled = PublicNetworkAccessType_Status("Disabled")
	PublicNetworkAccessType_StatusEnabled  = PublicNetworkAccessType_Status("Enabled")
)
