package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// hello is a component that displays a simple "Hello World!". A component is a
// customizable, independent, and reusable UI element. It is created by
// embedding app.Compo into a struct.
type home struct {
	app.Compo
}

type navbar struct {
	app.Compo
}

// The Render method is where the component appearance is defined. Here, a
// "Hello World!" is displayed as a heading.
func (h *home) Render() app.UI {
	return app.Div().Body(
		&navbar{},
		app.H1().Text("Hello World!"),
	)
}

func (n *navbar) Render() app.UI {
	return app.Div().Body(
		app.H1().
			Class("navbar").
			Text("André Karrlein"),
	)
}

// The main function is the entry point of the UI. It is where components are
// associated with URL paths and where the UI is started.
func main() {
	app.Route("/", &home{}) // hello component is associated with URL path "/".
	app.Run()               // Launches the PWA.
}
