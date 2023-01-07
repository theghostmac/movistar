// Package model contains model package of the metadata service having the Metadata structure for reuse.
package model

import "movistar/metadata/pkg/model"

// MovieDetails include movie metadata and it's aggregated rating.
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata"`
}
