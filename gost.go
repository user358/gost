package gost

import (
	"context"
	"errors"
)

// ErrNoEndpointUrls is returned if the list of urls to speed test was not received.
var ErrNoEndpointUrls = errors.New("no endpoint urls")

// Provider is the interface that wraps the basic Download and Upload methods.
type Provider interface {
	// Download should return the download speed in Mbit/s.
	Download(ctx context.Context) (float64, error)

	// Upload should return the upload speed in Mbit/s.
	Upload(ctx context.Context) (float64, error)
}

// Result uses as result for Measure method.
type Result struct {
	Download float64 // download speed in Mbit/s
	Upload   float64 // upload speed in Mbit/s
}

// Measure returns download and upload speed in Mbit/s.
func Measure(ctx context.Context, provider Provider) (*Result, error) {
	dlSpeed, err := provider.Download(ctx)
	if err != nil {
		return nil, err
	}

	ulSpeed, err := provider.Upload(ctx)
	if err != nil {
		return nil, err
	}

	result := &Result{
		Download: dlSpeed,
		Upload:   ulSpeed,
	}

	return result, nil
}
