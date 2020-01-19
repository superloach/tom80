package main

import (
	"flag"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"

	tom80 "github.com/superloach/tom80/lib"
)

var cons *tom80.Tom80
var auds []*audio.Player
var info tom80.ROMInfo

func init() {
	cons = tom80.MkTom80()

	game := flag.String("game", "", "`file` to load game from")
	flag.Parse()

	if *game == "" {
		println("please specify a game to load (-game file)")
		os.Exit(1)
	}

	err, info := cons.MEM.LoadROMFile(*game)
	if err != nil {
		print("unable to load rom ")
		println(*game)
		print("error: ")
		println(err.Error())
		os.Exit(1)
	} else {
		print("loaded ")
		println(info.Name())
	}
}

func opLoop() {
	for range time.Tick(time.Second / time.Duration(tom80.Cycles)) {
		if ebiten.IsForeground() {
			cons.CPU.DoOpcode()
		}
	}
}

func controlsLoop() {
	con := cons.IO.Controls[0]

	for {
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

	return nil
}

func main() {
	go opLoop()
	go controlsLoop()
	go audsLoop()
	err := ebiten.Run(
		update,
		int(tom80.VIDWidth),
		int(tom80.VIDHeight),
		8,
		"Tom80",
	)
	if err != nil {
		panic(err)
	}
}
