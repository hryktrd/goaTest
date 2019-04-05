package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/hryktrd/goaTest/app"
)

// PointController implements the point resource.
type PointController struct {
	*goa.Controller
}

// NewPointController creates a point controller.
func NewPointController(service *goa.Service) *PointController {
	return &PointController{Controller: service.NewController("PointController")}
}

// List runs the list action.
func (c *PointController) List(ctx *app.ListPointContext) error {
	// PointController_List: start_implement

	res := make(app.PointAreaCollection, 3)
	st := []app.PointArea{
		{
			ID:   1,
			Name: "a",
		},
		{
			ID:   2,
			Name: "b",
		},
		{
			ID:   3,
			Name: "c",
		},
	}

	for i := 0; i < 3; i++ {
		res[i] = &st[i]
	}

	return ctx.OK(res)

	// PointController_List: end_implement
}

// Show runs the show action.
func (c *PointController) Show(ctx *app.ShowPointContext) error {
	// PointController_Show: start_implement

	res := app.PointArea{
		ID:   ctx.PointID,
		Name: fmt.Sprintf("Point #%d", ctx.PointID),
	}

	return ctx.OK(&res)

	// PointController_Show: end_implement
}
