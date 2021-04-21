package main

import "github.com/maestre3d/quark-demo/notification-service/pkg/app"

func main() {
	b := app.NewSubscriber()
	b.ListenAndServe()
}
