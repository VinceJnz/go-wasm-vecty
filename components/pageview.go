package components

import (
	"log"
	"strconv"

	"github.com/VinceJnz/go-wasm-vecty/actions"
	"github.com/VinceJnz/go-wasm-vecty/store"
	"github.com/VinceJnz/go-wasm-vecty/store/model"
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"github.com/hexops/vecty/prop"
	"github.com/hexops/vecty/style"
)

// PageView is a vecty.Component which represents the entire page.
type PageView struct {
	vecty.Core

	Store        *store.Store
	newItemTitle string
}

func (p *PageView) onNewItemTitleInput(event *vecty.Event) {
	p.newItemTitle = event.Target.Get("value").String()
	log.Println("pageview onNewItemTitleInput ", "p.newItemTitle =", p.newItemTitle)
	vecty.Rerender(p)
}

func (p *PageView) onAdd(event *vecty.Event) {
	p.Store.Dispatcher.Dispatch(&actions.AddItem{
		Title: p.newItemTitle,
	})
	p.newItemTitle = ""
	vecty.Rerender(p)
}

func (p *PageView) onClearCompleted(event *vecty.Event) {
	p.Store.Dispatcher.Dispatch(&actions.ClearCompleted{})
}

func (p *PageView) onToggleAllCompleted(event *vecty.Event) {
	p.Store.Dispatcher.Dispatch(&actions.SetAllCompleted{
		Completed: event.Target.Get("checked").Bool(),
	})
}

// Render implements the vecty.Component interface.
func (p *PageView) Render() vecty.ComponentOrHTML {
	log.Println("pageview Render ", "p.newItemTitle =", p.newItemTitle)
	return elem.Body(
		elem.Section(
			vecty.Markup(
				vecty.Class("todoapp"),
			),

			p.renderHeader(),
			vecty.If(len(p.Store.Items) > 0,
				p.renderItemList(),
				p.renderFooter(),
			),
		),

		p.renderInfo(),
	)
}

func (p *PageView) renderHeader() *vecty.HTML {
	return elem.Header(
		vecty.Markup(
			vecty.Class("header"),
		),

		elem.Heading1(
			vecty.Text("todos"),
		),
		elem.Form(
			vecty.Markup(
				style.Margin(style.Px(0)),
				event.Submit(p.onAdd).PreventDefault(),
			),

			elem.Input(
				vecty.Markup(
					vecty.Class("new-todo"),
					prop.Placeholder("What needs to be done?"),
					prop.Autofocus(true),
					prop.Value(p.newItemTitle),
					event.Input(p.onNewItemTitleInput),
				),
			),
		),
	)
}

func (p *PageView) renderFooter() *vecty.HTML {
	count := p.Store.ActiveItemCount()
	itemsLeftText := " items left"
	if count == 1 {
		itemsLeftText = " item left"
	}

	return elem.Footer(
		vecty.Markup(
			vecty.Class("footer"),
		),

		elem.Span(
			vecty.Markup(
				vecty.Class("todo-count"),
			),

			elem.Strong(
				vecty.Text(strconv.Itoa(count)),
			),
			vecty.Text(itemsLeftText),
		),

		elem.UnorderedList(
			vecty.Markup(
				vecty.Class("filters"),
			),
			&FilterButton{Label: "All", Filter: model.All, Store: p.Store},
			vecty.Text(" "),
			&FilterButton{Label: "Active", Filter: model.Active, Store: p.Store},
			vecty.Text(" "),
			&FilterButton{Label: "Completed", Filter: model.Completed, Store: p.Store},
		),

		vecty.If(p.Store.CompletedItemCount() > 0,
			elem.Button(
				vecty.Markup(
					vecty.Class("clear-completed"),
					event.Click(p.onClearCompleted),
				),
				vecty.Text("Clear completed ("+strconv.Itoa(p.Store.CompletedItemCount())+")"),
			),
		),
	)
}

func (p *PageView) renderInfo() *vecty.HTML {
	return elem.Footer(
		vecty.Markup(
			vecty.Class("info"),
		),

		elem.Paragraph(
			vecty.Text("Double-click to edit a todo"),
		),
		elem.Paragraph(
			vecty.Text("Created by "),
			elem.Anchor(
				vecty.Markup(
					prop.Href("http://github.com/neelance"),
				),
				vecty.Text("Richard Musiol"),
			),
		),
		elem.Paragraph(
			vecty.Text("Part of "),
			elem.Anchor(
				vecty.Markup(
					prop.Href("http://todomvc.com"),
				),
				vecty.Text("TodoMVC"),
			),
		),
	)
}

func (p *PageView) renderItemList() *vecty.HTML {
	log.Println("pageview renderItemList1 ", "p.newItemTitle =", p.newItemTitle)
	var items vecty.List
	for i, item := range p.Store.Items {
		if (p.Store.Filter == model.Active && item.Completed) || (p.Store.Filter == model.Completed && !item.Completed) {
			continue
		}
		items = append(items, &ItemView{Index: i, Item: item, Store: p.Store})
	}
	log.Println("pageview renderItemList2 ", "p.newItemTitle =", p.newItemTitle)
	return elem.Section(
		vecty.Markup(
			vecty.Class("main"),
		),

		elem.Input(
			vecty.Markup(
				vecty.Class("toggle-all"),
				prop.ID("toggle-all"),
				prop.Type(prop.TypeCheckbox),
				prop.Checked(p.Store.CompletedItemCount() == len(p.Store.Items)),
				event.Change(p.onToggleAllCompleted),
			),
		),
		elem.Label(
			vecty.Markup(
				prop.For("toggle-all"),
			),
			vecty.Text("Mark all as complete"),
		),

		elem.UnorderedList(
			vecty.Markup(
				vecty.Class("todo-list"),
			),
			items,
		),
	)
}
