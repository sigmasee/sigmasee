package openapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sigmasee/sigmasee/customer/customer-api/mappers"
	openapiv1 "github.com/sigmasee/sigmasee/customer/customer-api/openapi/v1"
	"github.com/sigmasee/sigmasee/customer/customer-api/services"
	enterprisecontext "github.com/sigmasee/sigmasee/shared/enterprise/context"
	"go.uber.org/zap"
)

type OpenApiCustomerV1 interface {
	GetHttpHandler() http.Handler
}

type openApiCustomerV1 struct {
	logger          *zap.SugaredLogger
	contextHelper   enterprisecontext.ContextHelper
	customerService services.CustomerService
	outboxService   services.OutboxService
	mapper          mappers.Mapper
}

func NewOpenApiCustomerV1(
	logger *zap.SugaredLogger,
	contextHelper enterprisecontext.ContextHelper,
	customerService services.CustomerService,
	outboxService services.OutboxService,
	mapper mappers.Mapper) (OpenApiCustomerV1, error) {
	return &openApiCustomerV1{
		logger:          logger,
		contextHelper:   contextHelper,
		customerService: customerService,
		outboxService:   outboxService,
		mapper:          mapper,
	}, nil
}

func (s *openApiCustomerV1) GetHttpHandler() http.Handler {
	return openapiv1.Handler(s)
}

func (s *openApiCustomerV1) Liveness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiCustomerV1) Readiness(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiCustomerV1) UpsertCustomer(w http.ResponseWriter, r *http.Request, params openapiv1.UpsertCustomerParams) {
	start := time.Now()
	defer func(start time.Time) {
		s.logger.Infof("UpsertCustomer - Execution time: %s", time.Since(start))
	}(start)

	var err error

	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to excecute UpsertCustomer. Error: %v", err)
		}
	}()

	ctx := s.contextHelper.WithCorrelationId(r.Context(), params.XCorrelationId)

	var request openapiv1.CustomerUpsertRequest
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.sendError(w, http.StatusBadRequest, "bad request")

		return
	}

	customerToUpser, err := s.mapper.OpenApiV1CustomerUpsertRequestToCustomer(request)
	if err != nil {
		s.sendError(w, http.StatusInternalServerError, "failed to map customer")

		return
	}

	customer, err := s.customerService.Upsert(ctx, customerToUpser)
	if err != nil {
		s.sendError(w, http.StatusInternalServerError, "failed to upsert customer")

		return
	}

	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(s.mapper.CustomerToOpenApiV1CustomerUpsertResponse(customer)); err != nil {
		s.sendError(w, http.StatusInternalServerError, "failed to convert response to JSON")

		return
	}
}

func (s *openApiCustomerV1) ProcessOutboxChangeFeed(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func(start time.Time) {
		s.logger.Infof("ProcessOutboxChangeFeed - Execution time: %s", time.Since(start))
	}(start)

	var err error

	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to excecute ProcessOutboxChangeFeed. Error: %v", err)
		}
	}()

	if err = s.outboxService.HandleOutboxChangeFeedRequest(r.Context()); err != nil {
		s.sendError(w, http.StatusInternalServerError, "server internal error")

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *openApiCustomerV1) ProcessOutboxChangeFeedWithParams(w http.ResponseWriter, r *http.Request, date string, ndjsonId string) {
	start := time.Now()
	defer func(start time.Time) {
		s.logger.Infof("ProcessOutboxChangeFeedWithParams - Execution time: %s", time.Since(start))
	}(start)

	var err error

	defer func() {
		if err != nil {
			s.logger.Errorf("Failed to excecute ProcessOutboxChangeFeed. Error: %v", err)
		}
	}()

	if err = s.outboxService.HandleOutboxChangeFeedRequest(r.Context()); err != nil {
		s.sendError(w, http.StatusInternalServerError, "server internal error")

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *openApiCustomerV1) sendError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	errResponse := openapiv1.Error{
		Code:    int32(code),
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(errResponse); err != nil {
		s.logger.Errorf("Failed to convert error response to JSON. Error: %v", err)

		return
	}
}
