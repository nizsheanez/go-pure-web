package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

// App is a top-level app component.
type App struct {
	vecty.Core

	//session      *Session
	//settings     *Settings
	// any other stuff you need,
	// it's just a struct

	List      *List
}

func NewApp() *App {
	return &App{}
}

// Render implements the vecty.Component interface.
func (a *App) Render() vecty.ComponentOrHTML {
	return elem.Body(
		//a.header(),
		a.body(),
		vecty.Markup(
			//event.KeyDown(a.KeyListener),
		),
	)
}

func (a *App) body() vecty.ComponentOrHTML {
	a.List = &List{}

	return elem.Section(
		vecty.Markup(

			vecty.Class("section"),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("columns"),
			),
			// Left half
			elem.Div(
				vecty.Markup(
					vecty.Class("column", "is-half"),
				),
				NewInput(a.List.AddItem),
				a.List,
			),
			// Right half
			elem.Div(
				vecty.Markup(
					vecty.Class("column", "is-half"),
				),
				//vecty.If(!a.session.Started(), elem.Div(
				//    a.settings,
				//)),
				//vecty.If(a.session.Started(), elem.Div(
				//    a.resultsTable,
				//)),
			),
		),
	)
}
