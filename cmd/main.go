package main

import (
	"github.com/slovty/mo-gin/pkg/route"
)

func main() {

	r := route.InitRouter()

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
