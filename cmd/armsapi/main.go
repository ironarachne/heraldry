package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/ironarachne/heraldry"
	"github.com/kataras/iris"
	"github.com/patrickmn/go-cache"
)

func main() {
	app := iris.New()
	c := cache.New(5*time.Minute, 10*time.Minute)

	width := 320
	height := 420

	app.Get("/", func(ctx iris.Context) {
		ctx.Writef("armsapi")
	})

	app.Get("/blazon/{id:int64}", func(ctx iris.Context) {
		id, err := ctx.Params().GetInt64("id")
		if err != nil {
			ctx.Writef("error while trying to parse id parameter")
			ctx.StatusCode(iris.StatusBadRequest)
			return
		}
		blazon, found := c.Get("blazon_" + strconv.FormatInt(id, 10))
		if found {
			ctx.Writef(blazon.(string))
		} else {
			rand.Seed(id)
			device := heraldry.Generate()
			svg := device.RenderToSVG(width, height)
			blazon = device.RenderToBlazon()
			c.Set("blazon_"+strconv.FormatInt(id, 10), blazon, cache.DefaultExpiration)
			c.Set("svg_"+strconv.FormatInt(id, 10), svg, cache.DefaultExpiration)
			ctx.Writef(blazon.(string))
		}
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
		svg, found := c.Get("svg_" + strconv.FormatInt(id, 10))
		if found {
			ctx.Writef(svg.(string))
		} else {
			rand.Seed(id)
			device := heraldry.Generate()
			svg := device.RenderToSVG(width, height)
			blazon := device.RenderToBlazon()
			c.Set("blazon_"+strconv.FormatInt(id, 10), blazon, cache.DefaultExpiration)
			c.Set("svg_"+strconv.FormatInt(id, 10), svg, cache.DefaultExpiration)
			ctx.Writef(svg)
		}
	})

	app.Run(iris.Addr(":7476"))
}
