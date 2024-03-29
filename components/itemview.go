package components

import (
	"log"

	"github.com/VinceJnz/go-wasm-vecty/actions"
	"github.com/VinceJnz/go-wasm-vecty/store"
	"github.com/VinceJnz/go-wasm-vecty/store/model"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/hexops/vecty/style"
)

// ItemView is a vecty.Component which represents a single item in the TODO
// list.
type ItemView struct {
	vecty.Core

	Store     *store.Store `vecty:"prop"`
	Index     int          `vecty:"prop"`
	Item      *model.Item  `vecty:"prop"`
	editing   bool
	editTitle string
	input     *vecty.HTML
}

// Key implements the vecty.Keyer interface.
func (p *ItemView) Key() interface{} {
	log.Println("ItemView Key", "p.Index =", p.Index)
	return p.Index
}

func (p *ItemView) onDestroy(event *vecty.Event) {
	p.Store.Dispatcher.Dispatch(&actions.DestroyItem{
		Index: p.Index,
	})
}

func (p *ItemView) onToggleCompleted(event *vecty.Event) {
	p.Store.Dispatcher.Dispatch(&actions.SetCompleted{
		Index:     p.Index,
		Completed: event.Target.Get("checked").Bool(),
	})
}

func (p *ItemView) onStartEdit(event *vecty.Event) {
	p.editing = true
	p.editTitle = p.Item.Title
	vecty.Rerender(p)
	p.input.Node().Call("focus")
}

func (p *ItemView) onEditInput(event *vecty.Event) {
	log.Println("itemview onEditInput", "p.Index =", p.Index)
	p.editTitle = event.Target.Get("value").String()
	vecty.Rerender(p)
}

func (p *ItemView) onStopEdit(event *vecty.Event) {
	p.editing = false
	vecty.Rerender(p)
	p.Store.Dispatcher.Dispatch(&actions.SetTitle{
		Index: p.Index,
		Title: p.editTitle,
	})
}

// Render implements the vecty.Component interface.
func (p *ItemView) Render() vecty.ComponentOrHTML {
	log.Println("itemview Render", "p.Index =", p.Index)
	p.input = elem.Input(
		vecty.Markup(
			vecty.Class("edit"),
			prop.Value(p.editTitle),
			event.Input(p.onEditInput),
		),
	)

	return elem.ListItem(
		vecty.Markup(
			vecty.ClassMap{
				"completed": p.Item.Completed,
				"editing":   p.editing,
			},
		),

		elem.Div(
			vecty.Markup(
				vecty.Class("view"),
			),

			elem.Input(
				vecty.Markup(
					vecty.Class("toggle"),
					prop.Type(prop.TypeCheckbox),
					prop.Checked(p.Item.Completed),
					event.Change(p.onToggleCompleted),
				),
			),
			elem.Label(
				vecty.Markup(
					event.DoubleClick(p.onStartEdit),
				),
				vecty.Text(p.Item.Title),
			),
			elem.Button(
				vecty.Markup(
					vecty.Class("destroy"),
					event.Click(p.onDestroy),
				),
			),
		),
		elem.Form(
			vecty.Markup(
				style.Margin(style.Px(0)),
				event.Submit(p.onStopEdit).PreventDefault(),
			),
			p.input,
		),
	)
}
