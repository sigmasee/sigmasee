package openapi

import (
	"net/http"

	openapi_types "github.com/oapi-codegen/runtime/types"
	openapiv1 "github.com/sigmasee/sigmasee/apex/apex-api/openapi/v1"
	"go.uber.org/zap"
)

type OpenApiApexV1 interface {
	GetHttpHandler() http.Handler
}

type openApiApexV1 struct {
	logger *zap.SugaredLogger
}

func NewOpenApiApexV1(logger *zap.SugaredLogger) (OpenApiApexV1, error) {
	return &openApiApexV1{
		logger: logger,
	}, nil
}

func (s *openApiApexV1) GetHttpHandler() http.Handler {
	return openapiv1.Handler(s)
}

func (s *openApiApexV1) GetFmcaTopAssets(w http.ResponseWriter, r *http.Request, params openapiv1.GetFmcaTopAssetsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetFundRisk(w http.ResponseWriter, r *http.Request, code string) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetFundUnitPriceSeries(w http.ResponseWriter, r *http.Request, code string, params openapiv1.GetFundUnitPriceSeriesParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetFundVolatility(w http.ResponseWriter, r *http.Request, code string) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetGains(w http.ResponseWriter, r *http.Request, params openapiv1.GetGainsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInceptionDate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetPagedInvestments(w http.ResponseWriter, r *http.Request, params openapiv1.GetPagedInvestmentsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentsDocuments(w http.ResponseWriter, r *http.Request, params openapiv1.GetInvestmentsDocumentsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetPagedInvestmentTotalValues(w http.ResponseWriter, r *http.Request, params openapiv1.GetPagedInvestmentTotalValuesParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestment(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentBankAccounts(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentDocuments(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentDocumentsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentDocument(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, documentId openapi_types.UUID, fileName string) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetFmcaInvestmentMix(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetFmcaInvestmentMixParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentFmcaTopAssets(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentFmcaTopAssetsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentFunds(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentFundsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentGains(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentGainsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentInvestors(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentNetReturn(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentNetReturnParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentNetReturnSeries(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentNetReturnSeriesParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentPendingTransactions(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentPendingTransactionsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentPrescribedInvestorRate(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentPrescribedInvestorRateParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentsFundsPriceDates(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentStandingInstructions(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentTargetFundAllocation(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentTotalValue(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentTotalValueParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentTotalValueSeries(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentTotalValueSeriesParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetPagedInvestmentTransactions(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetPagedInvestmentTransactionsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentTransactionBalanceBreakdown(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentTransactionBalanceBreakdownParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentTransactionsCsv(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentTransactionsCsvParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestmentTransactionHighlights(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestmentTransactionHighlightsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetPagedInvestors(w http.ResponseWriter, r *http.Request, params openapiv1.GetPagedInvestorsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorsDocuments(w http.ResponseWriter, r *http.Request, params openapiv1.GetInvestorsDocumentsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetPagedInvestorTotalValues(w http.ResponseWriter, r *http.Request, params openapiv1.GetPagedInvestorTotalValuesParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestor(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorAddresses(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorDocuments(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestorDocumentsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorDocument(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, documentId openapi_types.UUID, fileName string) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorEmailAddresses(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorFmcaTopAssets(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestorFmcaTopAssetsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorGains(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestorGainsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorInceptionDate(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetPagedInvestorInvestments(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetPagedInvestorInvestmentsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorNetReturn(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestorNetReturnParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorPendingTransactions(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestorPendingTransactionsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorPhoneNumbers(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorPrescribedInvestorRate(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestorPrescribedInvestorRateParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorStandingInstructions(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorTotalValue(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestorTotalValueParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetPagedInvestorTransactions(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetPagedInvestorTransactionsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorTransactionBalanceBreakdown(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestorTransactionBalanceBreakdownParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetInvestorTransactionHighlights(w http.ResponseWriter, r *http.Request, id openapi_types.UUID, params openapiv1.GetInvestorTransactionHighlightsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetLatestPriceDate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetNetReturnSeries(w http.ResponseWriter, r *http.Request, params openapiv1.GetNetReturnSeriesParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetPendingTransactions(w http.ResponseWriter, r *http.Request, params openapiv1.GetPendingTransactionsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetStandingInstructions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetTotalValue(w http.ResponseWriter, r *http.Request, params openapiv1.GetTotalValueParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetTotalValueSeries(w http.ResponseWriter, r *http.Request, params openapiv1.GetTotalValueSeriesParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetPagedTransactions(w http.ResponseWriter, r *http.Request, params openapiv1.GetPagedTransactionsParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetTransactionBalanceBreakdown(w http.ResponseWriter, r *http.Request, params openapiv1.GetTransactionBalanceBreakdownParams) {
	w.WriteHeader(http.StatusOK)
}

func (s *openApiApexV1) GetTransactionHighlights(w http.ResponseWriter, r *http.Request, params openapiv1.GetTransactionHighlightsParams) {
	w.WriteHeader(http.StatusOK)
}
