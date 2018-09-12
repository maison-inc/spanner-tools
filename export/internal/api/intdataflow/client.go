package intdataflow

import dataflow "google.golang.org/api/dataflow/v1b3"

// Client represents a dataflow client for export.
type Client interface {
	Export(
		location string,
		serviceAccountEmail string,
		instanceID string,
		databaseID string,
		outputDir string,
	) (*dataflow.LaunchTemplateResponse, error)
}
