# HOW TO RUN

Create a `config.json` file. Then create an uploader config object, this will be imported into the run and unmarshalled.

## Sample Config object

You can verbatim copy paste this object into your `config.json` but remember to fill out proper credentials depending on your run environment.
Note that the port doesn't have to be 5004 but since this is the root sub service it's recomended to use 5004.

```json
{
    "UploaderCfg":{
        "StorageConfig":{
            "AWSConfig":{
                "Region":"us-east-2",
                "BucketName":"",
                "S3CredID":"",
                "S3CredSecret":"",
                "S3Token":""
            }
        }
    },
    "Port":5004
}
```
