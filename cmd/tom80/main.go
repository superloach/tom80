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

func opLoop() {
	for range time.Tick(time.Second / time.Duration(cons.IPF*60)) {
		if !cons.Paused {
			cons.CPU.DoOpcode()
		}
	}
}

func debugTextLoop() {
	for b := range cons.IO.Debug.Text {
		fmt.Printf("%q %d %X\n", string([]byte{b}), b, b)
	}
}

func controlsLoop() {
	con := cons.IO.Controls[0]

	for range time.Tick(time.Second / time.Duration(cons.IPF*5*60)) {
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

func audsLoop() {
	for a1 := range cons.IO.Audios[0] {
		if a1 != 0x00 {
			auds[0].Rewind()
			auds[0].Play()
		}
	}
}

func update(screen *ebiten.Image) error {
	cons.Paused = !ebiten.IsForeground()

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
	}

	return nil
}

func main() {
	go opLoop()
	go debugTextLoop()
	go controlsLoop()
	go audsLoop()
	ebiten.SetRunnableInBackground(true)
	err := ebiten.Run(update, int(tom80.VIDWidth), int(tom80.VIDHeight), 8, "Tom80")
	if err != nil {
		panic(err)
	}
}
