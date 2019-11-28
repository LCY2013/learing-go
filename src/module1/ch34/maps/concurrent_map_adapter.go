package maps

import (
	concurrent_map "github.com/easierway/concurrent_map"
)

type ConcurrentMapAdapter struct {
	cm *concurrent_map.ConcurrentMap
}

//读
func (cma *ConcurrentMapAdapter) Get(key interface{}) (interface{}, bool) {
	v, ok := cma.cm.Get(concurrent_map.StrKey(key.(string)))
	return v, ok
}

//写
func (cma *ConcurrentMapAdapter) Put(key interface{}, val interface{}) {
	cma.cm.Set(concurrent_map.StrKey(key.(string)), val)
}

//删除
func (cma *ConcurrentMapAdapter) Del(key interface{}) {
	cma.cm.Del(concurrent_map.StrKey(key.(string)))
}

func CreateConcurrentMapAdapter(limit int) *ConcurrentMapAdapter {
	return &ConcurrentMapAdapter{cm: concurrent_map.CreateConcurrentMap(limit)}
}
