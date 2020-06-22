package code

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type ReadOp struct {
	key  int
	resp chan int
}
type WriteOp struct {
	key  int
	val  int
	resp chan bool
}

func ProcessData(jobs <-chan []byte, stateReadChannel chan ReadOp, stateWriteChannel chan WriteOp) {

	outputChannel := make(chan []byte, 1)

	go WriteOutput(outputChannel)

	for job := range jobs {
		/* Does some processing which takes time */
		message := string(job)
		clusterIndex := strings.LastIndex(message, "CLUSTER:")
		clusterIndex = clusterIndex + len("CLUSTER:")
		viaIndex := strings.LastIndex(message, "via")
		clusterId, _ := strconv.Atoi(strings.TrimSpace(string(message[clusterIndex:viaIndex])))

		read := ReadOp{
			key:  clusterId,
			resp: make(chan int)}
		stateReadChannel <- read
		counter := <-read.resp
		counter++

		write := WriteOp{
			key:  clusterId,
			val:  counter,
			resp: make(chan bool)}
		stateWriteChannel <- write
		<-write.resp

		fmt.Println("-----------------------", ": BUSINESS-LOGIC: ", string(job), " Counter Value: ", counter)

		time.Sleep(time.Millisecond)
		outputChannel <- job
	}

	close(outputChannel)

}
