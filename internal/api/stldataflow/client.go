package stldataflow

import (
	"context"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/dataflow/v1b3"
)

// Client represents a dataflow client.
type Client struct {
	service *dataflow.Service
	projectID string
}

// NewClient creates a new client.
func NewClient(projectID string) (*Client, error) {
	ctx := context.Background()
	oauthClient, err := google.DefaultClient(ctx, dataflow.CloudPlatformScope)
	if err != nil {
		return nil, err
	}

	service, err := dataflow.New(oauthClient)
	if err != nil {
		return nil, err
	}
	return &Client{
		service: service,
		projectID: projectID,
	}, nil
}

