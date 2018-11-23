package main

import (
	// examples:
	// "github.com/divan/txqr/qr"
	// "github.com/gopherjs/vecty/example/todomvc/actions"
	// "github.com/iafan/goplayspace/client/js/console"

	"github.com/gopherjs/vecty"
	"time"
)

func main() {

	app := NewApp()

	vecty.SetTitle("My App")
	vecty.AddStylesheet("node_modules/bulma/css/bulma.css")
	vecty.RenderBody(app)

	time.AfterFunc(2 * time.Second, func() {
		app.List.SetItems([]Item{
			{
				Title: "alex",
			},
			{
				Title: "sharov",
			},
		})
	})

}
