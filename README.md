# Golang Packager

Builds a project, incluing all needed files to run the final application.


## How the go-packager works
- Reads all settings from the *go-packager.json* file
- Copies all folders and files to the destination
- Builds the application, according to the OS
- Compress the resulting folder

## Sample of pakager.json


```json
{
    "build_name": "first-build",
    "build_target": "C:\\",  
    "system": "windows",
    "packaged_items": [
        "static",
        "templates",
        "logs",
        ".env"
    ]
}
```



## Example

```sh
go-packager run
```