package arrays

import (
	"fmt"
	"math"
)

func pairGivenSumMain() {
	var array []int
	var sum int

	array = []int{1, 4, 45, 6, 10, 8, -8}
	sum = 16
	pairGivenSum(array, sum)
}

func pairGivenSum(array []int, sum int) {
	var min int = math.MaxInt32
	var data map[int]int = make(map[int]int)
	for index := 0; index < len(array); index++ {
		if min > array[index] {
			min = array[index]
		}
	}
	if min < 0 {
		min = min * -1
		for index := 0; index < len(array); index++ {
			array[index] = array[index] + min
			data[array[index]] = index
		}
		sum = sum + (2 * min)
	} else {
		for index := 0; index < len(array); index++ {
			data[array[index]] = index
		}
	}
	for index := 0; index < len(array); index++ {
		value, exists := data[(sum - array[index])]
		if exists {
			if index != value {
				fmt.Println(index, ",", value)
			}
		}
	}
}
