package slice

import (
	"errors"
	"reflect"
)

type slice struct {
	sliceAny interface{}
}

func New(a interface{}) slice {
	return slice{
		sliceAny: a,
	}
}

func (s *slice) sliceType() reflect.Type {
	return reflect.TypeOf(s.sliceAny)
}

func (s *slice) ValueAt(index int) (interface{}, error) {
	a, ok := s.sliceAny.([]int)
	if !ok {
		return nil, errors.New("slice is not of []int type")
	}
	return a[index], nil
}
