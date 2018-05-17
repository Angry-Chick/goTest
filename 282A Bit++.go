//初始值为0
//如果是++X或者X++则加1 --X或者X--则减1
package main

import (
	"fmt"
)

func main() {
	var num, result int
	fmt.Scanln(&num)
	array := make([]string, num+1)
	result = 0
	for i := 1; i <= num; i++ {
		fmt.Scanln(&array[i])
		if array[i] == "++X" || array[i] == "X++" {
			result++
		} else {
			result--
		}
	}
	fmt.Println(result)

}
