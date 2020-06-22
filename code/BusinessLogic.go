package code

import (
	"fmt"
	"time"
)

func ProcessData(jobs <-chan []byte) {
	outputChannel := make(chan []byte, 1)
	go WriteOutput(outputChannel)

	for job := range jobs {
		fmt.Println("-----------------------", "PROCESSING: ", string(job))

		/* Does some processing which takes time */
		time.Sleep(time.Millisecond)
		outputChannel <- job
	}

	close(outputChannel)

}
