package code

import "fmt"

func WriteOutput(outputChannel <-chan []byte) {
	for data := range outputChannel {
		fmt.Println("----------------------------------------OUTPUT: " + string(data))
	}
}
