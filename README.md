# kne.st
[![Go Report Card](https://goreportcard.com/badge/github.com/adcrn/kne.st)](https://goreportcard.com/report/github.com/adcrn/kne.st)

Web application codebase for a possible SaaS offering of [knest](https://github.com/adcrn/knest).

## Backend
### Setup
In order to use Auth0's API endpoint protection, sign up for an account and set
the following flags according to the credentials given:
```
export AUTH0_API_IDENTIFIER=<YOUR_AUTH0_API>
export AUTH0_DOMAIN=<YOUR_AUTH0_TENANT>.auth0.com
```

### Building
Navigate to the `cmd/webknest` directory and run `go build -o knest`.

### Usage
Start the server and then direct a browser to `localhost:8080`.

### Current Progress
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

## Frontend
### Setup
Install Vue.js and its corresponding tools along with Buefy, which provides UI components based on the Bulma CSS library; Bulma is also the library used throughout the project.
```
npm install axios buefy vue vuex vue-router
```

### Usage
Start the server by navigating to the `frontend` folder and running `npm run
serve`.

### Current Progress
- [ ] Authenticaton
- [ ] Folder upload flow
- [ ] Folder view logic
- [ ] Folder view page
- [x] Home page
- [ ] Login logic
- [x] Login page
- [ ] Registartion logic
- [x] Registration page
- [ ] Session management
- [ ] User profile page
- [ ] User profile page logic
- [ ] User reset password logic
