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
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name        string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method      string
	// Pattern is the pattern of the URI.
	Pattern     string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := gin.Default()
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/v1/",
		Index,
	},

	{
		"AuthLogonPost",
		http.MethodPost,
		"/v1/auth/logon",
		AuthLogonPost,
	},

	{
		"UePduSessionInfoGet",
		http.MethodGet,
		"/v1/ue-pdu-session-info",
		UePduSessionInfoGet,
	},

	{
		"UePduSessionInfoSmContextRefGet",
		http.MethodGet,
		"/v1/ue-pdu-session-info/:smContextRef",
		UePduSessionInfoSmContextRefGet,
	},

	{
		"RanGet",
		http.MethodGet,
		"/v1/ran",
		RanGet,
	},

	{
		"RanRanIdGet",
		http.MethodGet,
		"/v1/ran/:ranId",
		RanRanIdGet,
	},

	{
		"SampleSubscriberGet",
		http.MethodGet,
		"/v1/sample/subscriber",
		SampleSubscriberGet,
	},

	{
		"SubscriberDelUeIdServingPlmnIdDelete",
		http.MethodDelete,
		"/v1/subscriber/del/:ueId/:servingPlmnId",
		SubscriberDelUeIdServingPlmnIdDelete,
	},

	{
		"SubscriberGet",
		http.MethodGet,
		"/v1/subscriber",
		SubscriberGet,
	},

	{
		"SubscriberUeIdServingPlmnIdGet",
		http.MethodGet,
		"/v1/subscriber/:ueId/:servingPlmnId",
		SubscriberUeIdServingPlmnIdGet,
	},

	{
		"SubscriberUeIdServingPlmnIdPost",
		http.MethodPost,
		"/v1/subscriber/:ueId/:servingPlmnId",
		SubscriberUeIdServingPlmnIdPost,
	},

	{
		"RegisteredUeContextGet",
		http.MethodGet,
		"/v1/registered-ue-context",
		RegisteredUeContextGet,
	},

	{
		"RegisteredUeContextSupiGet",
		http.MethodGet,
		"/v1/registered-ue-context/:supi",
		RegisteredUeContextSupiGet,
	},
}