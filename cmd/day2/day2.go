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
			log.Fatal("whoopsie")
		}
		inputMatrix = append(inputMatrix, intArr)
	}
	counter := 0
	for _, v := range inputMatrix {
		if j := isReportSafeWithDamp(v); j {
			counter++
		}
	}
	spew.Dump(counter)
	// input := []int{8, 6, 4, 4, 1}
	// input2 := []int{1, 3, 2, 4, 5}
	// spew.Dump(isReportSafeWithDamp(input))
	// spew.Dump(isReportSafeWithDamp(input2))
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
			cstate = Decreasing
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

func triggerFalse(dampIsHit *bool) bool {
	if dampIsHit == nil {
		panic("bro give me something non-nil")
	}
	if *dampIsHit {
		return false
	}
	*dampIsHit = true
	return true
}

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func mydeleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func isReportSafeWithDamp(arr []int) bool {
	state := UNKNOWN
	for i := 0; i < len(arr)-1; i++ {
		current := arr[i]
		next := arr[i+1]
		if current == next {
			arrWithoutNext := deleteElement(arr, i+1)
			arrWithoutCurrent := deleteElement(arr, i)
			if isReportSafe(arrWithoutCurrent) || isReportSafe(arrWithoutNext) {
				return true
			}
			return false
		}
		var cstate Monotonicity
		n := current - next
		if n < 0 {
			cstate = Decreasing
		} else {
			cstate = Increasing
		}

		if state == UNKNOWN {
			state = cstate
		}
		if state != cstate {
			arrWithoutNext := deleteElement(arr, i+1)
			arrWithoutCurrent := deleteElement(arr, i)
			if isReportSafe(arrWithoutCurrent) && isReportSafe(arrWithoutNext) {
				return true
			}
			return false
		}
		absN := numutil.Abs(n)
		if absN < 1 || absN > 3 {
			arrWithoutNext := mydeleteElement(arr, i+1)
			arrWithoutCurrent := mydeleteElement(arr, i)
			spew.Dump(i + 1)
			spew.Dump(i)
			spew.Dump(arr)
			spew.Dump(arrWithoutCurrent)
			spew.Dump(arrWithoutNext)
			if isReportSafe(arrWithoutCurrent) || isReportSafe(arrWithoutNext) {
				return true
			}
			return false
		}
	}

	return true
}

const (
	Increasing Monotonicity = "increasing"
	Decreasing Monotonicity = "decreasing"
	UNKNOWN    Monotonicity = ""
)

type Monotonicity string
