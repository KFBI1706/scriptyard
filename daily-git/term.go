package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gdamore/tcell"
	colors "gopkg.in/go-playground/colors.v1"
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
	w, h := s.Size()
loop:
	for {
		if tmpW, tmpH := s.Size(); tmpW != w || tmpH != h {
			w, h = tmpW, tmpH
			s.Clear()
		}
		makeBorder(s)
		makeAvatar(s)
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

func makeAvatar(s tcell.Screen) {
	var j int
	img, err := getImage(AvatarURL)
	if err != nil {
		log.Fatal(err)
	}
	w, h := 24, 24/2
	st := tcell.StyleDefault.Background(tcell.ColorGreen)

	max := img.Bounds().Max

	xs, ys := max.X/w, max.Y/h

	for i := PAD; i < max.X/xs; i++ {
		for j = PAD; j < max.Y/ys; j++ {
			r, g, b, _ := img.At(i*xs, j*ys).RGBA()
			rgb, err := colors.RGB(uint8(r/0x101), uint8(g/0x101), uint8(b/0x101))
			if err != nil {
				log.Fatal(err)
			}
			s.SetCell(i, j, st.Background(tcell.GetColor(rgb.ToHEX().String())), ' ')
		}
	}

	j, i := j+1, PAD

	for inc, r := range Name {
		s.SetCell(i+inc, j, tcell.StyleDefault.Bold(true).Background(tcell.ColorWhite).Foreground(tcell.ColorGray), r)
	}

	s.Show()
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
				ii, jj := i, j-(h-7-PAD)-1 //Relative coordinates, when we know we're in the correct spot we act like it's starting from 0,0
				if ii >= PAD && ii < weekCnt+PAD && jj >= 0 && jj <= 7 {
					color := tcell.ColorWhiteSmoke
					if jj < len(weeks[ii-PAD].ContributionDays) {
						color = tcell.GetColor(string(weeks[ii-PAD].ContributionDays[jj].Color))
					}
					s.SetCell(i, j, st.Background(tcell.ColorWhiteSmoke).Foreground(color), '■')
				}
			}
		}
	}
	s.Show()

}
