package main

import (
	"flag"
	"log"
	"github.com/maison-inc/spanner-tools/internal/api/stldataflow"
	"github.com/maison-inc/spanner-tools/export/internal/api/intdataflow"
	"github.com/maison-inc/spanner-tools/internal/api/stldataflow/dfretrier"
)

var (
	projectID = flag.String("project_id", "", "your project ID")
	instanceID = flag.String("instance_id", "", "your Cloud Spanner instance ID to read")
	databaseID = flag.String("database_id", "", "your Cloud Spanner database ID to read")
	outputDir = flag.String("output_dir", "", "Cloud Storage path that the Avro files should be exported to")
	location = flag.String("location", "", "the region where you want the Cloud Dataflow job to run (such as us-central1)")
	serviceAccountEmail = flag.String("service_account_email", "", "Identity to run virtual machines as. Defaults to the default account")

	maxRetries = flag.Int("max_retries", 3, "retry count to export")
)

func export(df intdataflow.Client) error {
	resp, err := df.Export(
		*location,
		*serviceAccountEmail,
		*instanceID,
		*databaseID,
		*outputDir,
	)
	if err != nil {
		return err
	}

	marshalled, err := resp.MarshalJSON()
	if err != nil {
		return err
	}

	log.Println(string(marshalled))
	return nil
}

func run() error {
	df, err := stldataflow.NewClient(*projectID)
	if err != nil {
		return err
	}
	return export(
		dfretrier.NewRetriable(df, *maxRetries-1),
	)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
