package main


import (

    "gofr.dev/pkg/gofr"
	"t4nm4y/testProject/routes"
)

func main() {
    // initialise gofr object
    app := gofr.New()

    // register route greet
    app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
        return "Hello World!", nil
    })
	app.GET("/all", routes.all)

    // Starts the server, it will listen on the default port 8000.
    // it can be over-ridden through configs
    app.Start()
}