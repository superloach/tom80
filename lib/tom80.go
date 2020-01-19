package tom80

import "github.com/remogatto/z80"

const (
	Cycles int = 2500000 // 2.5 MHz
)

// A wrapper for the system components.
type Tom80 struct {
	MEM    *MEM
	IO     *IO
	CPU    *z80.Z80
	Paused bool
}

// Make a Tom80 system.
func MkTom80() *Tom80 {
	t := &Tom80{}
	t.MEM = MkMEM()
	t.IO = MkIO()
	t.CPU = z80.NewZ80(t.MEM, t.IO)
	t.CPU.SetPC(ROMStart)
	t.Paused = false
	return t
}
