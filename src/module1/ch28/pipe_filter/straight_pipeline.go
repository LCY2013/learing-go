package pipe_filter

import "errors"

var StraightProcessWrongError = errors.New("process error")

//StraightPipeline is composed of the filters,and the filters are piled as a process
type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

//CreateStraightPipeline create a new StraightPipeline
func CreateStraightPipeline(name string, filter ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filter,
	}
}

//Process is to process the coming data by the pipeline
func (sp *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	for _, filter := range *sp.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return nil, err
		}
		data = ret
	}
	return ret, err
}
