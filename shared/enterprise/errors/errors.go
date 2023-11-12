package errors

import (
	"github.com/life4/genesis/slices"
	"github.com/pkg/errors"
)

var ErrNotAuthroized = errors.New("not authorized")
var ErrCustomerNotFound = errors.New("customer not found")
var ErrOrganizationNotFound = errors.New("organization not found")
var ErrUnityNotFound = errors.New("unity not found")
var ErrLocationNotFound = errors.New("location not found")
var ErrApexNotFound = errors.New("apex not found")
var ErrFailedToVerifyToken = errors.New("failed to verify token")
var ErrTokenNotSupported = errors.New("token not supported")
var ErrNoVerifiablePieceFoundInClaim = errors.New("no verifiable piece found in claim")
var ErrNoMoreInteractionAllowed = errors.New("you have exceeded your free tier limit, please upgrade to premium tier to have full access to all features.")

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
