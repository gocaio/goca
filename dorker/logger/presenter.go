package logger

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorGreen
)

func (g *Game) render() error {
	termbox.Clear(defaultColor, defaultColor)

	var (
		w, h   = termbox.Size()
		midY   = h / 2
		left   = (w - g.arena.width) / 2
		right  = (w + g.arena.width) / 2
		top    = midY - (g.arena.height / 2)
		bottom = midY + (g.arena.height / 2) + 1
	)

	renderTitle(left, top)
	renderArena(g.arena, top, bottom, left)
	renderSnake(left, bottom, g.arena.snake)
	renderFood(left, bottom, g.arena.food)
	renderScore(left, bottom, g.score)
	renderQuitMessage(right, bottom)

	return termbox.Flush()
}

func renderSnake(left, bottom int, s *snake) {
	for _, b := range s.body {
		termbox.SetCell(left+b.x, bottom-b.y, ' ', snakeColor, snakeColor)
	}
}

func renderFood(left, bottom int, f *food) {
	termbox.SetCell(left+f.x, bottom-f.y, f.emoji, defaultColor, bgColor)
}

func renderArena(a *arena, top, bottom, left int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '│', defaultColor, bgColor)
		termbox.SetCell(left+a.width, i, '│', defaultColor, bgColor)
	}

	termbox.SetCell(left-1, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left-1, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+a.width, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+a.width, bottom, '┘', defaultColor, bgColor)

	fill(left, top, a.width, 1, termbox.Cell{Ch: '─'})
	fill(left, bottom, a.width, 1, termbox.Cell{Ch: '─'})
}

func renderScore(left, bottom, s int) {
	score := fmt.Sprintf("Score: %v", s)
	tbprint(left, bottom+1, defaultColor, defaultColor, score)
}

func renderQuitMessage(right, bottom int) {
	m := "Press ESC to quit"
	tbprint(right-17, bottom+1, defaultColor, defaultColor, m)
}

func renderTitle(left, top int) {
	tbprint(left, top-1, defaultColor, defaultColor, "Goca Snake")
}

func fill(x, y, w, h int, cell termbox.Cell) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			termbox.SetCell(x+lx, y+ly, cell.Ch, cell.Fg, cell.Bg)
		}
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
