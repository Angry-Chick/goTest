//第一行接收要输出的单词总个数
//然后输入每个单词
//输出要求：单词长度小于10原样输出，大于10输出“首字母+中间省略字母个数+尾字母”这样的格式
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var num int

	fmt.Scanln(&num)

	array := make([]string, num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&array[i])
	}

	for i := 0; i < num; i++ {
		if len(array[i]) <= 10 {
			fmt.Println(array[i]) //小于10原样输出
		} else {
			//大于10将首字母，省略字母个数，尾字母存到一个字符串数组中，最后通过Join函数拼接成一个字符串输出
			s := []string{array[i][0:1], strconv.Itoa(len(array[i]) - 2), array[i][len(array[i])-1 : len(array[i])]}
			fmt.Println(strings.Join(s, ""))
		}
	}

}
