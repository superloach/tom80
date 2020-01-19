package tom80

import (
	"os"
	"strings"
)

const (
	ROMStart uint16 = VIDEnd
	ROMSize uint16 = 0xffff - VIDSize // ~60 KiB
	ROMEnd uint16 = ROMStart + ROMSize
)

type ROMInfo map[string]string

func (r ROMInfo) Name() string {
	name, ok := r["name"]
	if !ok {
		return "UNKNOWN"
	}
	return name
}

func (r ROMInfo) Clear() bool {
	clear, ok := r["clear"]
	if !ok {
		return true
	}
	switch clear {
	case "false", "f", "no", "n":
		return false
	default:
		return true
	}
}

func (m *MEM) LoadROM(data []byte) ROMInfo {
	for i, b := range data {
		m.WriteByte(uint16(i), b)
	}

	info := m.GetROMInfo()

	if info.Clear() {
		m.ClearVID()
	}

	return info
}

func (m *MEM) LoadROMFile(name string) (error, ROMInfo) {
	f, err := os.Open(name)
	if err != nil {
		return err, map[string]string{}
	}

	data := make([]byte, ROMSize)

	_, err = f.Read(data[:])
	if err != nil {
		return err, map[string]string{}
	}

	return nil, m.LoadROM(data[:])
}

// Gather ROM info from the VRAM space.
func (m *MEM) GetROMInfo() ROMInfo {
	info := make(ROMInfo)

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
			info[kv[0]] = kv[1]
			bs = []byte{}
		} else {
			bs = append(bs, b)
		}
	}

	return info
}
