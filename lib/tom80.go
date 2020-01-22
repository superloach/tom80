package tom80

import "github.com/remogatto/z80"

// A wrapper for the system components.
type Tom80 struct {
	MEM    *MEM
	IO     *IO
	CPU    *z80.Z80

	IPF    int
	Paused bool
}

// Make a Tom80 system.
func MkTom80() *Tom80 {
	t := &Tom80{}
	t.MEM = MkMEM()
	t.IO = MkIO()
	t.CPU = z80.NewZ80(t.MEM, t.IO)
	t.CPU.SetPC(ROMStart)
	t.IPF = 50
	t.Paused = false
	return t
}
