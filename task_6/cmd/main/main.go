package main

import (
	"fmt"
	"math/rand"
	"sync"
	"task/internal/variables"
	"task/pkg/types"
	"time"
)

/*
	Using "Pipe" pattern
*/

var wg sync.WaitGroup

func baking(id, t1 int, bakeCh chan<- types.Cake, wg *sync.WaitGroup) chan<- types.Cake {
	// bakeTime - to generate a random number in the range, so that the baking time varies in the range i +- t1
	bakeTime := id + rand.Intn(2*t1+1) - t1

	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(t1) * time.Millisecond)
		backings := types.Cake{
			BakedBy:  id,
			BakeTime: bakeTime,
		}

		bakeCh <- backings
	}()

	return bakeCh
}

func main() {
	bakeCh := make(chan types.Cake)

	for i := 0; i < variables.K; i++ {
		wg.Add(1)
		baking(i, variables.T1, bakeCh, &wg)
	}

	for i := 0; i < variables.K; i++ {
		cake := <-bakeCh
		fmt.Printf("Cake baked by %d in %dms\n", cake.BakedBy, cake.BakeTime)
	}

	go func() {
		wg.Wait()
		close(bakeCh)
	}()
}
