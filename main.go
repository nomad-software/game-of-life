package main

import (
	"time"

	"github.com/nomad-software/game-of-life/term"
	"github.com/nomad-software/screensaver/screen/saver/game_of_life/colony"
)

func main() {
	term := term.New()
	colony := colony.New(term.Width, term.Height)

	destroy := make(chan bool)
	term.HandleInput(destroy)

lifecycle:
	for {
		colony.Incubate()
		term.Draw(colony.View())
		term.Update()

		select {
		case <-destroy:
			break lifecycle
		case <-time.After(time.Second / 20):
		}
	}
	term.Destroy()
}
