# guestbook-api

A dummy "guestbook" api to learn GoLang and Hexagonal architecture (and eventually all things backend). Currently only has two routes:

1. `GET - /v1/comments`
2. `POST - /v1/comments`

## Running Locally

...I committed the `local.yml` file. Still trying to understand settings... But this command  still probably won't work actually. Unless there's already an existing Postgres server on your local machine haha and you change the Postgres config. Whoops. 

```bash
GO_ENV=local go run cmd/server.go
```

## TODO

- Add validation to who gets to add a comment
- Add validation to see if the request body has the right stuff
- Docker
- Start Postgres server at boot