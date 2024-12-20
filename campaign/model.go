package campaign

import (
	"github.com/google/uuid"
	"github.com/jonesrussell/mp-emailer/shared"
	"github.com/jonesrussell/mp-emailer/user"
)

// Campaign represents an email campaign.
type Campaign struct {
	shared.BaseModel
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text;not null" json:"description"`
	Template    string    `gorm:"type:text;not null" json:"template"`
	OwnerID     uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
	Owner       user.User `gorm:"foreignKey:OwnerID" json:"-"`
	Tokens      []string  `gorm:"-" json:"tokens"`
}

// Representative represents a government representative.
type Representative struct {
	Name              string   `json:"name"`
	DistrictName      string   `json:"district_name"`
	ElectedOffice     string   `json:"elected_office"`
	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	Party             string   `json:"party_name"`
	Email             string   `json:"email"`
	URL               string   `json:"url"`
	PersonalURL       string   `json:"personal_url"`
	PhotoURL          string   `json:"photo_url"`
	Gender            string   `json:"gender"`
	Offices           []Office `json:"offices"`
	Extra             Extra    `json:"extra"`
	RepresentativeSet string   `json:"representative_set_name"`
}

// Office represents an office held by a representative.
type Office struct {
	Fax    string `json:"fax"`
	Tel    string `json:"tel"`
	Type   string `json:"type"`
	Postal string `json:"postal"`
}

// Extra contains additional information about a representative.
type Extra struct {
	Roles              []string `json:"roles"`
	PreferredLanguages []string `json:"preferred_languages"`
}

// APIResponse represents a response from the API containing representatives.
type APIResponse struct {
	RepresentativesCentroid []Representative `json:"representatives_centroid"`
}
