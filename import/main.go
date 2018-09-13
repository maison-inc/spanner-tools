package main

import (
	"context"
	"flag"
	"log"

	"github.com/maison-inc/spanner-tools/import/internal/api/intdataflow"
	"github.com/maison-inc/spanner-tools/import/internal/api/intdataflow/dfretrier"
	"github.com/maison-inc/spanner-tools/import/internal/api/intspndbadmin"
	"github.com/maison-inc/spanner-tools/import/internal/api/intspndbadmin/spndbadmretrier"
	"github.com/maison-inc/spanner-tools/internal/api/stldataflow"
	"github.com/maison-inc/spanner-tools/internal/api/stlspndbadmin"
)

var (
	projectID           = flag.String("project_id", "", "your project ID")
	instanceID          = flag.String("instance_id", "", "your Cloud Spanner instance ID to write")
	databaseID          = flag.String("database_id", "", "your Cloud Spanner database ID to write")
	inputDir            = flag.String("input_dir", "", "Cloud Storage path that the Avro files should be imported from")
	location            = flag.String("location", "", "the region where you want the Cloud Dataflow job to run (such as us-central1)")
	serviceAccountEmail = flag.String("service_account_email", "", "identity to run virtual machines as. Defaults to the default account")

	skipCreate = flag.Bool("skip_create_database", false, "skip to create a database before import")
	maxRetries = flag.Int("max_retries", 3, "retry count to export")
)

func createDatabase(spn intspndbadmin.Client) error {
	resp, err := spn.Create(
		context.Background(),
		*instanceID,
		*databaseID,
	)
	if err != nil {
		return err
	}

	log.Printf("database creation done: %s", resp)
	return nil
}

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

	log.Printf("import done: %s", string(marshalled))
	return nil
}

func run() (rerr error) {
	if !*skipCreate {
		spn, err := stlspndbadmin.NewClient(*projectID)
		if err != nil {
			return err
		}
		defer func() {
			closeErr := spn.Close()
			if closeErr == nil {
				return
			}
			if rerr == nil {
				rerr = closeErr
			}
			// only output the error, then ignore it.
			log.Println(closeErr)
		}()
		err = createDatabase(
			spndbadmretrier.NewRetriable(spn, *maxRetries-1),
		)
		if err != nil {
			return err
		}
	}

	df, err := stldataflow.NewClient(*projectID)
	if err != nil {
		return err
	}
	return doImport(
		dfretrier.NewRetriable(df, *maxRetries-1),
	)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
