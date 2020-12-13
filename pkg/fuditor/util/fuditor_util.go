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
	maxX, maxY := termbox.Size()
	for i := 0; i < maxX; i++ {
		termbox.SetCell(i, maxY-3, '-', termbox.ColorWhite, termbox.ColorBlack)
	}
	message := fmt.Sprint("INPUT KEY:Key=", _key.Key, ", Ch=", string(_key.Ch))
	data := []rune(message)
	for i := 0; i < maxX; i++ {
		if i < len(data) {
			termbox.SetCell(5+i, maxY-2, data[i], termbox.ColorWhite, termbox.ColorBlack)
		} else {
			termbox.SetCell(5+i, maxY-2, ' ', termbox.ColorWhite, termbox.ColorBlack)
		}
	}
	for i := 0; i < maxX; i++ {
		termbox.SetCell(i, maxY-1, '-', termbox.ColorWhite, termbox.ColorBlack)
	}
}
