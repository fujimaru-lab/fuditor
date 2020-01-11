package fuditor

import (
	"github.com/fujimaru-lab/f_logger/pkg/fujilog"
	"github.com/nsf/termbox-go"
)

// RunFuditor fuditorの本体を起動する
func RunFuditor(args []string) {
	err := termbox.Init()
	fujilog.LogInfo("start fuditor")
	if err != nil {
		fujilog.LogFatal(err.Error())
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	termbox.Flush()
	keepGOing() // ユーザーの入力を待ち受ける
	fujilog.LogInfo("end fuditor")
}

func keepGOing() {
acpt_loop:
	for {
		event := termbox.PollEvent()
		fujilog.LogDebug("key:", event.Key, "ch:", string(event.Ch))

		switch event.Type {
		case termbox.EventKey:
			if event.Key == termbox.KeyCtrlQ {
				break acpt_loop
			}
		}
		termbox.Flush()
	}
}
