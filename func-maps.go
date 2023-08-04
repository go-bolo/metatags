package metatags

import (
	"html/template"

	"github.com/go-bolo/bolo"
)

func renderMetatags(ctx *bolo.RequestContext) template.HTML {
	mt := ctx.Get("metatags").(*HTMLMetaTags)
	return mt.Render(ctx)
}
