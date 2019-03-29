package design // The convention consists of naming the design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("point", func() { // Resources group related API endpoints
	BasePath("/points")      // together. They map to REST resources for REST
	DefaultMedia(PointMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get point by id") // with its path, parameters (both path
		Routing(GET("/:pointId"))      // parameters and querystring values) and payload
		Params(func() {                // (shape of the request body).
			Param("pointId", Integer, "Bottle ID")
		})
		Response(OK, PointMedia) // Responses define the shape and status code
		Response(NotFound)       // of HTTP responses.
	})
	Action("list", func() {
		Description("point list")
		Routing(GET("/"))
		Response(OK, CollectionOf(PointMedia)) // Responses define the shape and status code
		Response(NotFound)                     // of HTTP responses.
	})
})
