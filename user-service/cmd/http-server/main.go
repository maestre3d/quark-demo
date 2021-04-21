package main

import "github.com/maestre3d/quark-demo/user-service/pkg/app"

func main() {
	api := app.NewHTTPAPI()
	api.ListenAndServe()
}
