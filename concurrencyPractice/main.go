package concurrencyPractice

import (
	"fmt"
	"sync"
)

func generateNumbers(total int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Generating number %d\n", idx)
	}
}

func printNumbers(waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for idx := 1; idx <= 3; idx++ {
		fmt.Printf("Printing number %d\n", idx)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)
	go printNumbers(&waitGroup)
	go generateNumbers(3)
}
