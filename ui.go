package main

import (
	"fmt"
	"os"

	"github.com/gizak/termui"
)

type UI struct {
	text *termui.Par
}

func NewUI(data []float64) (ui *UI) {
	ui = &UI{}

	// Initialize the UI
	err := termui.Init()
	if err != nil {
		panic(err)
	}

	// Setup FFT Graph
	spl0 := termui.NewLineChart()
	spl0.Data = data
	spl0.LineColor = termui.ColorGreen
	spl0.Height = 32
	spl0.Width = len(data)
	spl0.DataLabels = []string{""}
	spl0.LineColor = termui.ColorGreen | termui.AttrBold
	spl0.Border = true

	// Setup label
	par0 := termui.NewPar("")
	par0.Height = 1
	par0.Width = 512
	par0.X = 0
	par0.Y = 0
	par0.Border = false
	ui.text = par0

	// Compose UI
	termui.Body.AddRows(
		termui.NewRow(
			termui.NewCol(8, 0, spl0)),
		termui.NewRow(
			termui.NewCol(8, 0, par0)),
	)
	termui.Body.Align()

	// Setup key handlers
	termui.Handle("/sys/kbd", func(termui.Event) {
		termui.StopLoop()
		ui.Cleanup()
		os.Exit(0)
	})

	fmt.Printf("\033[H\033[2J")
	ui.Render()

	return
}

func (u *UI) Loop() {
	termui.Loop()
}

func (u *UI) Render() {
	termui.Render(termui.Body)
}

func (u *UI) SetText(text string) {
	u.text.Text = text
}

func (u *UI) Cleanup() {
	termui.Close()
}

// vim: set tabstop=4:
