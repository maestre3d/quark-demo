package main

import "github.com/maestre3d/quark-demo/analytics-service/pkg/app"

func main() {
	b := app.NewSubscriber()
	b.ListenAndServe()
}
