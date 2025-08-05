package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"goravel-blog/app/http/middleware"

	"goravel-blog/app/http/controllers"
)

func Api() {
	authController := controllers.NewAuthController()
	facades.Route().Prefix("auth").Group(func(r route.Router) {
		r.Post("login", authController.Login)
		r.Middleware(middleware.Jwt()).Group(func(rr route.Router) {
			rr.Get("info", authController.Info)
			rr.Post("logout", authController.Logout)
			rr.Post("refresh", authController.Refresh)
		})
	})
	facades.Route().Any("/test", controllers.NewTestController().Index)
	userController := controllers.NewUserController()
	facades.Route().Resource("users", userController)
	facades.Route().Middleware(middleware.Jwt()).Get("users/current", userController.Current)
	facades.Route().Resource("categories", controllers.NewCategoryController())
}
