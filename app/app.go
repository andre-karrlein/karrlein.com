package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type experience struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	City      string `json:"city"`
	Timeframe string `json:"timeframe"`
}

type home struct {
	app.Compo
}

type hero struct {
	app.Compo
}

type burger struct {
	app.Compo
}

type resume struct {
	app.Compo
}

func (b *burger) Render() app.UI {
	return app.Raw(`
	<span class="navbar-burger" data-target="navbarMenuHero">
		<span></span>
		<span></span>
		<span></span>
	</span>
	`)
}

func (h *hero) Render() app.UI {
	return app.Section().Body(
		app.Div().Body(
			app.Header().Body(
				app.Div().Body(
					app.Div().Body(
						app.A().Body(
							app.Img().Src("https://storage.googleapis.com/karrlein/karrlein.com/Logo%20AK_alpha.png"),
						).Class("navbar-item").Href("/"),
						&burger{},
					).Class("navbar-brand"),
					app.Div().Body(
						app.Div().Body(
							app.A().Text(
								"Resume",
							).Class("navbar-item").Href("#my-resume"),
							app.A().Text(
								"Contact",
							).Class("navbar-item").Href("mailto:andre@karrlein.com"),
							app.A().Body(
								app.Span().Body(
									app.I().Class("fab fa-twitter"),
								).Class("icon"),
								app.Span().Text(
									"Twitter",
								),
							).Class("navbar-item").Href("https://twitter.com/rb_ak1"),
							app.A().Body(
								app.Span().Body(
									app.I().Class("fab fa-linkedin"),
								).Class("icon"),
								app.Span().Text(
									"LinkedIn",
								),
							).Class("navbar-item").Href("https://linkedin.com/in/andré-karrlein-81879417a"),
							app.A().Body(
								app.Span().Body(
									app.I().Class("fab fa-github"),
								).Class("icon"),
								app.Span().Text(
									"Github",
								),
							).Class("navbar-item").Href("https://github.com/andre-karrlein"),
						).Class("navbar-end"),
					).Class("navbar-menu").ID("navbarMenuHero"),
				).Class("container"),
			).Class("navbar"),
		).Class("hero-head"),
		app.Div().Body(
			app.Div().Body(
				app.P().Text("Hello there, I'm").Class("subtitle"),
				app.P().Text("André Karrlein").Class("title"),
				app.P().Text("Solution Architect at Red Bull in Salzburg, Austria.").Class("subtitle profession"),
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
	).Class("hero is-danger is-large")
}

func (home *home) Render() app.UI {
	return app.Div().Body(
		&hero{},
		app.Div().Body(
			app.Div().Body(
				app.Div().Body(
					app.H1().Text("Resume").Class("title has-text-centered section-title").ID("my-resume"),
					&resume{},
				).Class("section-light about-me"),
			).Class("container"),
		).Class("main-content"),
	)
}

func (resume *resume) Render() app.UI {
	data := getData()

	return app.Div().Body(
		app.Range(data).Slice(func(i int) app.UI {
			s := fmt.Sprintf("%s in %s", data[i].Company, data[i].City)

			return app.Div().Body(
				app.P().Body(
					app.Strong().Text(data[i].Role),
					app.Small().Text("  "+data[i].Timeframe),
					app.Br(),
					app.Text(s),
				),
			).Class("box")
		}),
	)
}

func getData() []experience {
	resp, err := http.Get("https://api.karrlein.com/resume/v1/experience/")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	var experiences []experience
	json.Unmarshal([]byte(sb), &experiences)

	return experiences
}

// The main function is the entry point of the UI. It is where components are
// associated with URL paths and where the UI is started.
func main() {
	app.Route("/", &home{}) // hello component is associated with URL path "/".
	app.Run()               // Launches the PWA.
}
