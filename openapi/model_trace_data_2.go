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

type TraceData2 struct {

	TraceRef string `json:"traceRef"`

	TraceDepth TraceDepth2 `json:"traceDepth"`

	NeTypeList string `json:"neTypeList"`

	EventList string `json:"eventList"`

	CollectionEntityIpv4Addr string `json:"collectionEntityIpv4Addr,omitempty"`

	CollectionEntityIpv6Addr Ipv6Addr2 `json:"collectionEntityIpv6Addr,omitempty"`

	InterfaceList string `json:"interfaceList,omitempty"`
}
