package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "testinput"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	b, err := os.ReadFile(filename)
	errCheck(err)
	in := string(b)

	strS := strings.Split(in, ",")
	var intS []int
	for _, s := range strS {
		n, _ := strconv.Atoi(s)
		intS = append(intS, n)
	}

	for day := 0; day < 80; day++ {
		for i, n := range intS {
			if n == 0 {
				intS[i] = 6
				intS = append(intS, 8)
				continue
			}
			intS[i] -= 1
		}
		// fmt.Println(intS)
	}
	fmt.Println(len(intS))
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
