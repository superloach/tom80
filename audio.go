package tom80

const (
	AudioCount uint16 = 8
	AudioBuf   uint16 = 64
)

// A buffer of audio events, similar concept to MIDI.
type Audio chan AudioEvent

// Make an Audio.
func MkAudio() Audio {
	a := make(Audio, AudioBuf)
	return a
}

// Read an audio event from the buffer.
func (a Audio) Read() byte {
	return byte(<-a)
}

// Write an audio event to the buffer.
func (a Audio) Write(data byte) {
	a <- AudioEvent(data)
}
