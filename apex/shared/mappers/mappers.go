package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/apex/shared/entities"
	"github.com/sigmasee/sigmasee/apex/shared/models"
	"github.com/sigmasee/sigmasee/shared/enterprise/converters"
)

type Mapper interface {
	CustomerEntitiesToCustomers(src []*entities.ApexCustomer) []*models.Customer
	CustomerEntityToCustomer(src *entities.ApexCustomer) *models.Customer
	CustomerIdentityEntitiesToIdentities(src []*entities.ApexCustomerIdentity) []*models.Identity
	CustomerIdentityEntityToIdentity(src *entities.ApexCustomerIdentity) *models.Identity
}

type mapper struct {
}

func NewMapper() (Mapper, error) {
	return &mapper{}, nil
}

func (s *mapper) CustomerEntitiesToCustomers(src []*entities.ApexCustomer) []*models.Customer {
	return slices.Map(src, func(item *entities.ApexCustomer) *models.Customer {
		return s.CustomerEntityToCustomer(item)
	})
}

func (s *mapper) CustomerEntityToCustomer(src *entities.ApexCustomer) *models.Customer {
	if src == nil {
		return nil
	}

	return &models.Customer{
		ID:            src.ID,
		DeletedAt:     converters.TimeToPointerTime(src.DeletedAt),
		ModifiedAt:    src.ModifiedAt,
		EventRaisedAt: src.EventRaisedAt,
		Identities:    s.CustomerIdentityEntitiesToIdentities(src.Edges.Identities),
		Name:          converters.StrToPointerStr(src.Name),
		GivenName:     converters.StrToPointerStr(src.GivenName),
		MiddleName:    converters.StrToPointerStr(src.MiddleName),
		FamilyName:    converters.StrToPointerStr(src.FamilyName),
		PhotoURL:      converters.StrToPointerStr(src.PhotoURL),
		PhotoURL24:    converters.StrToPointerStr(src.PhotoURL24),
		PhotoURL32:    converters.StrToPointerStr(src.PhotoURL32),
		PhotoURL48:    converters.StrToPointerStr(src.PhotoURL48),
		PhotoURL72:    converters.StrToPointerStr(src.PhotoURL72),
		PhotoURL192:   converters.StrToPointerStr(src.PhotoURL192),
		PhotoURL512:   converters.StrToPointerStr(src.PhotoURL512),
	}
}

func (s *mapper) CustomerIdentityEntitiesToIdentities(src []*entities.ApexCustomerIdentity) []*models.Identity {
	return slices.Map(src, func(item *entities.ApexCustomerIdentity) *models.Identity {
		return s.CustomerIdentityEntityToIdentity(item)
	})
}

func (s *mapper) CustomerIdentityEntityToIdentity(src *entities.ApexCustomerIdentity) *models.Identity {
	if src == nil {
		return nil
	}

	return &models.Identity{
		ID:            src.ID,
		DeletedAt:     converters.TimeToPointerTime(src.DeletedAt),
		Email:         converters.StrToPointerStr(src.Email),
		EmailVerified: converters.BoolToPointerBool(src.EmailVerified),
	}
}
