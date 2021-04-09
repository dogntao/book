package main

import (
	"fmt"
	"math"
	"time"
)

func solution(nums []int) int {
	// str to intArr
	// lineArr := strings.Split(line, ",")
	// intArr := make([]int, len(lineArr))
	// for k, v := range lineArr {
	// 	intArr[k], _ = strconv.Atoi(v)
	// }

	// str to int
	// lineInt, _ := strconv.Atoi(line)

	// please write your code here

	return 0
}

func search(searchArr []int, search int) (ret int) {
	ret = 0
	left := 0
	right := len(searchArr) - 1
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if search == searchArr[mid] {
			ret = mid
			break
		} else if search > searchArr[mid] {
			left = mid + 1
		} else if search < searchArr[mid] {
			right = mid - 1
		}
	}
	return ret
}

func main() {
	expireYear := time.Now().Year() + 3
	expireTimeUnix := time.Date(expireYear, 1, 1, 0, 0, 0, 0, time.Local).Unix()
	fmt.Println(expireYear, expireTimeUnix-1)
	// r := bufio.NewReaderSize(os.Stdin, 20480)
	// for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
	// 	fmt.Println(solution(string(line)))
	// }
}
