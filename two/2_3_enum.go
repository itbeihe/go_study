package main

type color byte

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

const (
	black color = iota
	red
	blue
)

func test(c color) {
	println(c)
}
func main() {
	test(red)
	test(100)
	x := 3
	test(x)
}
