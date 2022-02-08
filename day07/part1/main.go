package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := "../testinput"
	if len(os.Args) == 2 {
		filename = "../" + os.Args[1]
	}
	b, err := os.ReadFile(filename)
	errCheck(err)
	in := string(b)

	strS := strings.Split(in, ",")

	m := make(map[int]int)
	var high = 0

	for _, s := range strS {
		n, _ := strconv.Atoi(s)
		if n > high {
			high = n
		}
		m[n] += 1
	}

	var lowestFuel = 0
	for i := 0; i <= high; i++ {
		var fuel = 0
		for k, v := range m {
			if i < k {
				fuel += (k - i) * v
			} else {
				fuel += (i - k) * v
			}
		}
		if i == 0 || fuel < lowestFuel {
			lowestFuel = fuel
		}
	}

	fmt.Println(lowestFuel)
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
