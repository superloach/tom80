package tom80

type IO struct {
	tom80    *Tom80
	Devices map[uint16]byte
}

func MkIO(tom80 *Tom80) *IO {
	i := IO{}
	i.tom80 = tom80
	i.Devices = make(map[uint16]byte)
	return &i
}

func (i *IO) ReadPort(address uint16) byte {
	address &= 0xFF
	return i.ReadPortInternal(address, false)
}

func (i *IO) WritePort(address uint16, value byte) {
	address &= 0xFF
	i.WritePortInternal(address, value, false)
}

func (i *IO) ReadPortInternal(address uint16, contend bool) byte {
	address &= 0xFF
	value := i.Devices[address]
	i.Devices[address] = 0
	return value
}

func (i *IO) WritePortInternal(address uint16, value byte, contend bool) {
	address &= 0xFF
	i.Devices[address] = value
}

func (i *IO) ContendPortPreio(address uint16)  { return }
func (i *IO) ContendPortPostio(address uint16) { return }
