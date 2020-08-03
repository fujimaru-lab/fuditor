package util

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

// IsFile 渡されたファイルパスが、ファイルパスを示す文字列であればtrue、そうでないならばfalse
func IsFile(_filepath string) bool {
	f, err := os.Open(_filepath)
	stat, err := f.Stat()
	defer f.Close()

	if os.IsNotExist(err) { // ファイルが存在しない
		return false
	} else if mode := stat.Mode(); mode.IsDir() {
		return false
	}
	return true
}

// ConfirmInput 開発時の入力確認用
func ConfirmInput(_key termbox.Event) {
	max_x, max_y := termbox.Size()
	for i := 0; i < max_x; i++ {
		termbox.SetCell(i, max_y-3, '-', termbox.ColorWhite, termbox.ColorBlack)
	}
	message := fmt.Sprint("INPUT KEY:Key=", _key.Key, ", Ch=", string(_key.Ch))
	data := []rune(message)
	for i := 0; i < max_x; i++ {
		if i < len(data) {
			termbox.SetCell(5+i, max_y-2, data[i], termbox.ColorWhite, termbox.ColorBlack)
		} else {
			termbox.SetCell(5+i, max_y-2, ' ', termbox.ColorWhite, termbox.ColorBlack)
		}
	}
	for i := 0; i < max_x; i++ {
		termbox.SetCell(i, max_y-1, '-', termbox.ColorWhite, termbox.ColorBlack)
	}
}
