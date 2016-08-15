package main

import (
	"math"

	"github.com/mjibson/go-dsp/fft"
	"github.com/mjibson/go-dsp/window"
)

var fin = make([]float64, width)

func doFFT(audio *Audio, out []float64) int {

	// Read remaining bits
	n, _ := audio.stream.AvailableToRead()
	if n > 0 {
		audio.stream.Read()
	}

	// Convert to float64
	for i := 0; i < width; i++ {
		fin[i] = float64(audio.buffer[i])
	}

	// Apply Hamming window
	window.Apply(fin, window.Hamming)

	// Compute FFT
	x := fft.FFTReal(fin)

	// Compute sqrt(x^2)
	var max = threshold
	var max_i = 0
	for i := 0; i < width; i++ {
		re := real(x[i])
		im := imag(x[i])
		fin[i] = math.Sqrt(re*re + im*im)

		// Find max
		if fin[i] > max {
			max = fin[i]
			max_i = i
		}

		// Scale down output
		out[i/32] = fin[i/2]
	}

	// Render
	return max_i
}

// vim: set tabstop=4:
