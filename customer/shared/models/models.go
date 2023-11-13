package models

import "time"

type IdentityEventTrack struct {
	ID         string
	IdentityID string
	CustomerID string
}

type Customer struct {
	ID           string
	DeletedAt    *time.Time
	ModifiedAt   time.Time
	CreatedAt    time.Time
	IntercomHash *string
	Identities   []*Identity
	Designation  *string
	Title        *string
	Name         *string
	GivenName    *string
	MiddleName   *string
	FamilyName   *string
	PhotoURL     *string
	PhotoURL24   *string
	PhotoURL32   *string
	PhotoURL48   *string
	PhotoURL72   *string
	PhotoURL192  *string
	PhotoURL512  *string
	Timezone     *string
	Locale       *string
	Settings     *CustomerSettings
}

type Identity struct {
	ID            string
	DeletedAt     *time.Time
	Email         *string
	EmailVerified *bool
}

type CustomerSettings struct {
	ID        string
	DeletedAt *time.Time
}
