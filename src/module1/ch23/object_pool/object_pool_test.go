package object_pool

import (
	"fmt"
	"testing"
	"time"
)

//测试对象池
func TestCreateReusableObjPool(t *testing.T) {
	num := 10
	if pool, err := CreateReusableObjPool(num); err == nil {
		//开始释放对象
		if err := pool.ReleaseReusableObj(&ReusableObj{}); err != nil {
			fmt.Println(err)
		}

		for i := 0; i < num; i++ {
			if ret, err := pool.GetReusableObj(time.Millisecond * 1); err == nil {
				fmt.Printf("%T \n", ret)
			} else {
				fmt.Println(err)
			}
		}
		if ret, err := pool.GetReusableObj(time.Millisecond * 1); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%T \n", ret)
		}
	} else {
		fmt.Println(err)
	}

	fmt.Println("Done")
}
