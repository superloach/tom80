//+build wasm

package main

import (
	"strconv"
	"syscall/js"

	"github.com/superloach/tom80"
)

func init() {
	cons = tom80.MkTom80()

	version_info()

	window := js.Global()
	document := window.Get("document")
	body := document.Get("body")

	ready := make(chan []byte)

	rom_select := document.Call("createElement", "input")
	rom_select.Set("type", "file")
	rom_select.Set("onchange", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
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

	set_clock := document.Call("createElement", "input")
	set_clock.Set("type", "button")
	set_clock.Set("value", "Set Clock Speed")
	set_clock.Set("onclick", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		cur_clock := strconv.Itoa(cons.Clock)
		prompt := window.Call("prompt", "enter new clock speed", cur_clock).String()
		new_clock, err := strconv.Atoi(prompt)
		if err != nil {
			window.Call("alert", "could not set clock speed to " + prompt)
		} else {
			cons.Clock = new_clock
		}
		return nil
	}))

	body.Call("prepend", set_clock)
	body.Call("prepend", rom_select)

	info := cons.MEM.LoadROM(<-ready)
	window.Call("alert", "loaded "+info.Name())
}
