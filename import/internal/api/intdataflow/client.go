package intdataflow

import dataflow "google.golang.org/api/dataflow/v1b3"

// Client represents a dataflow client for import.
type Client interface {
	Import(
		location string,
		serviceAccountEmail string,
		instanceID string,
		databaseID string,
		inputDir string,
	) (*dataflow.LaunchTemplateResponse, error)
}
