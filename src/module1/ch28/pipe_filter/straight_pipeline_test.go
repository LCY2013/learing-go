package pipe_filter

import (
	"fmt"
	"testing"
)

//测试流线型过滤器处理
func TestStraightPipeline(t *testing.T) {
	splitFilter := NewSplitFilter(",")
	toIntFilter := NewToIntFilter()
	sumFilter := NewSumFilter()
	sp := CreateStraightPipeline("csp", splitFilter, toIntFilter, sumFilter)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
}
