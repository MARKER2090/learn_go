package main

func main() {
	a1()
	a2()
}

func a1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}

func a2() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}
