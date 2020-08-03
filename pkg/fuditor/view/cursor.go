package view

import (
	"github.com/nsf/termbox-go"
)

type cursor struct {
	tab  int // タブ番号
	x, y int // カーソル位置
}

const (
	min_x = 1
	min_y = 1
)

var c *cursor

func init() {
	c = &cursor{x: min_x, y: min_y}
}

// setCursor x,yで指定したウィンドウ位置にカーソルを設定する
func setCursor(x, y int) {
	c.x = x
	c.y = y
	termbox.SetCursor(c.x, c.y)
}

func move_cursor(_key termbox.Event) {
	switch _key.Key {
	case termbox.KeyArrowUp:
		cursor_move_up()
	case termbox.KeyArrowDown:
		cursor_move_down()
	case termbox.KeyArrowRight:
		cursor_move_right()
	case termbox.KeyArrowLeft:
		cursor_move_left()
	}
	setCursor(c.x, c.y)
}

func cursor_move_up() {
	if c.y > min_y {
		c.y--
	}
}

func cursor_move_down() {
	_, max_y := termbox.Size()
	if c.y < max_y {
		c.y++
	}
}

func cursor_move_right() {
	max_x, _ := termbox.Size()
	if c.x < max_x {
		c.x++
	}
}

func cursor_move_left() {
	if c.x > min_x {
		c.x--
	}
}

func delete(_key termbox.Event) {
	if c.x > min_x {
		c.x--
		termbox.SetCell(c.x, c.y, ' ', termbox.ColorWhite, termbox.ColorBlack)
	}
	setCursor(c.x, c.y)
}

func output(_key termbox.Event) {
	if _key.Key == termbox.KeyEnter {
		setCursor(min_x, c.y+1)
	} else {
		termbox.SetCell(c.x, c.y, _key.Ch, termbox.ColorWhite, termbox.ColorBlack)
		setCursor(c.x+1, c.y)
	}
}
