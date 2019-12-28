package tom80

const AudioCount uint16 = 8
const AudioBuf uint16 = 8

type Audio chan byte

func MkAudio() Audio {
	a := make(Audio, AudioBuf)
	return a
}

func (a Audio) Read() byte {
	return <-a
}

func (a Audio) Write(data byte) {
	a <- data
}
