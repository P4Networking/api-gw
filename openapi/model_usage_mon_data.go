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

// UsageMonData - Contains remain allowed usage data for a subscriber.
type UsageMonData struct {

	LimitId string `json:"limitId"`

	Scopes map[string]UsageMonDataScope `json:"scopes,omitempty"`

	UmLevel UsageMonLevel `json:"umLevel,omitempty"`

	AllowedUsage UsageThreshold `json:"allowedUsage,omitempty"`

	ResetTime TimePeriod `json:"resetTime,omitempty"`
}