# Go request compose

[![Test](https://github.com/lambovg/go-request-compose/actions/workflows/test.yml/badge.svg)](https://github.com/lambovg/go-request-compose/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/lambovg/go-request-compose)](https://goreportcard.com/report/github.com/lambovg/go-request-compose)

* Async by default
* Future/promise support
* No 3rd party dependencies 
* No State

## Description
Build request object with parameters, headers, url and endpoints. 

* ~~Current state is GET requests.~~
* ~~Current state is POST requests.~~
* ~~Current state is PUT requests.~~
* ~~Current state is DELETE requests.~~
* ~~Current state is HEAD requests.~~
* Check payload options for POST and PUT

## TODO

* Add body in POST,PUT
* Re-work examples

## Installation
Clone this repo.

## Usage

```go
// Async request
import "cr "github.com/lambovg/go-request-compose/pkg/request"

cr.Params{Url: "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"}.Get()
```

```go
// Future/promise request
import "cr "github.com/lambovg/go-request-compose/pkg/request"

future := cr.Params{Url: "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"}.Get()
future()
```

```go
// extend client to call multiple endpoints
import "cr "github.com/lambovg/go-request-compose/pkg/request"

var client = new(cr.Params)
client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

client.Get()

client.Url = "https://d2kgi8nio2h9bn.cloudfront.net/ping.json"
client.Get()
```

```go
// create client and extend request with different, paths , url
var client = new(request.Params)
client.Hostname = "d2kgi8nio2h9bn.cloudfront.net"
client.Protocol = "https"
client.Path = "/hello-world.json"

// GET request to https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json
client.Get()

// GET request to https://d2kgi8nio2h9bn.cloudfront.net/ping.json
client.Path = "/ping.json"
client.Get()


// GET request to http://localhost:8080/zen, overrides the whole url
client.Url = "http://localhost:8080/zen"
client.Get()

// GET request to https://d2kgi8nio2h9bn.cloudfront.net/ping.json
client.Url = ""
client.Get()
```

```go
// attach request headers
// start test server with go run tools/server/server.go 
import "cr "github.com/lambovg/go-request-compose/pkg/request"

setHeaders := map[string][]string{
    "Accept":          {"application/json"},
    "Accept-Language": {"en-us"},
    "Cache-Control":   {"must-revalidate"},
}

future := cr.Params{
    Url:     "http://localhost:8080/zen",
    Headers: cr.Headers{Set: setHeaders}}.Get()
future()
```

```go
// Request with net/http client
import (
    cr "github.com/lambovg/go-request-compose/pkg/request
    "net/http"
)

client := http.Client{Timeout: 10 * time.Second}
params := cr.Params{Url: "http://localhost:8080/timeout"}

future := cr.HttpClient{Client: client}.Get(params)
future()
```

Detail usage is described in ```/examples``` folder.

## Contributing
Contributions are welcomed

## Roadmap

* Post request
* Put requests
* Delete requests
* Options requests
* Head requests
* Connect requests

## License
For open source projects, say how it is licensed.
