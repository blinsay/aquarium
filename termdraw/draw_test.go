package termdraw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimtedSprite_Advance(t *testing.T) {
	sprite := NewAnimatedSprite([]string{
		"-=-=",
		"=-=-",
	})

	assert.Equal(t, 0, sprite.currentSprite)
	sprite.Advance()
	assert.Equal(t, 1, sprite.currentSprite)
}
