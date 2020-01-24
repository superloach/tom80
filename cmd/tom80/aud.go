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
	sine := &sampler.Sampler{
		Rate: sampleRate,
		Func: sampler.Sine,
	}
	player, err := audio.NewPlayer(ctx, sine)
	if err != nil {
		panic(err)
	}
	auds[0] = Aud{sine, player}
}
