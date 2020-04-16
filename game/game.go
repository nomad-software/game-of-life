package game

import (
	"math/rand"
)

var (
	neighbourhood [8][2]int = [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
)

const (
	alive = '#'
	dead  = ' '
)

// Game is the main game.
type Game struct {
	width     int
	height    int
	substrate [][]rune
	output    [][]rune
}

// NewGame contructs a new game.
func NewGame(width int, height int) Game {
	g := Game{
		width:     width,
		height:    height,
		substrate: make([][]rune, width),
		output:    make([][]rune, width),
	}

	for i := 0; i < width; i++ {
		g.substrate[i] = make([]rune, height)
		g.output[i] = make([]rune, height)
	}

	g.Seed()

	return g
}

// Incubate creates the next generation.
func (g *Game) Incubate() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			neighbours := 0

			// Check the neighbourhood for alive cells.
			for _, pos := range neighbourhood {
				x2 := x + pos[0]
				y2 := y + pos[1]

				if x2 < 0 {
					x2 = g.width - 1
				}

				if x2 >= g.width {
					x2 = 0
				}

				if y2 < 0 {
					y2 = g.height - 1
				}

				if y2 >= g.height {
					y2 = 0
				}

				if g.output[x2][y2] == alive {
					neighbours++
				}
			}

			// The rules of survival.
			if neighbours == 3 {
				g.substrate[x][y] = alive
			} else if g.output[x][y] == alive && neighbours == 2 {
				g.substrate[x][y] = alive
			} else {
				g.substrate[x][y] = dead
			}
		}
	}

	// Swap the next generation to the output buffer.
	g.output, g.substrate = g.substrate, g.output
}

// View returns the current game view.
func (g *Game) View() [][]rune {
	return g.output
}

// Seed randomises the game cells.
func (g *Game) Seed() {
	for i := 0; i < (g.width * g.height / 4); i++ {
		g.output[rand.Intn(g.width)][rand.Intn(g.height)] = alive
	}
}
