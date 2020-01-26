package main

import "github.com/superloach/tom80"

// hold 64 cached color values
var colors [1 << 6]cacheRGBA

// type for cached color values
type cacheRGBA [4]uint32

// use cacheRGBA as a color.Color
func (cp cacheRGBA) RGBA() (r, g, b, a uint32) {
	r = cp[0]
	g = cp[1]
	b = cp[2]
	a = 0xFFFF
	return
}

// fill colors table
func init() {
	i := 0
	for i < 1<<6 {
		r, g, b, a := tom80.Pixel(i).RGBA()
		colors[i] = cacheRGBA{r, g, b, a}
		i++
	}
}
