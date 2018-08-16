package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 插入排序
func inSort(arr []string) []int {
	sortArr := make([]int, len(arr))
	// 字符串数组转整形
	newArr := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		newArr[i], _ = strconv.Atoi(arr[i])
	}
	// 插入排序
	sortArr[0] = newArr[0]
	for i := 1; i < len(newArr); i++ {
		for j := i - 1; j >= 0; j-- {
			if newArr[i] < sortArr[j] {
				sortArr[j+1] = sortArr[j]
				sortArr[j] = newArr[i]
			} else {
				sortArr[j+1] = newArr[i]
				break
			}
		}
	}
	return sortArr
}

// 多点之间最少成本->相邻两点最少成本
func solution(line string) string {
	// 处理数据
	line = strings.TrimSpace(line)
	arr := strings.Split(line, ";")
	cost := strings.Split(arr[0], " ")
	// 建站成本
	costA := cost[0]
	costAF, _ := strconv.ParseFloat(costA, 64)
	// 单位费用
	costB := cost[1]
	costBF, _ := strconv.ParseFloat(costB, 64)
	distance := strings.Split(arr[1], " ")
	distanceSort := inSort(distance)
	// 花费
	spend := costAF
	// 下一个和上一个距离x单位费用大于建站成本*2,建站，否则拉线
	for i := 1; i < len(distanceSort); i++ {
		if float64(distanceSort[i]-distanceSort[i-1]) > (costAF * 2 / costBF) {
			spend += costAF
		} else {
			spend += float64(distanceSort[i]-distanceSort[i-1]) / 2 * costBF
		}
	}
	// 小数点后补零
	spendStr := strconv.FormatFloat(spend, 'f', 1, 64)
	return spendStr
}

func main() {
	r := bufio.NewReaderSize(os.Stdin, 20480)
	for line, _, err := r.ReadLine(); err == nil; line, _, err = r.ReadLine() {
		fmt.Println(solution(string(line)))
	}
}
