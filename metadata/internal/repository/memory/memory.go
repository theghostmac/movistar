/*
this package hash a memory.Repository function for ease of reference.
there is a New, Get, and Put function in the package for use in the business logic layer.
Get and Put accepts context and does so as the first argument, which is idiomatic Go code.
the sync.RWMutex structure is used to protect against concurrent writes and reads operations.
*/

package memory

import (
	"context"
	"sync"

	"movistar/metadata/internal/repository"
	"movistar/metadata/pkg/model"
)

// Repository defines a memory movie metadata repository.
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new memory repository.
func New() *Repository {
	return &Repository{
		data: map[string]*model.Metadata{}}
}

// Get retrieves movie metadata by movie id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

// Put adds movie metadata for a given movie id.
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = metadata
	return nil
}
