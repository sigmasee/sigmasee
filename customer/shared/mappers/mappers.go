package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/customer/shared/entities"
	"github.com/sigmasee/sigmasee/customer/shared/models"
	customereventv1 "github.com/sigmasee/sigmasee/shared/clients/events/sigmasee/customer/v1"
	"github.com/sigmasee/sigmasee/shared/enterprise/converters"
)

type Mapper interface {
	CustomerEntitiesToCustomers(src []*entities.Customer) []*models.Customer
	CustomerToEventCustomer(src *models.Customer) *customereventv1.Customer
	CustomerEntityToCustomer(src *entities.Customer) *models.Customer
	IdentityEntitiesToIdentities(src []*entities.Identity) []*models.Identity
	IdentityEntityToIdentity(src *entities.Identity) *models.Identity
	CustomerSettingsEntityToCustomerSettings(src *entities.CustomerSetting) *models.CustomerSettings
}

type mapper struct {
}

func NewMapper() (Mapper, error) {
	return &mapper{}, nil
}

func (s *mapper) CustomerEntitiesToCustomers(src []*entities.Customer) []*models.Customer {
	return slices.Map(src, func(item *entities.Customer) *models.Customer {
		return s.CustomerEntityToCustomer(item)
	})
}

func (s *mapper) CustomerToEventCustomer(src *models.Customer) *customereventv1.Customer {
	return &customereventv1.Customer{
		Id:        src.ID,
		DeletedAt: converters.PointerTimeToPointerProtoTimestamp(src.DeletedAt),
		Identities: slices.Map(src.Identities, func(item *models.Identity) *customereventv1.Identity {
			return &customereventv1.Identity{
				Id:            item.ID,
				Email:         converters.PointerStrToStr(item.Email),
				EmailVerified: converters.PointerBoolToBool(item.EmailVerified),
			}
		}),
		Designation: converters.PointerStrToStr(src.Designation),
		Title:       converters.PointerStrToStr(src.Title),
		Name:        converters.PointerStrToStr(src.Name),
		GivenName:   converters.PointerStrToStr(src.GivenName),
		MiddleName:  converters.PointerStrToStr(src.MiddleName),
		FamilyName:  converters.PointerStrToStr(src.FamilyName),
		Timezone:    converters.PointerStrToStr(src.Timezone),
		Locale:      converters.PointerStrToStr(src.Locale),
		PhotoUrl:    converters.PointerStrToStr(src.PhotoURL),
		PhotoUrl24:  converters.PointerStrToStr(src.PhotoURL24),
		PhotoUrl32:  converters.PointerStrToStr(src.PhotoURL32),
		PhotoUrl48:  converters.PointerStrToStr(src.PhotoURL48),
		PhotoUrl72:  converters.PointerStrToStr(src.PhotoURL72),
		PhotoUrl192: converters.PointerStrToStr(src.PhotoURL192),
		PhotoUrl512: converters.PointerStrToStr(src.PhotoURL512),
	}
}

func (s *mapper) CustomerEntityToCustomer(src *entities.Customer) *models.Customer {
	if src == nil {
		return nil
	}

	return &models.Customer{
		ID:          src.ID,
		DeletedAt:   converters.TimeToPointerTime(src.DeletedAt),
		ModifiedAt:  src.ModifiedAt,
		CreatedAt:   src.CreatedAt,
		Identities:  s.IdentityEntitiesToIdentities(src.Edges.Identities),
		Designation: converters.StrToPointerStr(src.Designation),
		Title:       converters.StrToPointerStr(src.Title),
		Name:        converters.StrToPointerStr(src.Name),
		GivenName:   converters.StrToPointerStr(src.GivenName),
		MiddleName:  converters.StrToPointerStr(src.MiddleName),
		FamilyName:  converters.StrToPointerStr(src.FamilyName),
		PhotoURL:    converters.StrToPointerStr(src.PhotoURL),
		PhotoURL24:  converters.StrToPointerStr(src.PhotoURL24),
		PhotoURL32:  converters.StrToPointerStr(src.PhotoURL32),
		PhotoURL48:  converters.StrToPointerStr(src.PhotoURL48),
		PhotoURL72:  converters.StrToPointerStr(src.PhotoURL72),
		PhotoURL192: converters.StrToPointerStr(src.PhotoURL192),
		PhotoURL512: converters.StrToPointerStr(src.PhotoURL512),
		Timezone:    converters.StrToPointerStr(src.Timezone),
		Locale:      converters.StrToPointerStr(src.Locale),
		Settings:    s.CustomerSettingsEntityToCustomerSettings(src.Edges.CustomerSettings),
	}
}

func (s *mapper) IdentityEntitiesToIdentities(src []*entities.Identity) []*models.Identity {
	return slices.Map(src, func(item *entities.Identity) *models.Identity {
		return s.IdentityEntityToIdentity(item)
	})
}

func (s *mapper) IdentityEntityToIdentity(src *entities.Identity) *models.Identity {
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

func (s *mapper) CustomerSettingsEntityToCustomerSettings(src *entities.CustomerSetting) *models.CustomerSettings {
	if src == nil {
		return nil
	}

	return &models.CustomerSettings{
		ID:        src.ID,
		DeletedAt: converters.TimeToPointerTime(src.DeletedAt),
	}
}
