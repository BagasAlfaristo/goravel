package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version": support.Version,
		})
	})

	facades.Route().Get("/test", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("test.html", map[string]any{
			"version": support.Version,
		})
	})

	facades.Route().Get("/index", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("index.html", map[string]any{
			"version": support.Version,
		})
	})
}
