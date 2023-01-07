package http

import (
	"context"
	"encoding/json"
	"fmt"
	"movistar/metadata/pkg/model"
	"movistar/movie/internal/gateway"
	"net/http"
)

// Gateway defines a movie metadata HTTP gateway.
type Gateway struct {
	addr string
}

// New creates a new HTTP gateway for a movie metadata service.
func New(addr string) *Gateway {
	return &Gateway{
		addr: addr,
	}
}

// Get gets a movie metadata by the movie id.
func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	request, err := http.NewRequest(http.MethodGet, g.addr+"/metadata", nil)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)
	values := request.URL.Query()
	values.Add("id", id)
	request.URL.RawQuery = values.Encode()
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusNotFound {
		return nil, gateway.ErrNotFound
	} else if response.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non-2xx response: %v", response)
	}
	var v *model.Metadata
	if err := json.NewDecoder(response.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}
