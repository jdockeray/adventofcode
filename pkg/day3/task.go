package main

import (
	"bufio"
	"fmt"
	"github.com/jdockeray/adventofcode/pkg/day2"
)

type Slope struct {
	right int
	down  int
}

type Position struct {
	x     int
	y     int
	trees int
}

type Grid struct {
	rows   [][]int
	width  int
	height int
}

type Sled struct {
	slope    Slope
	position Position
}

type Simulation struct {
	grid  Grid
	sleds []Sled
}

func (receiver Grid) addRow(txt string) Grid {
	if receiver.width == 0 {
		receiver.width = len(txt)
	}
	receiver.height = receiver.height + 1
	var trees []int
	for index, t := range txt {
		if string(t) == "#" {
			trees = append(trees, index)
		}
	}
	receiver.rows = append(receiver.rows, trees)
	return receiver
}

func (receiver Simulation) run() {
	finished := []Sled{}
	for _, sled := range receiver.sleds {
		for sled.position.y < receiver.grid.height {
			sled = receiver.grid.move(sled)
		}
		finished = append(finished, sled)

	}
	receiver.sleds = finished

	fmt.Println(receiver.total())
}

func (receiver Grid) move(sled Sled) Sled {
	sled.position.x = sled.position.x + sled.slope.right
	sled.position.y = sled.position.y + sled.slope.down

	if receiver.isEnd(sled.position.y) {
		return sled
	}

	x := sled.position.x % receiver.width

	for _, tree := range receiver.rows[sled.position.y] {
		if tree == x {
			sled.position.trees = sled.position.trees + 1
		}
	}
	return sled
}

func (receiver Grid) isEnd(y int) bool {
	return y >= receiver.height
}

func (receiver Simulation) total() int {
	var total int = 1
	for _, sled := range receiver.sleds {
		total = total * sled.position.trees
	}
	return total
}

func main() {
	res, _ := day2.FetchData("http://127.0.0.1:8080/data/day3.txt")
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	sleds := []Sled{
		{
			slope: Slope{
				right: 1,
				down:  1,
			},
			position: Position{},
		},
		{
			slope: Slope{
				right: 3,
				down:  1,
			},
			position: Position{},
		},
		{
			slope: Slope{
				right: 5,
				down:  1,
			},
			position: Position{},
		},
		{
			slope: Slope{
				right: 7,
				down:  1,
			},
			position: Position{},
		},
		{
			slope: Slope{
				right: 1,
				down:  2,
			},
			position: Position{},
		},
	}

	var g = Grid{
		rows:   make([][]int, 0, 0),
		width:  0,
		height: 0,
	} // fill grid
	for scanner.Scan() {
		str := scanner.Text()
		g = g.addRow(str)
	}
	sim := Simulation{
		grid:  g,
		sleds: sleds,
	}

	sim.run()

}
