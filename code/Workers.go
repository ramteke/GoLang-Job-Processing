package code

import (
	"fmt"
)

type State struct {
	id         int
	inProgress bool
	counter    int
}

func Worker(workerId int, jobs <-chan []byte, stateReadChannel chan ReadOp, stateWriteChannel chan WriteOp) {
	fmt.Println("## WORKER: Created : ", workerId)
	counter := 0

	processChannel := make(chan []byte, 1)
	go ProcessData(processChannel, stateReadChannel, stateWriteChannel)

	for job := range jobs {
		fmt.Println("------------- WORKER [", workerId, "] DISPATCHING : ", string(job))
		message := string(job)
		message = message + fmt.Sprintf(" via Worker: %d", workerId)
		processChannel <- []byte(message)
		counter++
	}

	close(processChannel)
	fmt.Println(" ################ WORKER [", workerId, "] processed ", counter, "jobs  ###############")

}
