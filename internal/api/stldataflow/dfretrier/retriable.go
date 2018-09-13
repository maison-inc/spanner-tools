package dfretrier

import (
	"context"

	"github.com/lestrrat/go-backoff"
	"google.golang.org/api/dataflow/v1b3"
)

// Retriable represents the API to be subject to retry.
type Retriable interface {
	Export(
		location string,
		serviceAccountEmail string,
		instanceID string,
		databaseID string,
		outputDir string,
	) (*dataflow.LaunchTemplateResponse, error)
}

// NewRetriable creates a new Retriable.
func NewRetriable(
	inner Retriable,
	maxRetries int,
) Retriable {
	return &retrier{
		createPolicy: func() *backoff.Exponential {
			return backoff.NewExponential(backoff.WithMaxRetries(maxRetries))
		},
		inner: inner,
	}
}

type retrier struct {
	createPolicy func() *backoff.Exponential
	inner        Retriable
}

func (r *retrier) Export(
	location string,
	serviceAccountEmail string,
	instanceID string,
	databaseID string,
	outputDir string,
) (
	*dataflow.LaunchTemplateResponse,
	error,
) {
	b, cancel := r.createPolicy().Start(context.Background())
	defer cancel()

	for {
		resp, err := r.inner.Export(location, serviceAccountEmail, instanceID, databaseID, outputDir)
		if err == nil {
			return resp, nil
		}

		select {
		case <-b.Done():
			return nil, err
		case <-b.Next():
			return nil, err
		}
	}
}
