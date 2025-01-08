package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Code block for sequencial execution
	t1 := time.Now()
	sequencial()
	t1f := time.Now()
	t1t := t1f.Sub(t1)
	fmt.Println("Sequencial took", t1t)

	// Code block for routines. It cna apply parallelism due to free CPU
	t2 := time.Now()
	parallelism()
	t2f := time.Now()
	t2t := t2f.Sub(t2)
	fmt.Println("Parallelism took", t2t)
}

func sequencial() {
	for i := 0; i < 5; i++ {
		fmt.Println("Execution", i)

		waitTime()
	}
}

func parallelism() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		fmt.Println("Execution", i)

		go waitTimeParallel(&wg)
	}
	wg.Wait()
}

func waitTime() {
	time.Sleep(5 * time.Second)
}

func waitTimeParallel(wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	wg.Done()
}
