package logger

import "testing"

var pointsDouble = make(chan int)

func newDoubleArenaWithFoodFinder(h, w int, f func(*arena, coord) bool) *arena {
	a := newDoubleArena(h, w)
	a.hasFood = f
	return a
}

func newDoubleArena(h, w int) *arena {
	s := newSnake(RIGHT, []coord{
		coord{x: 1, y: 0},
		coord{x: 1, y: 1},
		coord{x: 1, y: 2},
		coord{x: 1, y: 3},
		coord{x: 1, y: 4},
	})

	return newArena(s, pointsDouble, h, w)
}

func TestArenaHaveFoodPlaced(t *testing.T) {
	if a := newDoubleArena(20, 20); a.food == nil {
		t.Fatal("Arena expected to have food placed")
	}
}

func TestMoveSnakeOutOfArenaHeightLimit(t *testing.T) {
	a := newDoubleArena(4, 10)
	a.snake.changeDirection(UP)

	if err := a.moveSnake(); err == nil || err.Error() != "Died" {
		t.Fatal("Expected Snake to die when moving outside the Arena height limits")
	}
}

func TestMoveSnakeOutOfArenaWidthLimit(t *testing.T) {
	a := newDoubleArena(10, 1)
	a.snake.changeDirection(LEFT)

	if err := a.moveSnake(); err == nil || err.Error() != "Died" {
		t.Fatal("Expected Snake to die when moving outside the Arena height limits")
	}
}

func TestPlaceNewFoodWhenEatFood(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*arena, coord) bool {
		return true
	})

	f := a.food

	a.moveSnake()

	if a.food.x == f.x && a.food.y == f.y {
		t.Fatal("Expected new food to have been placed on Arena")
	}
}

func TestIncreaseSnakeLengthWhenEatFood(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*arena, coord) bool {
		return true
	})

	l := a.snake.length

	a.moveSnake()

	if a.snake.length != l+1 {
		t.Fatal("Expected Snake to have grown")
	}
}

func TestAddPointsWhenEatFood(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*arena, coord) bool {
		return true
	})

	if p, ok := <-pointsDouble; ok && p != a.food.points {
		t.Fatalf("Value %d was expected but got %d", a.food.points, p)
	}

	a.moveSnake()
}

func TestDoesNotAddPointsWhenFoodNotFound(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*arena, coord) bool {
		return false
	})

	select {
	case p, _ := <-pointsChan:
		t.Fatalf("No point was expected to be received but received %d", p)
	default:
		close(pointsChan)
	}

	a.moveSnake()
}

func TestDoesNotPlaceNewFoodWhenFoodNotFound(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*arena, coord) bool {
		return false
	})

	f := a.food

	a.moveSnake()

	if a.food.x != f.x || a.food.y != f.y {
		t.Fatal("Food in Arena expected not to have changed")
	}
}

func TestDoesNotIncreaseSnakeLengthWhenFoodNotFound(t *testing.T) {
	a := newDoubleArenaWithFoodFinder(10, 10, func(*arena, coord) bool {
		return false
	})

	l := a.snake.length

	a.moveSnake()

	if a.snake.length != l {
		t.Fatal("Expected Snake not to have grown")
	}
}

func TestHasFood(t *testing.T) {
	a := newDoubleArena(20, 20)

	if !hasFood(a, coord{x: a.food.x, y: a.food.y}) {
		t.Fatal("Food expected to be found")
	}
}

func TestHasNotFood(t *testing.T) {
	a := newDoubleArena(20, 20)

	if hasFood(a, coord{x: a.food.x - 1, y: a.food.y}) {
		t.Fatal("No food expected to be found")
	}
}
