package term

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

// Term is the main terminal.
type Term struct {
	tcell tcell.Screen
}

// NewTerm contructs a new terminal.
func NewTerm() Term {
	tc, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	err = tc.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	tc.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorTeal))

	s := Term{
		tcell: tc,
	}

	return s
}

// Size returns the size of the terminal.
func (s *Term) Size() (int, int) {
	return s.tcell.Size()
}

// Display displays the passed runes onto the terminal.
func (s *Term) Display(buffer [][]rune) {
	width := len(buffer)
	height := len(buffer[0])

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			s.tcell.SetContent(x, y, buffer[x][y], nil, tcell.StyleDefault)
		}
	}
	s.tcell.Show()
}

// Destroy closes the terminal display and shows the original display.
func (s *Term) Destroy() {
	s.tcell.Fini()
}

// HandleInput waits for key presses and takes the appropriate action.
func (s *Term) HandleInput(signal chan bool) {
	go func() {
		for {
			ev := s.tcell.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(signal)
					return
				case tcell.KeyCtrlL:
					s.tcell.Sync()
				}
			case *tcell.EventResize:
				s.tcell.Sync()
			}
		}
	}()
}
