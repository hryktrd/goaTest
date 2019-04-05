//go:generate goagen bootstrap -d github.com/hryktrd/goaTest/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/goa/middleware/security/basicauth"
	"github.com/hryktrd/goaTest/app"
)

func main() {
	// Create service
	service := goa.New("area")
	app.UseBasicAuthMiddleware(service, basicauth.New("admin", "password"))
	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "point" controller
	c := NewPointController(service)
	app.MountPointController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
