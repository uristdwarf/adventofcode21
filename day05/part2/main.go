package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Vent struct {
	StartX int
	StartY int
	EndX   int
	EndY   int
}

func main() {
	filename := "testinput"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	f, err := os.Open(filename)
	errCheck(err)
	scanner := bufio.NewScanner(f)

	vents := []Vent{}

	var lx = 0
	var ly = 0
	lx, ly, vents = GetInput(scanner, vents)
	f.Close()

	board := Make2DSlice(lx, ly)

	for _, vent := range vents {

		recurseY := GetLower(vent.EndY, vent.StartY)
		recurseX := GetLower(vent.EndX, vent.StartX)
		recEndX := GetLargerNumber(vent.StartX, vent.EndX)
		recEndY := GetLargerNumber(vent.StartY, vent.EndY)

		if recurseX == recEndX {
			RecurseAdd(vent.StartX, recurseY, recEndY, &board, true)
		} else if recurseY == recEndY {
			RecurseAdd(vent.StartY, recurseX, recEndX, &board, false)
		} else {
			RecurseDiag(vent.StartX, vent.StartY, vent.EndX, vent.EndY, &board)
		}
	}
	var sum = 0
	for _, y := range board {
		fmt.Println(y)
		for _, n := range y {
			if n > 1 {
				sum++
			}
		}
	}
	fmt.Printf("Final answer: %v\n", sum)
}

func GetLower(startN, endN int) int {
	if startN > endN {
		return endN
	} else {
		return startN
	}
}

func RecurseAdd(anch, iter, end int, board *[][]int, y bool) {
	if iter > end {
		return
	}
	if y {
		(*board)[iter][anch] += 1 // This looks like shit
		RecurseAdd(anch, iter+1, end, board, y)
	} else {
		(*board)[anch][iter] += 1
		RecurseAdd(anch, iter+1, end, board, y)
	}
}

func RecurseDiag(x, y, endX, endY int, board *[][]int) {
	(*board)[y][x] += 1
	if x == endX {
		return
	}
	if y > endY {
		if x > endX {
			RecurseDiag(x-1, y-1, endX, endY, board)
		} else {
			RecurseDiag(x+1, y-1, endX, endY, board)
		}
	} else {
		if x > endX {
			RecurseDiag(x-1, y+1, endX, endY, board)
		} else {
			RecurseDiag(x+1, y+1, endX, endY, board)
		}
	}
}

func GetInput(scanner *bufio.Scanner, vents []Vent) (int, int, []Vent) {
	var lx = 0
	var ly = 0
	re := regexp.MustCompile(`(\d+),(\d+)`)

	for scanner.Scan() {
		match := re.FindAllStringSubmatch(scanner.Text(), -1)
		startX, _ := strconv.Atoi(match[0][1])
		startY, _ := strconv.Atoi(match[0][2])
		endX, _ := strconv.Atoi(match[1][1])
		endY, _ := strconv.Atoi(match[1][2])

		lx = CompareLarger(lx, startX, endX)
		ly = CompareLarger(ly, startY, endY)
		vent := Vent{
			StartX: startX,
			StartY: startY,
			EndX:   endX,
			EndY:   endY,
		}
		vents = append(vents, vent)
	}
	// +1 is needed otherwise the slice will be 1 short with the data
	return lx + 1, ly + 1, vents
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Make2DSlice(x, y int) [][]int {
	s := make([][]int, y)
	for i := range s {
		s[i] = make([]int, x)
	}
	return s
}

// Return is larger of two numbers. If equal returns n1
func GetLargerNumber(n1, n2 int) int {
	if n1 >= n2 {
		return n1
	} else {
		return n2
	}
}

// Returns original number if equal
func CompareLarger(ori, n1, n2 int) int {
	new := GetLargerNumber(n1, n2)
	if new > ori {
		return new
	}
	return ori
}
