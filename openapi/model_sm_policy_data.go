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

// SmPolicyData - Contains the SM policy data for a given subscriber.
type SmPolicyData struct {

	SmPolicySnssaiData map[string]SmPolicySnssaiData `json:"smPolicySnssaiData"`

	UmDataLimits map[string]UsageMonDataLimit `json:"umDataLimits,omitempty"`

	UmData map[string]UsageMonData `json:"umData,omitempty"`
}
