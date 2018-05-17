//在一个只有0和1的字符串中如果出现至少连续的7个1或0那么就输出YES，没有就输出NO
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	result := 0
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	for i := 1; i <= len(str1); i++ {
		if str1[i-1:i] == str2[i-1:i] {
			continue
		}
		if str1[i-1:i] > str2[i-1:i] {
			result = 1
			break
		}
		if str1[i-1:i] < str2[i-1:i] {
			result = -1
			break
		}
	}
	fmt.Println(result)
}
