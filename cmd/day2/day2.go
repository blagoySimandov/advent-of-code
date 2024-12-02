package main

import (
	"advent-of-code/internal/inpututil"
	"advent-of-code/internal/numutil"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	inputMatrix := [][]int{}
	for l := range inpututil.FileLines("input.txt") {
		innerArray := strings.Split(l, " ")
		intArr, err := inpututil.StringSliceToInt(innerArray)
		if err != nil {
			log.Fatal("ebalo si e maikata naistina")
		}
		inputMatrix = append(inputMatrix, intArr)
	}
	counter := 0
	for _, v := range inputMatrix {
		if j := isReportSafe(v); j {
			counter++
		}
	}
	spew.Dump(counter)
	// input := []int{1, 2, 7, 8, 9}
	//
	// input2 := []int{1, 3, 2, 4, 5}
	// spew.Dump(isReportSafe(input))
	// spew.Dump(isReportSafe(input2))
}

func isReportSafe(arr []int) bool {
	state := UNKNOWN
	for i := 0; i < len(arr)-1; i++ {
		current := arr[i]
		next := arr[i+1]
		if current == next {
			return false
		}
		var cstate Monotonicity
		n := current - next
		if n < 0 {
			cstate = Deacreasing
		} else {
			cstate = Increasing
		}

		if state == UNKNOWN {
			state = cstate
		}
		if state != cstate {
			return false
		}
		absN := numutil.Abs(n)
		if absN < 1 || absN > 3 {
			return false
		}

	}
	return true
}

const (
	Increasing  Monotonicity = "increasing"
	Deacreasing Monotonicity = "deacreasing"
	UNKNOWN     Monotonicity = ""
)

type Monotonicity string
