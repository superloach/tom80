package tom80

const (
	DebugTextBuf int = 64
)

// Abstraction of debug information.
type Debug struct {
	Mode int
	Text chan byte
}

// Make debug info.
func MkDebug() *Debug {
	d := &Debug{}
	d.Text = make(chan byte, DebugTextBuf)
	return d
}

// Read debug data.
func (d *Debug) Read() byte {
	var mode int = d.Mode
	d.Mode = 0

	switch mode {
	default:
		return 0x00
	}
}

// Write debug data.
//
// If mode is 0, set the mode to the data.
//
// If mode is 1, send data to the debug text buffer.
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
