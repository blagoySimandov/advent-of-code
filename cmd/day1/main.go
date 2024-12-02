package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		linearr := strings.Split(line, "   ")
		list1int, err := strconv.Atoi(linearr[0])
		if err != nil {
			log.Fatalf("error converting str %v", err)
		}

		list2int, err := strconv.Atoi(linearr[1])
		if err != nil {
			log.Fatalf("error converting str %v", err)
		}
		list1 = append(list1, list1int)
		list2 = append(list2, list2int)

	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	spew.Dump(getDistance(list1, list2))
	spew.Dump(calcSimilarityUnoptimized(list1, list2))
}

func getDistance(list1, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)
	distance := 0
	for i := 0; i < len(list1); i++ {
		distance += Abs(list1[i] - list2[i])
	}
	return distance
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// case0: check cache : if empty:
//
// case1: equal "=" -> counter++ go to next
// case2: less "=" -> go to next
// case3: bigger "=" or end of list ->  save counter/update cache if needed
// 1 1
// 1 1
// 1 1
// 2 1
// 2 2
func calcSimilarity(list1, list2 []int) int {
	cache := map[int]int{}
	similarity := 0
	for i := 0; i < len(list1); i++ {
		c := list1[i]
		if n, exists := cache[c]; exists {
			similarity += n
			continue
		}
		counter := 0
		for ii := 0; ii < len(list1); ii++ {
			c2 := list2[ii]
			if c2 == c {
				counter++
				continue
			}

			if c2 < c {
				continue
			}

			if c2 > c {
				similarity += counter
				cache[c] = counter
				break
			}
			// 3 3
		}
	}
	return similarity
}

func calcSimilarityUnoptimized(list1, list2 []int) int {
	similarity := 0
	for _, n1 := range list1 {
		for _, n2 := range list2 {
			if n1 == n2 {
				similarity += n1
			}
		}
	}
	return similarity
}
