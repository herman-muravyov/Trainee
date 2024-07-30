package main

import (
	"fmt"
	"sync"
	"task/internal"
	"task/internal/variables"
	"task/pkg/types"
)

func main() {
	wg := sync.WaitGroup{}
	mu := internal.SyncMutex{}

	bakeCh := make(chan types.Cake)
	packCh := make(chan types.Cake)

	for i := 0; i < variables.K; i++ {
		wg.Add(1)
		go mu.Baking(i, variables.T1, bakeCh, &wg)
		// go internal.Baking(i, variables.T1, bakeCh, &wg)
	}

	for j := 0; j < variables.K; j++ {
		wg.Add(1)
		go mu.Packing(j, variables.T2, bakeCh, packCh, &wg)
		// go internal.Packing(j, variables.T2, bakeCh, packCh, &wg)
	}

	go func() {
		wg.Wait()
		close(bakeCh)
		close(packCh)
	}()

	count := 0
	for cake := range packCh {
		fmt.Printf("Cake %d: BakedBy=%d, BakeTime=%dms, PackedBy=%d, PackTime=%dms\n",
			count, cake.BakedBy, cake.BakeTime, cake.PackedBy, cake.PackTime)
		count++
		if count >= variables.K {
			break
		}
	}

	fmt.Println("All cakes processed.")
}
