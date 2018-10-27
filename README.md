# kne.st
[![Go Report Card](https://goreportcard.com/badge/github.com/adcrn/kne.st)](https://goreportcard.com/report/github.com/adcrn/kne.st)

Backend for a possible SaaS offering of [knest](https://github.com/adcrn/knest).

## Setup
In order to use Auth0's API endpoint protection, sign up for an account and set
the following flags according to the credentials given:
```
export AUTH0_API_IDENTIFIER=<YOUR_AUTH0_API>
export AUTH0_DOMAIN=<YOUR_AUTH0_TENANT>.auth0.com
```

## Building
Navigate to the `cmd/webknest` directory and run `go build -o knest`.

## Usage
Start the server and then direct a browser to `localhost:8080`.

## Current Progress
- [ ] CORS middleware for frontend interaction
- [x] Folder metadata
- [x] Folder upload to local storage
- [ ] Key-value store implementation
- [x] Postgres implementation
- [x] Session management - using [Auth0](https://auth0.com)
- [ ] System call to knest desktop application
- [x] User authentication - using [Auth0](https://auth0.com)
- [x] User metadata
- [x] User registration
