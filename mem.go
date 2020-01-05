package tom80

const MEMSize uint16 = VIDSize + ROMSize

type MEM struct {
	tom80   *Tom80
	data    [MEMSize]byte
	ROMInfo map[string]string
}

func MkMEM(tom80 *Tom80) *MEM {
	m := MEM{}
	m.tom80 = tom80
	m.data = [MEMSize]byte{}
	return &m
}

func (m *MEM) ReadByte(address uint16) byte {
	return m.ReadByteInternal(address)
}

func (m *MEM) ReadByteInternal(address uint16) byte {
	return m.data[address]
}

func (m *MEM) WriteByte(address uint16, value byte) {
	m.WriteByteInternal(address, value)
}

func (m *MEM) WriteByteInternal(address uint16, value byte) {
	m.data[address] = value
}

func (m *MEM) ContendRead(address uint16, time int) { return }

func (m *MEM) ContendReadNoMreq(address uint16, time int)                  { return }
func (m *MEM) ContendReadNoMreq_loop(address uint16, time int, count uint) { return }

func (m *MEM) ContendWriteNoMreq(address uint16, time int)                  { return }
func (m *MEM) ContendWriteNoMreq_loop(address uint16, time int, count uint) { return }

func (m *MEM) Read(address uint16) byte {
	return m.ReadByte(address)
}

func (m *MEM) Write(address uint16, value byte, protectROM bool) {
	m.WriteByte(address, value)
}

func (m *MEM) Data() []byte {
	return m.data[:]
}
