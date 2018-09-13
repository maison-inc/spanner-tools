package stldataflow

import (
	"fmt"

	"google.golang.org/api/dataflow/v1b3"
)

// Export starts a job to export a spanner database to gcs avro.
// See https://cloud.google.com/dataflow/docs/templates/provided-templates#cloud_spanner_to_gcs_avro
func (c *Client) Export(
	location string,
	serviceAccountEmail string,
	instanceID string,
	databaseID string,
	outputDir string,
) (
	*dataflow.LaunchTemplateResponse,
	error,
) {
	templateSrv := dataflow.NewProjectsLocationsTemplatesService(c.service)
	call := templateSrv.Launch(c.projectID, location, &dataflow.LaunchTemplateParameters{
		Environment: exportEnvironment(serviceAccountEmail),
		JobName:     exportJobName(instanceID, databaseID),
		Parameters: map[string]string{
			"instanceId": instanceID,
			"databaseId": databaseID,
			"outputDir":  outputDir,
		},
	})
	call = call.GcsPath(`gs://dataflow-templates/2018-08-30-00/Cloud_Spanner_to_GCS_Avro`)
	return call.Do()
}

func exportJobName(instanceID, databaseID string) string {
	return fmt.Sprintf(
		`cloud-spanner-export-%s-%s`,
		instanceID,
		databaseID,
	)
}

func exportEnvironment(serviceAccount string) *dataflow.RuntimeEnvironment {
	return &dataflow.RuntimeEnvironment{
		ServiceAccountEmail: serviceAccount,
	}
}
