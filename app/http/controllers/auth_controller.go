package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel-blog/app/models"
	"goravel-blog/app/models/common/response"
)

type AuthController struct {
	// Dependent services
}

func NewAuthController() *AuthController {
	return &AuthController{
		// Inject services
	}
}

func (r *AuthController) Login(ctx http.Context) http.Response {
	var user models.User
	if err := facades.Orm().Query().Where(map[string]string{
		"email": ctx.Request().Input("username"),
		//"password": ctx.Request().Input("password"),
	}).First(&user); err != nil {
		return response.FailWithMessage(ctx, err.Error())
	}
	// 判断密码

	var (
		token string
		err   error
	)
	token, err = facades.Auth(ctx).Login(user)
	if err != nil {
		return response.InternalServerError(ctx, err.Error())
	}

	return response.OkWithData(ctx, http.Json{
		"token": token,
	})
}

func (r *AuthController) Info(ctx http.Context) http.Response {
	var (
		id   string
		user models.User
		err  error
	)
	id, err = facades.Auth(ctx).ID()
	if err != nil {
		return response.InternalServerError(ctx, err.Error())
	}
	err = facades.Auth(ctx).User(&user)
	if err != nil {
		return response.InternalServerError(ctx, err.Error())
	}
	return response.OkWithData(ctx, http.Json{
		"id":   id,
		"user": user,
	})
}

func (r *AuthController) Logout(ctx http.Context) http.Response {
	err := facades.Auth(ctx).Logout()
	if err != nil {
		return response.InternalServerError(ctx, err.Error())
	}

	return response.Ok(ctx)
}

func (r *AuthController) Refresh(ctx http.Context) http.Response {
	token, err := facades.Auth(ctx).Refresh()
	if err != nil {
		return response.InternalServerError(ctx, err.Error())
	}

	return response.OkWithData(ctx, http.Json{
		"token": token,
	})
}
