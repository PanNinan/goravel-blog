package middleware

import (
	"github.com/goravel/framework/auth"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/facades"
	"goravel-blog/app/models"
	"goravel-blog/app/models/common/response"
)

func Jwt() http.Middleware {
	return func(ctx http.Context) {
		guard := facades.Config().GetString("auth.defaults.guard")
		if ctx.Request().Header("Guard") != "" {
			guard = ctx.Request().Header("Guard")
		}

		token := ctx.Request().Header("Authorization", "")
		if token == "" {
			_ = ctx.Response().String(http.StatusUnauthorized, "Unauthorized").Abort()
			response.FailWithMessage(ctx, "Unauthorized")
			return
		}
		if _, err := facades.Auth(ctx).Guard(guard).Parse(token); err != nil {
			if errors.Is(err, auth.ErrorTokenExpired) {
				if "/auth/refresh" == ctx.Request().Path() {
					ctx.Request().Next()
					return
				} else {
					response.FailWithMessage(ctx, "Token is expired")
					return
				}
			} else {
				// Token is invalid
				response.FailWithMessage(ctx, "Token is invalid")
				return
			}
		} else {
			var user models.User
			if err = facades.Auth(ctx).User(&user); err != nil {
				response.InternalServerError(ctx, "User is invalid")
				return
			}
			ctx.WithValue("user", user)
		}

		ctx.Request().Next()
	}
}
