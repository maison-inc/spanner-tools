# spanner-tools

Assorted spanner-related commands.

## Installation

```
go get github.com/maison-inc/spanner-tools/export
go get github.com/maison-inc/spanner-tools/import
```

## Tools

### Export

```
$ go run export/main.go -h
  -database_id string
        your Cloud Spanner database ID to read
  -instance_id string
        your Cloud Spanner instance ID to read
  -location string
        the region where you want the Cloud Dataflow job to run (such as us-central1)
  -max_retries int
        retry count to export (default 3)
  -output_dir string
        Cloud Storage path that the Avro files should be exported to
  -project_id string
        your project ID
  -service_account_email string
        Identity to run virtual machines as. Defaults to the default account
```

### Import

```
$ go run import/main.go -h
  -database_id string
        your Cloud Spanner database ID to write
  -input_dir string
        Cloud Storage path that the Avro files should be imported from
  -instance_id string
        your Cloud Spanner instance ID to write
  -location string
        the region where you want the Cloud Dataflow job to run (such as us-central1)
  -max_retries int
        retry count to export (default 3)
  -project_id string
        your project ID
  -service_account_email string
        identity to run virtual machines as. Defaults to the default account
  -skip_create_database
        skip to create a database before import
```

## Development

### Install dependency

```
_devel/dep ensure -add dependency
```

### Test

```
make test/lint
```
