package http

import (
	"context"
	"encoding/json"
	"fmt"
	"movistar/movie/internal/gateway"
	"movistar/rating/pkg/model"
	"net/http"
)

// Gateway defines an HTTP gateway for the rating service.
type Gateway struct {
	addr string
}

// New creates a new HTTP gateway for the rating service.
func New(addr string) *Gateway {
	return &Gateway{addr: addr}
}

// GetAggregatedRating returns the aggregated rating for a record or ErrNotFound if there are no ratings for it.
func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	request, err := http.NewRequest(http.MethodGet, g.addr+"/rating", nil)
	if err != nil {
		return 0, err
	}
	request = request.WithContext(ctx)
	values := request.URL.Query()
	values.Add("type", fmt.Sprintf("%v", recordType))
	request.URL.RawQuery = values.Encode()
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusNotFound {
		return 0, gateway.ErrNotFound
	} else if response.StatusCode/100 != 2 {
		return 0, fmt.Errorf("non-2xx response: %v", response)
	}
	var v float64
	if err := json.NewDecoder(response.Body).Decode(&v); err != nil {
		return 0, err
	}
	return v, nil
}
