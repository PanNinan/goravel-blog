package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"goravel-blog/app/models"
	"goravel-blog/app/models/common/response"
)

type UserController struct {
	// Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		// Inject services
	}
}

func (r *UserController) Index(ctx http.Context) http.Response {
	return nil
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return nil
}

func (r *UserController) Current(ctx http.Context) http.Response {
	user, ok := ctx.Value("user").(models.User)
	if !ok {
		return response.FailWithMessage(ctx, "Unauthorized")
	}
	return response.OkWithData(ctx, user)
}

func (r *UserController) Store(ctx http.Context) http.Response {
	return nil
}

func (r *UserController) Update(ctx http.Context) http.Response {
	return nil
}

func (r *UserController) Destroy(ctx http.Context) http.Response {
	return nil
}
