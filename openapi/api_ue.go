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
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RegisteredUeContextGet - Get UE status
func RegisteredUeContextGet(c *gin.Context) {
	ueCtxs:=make([]UeContext,0)
	for _, v := range UEs {
		if v.Online {
			for _, subs := range Subscribers {
				if subs.UeId == v.UeId && subs.PlmnId == v.PlmnId {
					ueCtxs = append(ueCtxs, UeContext{
						AccessType:  "3GPP_ACCESS",
						Supi:        subs.UeId,
						Guti:        strconv.Itoa(int(v.PlmnId))+"cafe0000000"+v.UeId[len(v.UeId)-2:],
						Mcc:         strconv.Itoa(int(v.PlmnId))[:3],
						Mnc:         strconv.Itoa(int(v.PlmnId))[3:],
						Tac:         strconv.Itoa(int(subs.PlmnId)),
						PduSessions: []UeContextPduSessions{{
							PduSessionId: "1",
							SmContextRef: "urn:uuid:a37b276f-d8d2-4ad6-afd1-f3cb60ff1a80",
							Sst:          strconv.Itoa(int(subs.AccessAndMobilitySubscriptionData.Nssai.DefaultSingleNssais[0].Sst)),
							Sd:           subs.AccessAndMobilitySubscriptionData.Nssai.DefaultSingleNssais[0].Sd,
							Dnn:          subs.SmPolicyData.SmPolicySnssaiData["01010203"].SmPolicyDnnData["internet"].Dnn,
						}},
						CmState:     "CONNECTED",
					})
				}
			}
		}
	}
	c.JSON(http.StatusOK, ueCtxs)
}

// RegisteredUeContextSupiGet - Get specific UE status
func RegisteredUeContextSupiGet(c *gin.Context) {
	supi := c.Param("supi")
	fmt.Printf("Got SUPI: %v\n", supi)

	ueCtxs:=make([]UeContext,0)
	for _, v := range UEs {
		if v.Online {
			for _, subs := range Subscribers {
				if subs.UeId == v.UeId && subs.PlmnId == v.PlmnId {
					ueCtxs = append(ueCtxs, UeContext{
						AccessType:  "3GPP_ACCESS",
						Supi:        subs.UeId,
						Guti:        strconv.Itoa(int(v.PlmnId))+"cafe0000000"+v.UeId[len(v.UeId)-2:],
						Mcc:         strconv.Itoa(int(v.PlmnId))[:3],
						Mnc:         strconv.Itoa(int(v.PlmnId))[3:],
						Tac:         strconv.Itoa(int(subs.PlmnId)),
						PduSessions: []UeContextPduSessions{{
							PduSessionId: "1",
							SmContextRef: "urn:uuid:a37b276f-d8d2-4ad6-afd1-f3cb60ff1a80",
							Sst:          strconv.Itoa(int(subs.AccessAndMobilitySubscriptionData.Nssai.DefaultSingleNssais[0].Sst)),
							Sd:           subs.AccessAndMobilitySubscriptionData.Nssai.DefaultSingleNssais[0].Sd,
							Dnn:          subs.SmPolicyData.SmPolicySnssaiData["01010203"].SmPolicyDnnData["internet"].Dnn,
						}},
						CmState:     "CONNECTED",
					})
				}
			}
		}
	}

	for _,v := range ueCtxs{
		if v.Supi==supi {
			c.JSON(http.StatusOK, v)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{})
}
