package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type List struct {
	vecty.Core
	items []Item
}

func (l *List) Render() vecty.ComponentOrHTML {
	items := vecty.List{}
	for _, el := range l.items {
		items = append(items, item(el))
	}

	return elem.UnorderedList(
		vecty.Markup(
			vecty.Class("todo-list"),
		),
		items,
	)
}

func (l *List) SetItems(items []Item) {
	l.items = items
	vecty.Rerender(l)
}

func (l *List) AddItem(item Item) {
	l.items = append(l.items, item)
	vecty.Rerender(l)
}

func formattedTime() vecty.ComponentOrHTML {
	return elem.Time(
		vecty.Markup(
			vecty.Attribute("datetime", "2016-1-1"),
		),
		vecty.Text("11:09 PM - 1 Jan 2016"),
	)
}

func iconAngleDown() vecty.ComponentOrHTML {
	return elem.Span(
		vecty.Markup(
			vecty.Class("icon"),
		),
		elem.Italic(
			vecty.Markup(
				vecty.Class("fas", "fa-angle-down"),
				vecty.Attribute("aria-hidden", "true"),
			),
		),
	)
}

