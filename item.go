package main

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"github.com/iafan/goplayspace/client/js/console"
)

type Item struct {
	Title string
}

func item(it Item) vecty.ComponentOrHTML {
	console.Log("item: render")
	return elem.ListItem(
		vecty.Markup(
			vecty.Class("card"),
		),
		elem.Header(
			vecty.Markup(
				vecty.Class("card-header"),
			),
			elem.Paragraph(
				vecty.Markup(
					vecty.Class("card-header-Title"),
				),
				vecty.Text(it.Title),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("card-header-icon"),
					prop.Href("#"),
					vecty.Attribute("aria-label", "more options"),
				),
				iconAngleDown(),
			),
		),
		elem.Div(
			vecty.Markup(
				vecty.Class("card-content"),
			),
			elem.Div(
				vecty.Markup(
					vecty.Class("content"),
				),
				vecty.Text("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec iaculis mauris"),
				formattedTime(),
			),
		),
		elem.Footer(
			vecty.Markup(
				vecty.Class("card-footer"),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("card-footer-item"),
					prop.Href("#"),
				),
				vecty.Text("Save"),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("card-footer-item"),
					prop.Href("#"),
				),
				vecty.Text("Edit"),
			),
			elem.Anchor(
				vecty.Markup(
					vecty.Class("card-footer-item"),
					prop.Href("#"),
				),
				vecty.Text("Delete"),
			),
		),
	)
}
