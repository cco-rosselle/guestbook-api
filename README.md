# guestbook-api

A super basic/naive api that adds comments to a "guestbook" to ultimately learn GoLang and Hexagonal architecture (and eventually all things backend). Currently only has 3 routes:

1. `GET - /v1/comments`
2. `POST - /v1/comments`
3. `DELETE - /v1/comments/:commentid`

## Running Locally

To run the repo locally, you can create a ```local.yml``` file with the following settings (note: more relevant settings soon):

```yml
logging:
  level: debug
```

And then you can run this:

```bash
GO_ENV=local go run cmd/server.go
```

## TODO

- Add validation to who gets to add a comment
- Add validation to see if the request body has the right stuff
- Docker
- Start Postgres server at boot
- Scrap the db schema from Postgres and create one from code 
- Unit tests
- Find one and delete