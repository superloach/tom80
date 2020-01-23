package main

// 6-bit color type
type pixel byte

// Get RBGA values from 6-bit color
func (p pixel) RGBA() (r, g, b, a uint32) {
	q := uint32(p)
	r = ((q & (0b11 << 4)) >> 4) * 0x5555
	g = ((q & (0b11 << 2)) >> 2) * 0x5555
	b = ((q & (0b11 << 0)) >> 0) * 0x5555
	a = 0xffff
	return
}
