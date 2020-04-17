package main

import (
	"time"

	"github.com/nomad-software/game-of-life/colony"
	"github.com/nomad-software/game-of-life/term"
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
