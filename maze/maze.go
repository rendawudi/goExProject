package main

import (
	"fmt"
	"os"
)

func readMaze(fileName string) [][]int {
	var row, col int
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)

	for i := range maze{
		maze[i] = make([]int, col)
		for j := range maze[i]{
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	i, j int
}

func (p point) add(q point) point{
	return point{p.i + q.i, p.j + q.j}
}

func (p point) canAt(maze [][]int) bool  {
	if p.i < 0 || p.i >= len(maze) {
		return false
	}
	if p.j < 0 || p.j >= len(maze[p.i]) {
		return false
	}
	if maze[p.i][p.j] == 1 {
		return false
	}
	return true
}

func (p point) existAt(steps [][]int) bool {
	if steps[p.i][p.j] != 0 {
		return true
	}
	return false
}

var dirs = [4]point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0 , 1},
}

func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range maze{
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	steps[Q[0].i][Q[0].j] = 1

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			return steps
		}

		for i := range dirs{
			next := cur.add(dirs[i])

			if !next.canAt(maze) {
				continue
			}

			if next.existAt(steps) {
				continue
			}

			steps[next.i][next.j] = steps[cur.i][cur.j] + 1
			Q = append(Q, next)
		}
	}

	return steps

}


func main() {
	maze := readMaze("maze/maze.in")

	//for i := range maze{
	//	for j := range maze[i]{
	//		fmt.Printf("%3d", maze[i][j])
	//	}
	//	fmt.Println()
	//}

	steps :=walk(maze, point{0, 0},
	point{len(maze) - 1, len(maze[0]) - 1})

	for i := range steps{
		for j := range steps[i]{
			fmt.Printf("%3d", steps[i][j])
		}
		fmt.Println()
	}
}
