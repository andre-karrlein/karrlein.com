package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type climate struct {
	app.Compo
}

func (climate *climate) Render() app.UI {
	data := getClimateData()

	return app.Div().Class("body").Body(
		app.Header().Body(
			&navbar{},
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("resume").Body(
						app.H2().Class("title").Text("Klima"),
						app.Range(data).Slice(func(i int) app.UI {
							return app.Div().Class("").Body(
								app.P().Body(
									app.Strong().Text(data[i].heading),
									app.Br(),
									app.Text(data[i].content),
								),
							)
						}),
					),
				),
			),
		),
	)
}
