package metatags

import (
	"github.com/go-bolo/bolo"
	"github.com/go-bolo/system_settings"
	"github.com/gookit/event"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Plugin struct {
	bolo.Pluginer
	Name string
	App  bolo.App
}

func (r *Plugin) GetName() string {
	return r.Name
}

func (r *Plugin) Init(app bolo.App) error {
	logrus.Debug(r.GetName() + " Init")
	r.App = app

	app.GetEvents().On("bindMiddlewares", event.ListenerFunc(func(e event.Event) error {
		return r.bindMiddlewares(app)
	}), event.Normal)

	app.GetEvents().On("setTemplateFunctions", event.ListenerFunc(func(e event.Event) error {
		return r.setTemplateFunctions(app)
	}), event.Normal)

	return nil
}

func (p *Plugin) bindMiddlewares(app bolo.App) error {
	logrus.Debug("MMPlugin.bindMiddlewares " + p.GetName())

	router := app.GetRouter()
	router.Use(p.setDefaultMetatags())

	return nil
}

func (p *Plugin) setTemplateFunctions(app bolo.App) error {
	app.SetTemplateFunction("renderMetatags", renderMetatags)

	return nil
}

func (r *Plugin) BindRoutes(app bolo.App) error {
	return nil
}

func (r *Plugin) SetTemplateFuncMap(app bolo.App) error {
	return nil
}

func (r *Plugin) setDefaultMetatags() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cfg := r.App.GetConfiguration()

			mt := HTMLMetaTags{
				Title:       system_settings.GetD("SITE_NAME", cfg.Get("SITE_NAME")),
				Description: system_settings.GetD("SITE_DESCRIPTION", cfg.Get("SITE_DESCRIPTION")),
				ImageURL:    system_settings.GetD("SITE_IMAGE_URL", cfg.Get("SITE_IMAGE_URL")),
				SiteName:    system_settings.GetD("SITE_NAME", cfg.Get("SITE_NAME")),
			}

			c.Set("metatags", &mt)
			return next(c)
		}
	}
}

type PluginCfgs struct{}

func NewPlugin(cfg *PluginCfgs) *Plugin {
	p := Plugin{Name: "metatags"}
	return &p
}
