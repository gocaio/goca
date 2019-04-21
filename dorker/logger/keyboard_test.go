package logger

import (
	"testing"

	"github.com/nsf/termbox-go"
)

func TestKeyToDirectionDefault(t *testing.T) {
	d := keyToDirection(termbox.KeyEsc)

	if d != 0 {
		t.Fatalf("Expected direction to be 0 but got %v", d)
	}
}

func TestKeyToDirectionRight(t *testing.T) {
	d := keyToDirection(termbox.KeyArrowRight)

	if d != RIGHT {
		t.Fatalf("Expected direction to be RIGHT but got %v", d)
	}
}

func TestKeyToDirectionDown(t *testing.T) {
	d := keyToDirection(termbox.KeyArrowDown)

	if d != DOWN {
		t.Fatalf("Expected direction to be DOWN but got %v", d)
	}
}

func TestKeyToDirectionLeft(t *testing.T) {
	d := keyToDirection(termbox.KeyArrowLeft)

	if d != LEFT {
		t.Fatalf("Expected direction to be LEFT but got %v", d)
	}
}

func TestKeyToDirectionUp(t *testing.T) {
	d := keyToDirection(termbox.KeyArrowUp)

	if d != UP {
		t.Fatalf("Expected direction to be UP but got %v", d)
	}
}
