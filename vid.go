package tom80

const (
	VIDStart uint16 = 0x0000
	VIDWidth uint16 = 64
	VIDHeight uint16 = 48
	VIDSize uint16 = VIDWidth * VIDHeight
	VIDEnd uint16 = VIDStart + VIDSize
)

// Return a slice of memory corresponding to the display.
func (m *MEM) DumpVID() []byte {
	return m.data[VIDStart:VIDEnd]
}

// Set every pixel to 0x00.
func (m *MEM) ClearVID() {
	for i := VIDStart; i < VIDSize; i++ {
		m.WriteByte(VIDStart + i, 0x00)
	}
}
