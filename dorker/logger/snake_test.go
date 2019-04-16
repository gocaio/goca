package logger

import "testing"

func newDoubleSnake(d direction) *snake {
	return newSnake(d, []coord{
		coord{x: 1, y: 0},
		coord{x: 1, y: 1},
		coord{x: 1, y: 2},
		coord{x: 1, y: 3},
		coord{x: 1, y: 4},
	})
}

func TestSnakeBodyMove(t *testing.T) {
	snake := newDoubleSnake(RIGHT)
	snake.move()

	if snake.body[0].x != 1 || snake.body[0].y != 1 {
		t.Fatalf("Invalid body position %x", snake.body[0])
	}

	if snake.body[1].x != 1 || snake.body[1].y != 2 {
		t.Fatalf("Invalid body position %x", snake.body[1])
	}

	if snake.body[2].x != 1 || snake.body[2].y != 3 {
		t.Fatalf("Invalid body position %x", snake.body[2])
	}

	if snake.body[3].x != 1 || snake.body[3].y != 4 {
		t.Fatalf("Invalid body position %x", snake.body[3])
	}

	if snake.body[4].x != 2 || snake.body[4].y != 4 {
		t.Fatalf("Invalid body position %x", snake.body[4])
	}
}

func TestSnakeHeadMoveRight(t *testing.T) {
	snake := newDoubleSnake(RIGHT)
	snake.move()

	if snake.head().x != 2 || snake.head().y != 4 {
		t.Fatalf("Expected head to have moved to position [2 4], got %x", snake.head())
	}
}

func TestSnakeHeadMoveUp(t *testing.T) {
	snake := newDoubleSnake(UP)
	snake.move()

	if snake.head().x != 1 || snake.head().y != 5 {
		t.Fatalf("Expected head to have moved to position [1 5], got %x", snake.head())
	}
}

func TestSnakeHeadMoveDown(t *testing.T) {
	snake := newDoubleSnake(RIGHT)
	snake.move()

	snake.changeDirection(DOWN)
	snake.move()

	if snake.head().x != 2 || snake.head().y != 3 {
		t.Fatalf("Expected head to have moved to position [2 3], got %x", snake.head())
	}
}

func TestSnakeHeadMoveLeft(t *testing.T) {
	snake := newDoubleSnake(LEFT)
	snake.move()

	if snake.head().x != 0 || snake.head().y != 4 {
		t.Fatalf("Expected head to have moved to position [0 4], got %x", snake.head())
	}
}

func TestChangeDirectionToNotOposity(t *testing.T) {
	snake := newDoubleSnake(DOWN)
	snake.changeDirection(RIGHT)
	if snake.direction != RIGHT {
		t.Fatal("Expected to change Snake Direction to DOWN")
	}
}

func TestChangeDirectionToOposity(t *testing.T) {
	snake := newDoubleSnake(RIGHT)
	snake.changeDirection(LEFT)
	if snake.direction == LEFT {
		t.Fatal("Expected not to have changed Snake Direction to LEFT")
	}
}

func TestChangeDirectionToInvalidDirection(t *testing.T) {
	snake := newDoubleSnake(RIGHT)
	snake.changeDirection(5)
	if snake.direction != RIGHT {
		t.Fatal("Expected not to have changed Snake Direction")
	}
}

func TestSnakeDie(t *testing.T) {
	snake := newDoubleSnake(RIGHT)

	if err := snake.die(); err.Error() != "Died" {
		t.Fatal("Expected Snake die() to return error")
	}
}

func TestSnakeDieWhenMoveOnTopOfItself(t *testing.T) {
	snake := newDoubleSnake(RIGHT)
	snake.move()

	snake.changeDirection(DOWN)
	snake.move()

	snake.changeDirection(LEFT)

	if err := snake.die(); err.Error() != "Died" {
		t.Fatal("Expected Snake to die when moved on top of itself")
	}
}
