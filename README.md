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

