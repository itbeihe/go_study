package main

func main() {
	x := 100
	p := (*int)(&x)
	println(p)
}
