package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
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

	const DAYS = 240
	base := GetBase(DAYS)
	var length = len(intS)
	for _, n := range intS {
		for b, c := range base {
			if n == b+1 {
				length += c
			}
		}
	}
	fmt.Println(length)
}

func GetBase(d int) map[int]int {
	base := make(map[int]int)
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Mapping %v...\n", i+1)
			var val = 0
			Recurse(d, i+1, &val)
			base[i] = val
			fmt.Printf("Map %v is done: %v\n", i+1, val)
			wg.Done()
		}(i)
	}
	wg.Wait()

	return base
}

func Recurse(d, n int, sum *int) {
	for {
		// The days for this fish are over
		if d <= -1 {
			return
		}
		// New fish is born
		if n == -1 {
			*sum += 1
			// Get how many fish are born from the new one
			Recurse(d-6, 2, sum)
			n = 6
			continue
		}
		n--
		d--
	}
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
