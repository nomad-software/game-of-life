package main

import (
	"time"

	"github.com/nomad-software/game-of-life/game"
	"github.com/nomad-software/game-of-life/term"
)

func main() {
	term := term.NewTerm()
	width, height := term.Size()

	game := game.NewGame(width, height)

	signal := make(chan bool)
	term.HandleInput(signal)

lifecycle:
	for {
		game.Incubate()
		term.Display(game.View())

		select {
		case <-signal:
			break lifecycle
		case <-time.After(time.Millisecond * 100):
		}
	}
	term.Destroy()
}
