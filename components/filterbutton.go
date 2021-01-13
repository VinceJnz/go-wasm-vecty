package components

import (
	"github.com/VinceJnz/go-wasm-vecty/actions"
	"github.com/VinceJnz/go-wasm-vecty/store"
	"github.com/VinceJnz/go-wasm-vecty/store/model"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
)

// FilterButton is a vecty.Component which allows the user to select a filter
// state.
type FilterButton struct {
	vecty.Core

	Store  *store.Store
	Label  string            `vecty:"prop"`
	Filter model.FilterState `vecty:"prop"`
}

func (b *FilterButton) onClick(event *vecty.Event) {
	b.Store.Dispatcher.Dispatch(&actions.SetFilter{
		Filter: b.Filter,
	})
}

// Render implements the vecty.Component interface.
func (b *FilterButton) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		elem.Anchor(
			vecty.Markup(
				vecty.MarkupIf(b.Store.Filter == b.Filter, vecty.Class("selected")),
				prop.Href("#"),
				event.Click(b.onClick).PreventDefault(),
			),

			vecty.Text(b.Label),
		),
	)
}
