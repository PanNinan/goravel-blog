package providers

import (
	"github.com/goravel/framework/contracts/foundation"
	contractshttp "github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/http/limit"

	"goravel-blog/app/http"
	"goravel-blog/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register(app foundation.Application) {
}

func (receiver *RouteServiceProvider) Boot(app foundation.Application) {
	// Add HTTP middleware
	facades.Route().GlobalMiddleware(http.Kernel{}.Middleware()...)
	facades.Route().Recover(func(ctx contractshttp.Context, err any) {
		facades.Log().Error(err)
		_ = ctx.Response().String(contractshttp.StatusInternalServerError, "recover").Abort()
	})
	receiver.configureRateLimiting()

	// Add routes
	routes.Web()
	routes.Api()
}

func (receiver *RouteServiceProvider) configureRateLimiting() {
	facades.RateLimiter().For("global", func(ctx contractshttp.Context) contractshttp.Limit {
		return limit.PerMinute(1000)
	})
	facades.RateLimiter().ForWithLimits("ip", func(ctx contractshttp.Context) []contractshttp.Limit {
		return []contractshttp.Limit{
			limit.PerDay(1000),
			limit.PerMinute(2).By(ctx.Request().Ip()),
		}
	})
	facades.RateLimiter().For("web", func(ctx contractshttp.Context) contractshttp.Limit {
		return limit.PerMinute(60)
	})
	facades.RateLimiter().For("api", func(ctx contractshttp.Context) contractshttp.Limit {
		return limit.PerMinute(60)
	})
}
