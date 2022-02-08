package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := readInput()
	fmt.Println(Increases(input))
}

func Increases(numbers []int) int {
	first := true
	count := 0
	var n int
	var last int

	for i := len(numbers) - 1; i > 0; i-- {
		if first {
			numbers, last = threeSum(numbers)
			first = false
			continue
		}

		numbers, n = threeSum(numbers)

		if n > last {
			count++
		}
		last = n
	}

	return count
}

// First is the remainder, second is the three-measurement window
func threeSum(arr []int) ([]int, int) {
	if len(arr) < 3 {
		return arr, 0
	}

	var sum int
	for i := 0; i < 3; i++ {
		sum += arr[i]
	}
	arr = arr[1:]

	return arr, sum
}

func readInput() []int {
	var input []int
	file, err := os.Open("input")
	errCheck(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		errCheck(err)
		input = append(input, n)
	}
	return input
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
