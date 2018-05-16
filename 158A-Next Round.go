//第一行为参赛选手人数和第k位选手的分数
//如果第k位选手不是0分那么大于等于这个分数的人都通过，如果得了0分，就把所有不是0分的通过
//测试数据已经从大到小排好序了
package main

import "fmt"

func main() {
	var total, num, result int
	fmt.Scanln(&total, &num)
	var array = make([]int, total+1)
	for i := 1; i <= total; i++ {
		fmt.Scan(&array[i])
	}
	result = 0
	if array[num] == 0 {
		for i := 1; i <= num; i++ {
			if array[i] > 0 {
				result++
			}
		}
	} else {
		result = num
		for i := num; i < total; i++ {
			if array[i+1] >= array[num] {
				result++
			}
		}
	}
	fmt.Println(result)
}
