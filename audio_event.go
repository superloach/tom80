package tom80

type AudioEvent byte

func (a AudioEvent) Volume() int {
	return int(a >> 6)
}

func (a AudioEvent) Pitch() int {
	return int(a & 0x3F)
}
