package main

var pixels [1<<6]([3]uint32)

// Fill pixels table
func init() {
	i := 0
	for i < 1<<6 {
		pixels[i] = pixel(i).genRGBA()
		i++
	}
}

// 6-bit color type
type pixel byte

// Get RBGA values from 6-bit color
func (p pixel) RGBA() (r, g, b, a uint32) {
	rgba := pixels[int(p)]
	return rgba[0], rgba[1], rgba[2], 0xFFFF
}

func (p pixel) genRGBA() (rgba [3]uint32) {
	q := uint32(p)
	rgba[0] = ((q & (0b11 << 4)) >> 4) * 0x5555
	rgba[1] = ((q & (0b11 << 2)) >> 2) * 0x5555
	rgba[2] = ((q & (0b11 << 0)) >> 0) * 0x5555
	return
}
