package tom80

import "os"

const ROMStart uint16 = VIDEnd
const ROMSize uint16 = 0xffff - VIDSize // ~60 KiB
const ROMEnd uint16 = ROMStart + ROMSize

func (m *MEM) LoadROM(data []byte) {
	for i, b := range data {
		m.WriteByte(uint16(i), b)
	}
}

func (m *MEM) LoadROMFile(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}

	data := make([]byte, ROMSize)

	_, err = f.Read(data[:])
	if err != nil {
		return err
	}

	m.LoadROM(data[:])

	return nil
}
