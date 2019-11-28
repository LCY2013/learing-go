package pipe_filter

import (
	"errors"
	"fmt"
	"strconv"
)

var ToIntFilterWrongFormatError = errors.New("input data should be []string")

//第二步
type ToIntFilter struct {
}

func NewToIntFilter() *ToIntFilter {
	return &ToIntFilter{}
}

func (tf *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string) ////检查数据格式/类型，是否可以处理
	if !ok {
		return nil, ToIntFilterWrongFormatError
	}
	ret := []int{}
	for _, part := range parts {
		it, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = append(ret, it)
	}

	return ret, nil
}
