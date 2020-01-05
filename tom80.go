package tom80

import "github.com/remogatto/z80"

const Cycles int = 2500000 // 2.5 MHz
//const Cycles int = 25 // 25 Hz

type Tom80 struct {
	MEM *MEM
	IO  *IO
	CPU *z80.Z80
}

func MkTom80() *Tom80 {
	t := &Tom80{}
	t.MEM = MkMEM(t)
	t.IO = MkIO(t)
	t.CPU = z80.NewZ80(t.MEM, t.IO)
	t.CPU.SetPC(ROMStart)
	return t
}
