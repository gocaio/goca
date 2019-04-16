package logger

import (
	"testing"
	"time"
)

func TestDefaultGameScore(t *testing.T) {
	g := NewGame()

	if g.score != 0 {
		t.Fatalf("Initial Game Score expected to be 0 but it was %d", g.score)
	}
}

func TestGameMoveInterval(t *testing.T) {
	e := time.Duration(85) * time.Millisecond
	g := NewGame()
	g.score = 150

	if d := g.moveInterval(); d != e {
		t.Fatalf("Expected move interval to be %d but got %d", e, d)
	}
}

func TestAddPoints(t *testing.T) {
	g := NewGame()
	s := g.score
	g.addPoints(10)

	if s != 0 || g.score != 10 {
		t.Fatal("Expected 10 points to have been added to Game Score")
	}
}

func TestRetryGoBackToGameInitialState(t *testing.T) {
	g := NewGame()
	initScore := g.score
	initSnake := g.arena.snake

	g.arena.snake.changeDirection(UP)
	g.arena.moveSnake()
	g.addPoints(10)
	g.end()

	g.retry()

	if g.score != initScore {
		t.Fatal("Expected Score to have been reset")
	}

	for i, c := range g.arena.snake.body {
		if initSnake.body[i].x == c.x && initSnake.body[i].y == c.y {
			t.Fatal("Expected Snake body to have been reset")
		}
	}

	if g.arena.snake.direction == initSnake.direction {
		t.Fatal("Expected Snake direction to have been reset")
	}
}
