package arrays

import (
	"bufio"
	"fmt"
	"os"
)

func leftRotationMain() {
	reader := bufio.NewReader(os.Stdin)
	var n, r, temp int
	var array, newArray []int
	fmt.Fscanf(reader, "%d %d\n", &n, &r)
	for index := 0; index < n; index++ {
		fmt.Fscanf(reader, "%d ", &temp)
		array = append(array, temp)
	}

	newArray = array[r:]
	newArray = append(newArray, array[0:r]...)

	fmt.Println(newArray)
}
