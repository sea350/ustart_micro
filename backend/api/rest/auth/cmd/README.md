# HOW TO RUN

Create a `config.go` file. Then create an auth config pointer within it to be used within your `run.go` script

## Sample Config object

You can literally copy paste this into your `config.go` but remember to fill out proper credentials depending on your run environment

```go
package main

import (
  "github.com/sea350/ustart_mono/backend/api/rest/auth"
  sAuth "github.com/sea350/ustart_mono/backend/auth"
  "github.com/sea350/ustart_mono/backend/auth/storage"
  "github.com/sea350/ustart_mono/backend/auth/storage/sql"
)

var config = &auth.Config{
  AuthCfg: &sAuth.Config{
    StorageConfig: &storage.Config{
      SQLConfig: &sqlstore.Config{
        DriverName:         "postgres",
        Host:               "localhost",
        Port:               "5432",
        DBName:             "localhost",
        Username:           "postgres",
        Password:           "postgres",
        RegistryTable:      "auth",
        LoginTrackingTable: "logins",
      },
    },
  },
}

```
