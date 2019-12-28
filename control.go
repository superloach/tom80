package tom80

const ControlCount uint16 = 8

type Control byte

func MkControl() *Control {
	c := Control(0x00)
	return &c
}

func (c *Control) Read() byte {
	return byte(*c)
}

func (c *Control) Write(data byte) {
	*c = Control(data)
}
