package internal

import (
	"math/rand"
	"sync"
	"task/pkg/types"
)

type SyncMutex struct {
	mutex sync.Mutex
}

func (mu *SyncMutex) Baking(id, t1 int, bakeCh chan<- types.Cake, wg *sync.WaitGroup) {
	defer wg.Done()
	bakeTime := id + rand.Intn(t1+1)

	mu.mutex.Lock()
	defer mu.mutex.Unlock()

	backings := types.Cake{
		BakedBy:  id,
		BakeTime: bakeTime,
	}
	bakeCh <- backings
}

func (mu *SyncMutex) Packing(
	id,
	t2 int,
	bakeChan <-chan types.Cake,
	packChan chan<- types.Cake,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	packTime := id + rand.Intn(t2+1)

	for cake := range bakeChan {
		mu.mutex.Lock()
		cake.PackedBy = id
		cake.PackTime = packTime
		mu.mutex.Unlock()

		packChan <- cake
	}
}
