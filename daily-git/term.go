package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell"
	"github.com/y0ssar1an/q"
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
		makeBorder(s)
		makeCalendar(s)
		select {
		case <-quit:
			break loop
		case <-time.After(time.Second * 5):
		}
		start := time.Now()
		dur += time.Now().Sub(start)
	}

	s.Fini()
}

func makeBorder(s tcell.Screen) {
	w, h := s.Size()
	st := tcell.StyleDefault
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if i == 0 || j == 0 || i == w-1 || j == h-1 {
				s.SetCell(i, j, st, '*')
			}
		}
	}
}

func makeCalendar(s tcell.Screen) {
	w, h := s.Size()
	weekCnt := len(weeks)
	st := tcell.StyleDefault

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if j < h-PAD && j > h-7-PAD {
				if ii, jj := i, j-(h-7-PAD); ii >= PAD && ii < weekCnt+PAD && jj >= PAD && jj < 7+PAD {
					q.Q(ii, jj) //DEBUG
					color := tcell.ColorWhiteSmoke
					if jj-PAD < len(weeks[ii-PAD].ContributionDays) {
						color = tcell.GetColor(string(weeks[ii-PAD].ContributionDays[jj-PAD].Color))
					}
					s.SetCell(i, j, st.Background(tcell.ColorWhiteSmoke).Foreground(color), '■')
				}
			}
		}
	}
	s.Show()

}
