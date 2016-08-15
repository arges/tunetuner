package main

import "strconv"

var width = 8192
var threshold = 150.0

func main() {
	// Setup buffers.
	out := make([]float64, width/32)

	// Setup audio.
	audio := NewAudio(width, width)
	defer audio.Cleanup()

	// Setup UI.
	ui := NewUI(out)
	defer ui.Cleanup()

	// Start streaming audio.
	err := audio.stream.Start()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			// Do FFT.
			var max int
			max = doFFT(audio, out)

			// Render
			note := getClosestNote(float64(max))
			diff := float64(max) - note.Frequency
			ui.SetText("Frequency: " + strconv.Itoa(max) +
				"\tNote: " + note.Name +
				" (" + strconv.FormatFloat(diff, 'f', 2, 64) +
				")")
			ui.Render()
		}
	}()
	ui.Loop()
}

// vim: set tabstop=4:
