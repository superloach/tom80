package tom80

import "fmt"

const ControlCount uint16 = 8

type Control struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
	A     bool
	B     bool
	C     bool
	Menu  bool
}

func MkControl() *Control {
	return &Control{}
}

func (c *Control) Read() byte {
	var b byte = 0x00

	if c.Up {
		b = b | 0x80
	}
	if c.Down {
		b = b | 0x40
	}
	if c.Left {
		b = b | 0x20
	}
	if c.Right {
		b = b | 0x10
	}
	if c.A {
		b = b | 0x08
	}
	if c.B {
		b = b | 0x04
	}
	if c.C {
		b = b | 0x02
	}
	if c.Menu {
		b = b | 0x01
	}

	fmt.Printf("%0#b\n", b)
	return b
}

func (c *Control) Write(data byte) {
	return
}
