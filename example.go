package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/VinceJnz/go-wasm-vecty/actions"
	"github.com/VinceJnz/go-wasm-vecty/components"
	"github.com/VinceJnz/go-wasm-vecty/store"
	"github.com/VinceJnz/go-wasm-vecty/store/model"
	"github.com/hexops/vecty"
)

const debugTag = "main_"

func main() {
	var appStore *store.Store
	appStore = store.New()
	attachLocalStorage(appStore)

	vecty.SetTitle("GopherJS • TodoMVC")
	vecty.AddStylesheet("https://rawgit.com/tastejs/todomvc-common/master/base.css")
	vecty.AddStylesheet("https://rawgit.com/tastejs/todomvc-app-css/master/index.css")
	p := &components.PageView{Store: appStore}

	appStore.Listeners.Add(p, func() {
		vecty.Rerender(p)
	})
	vecty.RenderBody(p)
}

func attachLocalStorage(store *store.Store) {
	store.Listeners.Add(nil, func() { // Anonymous function stores data in the web browser local storage.
		data, err := json.Marshal(store.Items)
		if err != nil {
			println("failed to store items: " + err.Error())
		}
		js.Global().Get("localStorage").Set("items", string(data))
	})

	// gets data from the web browser local storage and puts it into store.Store.Items slice (a slice of model.Item structures)
	if data := js.Global().Get("localStorage").Get("items"); !data.IsUndefined() {
		var items []*model.Item
		if err := json.Unmarshal([]byte(data.String()), &items); err != nil {
			println("failed to load items: " + err.Error())
		}
		store.Dispatcher.Dispatch(&actions.ReplaceItems{
			Items: items,
		})
	}
}
