package logger

import (
	"os"
	"testing"
)

func TestFoodDefaultPoints(t *testing.T) {
	f := newFood(10, 10)

	if f.points != 10 {
		t.Fatalf("Expected Food default points to be 10 but got %v", f.points)
	}
}

func TestFoodEmoji(t *testing.T) {
	f := newFood(10, 10)

	if string(f.emoji) == "" {
		t.Fatal("Food emoji not expected to be blank")
	}
}

func TestFoodFallback(t *testing.T) {
	os.Setenv("LANG", "c")

	f := newFood(10, 10)

	if string(f.emoji) != "@" {
		t.Fatal("Food emoji expected to be '@'")
	}
}
