//3人组队参加比赛，第一行输入要做的题数，下来的n行代表每道题队员有思路的人数0代表没思路，1代表有思路
//若一道题有两个及以上的人有思路那么这道题可以做，否则先不做
//输出可以做的题数
package main

import (
	"fmt"
)

func main() {
	var num, result, count int
	fmt.Scanln(&num)
	array := make([][]int, num+1)
	for i := 1; i <= num; i++ {
		subarray := make([]int, 3+1)
		count = 0
		for j := 1; j <= 3; j++ {
			fmt.Scan(&subarray[j])
			if subarray[j] == 1 {
				count++
			}
		}
		if count >= 2 {
			result++
		}
		array[i] = subarray
	}
	fmt.Println(result)

}
