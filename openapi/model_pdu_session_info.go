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

type PduSessionInfo struct {

	Supi string `json:"supi,omitempty"`

	PDUSessionID string `json:"pDUSessionID,omitempty"`

	Dnn string `json:"dnn,omitempty"`

	Sst string `json:"sst,omitempty"`

	Sd string `json:"sd,omitempty"`

	AnType string `json:"anType,omitempty"`

	PDUAddress string `json:"pDUAddress,omitempty"`

	SessionRule *SessionRule `json:"sessionRule,omitempty"`

	UpCnxState string `json:"upCnxState,omitempty"`
}