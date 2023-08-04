package metatags

import (
	"html/template"
	"time"

	"github.com/go-bolo/bolo"
	"github.com/go-bolo/system_settings"
)

type HTMLMetaTags struct {
	Title       string
	Description string
	Canonical   string
	SiteName    string
	Type        string
	ImageURL    string
	ImageHeight string
	ImageWidth  string
	Author      string
	Keywords    string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	PublishedAt *time.Time
	TwitterSite string
}

func (r *HTMLMetaTags) Set(key string, value string) {

}

func (r *HTMLMetaTags) Get(key string) string {
	return ""
}

func (r *HTMLMetaTags) Render(ctx *bolo.RequestContext) template.HTML {
	html := ""
	pagePath := ""
	twitterSite := system_settings.Get("@_TWITTER")
	ogType := system_settings.GetD("METATAG_OG_TYPE", "website")

	if ctx.Param("id") != "" {
		ogType = system_settings.GetD("METATAG_OG_TYPE_ITEM", "article")
	}

	// Use the current url as canonical, if the url is a permanent with alias like /content/:id the backed will repply with redirect to the alias
	if r.Canonical == "" {
		pathBeforeAlias := ctx.GetString("pathBeforeAlias")
		if pathBeforeAlias != "" {
			pagePath = pathBeforeAlias
		} else {
			pagePath = ctx.Request().URL.Path
		}
	} else {
		pagePath = r.Canonical
	}

	pageUrl := ctx.AppOrigin + pagePath

	if pageUrl != "" {
		html += `<meta name="url" property="og:url" content="` + pageUrl + `">`
		html += `<link rel="canonical" href="` + pageUrl + `" />`
	}

	siteName := system_settings.Get("site_name")

	if siteName != "" {
		html += `<meta property="og:site_name" content="` + siteName + `">`
		html += `<meta itemprop="name" content="` + siteName + `">`
	}

	html += `<meta name="twitter:site" content="` + twitterSite + `">`
	html += `<meta property="og:type" content="` + ogType + `">`

	if r.Description != "" {
		html += `<meta name="description" property="og:description" content="` + r.Description + `">`
		html += `<meta name="description" content="` + r.Description + `">`
		html += `<meta name="twitter:description" content="` + r.Description + `">`
	}

	if r.Title != "" {
		html += `<meta name="twitter:title" content="` + r.Title + `">`
		html += `<meta name="title" property="og:title" content="` + r.Title + `">`
	}

	if r.ImageURL != "" {
		html += `<meta property="image" content="` + r.ImageURL + `">`
		html += `<meta property="og:image" content="` + r.ImageURL + `">`
		// html += `<meta property="og:image:width" content="1200">`
		// html += `<meta property="og:image:height" content="630">`
		html += `<meta name="robots" content="max-image-preview:large">`
		html += `<link rel="image_src" href="` + r.ImageURL + `"/>`
	}

	if r.Keywords != "" {
		html += `<meta name="keywords" content="` + r.Keywords + `">`
		html += `<meta name="news_keywords" content="` + r.Keywords + `">`
	}

	publisher := system_settings.Get("METATAG_ARTICLE_PUBLISHED")

	if publisher != "" {
		html += `<meta property="article:publisher" content="` + publisher + `">`
	}

	html += `<meta name="expires" content="never">`

	return template.HTML(html)
}
