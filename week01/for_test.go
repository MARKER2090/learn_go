package main

import (
	"testing"
)

func Test_Sum(t *testing.T) {
	if Sum(1, 2, 3) != 6 {
		t.Fatal("Sum err!")
	}
}
