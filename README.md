# go-wasm-vecty
todomvc example built using vecty and deployed via wasm.
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

