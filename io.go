package tom80

//import "fmt"

// Handler for port I/O.
//
// implements z80.PortAccessor
type IO struct {
	Debug    *Debug
	Controls [ControlCount]*Control
	Audios   [AudioCount]Audio
}

// Make a port I/O handler.
func MkIO() *IO {
	i := &IO{}
	i.Debug = MkDebug()
	for n := range i.Controls {
		i.Controls[n] = MkControl()
	}
	for n := range i.Audios {
		i.Audios[n] = MkAudio()
	}
	return i
}

// implements z80.PortAccessor
func (i *IO) ReadPort(address uint16) byte {
	return i.ReadPortInternal(address, false)
}

// implements z80.PortAccessor
func (i *IO) WritePort(address uint16, value byte) {
	i.WritePortInternal(address, value, false)
}

// implements z80.PortAccessor
func (i *IO) ReadPortInternal(address uint16, contend bool) byte {
	address &= 0x00FF
	switch {
	case address == 0x00:
		return i.Debug.Read()
	case address >= 0x01 && address < 0x01+ControlCount:
		p := address - 0x01
		v := i.Controls[p].Read()
		return v
	case address >= 0x01+ControlCount && address < 0x01+ControlCount+AudioCount:
		return i.Audios[address-(0x01+ControlCount)].Read()
	}
	return 0x00
}

// implements z80.PortAccessor
func (i *IO) WritePortInternal(address uint16, value byte, contend bool) {
	address &= 0x00FF
	switch {
	case address == 0x00:
		i.Debug.Write(value)
	case address >= 0x01 && address < 0x01+ControlCount:
		i.Controls[address-0x01].Write(value)
	case address >= 0x01+ControlCount && address < 0x01+ControlCount+AudioCount:
		i.Audios[address-(0x01+ControlCount)].Write(value)
	}
}

// implements z80.PortAccessor
func (i *IO) ContendPortPreio(address uint16) { return }

// implements z80.PortAccessor
func (i *IO) ContendPortPostio(address uint16) { return }
