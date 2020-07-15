package main

import (
	"github.com/gbin/goncurses"
	"log"
	"time"
)

type Point struct {
	x int
	y int
}

type Life struct {
	cells map[Point]int
}

func (l Life) get_cell(x int, y int) int {
	return l.cells[Point{x, y}]
}

func (l Life) check_cell(x int, y int) ([]Point, []Point) {
	x_coords := [3]int{x - 1, x, x + 1}
	y_coords := [3]int{y - 1, y, y + 1}
	total := 0

	for _, x_coord := range x_coords {
		for _, y_coord := range y_coords {
			total += l.get_cell(x_coord, y_coord)
		}
	}

	live, dead := make([]Point, 0), make([]Point, 0)
	cell := l.get_cell(x, y)
	if total == 3 {
		live = append(live, Point{x, y})
	} else if total < 3 || (total > 4 && cell != 0) {
		dead = append(dead, Point{x, y})
	}

	return live, dead
}

func (l Life) queue_cells() []Point {
	cells := make([]Point, 0)

	for i, _ := range l.cells {
		x, y := i.x, i.y
		for _, x_coord := range [3]int{x - 1, x, x + 1} {
			for _, y_coord := range [3]int{y - 1, y, y + 1} {
				cells = append(cells, Point{x_coord, y_coord})
			}
		}
	}

	return cells
}

func (l Life) play_game() {
	live, dead := make([]Point, 0), make([]Point, 0)
	for _, e := range l.queue_cells() {
		step_live, step_dead := l.check_cell(e.x, e.y)
		live = append(live, step_live...)
		dead = append(dead, step_dead...)
	}

	for _, e := range dead {
		if l.get_cell(e.x, e.y) != 0 {
			delete(l.cells, Point{e.x, e.y})
		}
	}

	for _, e := range live {
		l.cells[Point{e.x, e.y}] = 1
	}

}

func (game Life) run(screen *goncurses.Window) {

	screen.Timeout(0)
	adjust_x, adjust_y := 0, 0

	for {
		switch move := screen.GetChar(); move {
		case 'h':
			adjust_x -= 1
		case 'l':
			adjust_x += 1
		case 'k':
			adjust_y -= 1
		case 'j':
			adjust_y += 1
		case 'q':
			return
		}

		screen.Clear()
		game.play_game()
		max_y, max_x := screen.MaxYX()
		for i, _ := range game.cells {
			x, y := i.x, i.y
			visible_x := adjust_x < x && x < (max_x+adjust_x)
			visible_y := adjust_y < y && y < (max_y+adjust_y)
			if visible_x && visible_y {
				screen.MovePrint(y-adjust_y, x-adjust_x, "X")
			}
		}

		goncurses.Cursor(0)
		screen.Refresh()
		time.Sleep(100 * time.Millisecond)

	}
}

func main() {
	// Create starting positions for game
	points := make(map[Point]int)
	points[Point{25, 25}] = 1
	points[Point{26, 25}] = 1
	points[Point{25, 26}] = 1
	points[Point{24, 26}] = 1
	points[Point{25, 27}] = 1

	// Create game
	game := Life{points}

	// Init the ncurses interface
	screen, err := goncurses.Init()
	if err != nil {
		log.Fatal("init:", err)
	}

	// Run the game loop
	game.run(screen)

	// Destroy ncurses interface
	defer goncurses.End()

}
