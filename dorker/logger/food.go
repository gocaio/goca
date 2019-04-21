package logger

import (
	"math/rand"
	"os"
	"strings"
)

type food struct {
	emoji        rune
	points, x, y int
}

func newFood(x, y int) *food {
	return &food{
		points: 10,
		emoji:  getFoodEmoji(),
		x:      x,
		y:      y,
	}
}

func getFoodEmoji() rune {
	if hasUnicodeSupport() {
		return randomFoodEmoji()
	}

	return '@'
}

func randomFoodEmoji() rune {
	f := []rune{
		'G',
		// '🍒',
		// '🍍',
		// '🍑',
		// '🍇',
		// '🍏',
		// '🍌',
		// '🍫',
		// '🍭',
		// '🍕',
		// '🍩',
		// '🍗',
		// '🍖',
		// '🍬',
		// '🍤',
		// '🍪',
	}

	return f[rand.Intn(len(f))]
}

func hasUnicodeSupport() bool {
	return strings.Contains(os.Getenv("LANG"), "UTF-8")
}
