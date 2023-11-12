package random

import (
	"github.com/life4/genesis/slices"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/sigmasee/sigmasee/shared/enterprise/tuples"
)

type RandomHelper interface {
	Generate() (string, error)
	GenerateMany(count int) ([]string, error)
}

type randomHelper struct {
}

func NewRandomHelper() (RandomHelper, error) {
	return &randomHelper{}, nil
}

func (s *randomHelper) Generate() (string, error) {
	return gonanoid.New()
}

func (s *randomHelper) GenerateMany(count int) ([]string, error) {
	var indexes []int
	for i := 0; i < count; i++ {
		indexes = append(indexes, i)
	}

	result := slices.Map(indexes, func(item int) tuples.ValueErrorTuple[string] {
		val, err := s.Generate()
		if err != nil {
			return tuples.ValueErrorTuple[string]{
				Value: "",
				Error: err,
			}
		}

		return tuples.ValueErrorTuple[string]{
			Value: val,
			Error: nil,
		}
	})

	if err := tuples.ReduceErrors(result); err != nil {
		return nil, err
	}

	return tuples.GetValues(result), nil

}
