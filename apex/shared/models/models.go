package models

import (
	"time"
)

type Customer struct {
	ID            string
	DeletedAt     *time.Time
	ModifiedAt    time.Time
	EventRaisedAt time.Time
	Identities    []*Identity
	Name          *string
	GivenName     *string
	MiddleName    *string
	FamilyName    *string
	PhotoURL      *string
	PhotoURL24    *string
	PhotoURL32    *string
	PhotoURL48    *string
	PhotoURL72    *string
	PhotoURL192   *string
	PhotoURL512   *string
}

type Identity struct {
	ID            string
	DeletedAt     *time.Time
	Email         *string
	EmailVerified *bool
}
