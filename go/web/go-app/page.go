package web_site

import "github.com/maxence-charriere/go-app/v10/pkg/app"

type Page struct {
	app.Compo

	name string
}

func NewPage(name string) *Page {
	return &Page{name: name}
}

func (h *Page) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello, "),
			app.If(h.name != "", func() app.UI {
				return app.Text(h.name)
			}).Else(func() app.UI {
				return app.Text("World!")
			}),
		),
		app.P().Body(
			app.Input().
				Type("text").
				Value(h.name).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(h.ValueTo(&h.name)),
		),
	)
}
