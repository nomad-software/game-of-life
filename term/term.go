package term

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell"
)

// Term is the main terminal.
type Term struct {
	tcell  tcell.Screen
	Width  int
	Height int
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
	width, height := tc.Size()

	s := Term{
		tcell:  tc,
		Width:  width,
		Height: height,
	}

	return s
}

// Display displays the passed runes onto the terminal.
func (t *Term) Display(buffer [][]rune) {
	for y := 0; y < t.Height; y++ {
		for x := 0; x < t.Width; x++ {
			t.tcell.SetContent(x, y, buffer[x][y], nil, tcell.StyleDefault)
		}
	}
	t.tcell.Show()
}

// Destroy closes the terminal display and shows the original display.
func (t *Term) Destroy() {
	t.tcell.Fini()
}

// HandleInput waits for key presses and takes the appropriate action.
func (t *Term) HandleInput(signal chan bool) {
	go func() {
		for {
			ev := t.tcell.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter:
					close(signal)
					return
				case tcell.KeyCtrlL:
					t.tcell.Sync()
				}
			case *tcell.EventResize:
				t.tcell.Sync()
			}
		}
	}()
}
