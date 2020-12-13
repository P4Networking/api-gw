/*
 * Edgecore 5GC
 *
 * Edgecore 5GC API server.  You can find out more about Edgecore 5GC at [http://5gc.edge-core.cn](http://5gc.edge-core.cn). 
 *
 * API version: 1.0.0
 * Contact: jimmy_ou@edge-core.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type SmsManagementSubscriptionData struct {

	SupportedFeatures string `json:"supportedFeatures,omitempty"`

	MtSmsSubscribed bool `json:"mtSmsSubscribed,omitempty"`

	MtSmsBarringAll bool `json:"mtSmsBarringAll,omitempty"`

	MtSmsBarringRoaming bool `json:"mtSmsBarringRoaming,omitempty"`

	MoSmsSubscribed bool `json:"moSmsSubscribed,omitempty"`

	MoSmsBarringAll bool `json:"moSmsBarringAll,omitempty"`

	MoSmsBarringRoaming bool `json:"moSmsBarringRoaming,omitempty"`

	SharedSmsMngDataIds []string `json:"sharedSmsMngDataIds,omitempty"`
}