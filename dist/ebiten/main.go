package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
//	"github.com/hajimehoshi/ebiten/ebitenutil"
	raudio "github.com/hajimehoshi/ebiten/examples/resources/audio"

	"github.com/superloach/tom80"
)

var cons *tom80.Tom80
var auds []*audio.Player
var info tom80.ROMInfo

func init() {
	cons = tom80.MkTom80()

	audioContext, err := audio.NewContext(44100)
	if err != nil {
		panic(err)
	}

	aud, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Jab_wav))
	if err != nil {
		panic(err)
	}

	audioPlayer, err := audio.NewPlayer(audioContext, aud)
	if err != nil {
		panic(err)
	}

	auds = append(auds, audioPlayer)

	if len(os.Args) > 1 {
		err, info := cons.MEM.LoadROMFile(os.Args[1])
		if err != nil {
			fmt.Println("unable to load rom", os.Args[1])
			os.Exit(1)
		} else {
			fmt.Println("loaded", info.Name())
		}
	} else {
		fmt.Println("missing filename")
	}
}

func opLoop() {
	for range time.Tick(time.Second / time.Duration(tom80.Cycles)) {
		cons.CPU.DoOpcode()
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
