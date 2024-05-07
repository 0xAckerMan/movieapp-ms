package model

// RecordID defines a record id. Together with RecordType
// identifies unique recoeds across all types.
type RecordID string

// RecordType defines a record type. Together with RecordID
// identifies unique records across all types
type RecordType string

// Existing record types
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID defines a user id.
type UserID string

// RatingValue defines the value of a rating record
type RatingValue int

// Rating defines an individual rating defined by a user for // some record
type Rating struct {
	RecordID   string      `json:"recordid"`
	RecordType string      `json:"recordtype"`
	UserID     UserID      `json:"userid"`
	Value      RatingValue `json:"value"`
}
