package fuditor

import (
	"github.com/fujimaru-lab/f_logger/pkg/fujilog"
	"github.com/fujimaru-lab/fuditor/pkg/fuditor/util"
	"github.com/fujimaru-lab/fuditor/pkg/fuditor/view"
	"github.com/nsf/termbox-go"
)

// RunFuditor fuditorの本体を起動する
func RunFuditor() {
	err := termbox.Init()
	fujilog.LogInfo("start fuditor")
	if err != nil {
		fujilog.LogFatal(err.Error())
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	keepGOing() // ユーザーの入力を待ち受ける
	fujilog.LogInfo("end fuditor")
	termbox.Flush()
}

func keepGOing() {
acpt_loop:
	for {
		_key := termbox.PollEvent()
		fujilog.LogDebug("key:", _key.Key, "ch:", string(_key.Ch))
		util.ConfirmInput(_key)

		switch _key.Type {
		case termbox.EventKey:
			if _key.Key == termbox.KeyCtrlQ { // エディタの終了
				break acpt_loop
			} else {
				view.DrawView(_key)
			}
		}
		termbox.Flush()
	}
}
