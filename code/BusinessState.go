package code

func StateProcessor(stateReadChannel chan ReadOp, stateWriteChannel chan WriteOp) {
	var state = make(map[int]int)
	for {
		select {
		case read := <-stateReadChannel:
			read.resp <- state[read.key]
		case write := <-stateWriteChannel:
			state[write.key] = write.val
			write.resp <- true
		}
	}

}
