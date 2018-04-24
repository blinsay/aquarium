package main

import (
	"image"
	"log"
	"math/rand"
	"time"

	"github.com/blinsay/aquarium/termdraw"
	termbox "github.com/nsf/termbox-go"
)

const (
	tickTime  = 200 * time.Millisecond
	spawnTime = 2 * time.Second
)

func main() {
	canvas, err := newTermboxCanvas(termbox.ColorWhite, termbox.ColorBlack)
	if err != nil {
		log.Fatalf("error starting up: %s", err)
	}
	defer canvas.Close()

	termw, termh := termbox.Size()

	a := &aquarium{
		bounds:       image.Rectangle{Max: image.Point{X: termw + 10, Y: termh}},
		canvasOffset: image.Point{X: 5},
	}
	a.spawners = []spawner{
		// anglers
		{name: "angler (L)", chance: 0.1, speed: -2, sprite: &anglerL},
		{name: "angler (R)", chance: 0.1, speed: 2, sprite: &anglerR},
		// minnows
		{name: "minnow alone", chance: 0.2, speed: -1, sprite: &minnowL},
		{name: "minnow school (small)", chance: 0.1, speed: -1, sprite: &minnowSchoolSmallL},
		{name: "minnow school (med)", chance: 0.05, speed: -1, sprite: &minnowSchoolMedL},
		{name: "minnow school (lol)", chance: 0.01, speed: -1, sprite: &minnowUnionL},
		// mackrel
		{name: "mackrel (R)", chance: 0.5, speed: 4, sprite: &mackrelR},
	}
	for i := 0; i < a.bounds.Max.X; i += 5 {
		i += rand.Intn(5)
		a.fish = append(a.fish, &fish{
			name:     "seaweed",
			position: image.Point{X: i, Y: termh - (3 + rand.Intn(2))}, // FIXME: use the bounding box of the sprite
			sprite:   copySeaweed(seaweed),
		})
	}
	a.draw(canvas)
	canvas.flush()

	ticker := time.NewTicker(tickTime)
	defer ticker.Stop()
	spawner := time.NewTicker(spawnTime)
	defer spawner.Stop()

	events := pollTermboxEvents()

	for {
		select {
		case event := <-events:
			if shouldQuit(event) {
				return
			}
		case <-spawner.C:
			a.spawn()
		case <-ticker.C:
			a.move()
			canvas.clear()
			a.draw(canvas)
			canvas.flush()
		}
	}
}

// FIXME: this is extremely dumb. make a spawnSeaweed func or something.
func copySeaweed() *termdraw.AnimatedSprite {
	copy := seaweed
	return &copy
}

func pollTermboxEvents() chan termbox.Event {
	events := make(chan termbox.Event)
	go func() {
		for {
			events <- termbox.PollEvent()
		}
	}()
	return events
}

func shouldQuit(e termbox.Event) bool {
	if e.Type != termbox.EventKey {
		return false
	}

	return e.Key == termbox.KeyEsc || e.Ch == 'q'
}
