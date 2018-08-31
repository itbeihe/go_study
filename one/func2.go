package main

import "fmt"

type user struct {
	name string
	age  byte
}

type manager struct {
	user
	title string
}

func (mu manager) ToString() string {
	return fmt.Sprintf("%+v", mu)
}

func main() {
	var m manager
	m.name = "Tom"
	m.age = 29
	m.title = "CTO"

	fmt.Println(m.ToString())
}
