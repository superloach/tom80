package tom80

const DebugTextBuf int = 8

type Debug struct {
	Mode int
	Text chan byte
}

func MkDebug() *Debug {
	d := &Debug{}
	d.Text = make(chan byte, DebugTextBuf)
	return d
}

func (d *Debug) Read() byte {
	var mode int = d.Mode
	d.Mode = 0

	switch mode {
	default:
		return 0x00
	}
}

func (d *Debug) Write(data byte) {
	var mode int = d.Mode
	d.Mode = 0

	switch mode {
	case 0:
		d.Mode = int(data)
	case 1:
		d.Text <- data
	default:
		return
	}
}
