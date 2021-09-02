package reference

import (
	"context"
	"errors"
	reference "reference/pkg/grpc"
)

type Service struct {
	data map[string]int32
	reference.GetterServer
}

func (s Service) GetReferenceId(_ context.Context, name *reference.ReferenceName) (*reference.ReferenceId, error) {
	if name == nil {
		return nil, errors.New("запрос пустой")
	}

	ans, ok := s.data[name.GetName()]
	if !ok {
		return nil, errors.New("not found")
	}

	return &reference.ReferenceId{Id: ans}, nil
}

func NewService() reference.GetterServer {
	return &Service{map[string]int32{
		"electricity": 1,
		"temperature": 2,
		"power":       3,
	}, nil}
}
