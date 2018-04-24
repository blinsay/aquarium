package termdraw

import (
	"image"
	"strings"
	"unicode"
)

// A Canvas is anything that can be drawn on.
type Canvas interface {
	// Return the bounding rectangle for this canvas. Returns a (min, max) point
	// tuple.
	Bounds() image.Rectangle

	// Set places a character on the canvas. Set must silently ignore any
	// calls to set a point outside of its bounds.
	Set(x, y int, ch rune)
}

// A Sprite is a text graphic that can be drawn onto a canvas. See DrawSprite.
type Sprite struct {
	size image.Point
	rows []row
}

const newline = "\n"

// NewSprite converts a string to a Sprite.
//
// All trailing whitespace is removed from the string, including newlines. Empty
// lines are preserved.
func NewSprite(s string) Sprite {
	// TODO(benl): trim leading spaces and actually use row offsets
	lines := strings.Split(s, "\n")
	rows := make([]row, len(lines))

	var longestRow int
	for i := range rows {
		rows[i].chars = strings.TrimRightFunc(lines[i], unicode.IsSpace)

		if l := len(rows[i].chars); l > longestRow {
			longestRow = l
		}
	}

	return Sprite{size: image.Point{X: longestRow, Y: len(rows)}, rows: rows}
}

func not(f func(rune) bool) func(rune) bool {
	return func(r rune) bool {
		return !f(r)
	}
}

// Bounds calculates the bounding rectangle for a sprite when placed at the
// given point.
func (s Sprite) Bounds(at image.Point) image.Rectangle {
	return image.Rectangle{Min: at, Max: at.Add(s.size)}.Canon()
}

type row struct {
	offset int
	chars  string
}

// DrawSprite draws a sprite on the canvas with it's top left corner at (x, y)
func DrawSprite(c Canvas, at image.Point, s *Sprite) {
	if c.Bounds().Intersect(s.Bounds(at)).Empty() {
		return
	}

	for j, row := range s.rows {
		for i, ch := range row.chars {
			c.Set(at.X+row.offset+i, at.Y+j, ch)
		}
	}
}

// An AnimatedSprite is a series of sprites that should be drawn in sequence to
// a canvas.
type AnimatedSprite struct {
	sprites       []Sprite
	size          image.Point
	currentSprite int
}

// NewAnimatedSprite converts a series of strings into animated sprite frames.
func NewAnimatedSprite(frames []string) AnimatedSprite {
	sprites := make([]Sprite, len(frames))
	for i := range frames {
		sprites[i] = NewSprite(frames[i])
	}

	return AnimatedSprite{sprites: sprites}
}

// Current returns a pointer to the current sprite. This pointer will be invalid
// once Advance is called.
func (a *AnimatedSprite) Current() *Sprite {
	return &a.sprites[a.currentSprite]
}

// Advance moves the current frame of the sprite forward by one. Loops at the
// end.
func (a *AnimatedSprite) Advance() {
	a.currentSprite = ((a.currentSprite + 1) % len(a.sprites))
}
