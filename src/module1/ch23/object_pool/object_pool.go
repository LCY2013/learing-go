package object_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {
}

//定义对象常量池
type ObjectPool struct {
	obj chan *ReusableObj
}

//初始化常量池
func CreateReusableObjPool(num int) (*ObjectPool, error) {
	if num <= 0 {
		return nil, errors.New("num is allow to lower then zero")
	}
	//定义常量池信息
	objP := ObjectPool{obj: make(chan *ReusableObj, 10)}
	for i := 0; i < num; i++ {
		objP.obj <- &ReusableObj{}
	}
	return &objP, nil
}

//获取对象
func (op *ObjectPool) GetReusableObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-op.obj:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

//释放对象
func (op *ObjectPool) ReleaseReusableObj(rs *ReusableObj) error {
	select {
	case op.obj <- rs:
		return nil
	default:
		return errors.New("overFlow")
	}
}
