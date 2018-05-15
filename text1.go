// 已知某广场的长宽（都为整数），现有瓷砖（数量大于0小于10的9次方），请问最少需要多少块能把广场铺满,瓷砖不能打碎，瓷砖长宽与广场长宽平行
package main

import (
	"fmt"
)

func needNumber(m, n, a int32) int32 {
	var w, h int32
	if m%a > 0 {
		w = m/a + 1
	} else {
		w = m / a
	}
	if n%a > 0 {
		h = n/a + 1
	} else {
		h = n / a
	}
	return w * h
}
func main() {
	var m, n, a int32
	for {
		fmt.Println("请输入场地的长和宽以及瓷砖的长度(整数，用空格间隔）:")
		fmt.Scanln(&m, &n, &a)
		num := needNumber(m, n, a)
		fmt.Println("需要的瓷砖数目为:", num)
		continue
	}
}
