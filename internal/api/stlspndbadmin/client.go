package stlspndbadmin

import (
	spanner "cloud.google.com/go/spanner/admin/database/apiv1"
	"context"
	"fmt"
)

// Client represents a spanner instance admin client.
type Client struct {
	admin     *spanner.DatabaseAdminClient
	projectID string
}

// NewClient creates a new client.
func NewClient(projectID string) (*Client, error) {
	admin, err := spanner.NewDatabaseAdminClient(context.Background())
	if err != nil {
		return nil, err
	}
	return &Client{
		admin:     admin,
		projectID: projectID,
	}, nil
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.admin.Close()
}

func (c *Client) parent(instanceID string) string {
	return fmt.Sprintf(
		`projects/%s/instances/%s`,
		c.projectID,
		instanceID,
	)
}
