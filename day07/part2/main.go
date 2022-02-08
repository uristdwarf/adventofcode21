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
		var allfuel = 0 // All of the crabs fuel
		for k, v := range m {
			var fuel = 0
			var base = 0
			if i < k {
				base = (k - i)
			} else {
				base = (i - k)
			}
			for j := 0; j <= base; j++ {
				fuel += j
			}
			allfuel += fuel * v
		}
		if i == 0 || allfuel < lowestFuel {
			lowestFuel = allfuel
		}
	}

	fmt.Println(lowestFuel)
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
