package main

// 消费者
func consumer(data chan int, done chan bool) {
	for x := range data {
		println("recv:", x)
	}
	done <- true
}

//生产者
func producer(data chan int) {
	for i := 0; i < 4; i++ {
		println("send:", i)
		data <- i
	}
	close(data)
}

func main() {
	done := make(chan bool)
	data := make(chan int)

	go consumer(data, done)
	go producer(data)

	<-done
}
