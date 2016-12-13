package arrays

import (
	"bufio"
	"fmt"
	"os"
)

func sparseArrayMain() {
	reader := bufio.NewReader(os.Stdin)
	var n, q int
	var data map[string]int
	data = make(map[string]int, 0)
	var temp string

	fmt.Fscanf(reader, "%d\n", &n)

	for index := 0; index < n; index++ {
		fmt.Fscanf(reader, "%s\n", &temp)
		value, exists := data[temp]
		if !exists {
			data[temp] = 1
		} else {
			data[temp] = value + 1
		}
	}

	fmt.Fscanf(reader, "%d\n", &q)

	for index := 0; index < q; index++ {
		fmt.Fscanf(reader, "%s\n", &temp)
		value, exists := data[temp]
		if !exists {
			fmt.Println(0)
		} else {
			fmt.Println(value)
		}
	}

}
