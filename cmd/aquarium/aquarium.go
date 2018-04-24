package main

import (
	"image"
	"math/rand"

	"github.com/blinsay/aquarium/termdraw"
)

type aquarium struct {
	// an aquarium is a rectangle that is at least as large as a canvas. it has
	// edges outside the canvas so that fish can swim offscreen and then back on.
	bounds image.Rectangle

	// the canvas's min (x, y) should be offset a little bit from the aquarium's
	canvasOffset image.Point

	fish     []*fish
	spawners []spawner
}

type spawner struct {
	name   string
	chance float64
	speed  int
	sprite *termdraw.AnimatedSprite
}

func (s *spawner) spawn(bounds image.Rectangle) *fish {
	spriteCopy := *s.sprite
	f := &fish{
		name:   s.name,
		speed:  image.Point{X: s.speed},
		sprite: &spriteCopy,
	}

	// FIXME: use the bounding box of the sprite
	if s.speed < 0 {
		f.position = image.Point{X: bounds.Max.X, Y: rand.Intn(bounds.Dy() - 6)}
	} else {
		f.position = image.Point{X: 0, Y: rand.Intn(bounds.Dy() - 6)}
	}

	return f
}

// move all fish in the aquarium, advancing their sprites one frame each.
// fish swim out one side of the aquarium and disappear.
func (a *aquarium) move() {
	for i := len(a.fish) - 1; i >= 0; i-- {
		a.fish[i].move()

		if bounds := a.fish[i].spriteBounds(); bounds.Intersect(a.bounds).Empty() {
			a.fish = append(a.fish[:i], a.fish[i+1:]...)
		}
	}
}

// spawn new fish based on the amount of food in the aquarium, randomly placing
// any newly spawned fish near an edge.
//
// does not move anything in the aquarium.
func (a *aquarium) spawn() {
	// FIXME: do something with fish food
	for _, spawner := range a.spawners {
		if rand.Float64() > 1-spawner.chance {
			a.fish = append(a.fish, spawner.spawn(a.bounds))
		}
	}
}

// draw the aquarium's contents onto a canvas. redrawing the same state multiple
// times should be possible.
func (a *aquarium) draw(c termdraw.Canvas) {
	for _, fish := range a.fish {
		termdraw.DrawSprite(c, fish.position.Sub(a.canvasOffset), fish.sprite.Current())
	}
}

type fish struct {
	name     string
	sprite   *termdraw.AnimatedSprite
	position image.Point
	speed    image.Point
}

func (f *fish) spriteBounds() image.Rectangle {
	return f.sprite.Current().Bounds(f.position)
}

func (f *fish) move() {
	f.position = f.position.Add(f.speed)
	f.sprite.Advance()
}
