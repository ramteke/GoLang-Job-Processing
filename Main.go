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

	stateReadChannel := make(chan code.ReadOp)
	stateWriteChannel := make(chan code.WriteOp)

	inputChannel := make(chan []byte, numJobs)
	for w := 1; w <= 3; w++ {
		go code.Worker(w, inputChannel, stateReadChannel, stateWriteChannel)
	}
	go code.StateProcessor(stateReadChannel, stateWriteChannel)

	for r := 0; r < 100; r++ {
		message := fmt.Sprintf("SEQ: %d,  CLUSTER: %d", r, rand.Intn(5))
		fmt.Println("Generated: ", message)
		inputChannel <- ([]byte(message))

	}

	close(inputChannel)

	time.Sleep(time.Second * 5)

}
