package maps

import "sync"

type RWLockMap struct {
	m    map[interface{}]interface{}
	lock sync.RWMutex
}

//读
func (rw *RWLockMap) Get(key interface{}) (interface{}, bool) {
	rw.lock.RLock()
	v, ok := rw.m[key]
	rw.lock.RUnlock()
	return v, ok
}

//写
func (rw *RWLockMap) Put(key interface{}, val interface{}) {
	rw.lock.Lock()
	rw.m[key] = val
	rw.lock.Unlock()
}

//删除
func (rw *RWLockMap) Del(key interface{}) {
	rw.lock.Lock()
	delete(rw.m, key)
	rw.lock.Unlock()
}

func CreateRWMap(limit int) *RWLockMap {
	rwMap := make(map[interface{}]interface{}, limit)
	return &RWLockMap{m: rwMap}
}
