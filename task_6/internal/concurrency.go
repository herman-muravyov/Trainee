package internal

import (
	"math/rand"
	"sync"
	"task/pkg/types"
)

func Baking(id, t1 int, bakeCh chan<- types.Cake, wg *sync.WaitGroup) {
	bakeTime := id + rand.Intn(t1+1)

	defer wg.Done()
	backings := types.Cake{
		BakedBy:  id,
		BakeTime: bakeTime,
	}
	bakeCh <- backings
}

func Packing(
	id,
	t2 int,
	bakeChan <-chan types.Cake,
	packChan chan<- types.Cake,
	wg *sync.WaitGroup,
) {
	packTime := id + rand.Intn(t2+1)

	defer wg.Done()
	for cake := range bakeChan {
		cake.PackedBy = id
		cake.PackTime = packTime

		packChan <- cake
	}
}
