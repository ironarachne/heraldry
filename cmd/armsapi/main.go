package main

import (
	"math/rand"
	"time"

	"github.com/ironarachne/heraldry"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	width := 320
	height := 420

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("armsapi")
	})

	app.Get("/device", func(ctx iris.Context) {
		rand.Seed(time.Now().UnixNano())
		device := heraldry.Generate()
		svg := device.RenderToSVG(width, height)
		ctx.Writef(svg)
	})

	app.Get("/device/{id:int64}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			ctx.Writef("error while trying to parse id parameter")
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		rand.Seed(id)
		device := heraldry.Generate()
		svg := device.RenderToSVG(width, height)
		ctx.Writef(svg)
	})

	app.Run(iris.Addr(":7476"))
}
