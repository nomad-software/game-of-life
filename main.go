package main

import (
	"time"

	"github.com/nomad-software/game-of-life/game"
	"github.com/nomad-software/game-of-life/term"
)

func main() {
	term := term.NewTerm()
	game := game.NewGame(term.Width, term.Height)

	signal := make(chan bool)
	term.HandleInput(signal)

lifecycle:
	for {
		game.Incubate()
		term.Display(game.View())

		select {
		case <-signal:
			break lifecycle
		case <-time.After(time.Second / 20):
		}
	}
	term.Destroy()
}
