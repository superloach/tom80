//+build wasm

package main

import (
	"syscall/js"

	"github.com/superloach/tom80"
)

func init() {
	cons = tom80.MkTom80()

	window := js.Global()
	document := window.Get("document")
	body := document.Get("body")

	ready := make(chan []byte)

	input := document.Call("createElement", "input")
	input.Set("type", "file")
	input.Set("onchange", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return nil
		}
		files := args[0].Get("target").Get("files")
		length := files.Get("length").Int()
		if length != 1 {
			return nil
		}

		file := files.Index(0)
		reader := window.Get("FileReader").New()
		reader.Call("readAsArrayBuffer", file)
		reader.Set("onload", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if len(args) < 1 {
				return nil
			}
			event := args[0]
			content := event.Get("target").Get("result")
			src := window.Get("Uint8Array").New(content)
			dst := make([]byte, tom80.MEMSize)
			_ = js.CopyBytesToGo(dst, src)
			ready <- dst
			return nil
		}))
		return nil
	}))

	body.Call("prepend", input)

	select {
	case v := <-ready:
		info := cons.MEM.LoadROM(v)
		window.Call("alert", "loaded " + info.Name())
	}
}
