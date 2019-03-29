package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("area", func() { // API defines the microservice endpoint and
	Title("area API")                   // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:8080")
})
