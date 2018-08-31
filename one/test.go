package main

func main() {
	x := 100
	var s = "hello,world!"
	println(x, s)

	if x > 0 {
		println("x")
	} else if x < 0 {
		print("-x")
	} else {
		println("0")
	}

	switch {
	case x > 0:
		println("x")
	case x < 0:
		println("-x")
	default:
		println("0")
	}

	for i := 0; i < 5; i++ {
		println(i)
	}

	arr := []int{100, 101, 102}

	for i, n := range arr {
		println(i, ":", n)
	}

}
