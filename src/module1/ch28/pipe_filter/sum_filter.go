package pipe_filter

import "errors"

var SumFilterWrongFormatError = errors.New("input data should be []int")

//第三步
type SumFilter struct {
}

func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

func (sf *SumFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	sum := 0
	for _, part := range parts {
		sum += part
	}
	return sum, nil
}
