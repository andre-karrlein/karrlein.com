package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type navbar struct {
	app.Compo
}

func (navbar *navbar) Render() app.UI {
	return app.Nav().Class("navbar").Body(
		app.Div().Class("logo").Body(
			app.A().Href("/").Body(
				app.Img().Width(80).Src("/web/images/Logo_AK_alpha.png"),
			),
		),
		app.Ul().Class("menu").Body(
			app.Li().Body(
				app.A().Href("/resume.html").Text("Resume"),
			),
			app.Li().Body(
				app.A().Href("/climate.html").Text("Klima"),
			),
			app.Li().Body(
				app.A().Href("/imprint.html").Text("Impressum"),
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
