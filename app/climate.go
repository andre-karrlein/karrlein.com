package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type politics struct {
	app.Compo
}

func (politics *politics) Render() app.UI {
	data := getPoliticsData()

	return app.Div().Class("body").Body(
		app.Header().Body(
			&navbar{},
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("resume").Body(
						app.H2().Class("title").Text("Politik"),
						app.Range(data).Slice(func(i int) app.UI {
							return app.Div().Class("").Body(
								app.P().Body(
									app.Strong().Text(data[i].Heading),
									app.Br(),
									app.Text(data[i].Content),
								),
							)
						}),
					),
				),
			),
		),
	)
}
