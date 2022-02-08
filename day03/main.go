package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var lines []string

type Rates struct {
	GammaBin   string
	EpsilonBin string
	Gamma      int64
	Epsilon    int64
	Solution   int64
}

func main() {
	readInput()
	rates := Rates{}
	rates = Solve(rates)
	fmt.Println(rates)
	fmt.Printf("Solution is %v\n", rates.Solution)
}

func Solve(rates Rates) Rates {
	rates.GammaBin = GetBit(lines, false)
	rates.EpsilonBin = GetBit(lines, true)

	var err error
	rates.Gamma, err = strconv.ParseInt(rates.GammaBin, 2, 64)
	errCheck(err)
	rates.Epsilon, err = strconv.ParseInt(rates.EpsilonBin, 2, 64)
	errCheck(err)

	rates.Solution = rates.Gamma * rates.Epsilon

	return rates
}

func GetBit(bits []string, reverse bool) string {
	for i := 0; len(bits) != 1; i++ {

		zeroc, onec := CountCol(i, bits)

		if zeroc > onec {
			if !reverse {
				bits = FilterBits(i, '0', bits)
			} else {
				bits = FilterBits(i, '1', bits)
			}
		} else {
			if !reverse {
				bits = FilterBits(i, '1', bits)
			} else {
				bits = FilterBits(i, '0', bits)
			}
		}
	}
	return bits[0]
}

// First is zeroc, second onec
func CountCol(i int, bits []string) (int, int) {
	var zeroc, onec int

	for _, bin := range bits {
		switch bin[i] {
		case '0':
			zeroc += 1
		case '1':
			onec += 1
		}
	}
	return zeroc, onec
}

func FilterBits(i int, n rune, oldlines []string) []string {
	var repline []string
	for k, bin := range oldlines {
		if rune(bin[i]) == n {
			repline = append(repline, oldlines[k])
		}
	}
	return repline
}

func readInput() {
	file, err := os.Open("testinput")
	errCheck(err)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
