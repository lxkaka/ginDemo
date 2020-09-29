package main

//func main() {
//	var (
//		sign = make(chan int)
//		//done = make(chan bool)
//		wg = &sync.WaitGroup{}
//	)
//	wg.Add(2)
//	go do(sign, wg)
//	go command(sign, wg)
//	println("go...")
//
//	//<-done
//	wg.Wait()
//	println("done")
//}
//
//func do(s <-chan int, wg *sync.WaitGroup) {
//	//select {
//	//case <-s:
//	//	println("get sign and do --->")
//	//}
//	//defer close(d)
//	for i := range s {
//		println("get sign", i)
//	}
//	//d <- true
//	wg.Done()
//}
//
//func command(s chan<- int, wg *sync.WaitGroup) {
//	defer close(s)
//	println("send sign")
//	time.Sleep(3 * time.Second)
//	s <- 1
//	wg.Done()
//}

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("Hello, playground")
	//m := make(map[string]int64)
	m := &sync.Map{}
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(a int) {
			//m[strconv.FormatInt(c, 10)] = c
			m.Store(strconv.FormatInt(int64(a), 10), a)
			wg.Done()
		}(i)
	}
	wg.Wait()
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

}
