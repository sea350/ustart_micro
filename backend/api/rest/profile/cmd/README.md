# HOW TO RUN

Create a `config.json` file. Then create an profile config object, this will be imported into the run and unmarshalled.

## Sample Config object

You can verbatim copy paste this object into your `config.json` but remember to fill out proper credentials depending on your run environment.
Note that the port doesn't have to be 5001 but since this is the root sub service it's recomended to start at 5001 and count up.

```json
{"ProfCfg":{
    "StorageConfig":{
        "ElasticConfig":{
            "ElasticAddr":"localhost:9200",
            "EIndex":"PROFILE",
            "EType":"test-profile_data"
            }
        },
        "DefaultAvatar":"INSERT URL HERE",
        "DefaultBanner":"INSERT URL HERE"
    },
    "Port":5002
}

```
