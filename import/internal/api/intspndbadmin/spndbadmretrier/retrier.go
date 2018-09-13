package spndbadmretrier

import (
	"context"

	"github.com/lestrrat/go-backoff"
	"github.com/maison-inc/spanner-tools/import/internal/api/intspndbadmin"
	databasepb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

// NewRetriable creates a new Retriable client.
func NewRetriable(
	inner intspndbadmin.Client,
	maxRetries int,
) intspndbadmin.Client {
	return &retrier{
		createPolicy: func() *backoff.Exponential {
			return backoff.NewExponential(backoff.WithMaxRetries(maxRetries))
		},
		inner: inner,
	}
}

type retrier struct {
	createPolicy func() *backoff.Exponential
	inner        intspndbadmin.Client
}

func (r *retrier) Create(
	ctx context.Context,
	instanceID string,
	databaseID string,
) (
	*databasepb.Database,
	error,
) {
	b, cancel := r.createPolicy().Start(context.Background())
	defer cancel()

	for {
		resp, err := r.inner.Create(ctx, instanceID, databaseID)
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
