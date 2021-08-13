package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type navbar struct {
	app.Compo
}

func (navbar *navbar) Render() app.UI {
	return app.Nav().Class("navbar").Body(
		app.Div().Class("logo").Body(
			app.A().Href("/").Body(
				app.Img().Width(80).Src("https://storage.googleapis.com/karrlein/karrlein.com/Logo%20AK_alpha.png"),
			),
		),
		app.Ul().Class("menu").Body(
			app.Li().Body(
				app.A().Href("/resume").Text("Resume"),
			),
			app.Li().Body(
				app.A().Href("https://github.com/andre-karrlein").Text("Github"),
			),
			app.Li().Body(
				app.A().Href("https://twitter.com/rb_ak1").Text("Twitter"),
			),
			app.Li().Body(
				app.A().Href("https://linkedin.com/in/andré-karrlein-81879417a").Text("LinkedIn"),
			),
		),
	)
}
