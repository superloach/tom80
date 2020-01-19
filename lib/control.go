package tom80

import "sync"

const (
	ControlCount uint16 = 8
)

// Abstraction of a controller.
type Control struct {
	sync.Mutex
	Up    bool
	Down  bool
	Left  bool
	Right bool
	A     bool
	B     bool
	C     bool
	Menu  bool
}

// Make a controller.
func MkControl() *Control {
	return &Control{}
}

// Read from controller's buttons.
func (c *Control) Read() byte {
	c.Lock()
	defer c.Unlock()

	var b byte = 0x00

	if c.Up {
		b |= 1 << 7
	}
	if c.Down {
		b |= 1 << 6
	}
	if c.Left {
		b |= 1 << 5
	}
	if c.Right {
		b |= 1 << 4
	}
	if c.A {
		b |= 1 << 3
	}
	if c.B {
		b |= 1 << 2
	}
	if c.C {
		b |= 1 << 1
	}
	if c.Menu {
		b |= 1 << 0
	}

	return b
}

// Write to controller's buttons.
func (c *Control) Write(data byte) {
	c.Lock()
	defer c.Unlock()

	c.Up = (data&1<<7 != 0)
	c.Down = (data&1<<6 != 0)
	c.Left = (data&1<<5 != 0)
	c.Right = (data&1<<4 != 0)
	c.A = (data&1<<3 != 0)
	c.B = (data&1<<2 != 0)
	c.C = (data&1<<1 != 0)
	c.Menu = (data&1<<0 != 0)

	return
}
