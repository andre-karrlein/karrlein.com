package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type resume struct {
	app.Compo
}

func (resume *resume) Render() app.UI {
	data := getData()

	return app.Div().Class("body").Body(
		app.Header().Body(
			&navbar{},
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("resume").Body(
						app.H2().Class("title").Text("Resume"),
						app.Range(data).Slice(func(i int) app.UI {
							s := fmt.Sprintf("%s in %s", data[i].Company, data[i].City)

							return app.Div().Class("").Body(
								app.P().Body(
									app.Strong().Text(data[i].Role),
									app.Small().Text("  "+data[i].Timeframe),
									app.Br(),
									app.Text(s),
								),
							)
						}),
					),
				),
			),
		),
	)
}
