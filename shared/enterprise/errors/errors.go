package errors

import (
	"github.com/life4/genesis/slices"
	"github.com/pkg/errors"
)

var ErrNotAuthroized = errors.New("not authorized")
var ErrCustomerNotFound = errors.New("customer not found")
var ErrFailedToVerifyToken = errors.New("failed to verify token")
var ErrTokenNotSupported = errors.New("token not supported")
var ErrNoVerifiablePieceFoundInClaim = errors.New("no verifiable piece found in claim")

func ReduceErrors(src []error) error {
	return slices.Reduce(src, nil, func(item error, aggregatedErr error) error {
		if item == nil {
			return aggregatedErr
		}

		if aggregatedErr == nil {
			return item
		}

		return errors.Wrap(aggregatedErr, item.Error())
	})
}
