package main

import (
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/superloach/sampler"
	"github.com/superloach/tom80"
)

var auds [tom80.AudioCount]Aud

type Aud struct {
	*sampler.Sampler
	*audio.Player
}

var sampleRate int = 44100

func init() {
	var err error

	ctx := audio.CurrentContext()
	if ctx == nil {
		ctx, err = audio.NewContext(sampleRate)
		if err != nil {
			panic(err)
		}
	}
	sampleRate = ctx.SampleRate()

	sn := 0
	for sn < 3 {
		sine := &sampler.Sampler{Rate: sampleRate, Func: sampler.Sine}
		player, _ := audio.NewPlayer(ctx, sine)
		auds[sn] = Aud{sine, player}
		sn++
	}

	qn := 0
	for qn < 2 {
		square := &sampler.Sampler{Rate: sampleRate, Func: sampler.Square}
		player, _ := audio.NewPlayer(ctx, square)
		auds[3+qn] = Aud{square, player}
		qn++
	}

	wn := 0
	for wn < 2 {
		saw := &sampler.Sampler{Rate: sampleRate, Func: sampler.Saw}
		player, _ := audio.NewPlayer(ctx, saw)
		auds[5+wn] = Aud{saw, player}
		wn++
	}

	noise := &sampler.Sampler{Rate: sampleRate, Func: sampler.Noise}
	player, _ := audio.NewPlayer(ctx, noise)
	auds[7] = Aud{noise, player}
}
