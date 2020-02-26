# HOW TO RUN

Create a `config.json` file. Then create a session config object, this will be imported into the run and unmarshalled.

## Sample Config object

You can verbatim copy paste this object into your `config.json` but remember to fill out proper credentials depending on your run environment.

```json
{
    "StorageConfig":{
        "SQLConfig":{
            "DriverName":"postgres",
            "Host":"localhost",
            "Port":"5432",
            "DBName":"test",
            "Username":"postgres",
            "Password":"password",
            "SessionTableName":"sess"
        }
    },
    "SessionKey":"session_please"
}
```
