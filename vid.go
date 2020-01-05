package tom80

const VIDStart uint16 = 0x0000
const VIDWidth uint16 = 64
const VIDHeight uint16 = 48
const VIDSize uint16 = VIDWidth * VIDHeight
const VIDEnd uint16 = VIDStart + VIDSize

func (m *MEM) DumpVID() []byte {
	v := make([]byte, VIDSize)
	for i := uint16(0x00); i < VIDSize; i++ {
		v[i] = m.ReadByte(VIDStart + i)
	}
	return v
}

func (m *MEM) ClearVID() {
	for i := uint16(0x00); i < VIDSize; i++ {
		m.WriteByte(VIDStart + i, 0x00)
	}
}
