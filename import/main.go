package main

import (
	"flag"
	"log"
	"github.com/maison-inc/spanner-tools/internal/api/stldataflow"
	"github.com/maison-inc/spanner-tools/import/internal/api/intdataflow"
)

var (
	projectID = flag.String("project_id", "", "your project ID")
	instanceID = flag.String("instance_id", "", "your Cloud Spanner instance ID to write")
	databaseID = flag.String("database_id", "", "your Cloud Spanner database ID to write")
	inputDir = flag.String("input_dir", "", "Cloud Storage path that the Avro files should be imported from")
	location = flag.String("location", "", "the region where you want the Cloud Dataflow job to run (such as us-central1)")
	serviceAccountEmail = flag.String("service_account_email", "", "Identity to run virtual machines as. Defaults to the default account")
)

func doImport(df intdataflow.Client) error {
	resp, err := df.Import(
		*location,
		*serviceAccountEmail,
		*instanceID,
		*databaseID,
		*inputDir,
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
	return doImport(df)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
