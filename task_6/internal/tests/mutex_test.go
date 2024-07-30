package tests

import (
	"math/rand"
	"runtime"
	"sync"
	"task/internal/variables"
	"task/pkg/types"
	"testing"
)

type SyncMutex struct {
	mutex sync.RWMutex
}

func (mu *SyncMutex) BakingPerfTest(id, t1 int, bakeCh chan<- types.Cake, wg *sync.WaitGroup) {
	bakeTime := id + rand.Intn(t1+1)

	mu.mutex.Lock()
	defer mu.mutex.Unlock()

	backings := types.Cake{
		BakedBy:  id,
		BakeTime: bakeTime,
	}

	bakeCh <- backings
}

func (mu *SyncMutex) PackingPerfTest(
	id,
	t2 int,
	bakeChan <-chan types.Cake,
	packChan chan<- types.Cake,
	wg *sync.WaitGroup,
) {
	packTime := id + rand.Intn(t2+1)

	for cake := range bakeChan {
		mu.mutex.Lock()
		cake.PackedBy = id
		cake.PackTime = packTime
		mu.mutex.Unlock()

		packChan <- cake
	}
}

func BenchmarkTestMutexOptimization(b *testing.B) {
	wg := sync.WaitGroup{}
	bakeCh := make(chan types.Cake)
	packCh := make(chan types.Cake)

	wg.Add(runtime.NumCPU())

	mu := SyncMutex{}

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < variables.K; j++ {
				mu.BakingPerfTest(j, variables.T1, bakeCh, &wg)
			}
		}(i)
	}

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < variables.K; j++ {
				mu.PackingPerfTest(idx, variables.T2, bakeCh, packCh, &wg)
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(bakeCh)
		close(packCh)
	}()
}
