package main

//厄拉多塞筛法+6N+-1法求10000000以内素数

import (
	"fmt"
)

var num [10000000]bool

func initList() {
	//使用6N+-1法初始化数组
	for i := 0; i < 10000000; i++ {
		if i%6 == 5 || i%6 == 1 || i < 6 {
			num[i] = true
		}
	}
}

func sieve() {
	for i := 4; i < 10000000; i++ {
		if num[i] {
			for j := 2; j*i < 10000000; j++ {
				num[j*i] = false
			}
		}
	}
}

func main() {
	fmt.Println("Prime numbers in first 10M numbers")
	initList()
	sieve()
	//for i := 0; i < 10000000; i++ {
	//	if num[i] {
	//		fmt.Print(i)
	//		fmt.Print("  ")
	//	}
	//}
	fmt.Println("Complete!")

}
