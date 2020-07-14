package main

import "fmt"

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

	fmt.Println("x_coords:", x_coords)

	for _, x_coord := range x_coords {
		for _, y_coord := range y_coords {
			total += l.get_cell(x_coord, y_coord)
			fmt.Println("Total: ", total)
			fmt.Println("Coords: ", x_coord, y_coord)
			fmt.Println("Cell val: ", l.get_cell(x_coord, y_coord))
		}
	}

	live, dead := make([]Point, 0), make([]Point, 0)
	cell := l.get_cell(x, y)
	if total == 3 {
		live = append(live, Point{x, y})
	} else if total < 3 || (total > 4 && cell != 0) {
		dead = append(dead, Point{x, y})
	}

	fmt.Printf("Found %d live cells\n", len(live))
	fmt.Printf("Found %d dead cells\n", len(dead))

	return live, dead
}

func (l Life) queue_cells() []Point {
	cells := make([]Point, 0)

	for i, _ := range l.cells {
		x, y := i.x, i.y
		fmt.Printf("Queuing cell %d %d\n", x, y)
		for _, x_coord := range [3]int{x - 1, x, x + 1} {
			for _, y_coord := range [3]int{y - 1, y, y + 1} {
				cells = append(cells, Point{x_coord, y_coord})
			}
		}
	}

	fmt.Printf("Queued %d cells\n", len(cells))

	return cells
}

func (l Life) play_game() {
	live, dead := make([]Point, 0), make([]Point, 0)
	for _, e := range l.queue_cells() {
		fmt.Printf("Checking cell %d %d\n", e.x, e.y)
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

func main() {
	points := make(map[Point]int)
	points[Point{25, 25}] = 1
	points[Point{26, 25}] = 1
	points[Point{25, 26}] = 1
	points[Point{24, 26}] = 1
	points[Point{25, 27}] = 1
	game := Life{points}
	for {
		game.play_game()
		fmt.Println(game)
		game.play_game()
		fmt.Println(game)
		game.play_game()
		fmt.Println(game)
		break
	}

}
