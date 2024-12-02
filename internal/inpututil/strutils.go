package inpututil

import (
	"fmt"
	"strconv"
)

func StringSliceToInt(arr []string) ([]int, error) {
	intArray := []int{}
	for _, v := range arr {
		i, err := strconv.Atoi(v)
		if err != nil {
			return intArray, fmt.Errorf("err %w", err)
		}
		intArray = append(intArray, i)
	}
	return intArray, nil
}
