// this package defines the model for the rating service which other services will use.

package model

// RecordID defines a record id. It identifies unique records across all types, with RecordType.
type RecordID string

// RecordType defines a record type. It identifies unique records across all types, with RecordID.
type RecordType string

// RecordTypeMovie is an existing record types.
const (
	RecordTypeMovie = RecordType("movie")
)

// UserID defines a user id.
type UserID string

// RatingValue defines a value of a rating record.
type RatingValue int

// Rating defines and individual rating created by a user for some record.
type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}
