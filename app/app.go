package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type home struct {
	app.Compo
}

func (h *home) Render() app.UI {
	return app.Section().Body(
		app.Div().Body(
			app.Header().Body(
				app.Div().Body(
					app.Div().Body(
						app.A().Body(
							app.Img().Src("https://bulma.io/images/bulma-type-white.png"),
						).Class("navbar-item"),
					).Class("navbar-brand"),
				).Class("container"),
			).Class("navbar"),
		).Class("hero-head"),
		app.Div().Body(
			app.Div().Body(
				app.P().Text("André Karrlein.").Class("title"),
				app.P().Text("I'm a Solution Architect at Red Bull in Salzburg.").Class("subtitle"),
			).Class("container has-text-centered"),
		).Class("hero-body"),
		app.Div().Body(
			app.Nav().Body(
				app.Div().Body(
					app.Ul().Body(
						app.Li().Body(
							app.P().Text("André Karrlein ® 2021"),
						),
					),
				).Class("container has-text-centered"),
			).Class("tabs is-boxed is-fullwidth"),
		).Class("hero-foot"),
	).Class("hero is-danger is-fullheight-with-navbar")
}

// The main function is the entry point of the UI. It is where components are
// associated with URL paths and where the UI is started.
func main() {
	app.Route("/", &home{}) // hello component is associated with URL path "/".
	app.Run()               // Launches the PWA.
}
