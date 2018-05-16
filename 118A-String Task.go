//输入一个字符串，都是字母可以大写可以小写
//输出时去掉所有的元音a e i o u y，所有的大写变为小写，所有的辅音前加.
package main

import (
	"fmt"
	"strings"
)

func main() {
	var oldStr string
	fmt.Scanln(&oldStr)
	var newStr = make([]string, 0)
	oldStr = strings.ToLower(oldStr) //将字符串变为小写
	for index := 0; index < len(oldStr); index++ {
		str := oldStr[index : index+1]
		if str != "a" && str != "e" && str != "i" && str != "o" && str != "u" && str != "y" {
			newStr = append(newStr, str)
		}
	}
	fmt.Println("." + strings.Join(newStr, "."))

}
