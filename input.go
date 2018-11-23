package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

type input struct {
	vecty.Core

	onNewItem func(Item)
	val       string
	input     *vecty.HTML
}

func NewInput(onNewItem func(Item)) *input {
	return &input{
		onNewItem: onNewItem,
	}
}

func (i *input) Render() vecty.ComponentOrHTML {
	input := elem.Input(
		vecty.Markup(
			vecty.Class("input"),
			prop.Type("text"),
			prop.Value(i.val),
			event.Input(i.onInput),
		),
	)

	return elem.Form(
		elem.Div(
			vecty.Markup(
				vecty.Class("field", "has-addons"),
			),

			elem.Div(
				vecty.Markup(
					vecty.Class("control"),
				),
				input,
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("control"),
				),
				elem.Button(
					vecty.Markup(
						vecty.Class("button", "is-primary"),
						prop.Type("submit"),
					),
					vecty.Text("Add"),
				),
			),
		),
		vecty.Markup(
			event.Submit(i.onSubmit).PreventDefault(),
		),
	)
}


func (i *input) onStartType(event *vecty.Event) {
	vecty.Rerender(i)
	i.Focus()
}

func (i *input) Focus() {
	i.input.Node().Call("focus")
}

func (i *input) onInput(event *vecty.Event) {
	i.val = event.Target.Get("value").String()
	vecty.Rerender(i)
}

func (i *input) onSubmit(event *vecty.Event) {
	if i.val == "" {
		return
	}

	i.onNewItem(Item{Title: i.val})

	i.val = ""
	vecty.Rerender(i)
}
