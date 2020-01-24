package tom80

import (
	"os"
	"path"
	"strings"
)

const (
	ROMStart uint16 = VIDEnd
	ROMSize  uint16 = 0xffff - VIDSize // ~60 KiB
	ROMEnd   uint16 = ROMStart + ROMSize
)

// Abstraction of ROM information.
type ROMInfo map[string]string

// Name of the ROM.
func (r ROMInfo) Name() string {
	name, ok := r["name"]
	if !ok {
		return "UNKNOWN"
	}
	return name
}

// Whether VRAM should be cleared.
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

// Load a ROM from data.
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

// Load a ROM from a file.
func (m *MEM) LoadROMFile(name string) (ROMInfo, error) {
	name = path.Join(".", "prgm", name, name+".bin")

	f, err := os.Open(name)
	if err != nil {
		return map[string]string{}, err
	}

	data := make([]byte, ROMSize)

	_, err = f.Read(data[:])
	if err != nil {
		return map[string]string{}, err
	}

	return m.LoadROM(data[:]), nil
}

// Gather ROM info from the VRAM region.
func (m *MEM) GetROMInfo() ROMInfo {
	info := make(ROMInfo)

	bs := []byte{}
	for i := uint16(0); i < VIDEnd; i++ {
		b := m.ReadByte(i)
		if b == 0x00 || b == '\n' {
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
