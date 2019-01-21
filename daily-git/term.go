package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

func termInit() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
	s, e := tcell.NewScreen()
	if e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}
	if e = s.Init(); e != nil {
		fmt.Fprintf(os.Stderr, "%v\n", e)
		os.Exit(1)
	}

	s.SetStyle(tcell.StyleDefault.
		Foreground(tcell.ColorBlack).
		Background(tcell.ColorWhite))
	s.Clear()

	quit := make(chan struct{})
	go func() {
		for {
			ev := s.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(quit)
					return
				case tcell.KeyCtrlL:
					s.Sync()
				}
			case *tcell.EventResize:
				s.Sync()
			}
		}
	}()

	dur := time.Duration(0)
loop:
	for {
		makeCalendar(s)
		select {
		case <-quit:
			break loop
		case <-time.After(time.Second * 60):
		}
		start := time.Now()
		dur += time.Now().Sub(start)
	}

	s.Fini()
}

func makeCalendar(s tcell.Screen) {
	w, h := s.Size()
	weekCnt := len(weeks)
	st := tcell.StyleDefault

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if i == w-weekCnt-PAD || j == h-7-PAD || i == w-PAD || j == h-PAD {
				s.SetCell(i, j, st, '*')
			} else if i >= PAD && i < weekCnt+PAD && j >= PAD && j < 7+PAD {
				color := tcell.ColorWhiteSmoke
				if j-PAD < len(weeks[i-PAD].ContributionDays) {
					color = tcell.GetColor(string(weeks[i-PAD].ContributionDays[j-PAD].Color))
				}
				s.SetCell(i, j, st.Background(tcell.ColorWhiteSmoke).Foreground(color), '■')
			}
		}
	}
	s.Show()

}
