package tom80

const (
	MEMSize uint16 = VIDSize + ROMSize
)

// System memory.
//
// implements z80.MemoryAccessor
type MEM struct {
	data    [MEMSize]byte
}

// Make system memory.
func MkMEM() *MEM {
	m := MEM{}
	m.data = [MEMSize]byte{}
	return &m
}

// implements z80.MemoryAccessor
func (m *MEM) ReadByte(address uint16) byte {
	return m.ReadByteInternal(address)
}

// implements z80.MemoryAccessor
func (m *MEM) ReadByteInternal(address uint16) byte {
	return m.data[address]
}

// implements z80.MemoryAccessor
func (m *MEM) WriteByte(address uint16, value byte) {
	m.WriteByteInternal(address, value)
}

// implements z80.MemoryAccessor
func (m *MEM) WriteByteInternal(address uint16, value byte) {
	m.data[address] = value
}

// implements z80.MemoryAccessor
func (m *MEM) ContendRead(address uint16, time int) { return }

// implements z80.MemoryAccessor
func (m *MEM) ContendReadNoMreq(address uint16, time int)                  { return }

// implements z80.MemoryAccessor
func (m *MEM) ContendReadNoMreq_loop(address uint16, time int, count uint) { return }

// implements z80.MemoryAccessor
func (m *MEM) ContendWriteNoMreq(address uint16, time int)                  { return }

// implements z80.MemoryAccessor
func (m *MEM) ContendWriteNoMreq_loop(address uint16, time int, count uint) { return }

// implements z80.MemoryAccessor
func (m *MEM) Read(address uint16) byte {
	return m.ReadByte(address)
}

// implements z80.MemoryAccessor
func (m *MEM) Write(address uint16, value byte, protectROM bool) {
	m.WriteByte(address, value)
}

// implements z80.MemoryAccessor
func (m *MEM) Data() []byte {
	return m.data[:]
}
