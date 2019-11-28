package maps

import "sync"

type SyncMapAdapter struct {
	sm sync.Map
}

//读
func (sma *SyncMapAdapter) Get(key interface{}) (interface{}, bool) {
	v, ok := sma.sm.Load(key)
	return v, ok
}

//写
func (sma *SyncMapAdapter) Put(key interface{}, val interface{}) {
	sma.sm.Store(key, val)
}

//删除
func (sma *SyncMapAdapter) Del(key interface{}) {
	sma.sm.Delete(key)
}

func CreateSyncMap() *SyncMapAdapter {
	return &SyncMapAdapter{}
}
