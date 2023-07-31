## Steps son help, mom stuck.

- pastikan telah mengunduh golang
<!--- unduh tinygo -->
- buat project folder
- buat file main.go
- copy wasm_exec.js menggunakan command :

  `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`

- buka settings di vscode

  - Search `go tools env`
  - klik `edit in settings.json`
  - lalu masukan seperti ini :

  ```json
  {
    "go.toolsEnvVars": {
      "GOOS": "js",
      "GOARCH": "wasm"
    }
  }
  ```

  - save.

- coding. 😁

## Note kode penting

```go
	document.Call("addEventListener", "keydown", js.FuncOf(
		func(this js.Value, args []js.Value) any {
			key := args[0].Get("key").String()
			rect := div.Call("getBoundingClientRect")
			go func() {
				switch key {
				case "ArrowUp":
					top := rect.Get("top").Float() - 100
					style.Set("top", js.ValueOf(fmt.Sprintf("%.2fpx", top)))
				case "ArrowDown":
					top := rect.Get("top").Float() + 100
					style.Set("top", js.ValueOf(fmt.Sprintf("%.2fpx", top)))
				case "ArrowLeft":
					left := rect.Get("left").Float() - 100
					style.Set("left", js.ValueOf(fmt.Sprintf("%.2fpx", left)))
				case "ArrowRight":
					left := rect.Get("left").Float() + 100
					style.Set("left", js.ValueOf(fmt.Sprintf("%.2fpx", left)))
				}
			}()
			return nil
		},
	))
```
# pingpongwasm
