package main

import (
	"fmt"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	raudio "github.com/hajimehoshi/ebiten/examples/resources/audio"

	"github.com/superloach/tom80"
)

type Dist struct {
	tom80  *tom80.Tom80
	audios []*audio.Player
}

func MkDist() *Dist {
	d := Dist{}

	t := tom80.MkTom80()
	d.tom80 = t

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

	d.audios = append(d.audios, audioPlayer)

	if len(os.Args) > 1 {
		err := d.tom80.MEM.LoadROMFile(os.Args[1])
		if err != nil {
			fmt.Println("unable to load rom", os.Args[1])
		} else {
			fmt.Println("loaded rom", os.Args[1])
		}
	}

	return &d
}

func (d *Dist) OpLoop() {
	for range time.Tick(time.Second / time.Duration(tom80.Cycles)) {
		d.tom80.CPU.DoOpcode()
	}
}

func (d *Dist) ControlsLoop() {
	for {
		d.tom80.IO.Controls[0].Up = ebiten.IsKeyPressed(ebiten.KeyW)
		d.tom80.IO.Controls[0].Down = ebiten.IsKeyPressed(ebiten.KeyA)
		d.tom80.IO.Controls[0].Left = ebiten.IsKeyPressed(ebiten.KeyS)
		d.tom80.IO.Controls[0].Right = ebiten.IsKeyPressed(ebiten.KeyD)
		d.tom80.IO.Controls[0].A = ebiten.IsKeyPressed(ebiten.KeyComma)
		d.tom80.IO.Controls[0].B = ebiten.IsKeyPressed(ebiten.KeyPeriod)
		d.tom80.IO.Controls[0].C = ebiten.IsKeyPressed(ebiten.KeySlash)
		d.tom80.IO.Controls[0].Menu = ebiten.IsKeyPressed(ebiten.KeyEscape)
	}
}

func (d *Dist) DebugTextLoop() {
	for v := range d.tom80.IO.Debug.Text {
		print(string([]byte{v}))
	}
}

func (d *Dist) AudiosLoop() {
	for a1 := range d.tom80.IO.Audios[0] {
		if a1 != 0x00 {
			d.audios[0].Rewind()
			d.audios[0].Play()
		}
	}
}

func (d *Dist) Run() error {
	go d.OpLoop()
	go d.ControlsLoop()
	go d.DebugTextLoop()
	go d.AudiosLoop()
	return ebiten.RunGame(d)
}

func (d *Dist) Update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		// dump video memory
		v := d.tom80.MEM.DumpVID()

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

func (d *Dist) Layout(ow, oh int) (int, int) {
	return int(tom80.VIDWidth), int(tom80.VIDHeight)
}

func main() {
	d := MkDist()
	err := d.Run()
	if err != nil {
		panic(err)
	}
}
