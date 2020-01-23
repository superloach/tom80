package tom80

type AudioEvent byte

func (a AudioEvent) Volume() int {
	return 0
}

func (a AudioEvent) Pitch() int {
	return 0
}

func (a AudioEvent) Hold() int {
	return 0
}
