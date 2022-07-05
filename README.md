# oauth-package-api-golang
To help you integrate with Dock's Caradhras APIs, we developed a series of code samples considering best practices and the best way to manage the data returned by the endpoints. So, feel free to use this code as you see fit.
## Description
This code sample demonstrates how to implement OAuth 2.0 authorization using caching so that no unnecessary extra calls are made.
To make this possible, we use the information of the field returned with the expiration time. Every time we use a token, we check if the expiration time has passed. We only perform a new authentication if necessary.
## Getting Started
### Prerequisites
* Some programming experience. The code here is simple, but it helps if you know about functions.
* A tool to edit your code. You can use any text editor you have. Most text editors have good support for Go. The most popular are VSCode (free), GoLand (paid), and Vim (free).
* A command terminal. Go works well using any terminal on Linux and Mac. On Windows, use PowerShell or cmd.
* Basic git knowledge to install the sample code.
### Dependencies
* Install Go - Click <a href="https://go.dev/doc/install">here</a> for the steps to download and install.
### Installing
To install the sample, clone the following git repository:
```
$ git clone https://github.com/dock-samples/oauth-package-api-golang
```
### Configuring
Before you use the sample, you must configure the following global variables:
```go
var (
    username = "username"
    password = "password"
    url      = authorization.Hml_url
)
```
### Executing program
To use the sample, run the following command from the root folder of the git project:
```
$ go run cmd/main.go
```
## Authors
- Experiences team

## License
[![License](https://img.shields.io/badge/License-Apache_2.0-yellowgreen.svg)](https://opensource.org/licenses/Apache-2.0)  