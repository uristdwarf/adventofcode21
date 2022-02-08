// Just use regex lol
package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	filename := "../testinput"
	if len(os.Args) == 2 {
		filename = "../" + os.Args[1]
	}
	var re = regexp.MustCompile(`(?m)\b(?:\w{2}|\w{4}|\w{3}|\w{7})\b`)
	b, err := os.ReadFile(filename)
	errCheck(err)
	in := string(b)
	inslc := strings.Split(in, "\n")
	// 0 is in, 1 is out

	var sum = 0
	for _, s := range inslc {
		inout := strings.Split(s, "|")
		for range re.FindAllString(inout[1], -1) {
			sum++
		}
	}

	fmt.Println(sum)
}

func errCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
