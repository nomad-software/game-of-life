package game

import (
	"math/rand"
	"time"
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
	buffer    [][]rune
}

// NewGame contructs a new game.
func NewGame(width int, height int) Game {
	g := Game{
		width:     width,
		height:    height,
		substrate: make([][]rune, width),
		buffer:    make([][]rune, width),
	}

	for i := 0; i < width; i++ {
		g.substrate[i] = make([]rune, height)
		g.buffer[i] = make([]rune, height)
	}

	g.seed()

	return g
}

// Incubate creates the next generation.
func (g *Game) Incubate() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			cell := g.buffer[x][y]
			neighbours := 0

			for _, pos := range neighbourhood {
				x2 := x + pos[0]
				y2 := y + pos[1]

				if x2 >= 0 && x2 < g.width && y2 >= 0 && y2 < g.height {
					neighbour := g.buffer[x2][y2]
					if neighbour == alive {
						neighbours++
					}
				}
			}

			if cell == alive && (neighbours == 2 || neighbours == 3) {
				g.substrate[x][y] = alive
			} else if cell == dead && neighbours == 3 {
				g.substrate[x][y] = alive
			} else {
				g.substrate[x][y] = dead
			}
		}
	}

	g.buffer, g.substrate = g.substrate, g.buffer
}

// View returns the current game view.
func (g *Game) View() [][]rune {
	return g.buffer
}

// Seed randomises the game cells.
func (g *Game) seed() {
	for y := 0; y < g.height; y++ {
		for x := 0; x < g.width; x++ {
			rand.Seed(time.Now().UnixNano())
			if rand.Intn(4) == 0 {
				g.buffer[x][y] = alive
			}
		}
	}
}
