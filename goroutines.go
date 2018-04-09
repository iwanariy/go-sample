package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepN(t int) {
}

func main() {
	wg := &sync.WaitGroup{}

	// working
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// working
	for i := 0; i < 5; i++ {
		wg.Add(1)
		j := i
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(j) * time.Second)
			fmt.Println("Sleep " + fmt.Sprint(j) + " seconds.")
		}()
	}

	// not working
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println("Sleep " + fmt.Sprint(i) + " seconds.")
		}()
	}

	wg.Wait()
}
