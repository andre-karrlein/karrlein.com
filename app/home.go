package main

import (
	"github.com/maxence-charriere/go-app/v8/pkg/app"
)

type home struct {
	app.Compo
}

func (home *home) Render() app.UI {
	return app.Body().Body(
		app.Header().Body(
			&navbar{},
			app.Div().Class("container").Body(
				app.Div().Class("row").Body(
					app.Div().Class("text-content").Body(
						app.H2().Text("Hi, I'm"),
						app.H2().Class("title").Text("andre karrlein"),
						app.P().Text("SOLUTION ARCHITECT AT RED BULL IN SALZBURG, AUSTRIA."),
					),
					app.Div().Class("img-box").Body(
						app.Img().Src("https://storage.googleapis.com/karrlein/karrlein.com/me_512_512.png"),
					),
				),
			),
		),
	)
}
