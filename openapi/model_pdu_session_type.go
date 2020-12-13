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

type PduSessionType string

// List of PduSessionType
const (
	IPV4 PduSessionType = "IPV4"
	IPV6 PduSessionType = "IPV6"
	IPV4_V6 PduSessionType = "IPV4V6"
	UNSTRUCTURED PduSessionType = "UNSTRUCTURED"
	ETHERNET PduSessionType = "ETHERNET"
)
