package mappers

import (
	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/customer/shared/models"
	customereventv1 "github.com/sigmasee/sigmasee/shared/clients/events/sigmasee/customer/v1"
	"github.com/sigmasee/sigmasee/shared/enterprise/converters"
)

type Mapper interface {
	CustomerEventToCustomer(src *customereventv1.Event) (*models.Customer, error)
}

type mapper struct {
}

func NewMapper() (Mapper, error) {
	return &mapper{}, nil
}

func (s *mapper) CustomerEventToCustomer(src *customereventv1.Event) (*models.Customer, error) {
	deletedAt, err := converters.PointerProtoTimestampToPointerTime(src.Data.AfterState.DeletedAt)
	if err != nil {
		return nil, err
	}

	eventRaisedAt, err := converters.PointerProtoTimestampToTime(src.Metadata.Time)
	if err != nil {
		return nil, err
	}

	customer := src.Data.AfterState

	return &models.Customer{
		ID:         customer.Id,
		DeletedAt:  deletedAt,
		ModifiedAt: eventRaisedAt,
		Identities: slices.Map(customer.Identities, func(item *customereventv1.Identity) *models.Identity {
			return &models.Identity{
				ID:            item.Id,
				Email:         converters.StrToPointerStr(item.Email),
				EmailVerified: converters.BoolToPointerBool(item.EmailVerified),
			}
		}),
		Designation: converters.StrToPointerStr(customer.Designation),
		Title:       converters.StrToPointerStr(customer.Title),
		Name:        converters.StrToPointerStr(customer.Name),
		GivenName:   converters.StrToPointerStr(customer.GivenName),
		MiddleName:  converters.StrToPointerStr(customer.MiddleName),
		FamilyName:  converters.StrToPointerStr(customer.FamilyName),
		PhotoURL:    converters.StrToPointerStr(customer.PhotoUrl),
		PhotoURL24:  converters.StrToPointerStr(customer.PhotoUrl24),
		PhotoURL32:  converters.StrToPointerStr(customer.PhotoUrl32),
		PhotoURL48:  converters.StrToPointerStr(customer.PhotoUrl48),
		PhotoURL72:  converters.StrToPointerStr(customer.PhotoUrl72),
		PhotoURL192: converters.StrToPointerStr(customer.PhotoUrl192),
		PhotoURL512: converters.StrToPointerStr(customer.PhotoUrl512),
		Timezone:    converters.StrToPointerStr(customer.Timezone),
		Locale:      converters.StrToPointerStr(customer.Locale),
	}, nil
}
