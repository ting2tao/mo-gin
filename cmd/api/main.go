package main

import (
	"github.com/slovty/mo-gin/internal/app"
	"github.com/slovty/mo-gin/pkg/route"
)

func main() {
	injector, _, err := app.BuildInjector()
	if err != nil {
		return
	}
	r := route.InitRouter()
	injector.Router.RegisterRouter(r)
	_ = r.Run()
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
