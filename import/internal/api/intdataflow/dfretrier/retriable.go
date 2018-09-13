package dfretrier

import (
	"context"

	"github.com/lestrrat/go-backoff"
	"github.com/maison-inc/spanner-tools/import/internal/api/intdataflow"
	"google.golang.org/api/dataflow/v1b3"
)

// NewRetriable creates a new Retriable client.
func NewRetriable(
	inner intdataflow.Client,
	maxRetries int,
) intdataflow.Client {
	return &retrier{
		createPolicy: func() *backoff.Exponential {
			return backoff.NewExponential(backoff.WithMaxRetries(maxRetries))
		},
		inner: inner,
	}
}

type retrier struct {
	createPolicy func() *backoff.Exponential
	inner        intdataflow.Client
}

func (r *retrier) Import(
	location string,
	serviceAccountEmail string,
	instanceID string,
	databaseID string,
	inputDir string,
) (
	*dataflow.LaunchTemplateResponse,
	error,
) {
	b, cancel := r.createPolicy().Start(context.Background())
	defer cancel()

	for {
		resp, err := r.inner.Import(location, serviceAccountEmail, instanceID, databaseID, inputDir)
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
