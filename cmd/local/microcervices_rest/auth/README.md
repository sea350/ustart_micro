# HOW TO RUN

Create a `config.json` file. Then create an auth config object, this will be imported into the run and unmarshalled. Note that the environment variable by the name `USTART_AUTH_PORT` will take priority as the select port the api will run on if it's populated.

## Sample Config object

You can verbatim copy paste this object into your `config.json` but remember to fill out proper credentials depending on your run environment.
Note that the port doesn't have to be 5001 but it is recomended so as to not conflict with other APIs.

```json
{
    "AuthCfg":{
        "StorageConfig":{
            "ElasticConfig":null,
            "SQLConfig":{
                "DriverName":"postgres",
                "Host":"",
                "Port":"5432",
                "DBName":"",
                "Username":"",
                "Password":"",
                "RegistryTable":"registry",
                "LoginTrackingTable":"login"
            }
        }
    },
    "Port":5001
}

```
