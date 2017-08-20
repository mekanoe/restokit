package middleware

import (
	"regexp"

	"github.com/valyala/fasthttp"
)

var (
	// ShortName is used for parsing and outputting versioned/vendored headers.
	ShortName = "resto"

	versionRegex    = regexp.MustCompile(`^application/vnd.` + ShortName + `.(v[0-9]+(?:[a-z]+)?)\+json$`)
	originShortName = ShortName
)

// VersionedRouteMap is a container class for storing handlers for versions
type VersionedRouteMap map[string]fasthttp.RequestHandler

// VersionedRoute returns a middleware handler that switches API routes to specific version of said API
func VersionedRoute(i VersionedRouteMap) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		v := parseAcceptVersion(ctx.Request.Header.Peek("Accept"))

		route, ok := i[v]
		if !ok {
			v = "default"
			route = i[v]
		}

		ctx.Response.Header.Set(ShortName+"-API-Version", v)

		route(ctx)
	}
}

// application/vnd.slpn[.version]+json
func parseAcceptVersion(h []byte) string {

	if ShortName != originShortName {
		versionRegex = regexp.MustCompile(`^application/vnd.` + ShortName + `.(v[0-9]+(?:[a-z]+)?)\+json$`)
		originShortName = ShortName
	}

	s := versionRegex.FindSubmatch(h)

	if len(s) == 2 {
		return string(s[1])
	}

	return "default"
}
