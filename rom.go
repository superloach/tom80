package tom80

import "os"
import "strings"

const ROMStart uint16 = VIDEnd
const ROMSize uint16 = 0xffff - VIDSize // ~60 KiB
const ROMEnd uint16 = ROMStart + ROMSize

func (m *MEM) LoadROM(data []byte) {
	for i, b := range data {
		m.WriteByte(uint16(i), b)
	}

	m.GetROMInfo()
	m.ClearVID()
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

func (m *MEM) GetROMInfo() {
	m.ROMInfo = make(map[string]string)

	bs := []byte{}
	for i := uint16(0); i < VIDEnd; i++ {
		b := m.ReadByte(i)
		if b == 0x00 {
			str := string(bs)
			kv := strings.Split(str, ":")
			if len(kv) == 0 {
				break
			} else if len(kv) == 1 {
				kv = append(kv, "")
			}
			m.ROMInfo[kv[0]] = kv[1]
			bs = []byte{}
		} else {
			bs = append(bs, b)
		}
	}
}
