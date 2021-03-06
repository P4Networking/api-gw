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

type GlobalRanNodeId struct {

	PlmnId PlmnId `json:"plmnId"`

	N3IwfId string `json:"n3IwfId,omitempty"`

	GNbId GNbId `json:"gNbId,omitempty"`

	NgeNbId string `json:"ngeNbId,omitempty"`

	WagfId string `json:"wagfId,omitempty"`

	TngfId string `json:"tngfId,omitempty"`

	Nid string `json:"nid,omitempty"`

	ENbId string `json:"eNbId,omitempty"`
}
