package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kolide/fleet/server/kolide"
)

type statusResponse struct {
	Err error `json:"error,omitempty"`
}

func (m statusResponse) error() error { return m.Err }

func makeStatusLiveQueryEndpoint(svc kolide.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		var resp statusResponse
		if err := svc.StatusLiveQuery(ctx); err != nil {
			resp.Err = err
		}
		return resp, nil
	}
}

func makeStatusResultStoreEndpoint(svc kolide.Service) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		var resp statusResponse
		if err := svc.StatusResultStore(ctx); err != nil {
			resp.Err = err
		}
		return resp, nil
	}
}
