package main

type X int

func (x *X) inc() {
	*x++
}

func main() {
	var z X
	z.inc()
	println(z)
}
