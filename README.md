# Golang Packager

Builds a Golang project, incluing all needed files to run the final application.


## How the go-packager works
- Reads all settings from the *go-packager.json* file
- Copies all folders and files to the destination
- Builds the application, according to the OS
- Compress the resulting folder

## Installation
- Download binary: 
    - [Windows](bin/windows/go-packager.exe)
    - [Linux](bin/linux/go-packager)
    - [MacOS](bin/macos/go-packager)
    - [FreeBsd](bin/freebsd/go-packager)
    - [Solaris](bin/solaris/go-packager)
- Add to PATH evironment variable
- Configure *go-packager.json* file in application root
- Run **go-packager** in application root
``
go-packager run
``

## Sample of *go-pakager.json* configuration

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
