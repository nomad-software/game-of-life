package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell"
)

var (
	substrate     [][]rune
	neighbourhood [8][2]int = [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
)

const (
	alive = '#'
	dead  = ' '
)

func main() {
	dish, err := tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	err = dish.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	dish.SetStyle(tcell.StyleDefault.Foreground(tcell.ColorTeal))

	width, height := dish.Size()

	initSubstrate(width, height)
	seed(width, height)
	replace(dish, width, height)

	life := make(chan bool)
	go handleKeys(dish, life)

cycle:
	for {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				incubate(dish, x, y)
			}
		}
		replace(dish, width, height)

		select {
		case <-life:
			break cycle
		case <-time.After(time.Millisecond * 100):
		}
	}
	dish.Fini()
}

func handleKeys(dish tcell.Screen, life chan bool) {
	for {
		ev := dish.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape, tcell.KeyEnter:
				close(life)
				return
			case tcell.KeyCtrlL:
				dish.Sync()
			}
		case *tcell.EventResize:
			dish.Sync()
		}
	}
}

func initSubstrate(width int, height int) {
	substrate = make([][]rune, width)
	for i := range substrate {
		substrate[i] = make([]rune, height)
	}
}

func seed(width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rand.Seed(time.Now().UnixNano())
			if rand.Intn(4) == 0 {
				substrate[x][y] = alive
			}
		}
	}
}

func incubate(dish tcell.Screen, x int, y int) {
	cell, _, _, _ := dish.GetContent(x, y)
	neighbours := 0

	for _, pos := range neighbourhood {
		neighbour, _, _, _ := dish.GetContent(x+pos[0], y+pos[1])
		if neighbour == alive {
			neighbours++
		}
	}

	if cell == alive && (neighbours == 2 || neighbours == 3) {
		substrate[x][y] = alive
	} else if cell == dead && neighbours == 3 {
		substrate[x][y] = alive
	} else {
		substrate[x][y] = dead
	}
}

func replace(dish tcell.Screen, width int, height int) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			dish.SetContent(x, y, substrate[x][y], nil, tcell.StyleDefault)
		}
	}
	dish.Show()
}
