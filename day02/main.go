package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	posit int
	depth int
	aim   int
)

func main() {
	readInput()
	fmt.Printf("posit: %v\n", posit)
	fmt.Printf("depth: %v\n", depth)
	fmt.Printf("multi: %v\n", posit*depth)
}

func readInput() map[string]string {
	m := make(map[string]string)
	file, err := os.Open("input")
	errCheck(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sep := strings.Split(scanner.Text(), " ")
		val, err := strconv.Atoi(sep[1])
		errCheck(err)
		decideCase(sep[0], val)
	}

	return m
}

func decideCase(com string, val int) {
	switch com {
	case "forward":
		posit += val
		depth += val * aim
	case "down":
		aim += val
	case "up":
		aim -= val
	}
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
