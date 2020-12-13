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

type DnnConfiguration struct {

	PduSessionTypes PduSessionTypes `json:"pduSessionTypes"`

	SscModes SscModes `json:"sscModes"`

	IwkEpsInd bool `json:"iwkEpsInd,omitempty"`

	LadnIndicator bool `json:"ladnIndicator,omitempty"`

	Var5gQosProfile SubscribedDefaultQos `json:"5gQosProfile,omitempty"`

	SessionAmbr Ambr `json:"sessionAmbr,omitempty"`

	Var3gppChargingCharacteristics string `json:"3gppChargingCharacteristics,omitempty"`

	StaticIpAddress []IpAddress `json:"staticIpAddress,omitempty"`

	UpSecurity UpSecurity `json:"upSecurity,omitempty"`
}