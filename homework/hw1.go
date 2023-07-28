/*
实现切片的删除操作
实现删除切片特定下标元素的方法。
• 要求一：能够实现删除操作就可以。
• 要求二：考虑使用比较高性能的实现。
• 要求三：改造为泛型方法。
• 要求四：支持缩容，并且设计缩容机制。
*/
package main

import (
	"errors"
	"fmt"
)

type myslice[T comparable] []T

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//要求一和二
	fmt.Println(s)
	fmt.Println("len(s),cap(s):", len(s), cap(s))
	s1 := delWordsV1(3, s)
	fmt.Println(s1)
	fmt.Println("len(s1),cap(s1):", len(s1), cap(s1)) //9,10

	//要求三
	s2, err := delWordsV3(5, s1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s2)
		fmt.Println("len(s2),cap(s2):", len(s2), cap(s2)) //8,10

	}

	//要求四
	s3, err := delWordsV4(5, s2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s3)
		fmt.Println("len(s3),cap(s3):", len(s3), cap(s3)) //7,7
	}
}

// 要求一：实现简单的删除操作
func delWordsV1(n int, s []int) []int {
	return append(s[:n-1], s[n:]...)
}

// 要求二：考虑使用比较高性能的实现。(我觉得delWordsV1已经很简约很高性能了)

// 要求三：改造为泛型方法
func delWordsV3[T comparable](n int, s []T) ([]T, error) {
	if n < 0 || n > len(s) {
		return s, errors.New("您的下标超出切片元素范围")
	} else {
		return append(s[:n-1], s[n:]...), nil
	}
}

//要求四：支持缩容，并且设计缩容机制。
func delWordsV4[T comparable](n int, s []T) ([]T, error) {
	if n < 0 || n > len(s) {
		return s, errors.New("您的下标超出切片元素范围")
	} else {
		var ss []T
		ss = append(s[:n-1], s[n:]...)
		return ss[:len(ss):len(ss)], nil

	}
}
