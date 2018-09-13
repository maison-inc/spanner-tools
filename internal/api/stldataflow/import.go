package stldataflow

import (
	"fmt"

	"google.golang.org/api/dataflow/v1b3"
)

// Import starts a job to import gcs avro to a spanner database.
// See https://cloud.google.com/dataflow/docs/templates/provided-templates#gcs_avro_to_cloud_spanner
func (c *Client) Import(
	location string,
	serviceAccountEmail string,
	instanceID string,
	databaseID string,
	inputDir string,
) (
	*dataflow.LaunchTemplateResponse,
	error,
) {
	templateSrv := dataflow.NewProjectsLocationsTemplatesService(c.service)
	call := templateSrv.Launch(c.projectID, location, &dataflow.LaunchTemplateParameters{
		Environment: importEnvironment(serviceAccountEmail),
		JobName:     importJobName(instanceID, databaseID),
		Parameters: map[string]string{
			"instanceId": instanceID,
			"databaseId": databaseID,
			"inputDir":   inputDir,
		},
	})
	call = call.GcsPath(`gs://dataflow-templates/2018-08-30-00/GCS_Avro_to_Cloud_Spanner`)
	return call.Do()
}

func importJobName(instanceID, databaseID string) string {
	return fmt.Sprintf(
		`cloud-spanner-import-%s-%s`,
		instanceID,
		databaseID,
	)
}

func importEnvironment(serviceAccount string) *dataflow.RuntimeEnvironment {
	return &dataflow.RuntimeEnvironment{
		ServiceAccountEmail: serviceAccount,
	}
}
