package maps

import (
	"strconv"
	"sync"
	"testing"
)

const (
	NumOfReader = 200
	NumOfWriter = 100
)

type Map interface {
	Put(key interface{}, val interface{})
	Get(key interface{}) (interface{}, bool)
	Del(key interface{})
}

func benchmarkMap(b *testing.B, hm Map) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for i := 0; i < NumOfWriter; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Put(strconv.Itoa(i), i*i)
					hm.Put(strconv.Itoa(i), i*i)
					hm.Del(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
		for i := 0; i < NumOfReader; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 100; i++ {
					hm.Get(strconv.Itoa(i))
				}
				wg.Done()
			}()
		}
	}
}

func BenchmarkMap(b *testing.B) {
	b.Run("map run with rwLock", func(b *testing.B) {
		hm := CreateRWMap(0)
		benchmarkMap(b, hm)
	})
	b.Run("map run with sync", func(b *testing.B) {
		hm := CreateSyncMap()
		benchmarkMap(b, hm)
	})
	b.Run("map run with concurrent", func(b *testing.B) {
		hm := CreateConcurrentMapAdapter(199)
		benchmarkMap(b, hm)
	})
}
