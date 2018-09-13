package intspndbadmin

import (
	databasepb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	"context"
)

// Client represents a spanner admin database client for create.
type Client interface {
	Create(
		ctx context.Context,
		instanceID string,
		databaseID string,
	) (*databasepb.Database, error)
}
