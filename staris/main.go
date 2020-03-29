package main

import "fmt"

//n个台阶， 一次可以走1步 也可以走两步， 有多少种走法

func walk(n int) int {
	if n == 1 {
		//如果只有一个台阶  只有一种走法
		return 1
	}
	if n == 2 {
		//如果有两个台阶  有两种走法  1，1 或  2
		return 2
	}
	return walk(n-1) + walk(n-2)
}
func main() {
	total := 30
	n := walk(total)
	fmt.Printf("若有%d个台阶，一共有%d个走法", total, n)
}
