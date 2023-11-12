package services

import (
	"context"

	"github.com/sigmasee/sigmasee/apex/shared/repositories"
	enterprisecontext "github.com/sigmasee/sigmasee/shared/enterprise/context"
	"go.uber.org/zap"
)

type CustomerService interface {
	DoesCustomerExist(ctx context.Context) (bool, error)
}

type customerService struct {
	logger             *zap.SugaredLogger
	contextHelper      enterprisecontext.ContextHelper
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(
	logger *zap.SugaredLogger,
	contextHelper enterprisecontext.ContextHelper,
	customerRepository repositories.CustomerRepository) (CustomerService, error) {
	return &customerService{
		logger:             logger,
		contextHelper:      contextHelper,
		customerRepository: customerRepository,
	}, nil
}

func (s *customerService) DoesCustomerExist(ctx context.Context) (bool, error) {
	verifiableToken := s.contextHelper.GetVerifiableToken(ctx)
	customer, err := s.customerRepository.GetByVerifiableToken(ctx, nil, true, verifiableToken)
	if err != nil {
		return false, err
	}

	return customer != nil, nil
}
