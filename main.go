package main

import (
	"sync"
)

const maxCount = 500000

func once(onceBody func(), done chan bool, once *sync.Once) {
	for i := 0; i < maxCount; i++ {
		go func() {
			once.Do(onceBody)

			done <- true
		}()
	}

	for i := 0; i < maxCount; i++ {
		<-done
	}
}

func branch(onceBody func(), done chan bool) {
	for i := 0; i < maxCount; i++ {
		go func(i int) {
			if i == 0 {
				onceBody()
			}

			done <- true
		}(i)
	}

	for i := 0; i < maxCount; i++ {
		<-done
	}
}

func main() {}
