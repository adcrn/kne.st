# kne.st
[![Go Report Card](https://goreportcard.com/badge/github.com/adcrn/kne.st)](https://goreportcard.com/report/github.com/adcrn/kne.st)

Backend for a possible SaaS offering of [knest](https://github.com/adcrn/knest).

## Building
`go build -o knest`

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
[x] Postgres implementation
[ ] Key-value store implementation
[x] User registration
[x] User metadata
[ ] Folder upload
[x] Folder metadata
