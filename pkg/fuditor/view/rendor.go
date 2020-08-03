package view

import (
	"github.com/nsf/termbox-go"
)

// DrawView 画面を表示する
// cursor位置の設定、bufferの変更と画面への反映
func DrawView(_key termbox.Event) {
	switch _key.Key {
	case termbox.KeyArrowUp, termbox.KeyArrowDown, termbox.KeyArrowRight, termbox.KeyArrowLeft:
		move_cursor(_key)
	case termbox.KeyDelete, termbox.KeyBackspace:
		delete(_key)
	default:
		output(_key)
	}
}
