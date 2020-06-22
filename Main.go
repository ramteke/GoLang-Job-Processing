package main

import (
	"Threads/code"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Test output channel
	numJobs := 10

	inputChannel := make(chan []byte, numJobs)
	for w := 1; w <= 3; w++ {
		go code.Worker(w, inputChannel)
	}

	for r := 0; r < 100; r++ {
		message := fmt.Sprintf("SEQ: %d,  CLUSTER: %d", r, rand.Intn(5))
		fmt.Println("Generated: ", message)
		inputChannel <- ([]byte(message))

	}

	close(inputChannel)

	time.Sleep(time.Second * 5)

}
