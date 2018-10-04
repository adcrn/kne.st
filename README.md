# kne.st
[![Go Report Card](https://goreportcard.com/badge/github.com/adcrn/kne.st)](https://goreportcard.com/report/github.com/adcrn/kne.st)
A web application codebase for a possible subscription service-based version of knest.

## Building
`go build -o knest`

## Testing
`go test -v ./...`

## Usage
Start the server by running `./knest` and then direct a browser to
`localhost:8080`

## Minimally Viable Components
+ Landing page: offers a signup/login option and describes what the app
  does
+ Authentication: only authorized users should be able to use the service
+ Folder upload: allow for folder upload so users don't have to do a photo at a
  time
+ Presentation of processed photos: previews photos that passed the stages
+ Compression and download of processed photos: offers a zip file of finished
  photos

## Current Progress
The backend API is being hashed out in piecemeal fashion through test-driven
development. A database engine and schema still need to be chosen for
the act of storing folders and their photos. The frontend will consume the API; it will probably be Vue.js with Bulma CSS.
