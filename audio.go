package main

import (
	"github.com/gordonklaus/portaudio"
)

type Audio struct {
	stream *portaudio.Stream
	buffer []float32
}

func NewAudio(bitrate int, bufferSize int) (audio *Audio) {

	// Create new Audio structure
	audio = &Audio{}
	audio.buffer = make([]float32, bufferSize)

	// Initialize portaudio
	portaudio.Initialize()

	// Open stream for recording
	stream, err := portaudio.OpenDefaultStream(1, 0, float64(bitrate), bufferSize, audio.buffer)
	if err != nil {
		panic(err)
	}
	audio.stream = stream

	return audio
}

func (a *Audio) Cleanup() {
	err := a.stream.Stop()
	if err != nil {
		panic(err)
	}
	a.stream.Close()
	portaudio.Terminate()
}

// vim: set tabstop=4:
