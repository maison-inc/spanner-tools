package stlspndbadmin

import (
	"context"

	adminpb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
	databasepb "google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

// Create creates a database.
func (c *Client) Create(
	ctx context.Context,
	instanceID string,
	databaseID string,
) (*databasepb.Database, error) {
	op, err := c.admin.CreateDatabase(ctx, &adminpb.CreateDatabaseRequest{
		Parent:          c.parent(instanceID),
		CreateStatement: "CREATE DATABASE `" + databaseID + "`",
	})
	if err != nil {
		return nil, err
	}
	return op.Wait(ctx)
}
