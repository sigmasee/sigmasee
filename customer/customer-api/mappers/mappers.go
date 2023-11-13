package mappers

import (
	"github.com/life4/genesis/slices"
	graphqlmodels "github.com/sigmasee/sigmasee/customer/customer-api/graphql/v1/models"
	openapiv1 "github.com/sigmasee/sigmasee/customer/customer-api/openapi/v1"
	"github.com/sigmasee/sigmasee/customer/shared/entities"
	"github.com/sigmasee/sigmasee/customer/shared/models"
	"github.com/sigmasee/sigmasee/shared/enterprise/converters"
	"github.com/sigmasee/sigmasee/shared/enterprise/random"
)

type Mapper interface {
	CustomerToGraphQLCustomerDetails(src *models.Customer) *graphqlmodels.CustomerDetails
	OpenApiV1CustomerUpsertRequestToCustomer(src openapiv1.CustomerUpsertRequest) (*models.Customer, error)
	CustomerToOpenApiV1CustomerUpsertResponse(src *models.Customer) openapiv1.CustomerUpsertResponse
}

type mapper struct {
	randomHelper random.RandomHelper
}

func NewMapper(randomHelper random.RandomHelper) (Mapper, error) {
	return &mapper{
		randomHelper: randomHelper,
	}, nil
}

func (s *mapper) CustomerToGraphQLCustomerDetails(src *models.Customer) *graphqlmodels.CustomerDetails {
	// TODO: 20230802 - Morteza: Returning first of everything until we implement the feature to let user decide what the active name, email and photo should be

	identitiesWithValidEmail := slices.Filter(src.Identities, func(item *models.Identity) bool {
		return item.Email != nil
	})

	var email *graphqlmodels.CustomerEmail
	if len(identitiesWithValidEmail) > 0 {
		firstEmail := identitiesWithValidEmail[0]

		email = &graphqlmodels.CustomerEmail{
			ID:       firstEmail.ID,
			Email:    converters.PointerStrToStr(firstEmail.Email),
			Verified: converters.PointerBoolToBool(firstEmail.EmailVerified),
		}
	}

	return &graphqlmodels.CustomerDetails{
		ID:           src.ID,
		CreatedAt:    src.CreatedAt,
		IntercomHash: src.IntercomHash,
		Email:        email,
		Designation:  src.Designation,
		Title:        src.Title,
		Name:         src.Name,
		GivenName:    src.GivenName,
		MiddleName:   src.MiddleName,
		FamilyName:   src.FamilyName,
		PhotoURL:     src.PhotoURL,
		PhotoURL24:   src.PhotoURL24,
		PhotoURL32:   src.PhotoURL32,
		PhotoURL48:   src.PhotoURL48,
		PhotoURL72:   src.PhotoURL72,
		PhotoURL192:  src.PhotoURL192,
		PhotoURL512:  src.PhotoURL512,
		Timezone:     src.Timezone,
		Locale:       src.Locale,
		Settings: &entities.CustomerSetting{
			ID: src.Settings.ID,
		},
	}
}

func (s *mapper) OpenApiV1CustomerUpsertRequestToCustomer(src openapiv1.CustomerUpsertRequest) (*models.Customer, error) {
	var email *string

	if src.Email != nil {
		email = converters.StrToPointerStr(string(*src.Email))
	}

	id, err := s.randomHelper.Generate()
	if err != nil {
		return nil, err
	}

	return &models.Customer{
		ID: id,
		Identities: []*models.Identity{
			{
				ID:            src.VerifiableToken,
				Email:         email,
				EmailVerified: src.EmailVerified,
			},
		},
		Designation: src.Designation,
		Title:       src.Title,
		Name:        src.Name,
		GivenName:   src.GivenName,
		MiddleName:  src.MiddleName,
		FamilyName:  src.FamilyName,
		PhotoURL:    src.PhotoUrl,
		PhotoURL24:  src.PhotoUrl24,
		PhotoURL32:  src.PhotoUrl32,
		PhotoURL48:  src.PhotoUrl48,
		PhotoURL72:  src.PhotoUrl72,
		PhotoURL192: src.PhotoUrl192,
		PhotoURL512: src.PhotoUrl512,
		Timezone:    src.Timezone,
		Locale:      src.Locale,
	}, nil
}

func (s *mapper) CustomerToOpenApiV1CustomerUpsertResponse(src *models.Customer) openapiv1.CustomerUpsertResponse {
	return openapiv1.CustomerUpsertResponse{
		Id: src.ID,
	}
}
