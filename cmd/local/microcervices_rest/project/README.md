# HOW TO RUN

Create a `config.json` file. Then create an project config object, this will be imported into the run and unmarshalled. Note that the environment variable by the name `USTART_PROJ_PORT` will take priority as the select port the api will run on if it's populated.

## Sample Config object

You can verbatim copy paste this object into your `config.json` but remember to fill out proper credentials depending on your run environment.
Note that the port doesn't have to be 5005 but it is recomended so as to not conflict with other APIs.

```json
{
    "ProjCfg":{
        "StorageConfig":{
            "ElasticConfig":{
                "ElasticAddr":"http://localhost:9200/",
                "EIndex":"project",
                "EType":"test-project_data"
            }
        },
        "AuthConfig":{
            "StorageConfig":{
                "ElasticConfig":null,
                "SQLConfig":{
                    "DriverName":"postgres",
                    "Host":"localhost",
                    "Port":"5432",
                    "DBName":"",
                    "Username":"postgres",
                    "Password":"",
                    "MemberTableName":"project_members",
                    "RoleTableName":"project_roles",
                    "RequestTableName":"project_requests"
                }
            }
        },
        "DefaultAvatar":"https://ustart-default.s3.amazonaws.com/Defult_Project_Page_Logo.png",
        "DefaultBanner":"https://ustart-default.s3.amazonaws.com/Defult_project_Banner_Logo.png"
    },
    "Port":5005
}

```
