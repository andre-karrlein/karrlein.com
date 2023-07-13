package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type imprint struct {
	app.Compo
}

func (imprint *imprint) Render() app.UI {
	return app.Div().Class("body").Body(
		app.Header().Body(
			&navbar{},
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("resume").Body(
						app.H2().Class("title").Text("Impressum"),
							app.P().Body(
								app.Text("Angaben gemäß § 5 TMG"),
								app.Br(),
								app.Br(),
								app.Strong().Text("André Karrlein"),
								app.Br(),
								app.Text("Münchener Str. 25 F"),
								app.Br(),
								app.Text("83395 Freilassing"),
							),
					),
				),
			),
		),
	)
}

