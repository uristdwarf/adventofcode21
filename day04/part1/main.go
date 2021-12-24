package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	rows [][]BoardNumber
}

type BoardNumber struct {
	number string
	marked bool
}

func main() {
	filename := "testinput"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	f, err := os.Open(filename)
	errCheck(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	numbRaw := scanner.Text()
	numbers := strings.Split(numbRaw, ",")

	re := regexp.MustCompile(`\d+`)
	boards := []Board{}
	var boardNumbers [][]BoardNumber
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		for i := 0; i < 5; i++ {
			rowNums := re.FindAllString(scanner.Text(), -1)
			boardNumbers = append(boardNumbers, MakeBoardNumbers(rowNums))
			scanner.Scan()
		}
		boards = append(boards, Board{rows: boardNumbers})
		boardNumbers = nil
	}

	fmt.Println(numbers)

out:
	for _, num := range numbers {
		for boardIndex, board := range boards {
			board.MarkNumbers(num)
			if board.CheckRow() || board.CheckCol() {
				fmt.Printf("Board No. %v is the winner!\n", boardIndex+1)
				unmarked := board.GetUnmarked()
				sum := 0
				for _, n := range unmarked {
					sum += n
				}
				fmt.Printf("Unmarked sum: %v\n", sum)
				intnum, _ := strconv.Atoi(num)
				fmt.Printf("Winning number: %v\n", intnum)
				fmt.Printf("Answer: %v * %v = %v\n", sum, intnum, sum*intnum)
				break out
			}
		}
	}
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MakeBoardNumbers(nums []string) []BoardNumber {
	var boardNumbers []BoardNumber
	for _, num := range nums {
		boardNum := BoardNumber{number: num, marked: false}
		boardNumbers = append(boardNumbers, boardNum)
	}
	return boardNumbers
}

func (b *Board) MarkNumbers(num string) {
	for colIndex, row := range b.rows {
		for rowIndex, boardNum := range row {
			if boardNum.number == num {
				b.rows[colIndex][rowIndex].marked = true
			}
		}
	}
}

func (b *Board) CheckRow() bool {
	count := 0
	for _, row := range b.rows {
		for _, boardNum := range row {
			if boardNum.marked {
				count++
			} else {
				count = 0
				break
			}
		}
		if count == 5 {
			return true
		}
	}
	return false
}

func (b *Board) CheckCol() bool {
	for _, row := range b.rows {
		for rowIndex := range row {
			if b.recurseCol(0, rowIndex) {
				return true
			}
		}
	}
	return false
}

func (b *Board) recurseCol(colIndex, rowIndex int) bool {
	if colIndex == 4 {
		return b.rows[colIndex][rowIndex].marked
	}
	if b.rows[colIndex][rowIndex].marked {
		return b.recurseCol(colIndex+1, rowIndex)
	} else {
		return false
	}
}

func (b *Board) GetUnmarked() []int {
	var unNums []int
	for _, row := range b.rows {
		for _, boardNum := range row {
			if !boardNum.marked {
				intNum, err := strconv.Atoi(boardNum.number)
				errCheck(err)
				unNums = append(unNums, intNum)
			}
		}
	}
	return unNums
}
