package main

import (
	"fmt"
	"time"
)

func main() {
	// c := make(chan string)

	// //Mengirim data kepada channel
	// c <- value

	// //Menerima data dari channel
	// result := <-c

	fmt.Println("======================================")
	//Channels (Implementing channel)

	c := make(chan string)

	go introduce("Tri", c)

	go introduce("Wah", c)

	go introduce("Yono", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	close(c)
}

func introduce(student string, c2 chan string) {
	result := fmt.Sprintf("Hello... Namaku %s", student)

	c2 <- result

	fmt.Println("======================================")

	// UnBuffered Channel (Buffered vs unbuffered channel)
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleep for 2 seconds")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	close(c1)
	time.Sleep(time.Second)

	fmt.Println("======================================")

	//Buffered Channel (Buffered vs unbuffered channel)
	c3 := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("fun goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("fun goroutine #%d after sending data into the channel\n", i)
		}

		close(c)
	}(c3)

	fmt.Println("Main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c2 { //v = <- c2
		fmt.Println("main goroutine received value from channel:", v)
	}

	fmt.Println("======================================")

	//Channel (Select)
	c4 := make(chan string)
	c5 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)

		c4 <- "Yooo!"
	}()

	go func() {
		time.Sleep(1 * time.Second)

		c5 <- "Waaw"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg4 := <-c4:
			fmt.Println("Received", msg4)
		case msg5 := <-c5:
			fmt.Println("Received", msg5)
		}
	}
}
