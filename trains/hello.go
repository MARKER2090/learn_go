package main

import (
	"fmt"
	"hello/pa"
	"hello/pa/pb"
)

func main() {
	a := "hello?"
	fmt.Println(a)
	pb.PbFunc()
	pa.PaFunc()
}
