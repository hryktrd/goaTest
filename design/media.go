package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var PointMedia = MediaType("application/vnd.point.area+json", func() {
	Description("A station of mine")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique station ID")
		Attribute("name", String, "Name of station")
		Required("id", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id") // Media types may have multiple views and must
		Attribute("name")
	})
})
