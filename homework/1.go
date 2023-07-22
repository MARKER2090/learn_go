/*
实现切片的删除操作
实现删除切片特定下标元素的方法。
• 要求一：能够实现删除操作就可以。
• 要求二：考虑使用比较高性能的实现。
• 要求三：改造为泛型方法。
• 要求四：支持缩容，并且设计缩容机制。
*/
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(delWords(3, s))
}

func delWords(n int, s []int) []int {
	return append(s[:n-1], s[n:]...)
}
