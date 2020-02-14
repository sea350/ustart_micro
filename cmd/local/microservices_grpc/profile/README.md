# HOW TO RUN

Create a `config.json` file. Then create an profile config object, this will be imported into the run and unmarshalled. Note that the environment variable by the name `USTART_PROF_PORT` will take priority as the select port the api will run on if it's populated.

## Sample Config object

You can verbatim copy paste this object into your `config.json` but remember to fill out proper credentials depending on your run environment.
Note that the port doesn't have to be 5104 but it is recomended so as to not conflict with other APIs.

```json
{
    "ProfileCfg":{
        "StorageConfig":{
            "ElasticConfig":{
            "ElasticAddr":"http://localhost:9200",
            "EIndex":"profile",
                "EType":"test-profile_data"
            }
        },
        "DefaultAvatar":"https://ustart-default.s3.amazonaws.com/Defult_Profile_Page_Logo.png",
        "DefaultBanner":"https://ustart-default.s3.amazonaws.com/Defult_Profile_Banner_Logo.png"
    },
    "Port":5004
}
```
