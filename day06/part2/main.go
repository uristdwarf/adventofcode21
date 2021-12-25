package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Original solution, see below function for a faster one
func main() {
	filename := "testinput"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	b, err := os.ReadFile(filename)
	errCheck(err)
	in := string(b)

	// Faster solution, see function comment for details
	Faster(in)
	os.Exit(0)

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

/*
 The original solution was above, however it was quite slow
 (took me ten minutes on my toaster), so I had to think of a
 faster solution. Thanks to @Olaroll for making me think about
 duplicate values.
*/
func Faster(in string) {
	strS := strings.Split(in, ",")

	m := make(map[int]int)
	for i := 0; i <= 8; i++ {
		m[i] = 0
	}
	for _, s := range strS {
		n, _ := strconv.Atoi(s)
		m[n] += 1
	}

	var result = len(strS) // Base result on input amount
	// HOLY SHIT HOW FAST THIS IS
	for day := 0; day < 256; day++ {
		v := m[0] // Save value on start
		// Shuffle the map
		for i := 0; i <= 8; i++ {
			m[i] = m[i+1]
		}
		m[6] += v
		m[8] += v
		result += v
	}
	fmt.Println(result)
}
