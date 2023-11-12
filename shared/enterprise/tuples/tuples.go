package tuples

import (
	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/pkg/errors"
)

type ValueErrorTuple[T any] struct {
	Value T
	Error error
}

type TwoValuesTuple[T1 any, T2 any] struct {
	Value1 T1
	Value2 T2
}

func ReduceErrors[T any](src []ValueErrorTuple[T]) error {
	return slices.Reduce(src, nil, func(item ValueErrorTuple[T], aggregatedErr error) error {
		if item.Error == nil {
			return aggregatedErr
		}

		if aggregatedErr == nil {
			return item.Error
		}

		return errors.Wrap(aggregatedErr, item.Error.Error())
	})
}

func ReduceMapErrors[K comparable, T any](src map[K]ValueErrorTuple[T]) error {
	var err error

	for _, item := range src {
		if item.Error != nil {
			if err == nil {
				err = item.Error
			} else {
				err = errors.Wrap(err, item.Error.Error())
			}
		}
	}

	return err
}

func GetValues[T any](src []ValueErrorTuple[T]) []T {
	return slices.Map(src, func(item ValueErrorTuple[T]) T {
		return item.Value
	})
}

func GetMapValues[K comparable, T any](src map[K]ValueErrorTuple[T]) map[K]T {
	return maps.MapValues(src, func(item ValueErrorTuple[T]) T {
		return item.Value
	})
}
