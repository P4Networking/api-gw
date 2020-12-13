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

type Arp struct {

	PriorityLevel *int32 `json:"priorityLevel"`

	PreemptCap PreemptionCapability `json:"preemptCap"`

	PreemptVuln PreemptionVulnerability `json:"preemptVuln"`
}
