# go-wasm-vecty
A copy of the vecty todomvc example built using golang/vecty and deployed as a wasm.

This update removes all the package global variables from the origional example.
## WASM Notes

Environment vars need to be set in at a CMD line. (The VS Terminal does not work in this instance)

Important not to have spaces in the set statements as it will not work
```
set GOOS=js
set GOARCH=wasm
go build -o main.wasm
```

## Modules
https://github.com/hexops/vecty



## App flow
The way this app works...


When the aplication is started
It creates a new store: Which registers 1 callback with the dispatcher. This is the store action callback function. When this is actioned it makes changes to the store then fires the listenerRegistary.
It then adds 1 listener to the storeutil registry. This is the pageview rerender function. When called, it updates the whole page using the vecty.rerender function.
It then renders the page: Various page events dispatch actions to the dispatcher which then dispatches the given action to all registered callbacks.

Dispatcher implementation
=========================
This manages callbacks by registering/unregistering them in a store. A callback is a function that responds to actions.
When Dispatch() is called it iterates over the callbacks in the store and provides the action to the callback functions.

Store implementation
====================
This creates a Store structure in memory that contains pointers to each type of item that needs to be stored.
For example: Dispatcher, Listner, Data items, etc.

It uses actions to make changes to stored items.
When the store is created it registers its onAction() function with the dispatcher.
When called by the dispatcher, the onAction() function chooses which actions to respond to.

Listner implementation
======================
This is a storeutil
It provides a structure for storing listners: Functions that don't accept or return parameters.
When Fire() is called it iterates through the items in the listner structure and calls the item function.

For example
After an Action is complete, listners can be fired to update long-term stored data (e.g. write data to a database)

Footnote:
=========
In theory it is possible to have multiple Listerner and Dispatcher instances. Not sure if this is a godd idea???
