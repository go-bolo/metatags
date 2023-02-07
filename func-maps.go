package metatags

import (
	"html/template"

	"github.com/go-catupiry/catu"
)

func renderMetatags(ctx *catu.RequestContext) template.HTML {
	mt := ctx.Get("metatags").(*HTMLMetaTags)
	return mt.Render(ctx)
}
