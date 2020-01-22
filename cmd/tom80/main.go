package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"

	"github.com/superloach/tom80"
)

var cons *tom80.Tom80
var auds []*audio.Player
var info tom80.ROMInfo

func controlsLoop() {
	con := cons.IO.Controls[0]

	// update controls twice per frame
	for range time.Tick(time.Second / time.Duration(60 * 2)) {
		con.Lock()

		con.Up = ebiten.IsKeyPressed(ebiten.KeyW)
		con.Left = ebiten.IsKeyPressed(ebiten.KeyA)
		con.Down = ebiten.IsKeyPressed(ebiten.KeyS)
		con.Right = ebiten.IsKeyPressed(ebiten.KeyD)
		con.A = ebiten.IsKeyPressed(ebiten.KeyComma)
		con.B = ebiten.IsKeyPressed(ebiten.KeyPeriod)
		con.C = ebiten.IsKeyPressed(ebiten.KeySlash)
		con.Menu = ebiten.IsKeyPressed(ebiten.KeyEscape)

		con.Unlock()
	}
}

func update(screen *ebiten.Image) error {
	cons.Paused = !ebiten.IsForeground()

	select {
	case b, ok := <-cons.IO.Debug.Text:
		if ok {
			fmt.Printf("%q %d %X\n", string([]byte{b}), b, b)
		}
	default:
	}

	if !ebiten.IsDrawingSkipped() {
		// dump video memory
		v := cons.MEM.DumpVID()

		// display video memory
		for i, b := range v {
			screen.Set(
				i%int(tom80.VIDWidth),
				i/int(tom80.VIDWidth),
				tom80.Pixel(b),
			)
		}
	}

	if cons.Paused {
		white := tom80.Pixel(0b11111111)
		screen.Set(0, 0, white)
		screen.Set(2, 0, white)
		screen.Set(0, 1, white)
		screen.Set(2, 1, white)
		screen.Set(0, 2, white)
		screen.Set(2, 2, white)
	} else {
		i := 0
		for i < cons.IPF {
			cons.CPU.DoOpcode()
			i++
		}
	}

	return nil
}

func main() {
	go controlsLoop()
	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(update, int(tom80.VIDWidth), int(tom80.VIDHeight), 8, "Tom80")
	if err != nil {
		panic(err)
	}
}
