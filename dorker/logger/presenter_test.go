package logger

import "testing"

func TestPresenterRendersSuccessfully(t *testing.T) {
	g := NewGame()

	if err := g.render(); err != nil {
		t.Fatal("Expected Game to have been rendered successfully")
	}
}
