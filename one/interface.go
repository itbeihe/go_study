package main

type user struct {
	name string
	age  byte
}

func (u user) Print() string {
	//fmt.Printf("%+v\n",u.name)
	return u.name
}

type Printer interface {
	Print() string
}

func main() {
	var u user
	u.name = "Tom"
	u.age = 26

	var p Printer = u
	var z = p.Print()
	print(z)
}
