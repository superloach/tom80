package main

import (
	"os"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	raudio "github.com/hajimehoshi/ebiten/examples/resources/audio"

	"github.com/superloach/tom80"
)

type Dist struct {
	tom80 *tom80.Tom80
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

	d.tom80.IO.Devices[0] = 0
	d.tom80.IO.Devices[1] = 0
	d.tom80.IO.Devices[64] = 0

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

func (d *Dist) Loop() {
	for range time.Tick(time.Second / time.Duration(tom80.Cycles)) {
		d.tom80.CPU.DoOpcode()

		b := byte(0)

		switch {
		case ebiten.IsKeyPressed(ebiten.KeyW):
			b |= 1<<7
		case ebiten.IsKeyPressed(ebiten.KeyA):
			b |= 1<<6
		case ebiten.IsKeyPressed(ebiten.KeyS):
			b |= 1<<5
		case ebiten.IsKeyPressed(ebiten.KeyD):
			b |= 1<<4
		case ebiten.IsKeyPressed(ebiten.KeyComma):
			b |= 1<<3
		case ebiten.IsKeyPressed(ebiten.KeyPeriod):
			b |= 1<<2
		case ebiten.IsKeyPressed(ebiten.KeySlash):
			b |= 1<<1
		case ebiten.IsKeyPressed(ebiten.KeyEscape):
			b |= 1<<0
		}

		d.tom80.IO.Devices[1] = b

		v := d.tom80.IO.ReadPortInternal(0, false)
		if v != 0 {
			print(string([]byte{v}))
		}

		a1 := d.tom80.IO.ReadPortInternal(64, false)
		if a1 != 0 {
			d.audios[0].Rewind()
			d.audios[0].Play()
		}
	}
}

func (d *Dist) Run() error {
	go d.Loop()
	return ebiten.RunGame(d)
}

func (d *Dist) Update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		// dump video memory
		v := d.tom80.MEM.DumpVID()

		// display video memory
		for i, b := range v {
			screen.Set(
				i % int(tom80.VIDWidth),
				i / int(tom80.VIDWidth),
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
