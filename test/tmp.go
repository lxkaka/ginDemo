package main

func main() {
	var (
		sign = make(chan int)
		done = make(chan struct{})
	)
	go do(sign, done)
	go command(sign)
	println("go...")

	<-done
	println("done")
}

func do(s <-chan int, d chan<- struct{}) {
	//select {
	//case <-s:
	//	println("get sign and do --->")
	//}
	//defer close(d)
	for i := range s {
		println("get sign", i)
	}
	d <- struct{}{}
}

func command(s chan<- int) {
	defer close(s)
	println("send sign")
	//time.Sleep(3 * time.Second)
	s <- 1
	s <- 2
}
