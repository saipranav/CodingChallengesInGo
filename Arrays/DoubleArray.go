package arrays

import (
	"bufio"
	"fmt"
	"os"
)

func doubleArrayMain() {
	var n, q, lastAns, selectedIndex, seq int
	var queryType, queryX, queryY int
	var seqList [][]int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &n, &q)
	seqList = make([][]int, n)
	for index := 0; index < n; index++ {
		seqList[index] = make([]int, 0)
	}
	for indexQ := 0; indexQ < q; indexQ++ {
		fmt.Fscanf(reader, "%d %d %d\n", &queryType, &queryX, &queryY)
		if queryType == 1 {
			seq = ((queryX ^ lastAns) % n)
			seqList[seq] = append(seqList[seq], queryY)
		} else {
			seq = ((queryX ^ lastAns) % n)
			selectedIndex = ((queryY) % len(seqList[seq]))
			lastAns = seqList[seq][selectedIndex]
			fmt.Println(lastAns)
		}
	}
}
