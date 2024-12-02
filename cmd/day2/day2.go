package main

import (
	"advent-of-code/internal/inpututil"
	"advent-of-code/internal/numutil"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

const (
	Increasing Monotonicity = "increasing"
	Decreasing Monotonicity = "decreasing"
	UNKNOWN    Monotonicity = ""
)

type Monotonicity string

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
		if isReportSafeWithDamp(v) {
			counter++
		}
	}

	spew.Dump(counter)
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

func deleteElement(slice []int, index int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return append(newSlice[:index], newSlice[index+1:]...)
}

func isReportSafeWithDamp(arr []int) bool {
	state := UNKNOWN
	for i := 0; i < len(arr)-1; i++ {
		current := arr[i]
		next := arr[i+1]
		if current == next {
			arrWithoutNext := deleteElement(arr, i+1)
			arrWithoutCurrent := deleteElement(arr, i)
			prevBool := false
			if i != 0 {
				arrWithoutPrev := deleteElement(arr, i-1)
				prevBool = isReportSafe(arrWithoutPrev)
			}
			if isReportSafe(arrWithoutCurrent) || isReportSafe(arrWithoutNext) || prevBool {
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
			prevBool := false
			if i != 0 {
				arrWithoutPrev := deleteElement(arr, i-1)
				prevBool = isReportSafe(arrWithoutPrev)
			}
			if isReportSafe(arrWithoutCurrent) || isReportSafe(arrWithoutNext) || prevBool {
				return true
			}
			return false
		}
		absN := numutil.Abs(n)
		if absN < 1 || absN > 3 {
			arrWithoutNext := deleteElement(arr, i+1)
			arrWithoutCurrent := deleteElement(arr, i)
			prevBool := false
			if i != 0 {
				arrWithoutPrev := deleteElement(arr, i-1)
				prevBool = isReportSafe(arrWithoutPrev)
			}
			if isReportSafe(arrWithoutCurrent) || isReportSafe(arrWithoutNext) || prevBool {
				return true
			}
			return false
		}
	}
	return true
}
