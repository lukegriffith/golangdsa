package workerpools

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type WorkItem struct {
	ID      int
	Payload int
}

type Result struct {
	WorkID int
	Value  int
}

func Worker(ctx context.Context, id int, tasks <-chan WorkItem, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[Worker %d] Context done, shutting down.\n", id)
			return
		case task, ok := <-tasks:
			if !ok {
				return
			}
			fmt.Printf("[Worker %d] Processing task %d...\n", id, task.ID)
			processedValue := task.Payload * task.Payload
			time.Sleep(time.Second * 3)
			results <- Result{
				WorkID: task.ID,
				Value:  processedValue,
			}
		}
	}
}

func Main() {
	const numWorkers = 3
	const numTasks = 25

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	tasks := make(chan WorkItem, numTasks)
	results := make(chan Result, numTasks)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(ctx, i, tasks, results, &wg)
	}

	for i := 1; i <= numTasks; i++ {
		tasks <- WorkItem{
			ID:      i,
			Payload: i,
		}
	}

	close(tasks)
	wg.Wait()
	close(results)
	for res := range results {
		fmt.Printf("Task %d processed, result value: %d\n", res.WorkID, res.Value)
	}
}
