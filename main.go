package main

import (
	"fmt"
	"strconv"
)

var width = 8192
var threshold = 150.0

func main() {
	// Setup audio.
	audio := NewAudio(width, width)
	defer audio.Cleanup()

	// Start streaming audio.
	err := audio.stream.Start()
	if err != nil {
		panic(err)
	}
	for {
		// Do FFT.
		var max int
		max = doFFT(audio)

		// Render
		note := getClosestNote(float64(max))
		diff := float64(max) - note.Frequency
		fmt.Println("Frequency: " + strconv.Itoa(max) +
			"\tNote: " + note.Name +
			" (" + strconv.FormatFloat(diff, 'f', 2, 64) +
			")")
	}
}

// vim: set tabstop=4:
