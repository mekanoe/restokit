// Package api is for API routes. Everything here is ideally unexported.
package api

//go:generate restokit-codegen $CWD

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// GET /test v2 default
func testGet(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello world! v2")
}

// GET /test v1
func testGetv1(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Hello world! v1")
}

// GET /hello/:name
// NoLogging
func hello(ctx *fasthttp.RequestCtx) {
	ctx.WriteString(fmt.Sprintf("Hello, %s!", ctx.UserValue("name")))
}

// GET /json
// JSON
func jsonTest(ctx *fasthttp.RequestCtx) {
	var intest map[string]interface{}
	ctx.UserValue("json:in").(func(interface{}) error)(&intest)

	out := map[string]interface{}{
		"in": intest,
		"hi": "hello!",
	}

	ctx.UserValue("json:out").(func(interface{}) error)(out)
}

// GET /localmw
// JSON @localMw
func getlocalmw(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("hello")
}

// GET /localmw2
// @mw.NoLogging JSON
func getlocalmw2(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("hello")
}

func localMw(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Add("localmw", "true")
	}
}

// // POST /test v1
// // Inject
// func testPost(ctx *fasthttp.RequestCtx) {
// 	ctx.WriteString("Hello world!")
// }

// // POST /test v2
// // Inject Inject2(Ole) Other
// func testPost(ctx *fasthttp.RequestCtx) {
// 	ctx.WriteString("Hello world!")
// }
