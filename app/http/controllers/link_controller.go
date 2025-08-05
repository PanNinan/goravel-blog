package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type LinkController struct {
	// Dependent services
}

func NewLinkController() *LinkController {
	return &LinkController{
		// Inject services
	}
}

func (r *LinkController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *LinkController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *LinkController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *LinkController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *LinkController) Destroy(ctx http.Context) http.Response {
	return nil
}
