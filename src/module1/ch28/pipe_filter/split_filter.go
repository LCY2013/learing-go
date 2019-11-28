package pipe_filter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data should be string")

//第一步
type SplitFilter struct {
	delimiter string
}

//新建一个切割过滤器,构造方法
func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

//结构体处理器
func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string) //检查数据格式/类型，是否可以处理
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
