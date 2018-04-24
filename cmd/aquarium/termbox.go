package main

import (
	"image"

	termbox "github.com/nsf/termbox-go"
)

type termboxCanvas struct {
	fg termbox.Attribute
	bg termbox.Attribute

	cells  []termbox.Cell
	width  int
	height int
}

func newTermboxCanvas(fg, bg termbox.Attribute) (*termboxCanvas, error) {
	if err := termbox.Init(); err != nil {
		return nil, err
	}

	canvas := termboxCanvas{fg: fg, bg: bg}
	canvas.clear()

	return &canvas, nil
}

func (t *termboxCanvas) Close() {
	termbox.Close()
}

func (t *termboxCanvas) clear() {
	termbox.Clear(termbox.ColorDefault, t.bg)

	t.cells = termbox.CellBuffer()
	t.width, t.height = termbox.Size()
}

func (t *termboxCanvas) flush() {
	termbox.Flush()

	t.cells = termbox.CellBuffer()
	t.width, t.height = termbox.Size()
}

func (t *termboxCanvas) Bounds() image.Rectangle {
	return image.Rectangle{
		Max: image.Point{X: t.width, Y: t.height},
	}
}

func (t *termboxCanvas) Set(x, y int, ch rune) {
	if x < 0 || x >= t.width {
		return
	}
	if y < 0 || y >= t.height {
		return
	}
	t.cells[y*t.width+x] = termbox.Cell{ch, t.fg, t.bg}
}
