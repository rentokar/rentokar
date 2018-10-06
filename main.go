package main

import (
	// "fmt"

	"github.com/richard/rentokar/model"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main(){
	app := iris.Default()

	// Simple group: v1.
	v1 := app.Party("/v1")
	{
		user := v1.Party("/user")
		{
			user.Get("/login", func(ctx context.Context) {
				ctx.Writef("Indirect")
			})
		}

		// v1.Post("/submit", submitEndpoint)
		// v1.Post("/read", readEndpoint)
	}

	model.InitDb()
	// Simple group: v2.
	// v2 := app.Party("/v2")
	// {
	// 	v2.Post("/login", loginEndpoint)
	// 	v2.Post("/submit", submitEndpoint)
	// 	v2.Post("/read", readEndpoint)
	// }

	app.Run(iris.Addr(":8080"))
}