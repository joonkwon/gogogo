package main

import (
	"fmt"
	"time"
)

func count(thing string) {
	for i := 1; i <= 5; i++ {
		// for i := 1; true; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	// 1. sheep never finishes  -------------------
	// count("sheep")
	// count("fish")

	// 2. go routines  -------------------
	// go count("sheep")
	// count("fish")

	// 3. go routines  -------------------
	// go count("sheep")
	// go count("fish")

	// 4. 2 seconds  -------------------
	// time.Sleep(time.Second * 2)

	//5. Add Scanln  -------------------
	// fmt.Scanln()

	// 6. using wait group  -------------------
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	count("sheep")
	// 	wg.Done() //That will descrease wg i 1
	// }()

	// wg.Wait()

	// 7 . Communication back to the main routineL: channel
	// c := make(chan string)
	// go countCh("sheep", c)

	// msg := <-c //sending / receving from channel are blocking operation. it wait until something gets to channel and receiver is ready to receive
	// fmt.Println(msg)

	// 8. wrap in a for loop
	// for { //fatal
	// 	msg := <-c //sending / receving from channel are blocking operation. it wait until something gets to channel and receiver is ready to receive
	// 	fmt.Println("z" + msg)
	// }

	// 9. close the channel

	// 10. check if the cahnnel is open
	// for {
	// 	msg, isOpen := <-c //sending / receving from channel are blocking operation. it wait until something gets to channel and receiver is ready to receive
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	// 11. Loop range
	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	// 12. Deadlock! //never receives blocked at send. this is the reason we haveto receive in a differente go routine
	// x := make(chan string)
	// x <- "hello"
	// msg := <-x
	// fmt.Println(msg)

	// 13. buffered channel
	x := make(chan string, 2) //we say do not block undtil the channel has 2 strings
	x <- "hello"
	msg := <-x
	fmt.Println(msg)

	// 14. buffered channel
	// x := make(chan string, 2) //we say do not block undtil the channel has 2 strings
	// x <- "hello"
	// x <- "world"
	// msg := <-x
	// fmt.Println(msg)
	// msg = <-x
	// fmt.Println(msg)

	// 15. 3rd value //channel is going to be full
	// x <- "three"

	//16. Another blocking example
	// c1 := make(chan string)
	// c2 := make(chan string)

	// go func() {
	// 	for {
	// 		c1 <- "Every 500ms"
	// 		time.Sleep(time.Millisecond * 500)
	// 	}
	// }()

	// go func() {
	// 	for {
	// 		c2 <- "Every 2s"
	// 		time.Sleep(time.Millisecond * 2000)
	// 	}
	// }()

	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	// 17. using select: being able to receive whatever channel has a message
	// for {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println(msg2)
	// 	}
	// }

	//18. multiple workers pulling out from the channel
	// jobs := make(chan int, 100)
	// results := make(chan int, 100)

	// go worker(jobs, results)

	// for i := 0; i < 100; i++ {
	// 	jobs <- i
	// }
	// close(jobs)

	// for j := 0; j < 100; j++ {
	// 	fmt.Println(<-results)
	// }

	//19. Add workers

	//20. channel directions in workers
}

func countCh(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		//fmt.Println(i, thing)
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

// // func worker(jobs <-chan int, results chan<- int) {
// func worker(jobs chan int, results chan int) {
// 	for n := range jobs {
// 		results <- fib(n)
// 	}
// }

// func fib(n int) int {
// 	if n <= 1 {
// 		return n
// 	}

// 	return fib(n-1) + fib(n-2)
// }
