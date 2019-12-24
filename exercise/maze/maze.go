package main

import (
	"fmt"
	"os"
)

func main() {
	maze := readMaze("exercise/maze/maze")

	//steps := walk(maze, point{0 ,0}, point{len(maze) - 1, len(maze[0]) - 1})
	steps := walk(maze, point{0 ,0}, point{2, 2})

	fmt.Println("maze:")
	print(maze)
	fmt.Println("steps:")
	print(steps)
}

func print(arr [][]int) {
	for _, row := range arr {
		for _, item := range row {
			fmt.Printf("%5d", item)
		}
		fmt.Println()
	}
}

type point struct {
	i, j int
}

func (s *point) add(d point) point {
	return point{s.i + d.i, s.j + d.j}
}

var directions = [4]point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0,1},
}

func walk(maze [][]int, start point, end point) [][]int {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}

	queue := []point{start}
	arrived := false

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, direction := range directions {
			next := cur.add(direction)

			// 越界
			if next.i < 0 || next.i >= len(maze) || next.j < 0 || next.j >= len(maze[0]) {
				continue
			}

			// 碰墙
			if maze[next.i][next.j] == 1 {
				continue
			}

			// 已经经过
			if steps[next.i][next.j] != 0 {
				continue
			}

			// 回到起点
			if next == start {
				continue
			}

			queue = append(queue, next)
			steps[next.i][next.j] = steps[cur.i][cur.j] + 1

			// 到达终点
			if next == end {
				arrived = true
			}
		}

		if arrived {
			break
		}
	}

	return steps
}

func readMaze(path string) [][]int {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer file.Close()

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)

	for i := range maze {
		maze[i] = make([]int, col)

		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}
