package tom80

//import "fmt"

type IO struct {
	tom80    *Tom80
	Debug    *Debug
	Controls [ControlCount]*Control
	Audios   [AudioCount]Audio
}

func MkIO(tom80 *Tom80) *IO {
	i := &IO{}
	i.tom80 = tom80
	i.Debug = MkDebug()
	for n := range i.Controls {
		i.Controls[n] = MkControl()
	}
	for n := range i.Audios {
		i.Audios[n] = MkAudio()
	}
	return i
}

func (i *IO) ReadPort(address uint16) byte {
	return i.ReadPortInternal(address, false)
}

func (i *IO) WritePort(address uint16, value byte) {
	i.WritePortInternal(address, value, false)
}

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

func (i *IO) ContendPortPreio(address uint16)  { return }
func (i *IO) ContendPortPostio(address uint16) { return }
