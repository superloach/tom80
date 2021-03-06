package tom80

import "github.com/remogatto/z80"

const DefaultClock int = 1000000 // 1 MHz

// A wrapper for the system components.
type Tom80 struct {
	MEM *MEM
	IO  *IO
	CPU *z80.Z80

	Clock  int
	Tick   int
	Paused bool
}

// Make a Tom80 system.
func MkTom80() *Tom80 {
	t := &Tom80{}
	t.MEM = MkMEM()
	t.IO = MkIO()
	t.CPU = z80.NewZ80(t.MEM, t.IO)
	t.CPU.SetPC(ROMStart)
	t.Clock = DefaultClock
	t.Tick = 0
	t.Paused = false
	return t
}
