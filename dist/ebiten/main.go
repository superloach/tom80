package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	raudio "github.com/hajimehoshi/ebiten/examples/resources/audio"

	"github.com/superloach/tom80"
)

var cons *tom80.Tom80
var auds  []*audio.Player

func init() {
	cons = tom80.MkTom80()

	audioContext, err := audio.NewContext(44100)
	if err != nil { panic(err) }

	aud, err := wav.Decode(audioContext, audio.BytesReadSeekCloser(raudio.Jab_wav))
	if err != nil { panic(err) }

	audioPlayer, err := audio.NewPlayer(audioContext, aud)
	if err != nil { panic(err) }

	auds = append(auds, audioPlayer)

	if len(os.Args) > 1 {
		err := cons.MEM.LoadROMFile(os.Args[1])
		if err != nil {
			fmt.Println("unable to load rom", os.Args[1])
			os.Exit(1)
		} else {
			fmt.Println("loaded rom", os.Args[1])
		}
	}
}

func opLoop() {
	for range time.Tick(time.Second / time.Duration(tom80.Cycles)) {
		cons.CPU.DoOpcode()
	}
}

func controlsLoop() {
	for {
		cons.IO.Controls[0].Up = ebiten.IsKeyPressed(ebiten.KeyW)
		cons.IO.Controls[0].Left = ebiten.IsKeyPressed(ebiten.KeyA)
		cons.IO.Controls[0].Down = ebiten.IsKeyPressed(ebiten.KeyS)
		cons.IO.Controls[0].Right = ebiten.IsKeyPressed(ebiten.KeyD)
		cons.IO.Controls[0].A = ebiten.IsKeyPressed(ebiten.KeyComma)
		cons.IO.Controls[0].B = ebiten.IsKeyPressed(ebiten.KeyPeriod)
		cons.IO.Controls[0].C = ebiten.IsKeyPressed(ebiten.KeySlash)
		cons.IO.Controls[0].Menu = ebiten.IsKeyPressed(ebiten.KeyEscape)
	}
}

func debugTextLoop() {
	for v := range cons.IO.Debug.Text {
		print(string([]byte{v}))
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

	name, ok := cons.MEM.ROMInfo["name"]
	if ok {
		ebitenutil.DebugPrint(screen, name)
	}

	return nil
}

func main() {
	go opLoop()
	go controlsLoop()
	go debugTextLoop()
	go audsLoop()
	err := ebiten.Run(
		update,
		int(tom80.VIDWidth),
		int(tom80.VIDHeight),
		8,
		"Tom80",
	)
	if err != nil { panic(err) }
}
