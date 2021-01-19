package main

func main() {
	var (
		sign = make(chan int)
		done = make(chan bool)
	)
	go do(sign, done)
	go command(sign)
	println("go...")

	<-done
	println("done")
}

func do(s <-chan int, d chan<- bool) {
	//select {
	//case <-s:
	//	println("get sign and do --->")
	//}
	//defer close(d)
	for i := range s {
		println("get sign", i)
	}
	d <- true
}

func command(s chan<- int) {
	defer close(s)
	println("send sign")
	//time.Sleep(3 * time.Second)
	s <- 1
	s <- 2
}
