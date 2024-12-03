package main

import (
	"advent-of-code/internal/inpututil"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	fullDoc := ""
	for v := range inpututil.FileLines("input.txt") {
		fullDoc += v
	}

	pattern := `mul\((\d+),(\d+)\)`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(fullDoc, -1)
	sum := 0
	for _, match := range matches {
		if len(match) == 3 {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			sum += num1 * num2
			fmt.Printf("First number: %d, Second number: %d\n", num1, num2)
		}
	}

	fmt.Printf("%d", sum)
}
