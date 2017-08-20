package middleware

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// JSON adds ctx mix-ins for unmarshalling/marshalling data in one line vs a few.
func JSON(f fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		ctx.SetUserValue("json:in", func(i interface{}) error {
			return json.Unmarshal(ctx.Request.Body(), i)
		})

		ctx.SetUserValue("json:out", func(i interface{}) error {
			out, err := json.Marshal(i)
			if err != nil {
				return err
			}

			ctx.Response.Header.Add("Content-Type", "application/json")
			_, err = ctx.Write(out)
			return err
		})
		f(ctx)
	}
}
