package repository

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("worker", id, "started")
	time.Sleep(2 * time.Second)
	fmt.Println("Worker", id, "finished")

}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(1, &wg)
	go worker(2, &wg)

	wg.Wait()
	fmt.Println("All workers completed")
}
