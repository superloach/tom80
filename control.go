package tom80

const (
	ControlCount uint16 = 8
)

const (
	BtnUp	int = (7-iota)
	BtnDown
	BtnLeft
	BtnRight
	BtnA
	BtnB
	BtnC
	BtnMenu
)

// Abstraction of a controller.
type Control byte

func MkControl() *Control {
	c := Control(0)
	return &c
}

// Read from controller's buttons.
func (c *Control) Read() byte {
	return byte(*c)
}

// Write to controller's buttons.
func (c *Control) Write(data byte) {
	*c = Control(data)
}

// Check if a button is pressed
func (c *Control) Pressed(btn int) bool {
	return *c&1<<btn>0
}

// Set a button's state
func (c *Control) Press(btn int, state bool) {
	mask := Control(1<<btn)
	if state {
		*c |= mask
	} else {
		*c &^= mask
	}
}
