//在一个只有0和1的字符串中如果出现至少连续的7个1或0那么就输出YES，没有就输出NO
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	result := "NO"
	fmt.Scanln(&str)
	if strings.Contains(str, "0000000") {
		result = "YES"
	}
	if strings.Contains(str, "1111111") {
		result = "YES"
	}
	fmt.Println(result)

}
