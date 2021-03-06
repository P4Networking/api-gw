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

import (
	"time"
)

// UsageMonDataLimit - Contains usage monitoring control data for a subscriber.
type UsageMonDataLimit struct {

	LimitId string `json:"limitId"`

	Scopes map[string]UsageMonDataScope `json:"scopes,omitempty"`

	UmLevel UsageMonLevel `json:"umLevel,omitempty"`

	StartDate time.Time `json:"startDate,omitempty"`

	EndDate time.Time `json:"endDate,omitempty"`

	UsageLimit UsageThreshold `json:"usageLimit,omitempty"`

	ResetPeriod time.Time `json:"resetPeriod,omitempty"`
}
