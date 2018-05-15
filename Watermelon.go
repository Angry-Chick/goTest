//有一个西瓜，重量在1到100公斤，切一刀分给两个人，可以不均分，但是两个人的重量必须是偶数
package main

import "fmt"

func main() {
	var k uint8
	fmt.Scanln(&k)
	if k == 2 {
		fmt.Println("NO")
	} else {
		if k%2 != 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("yes")
		}
	}
}

//
