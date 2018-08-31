package main

func main() {
	//data := [3]string{"a","b","c"}
	data := "ab12"
	for i, s := range data {
		println(i, string(s))
	}
}
