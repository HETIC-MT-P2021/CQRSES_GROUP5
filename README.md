# MT4 Exam - Form

This is a WIP for the MT4 Exam, once finished it should provide the possibility to create and answer forms with multiples choices, boolean choices and slider values.

## Requirements

If you use docker you will only need:

- Docker;
- Docker-Compose;

Refer to [Docker-Setup](#docker-setup) to install with docker.

To run this project, you will also need to install the following dependencies on your system:

- [go](https://golang.org/doc/install)

## How to launch the project

- To run the project
  `docker-compose up --build`

- On another terminal, after the docker is up
  `docker-compose exec go /bin/sh`
  `go run main.go`

  Since you cannot directly register at the moment, you will need to manually insert a new user in the database

- Once the docker is up, you can acces the database by typing
  `docker-compose exec api-databise /bin/sh`
  `psql`

  Then insert a new entry to the user table
  `INSERT INTO "users" (id, username, password, mail) VALUES (1, 'I', 'Need', 'Help');`

  This can also be done by using adminer on port 3000

## Database Config

- The different paramaters of the database are set to default values inside the docker-compose

You can change them by creating a .env and editing the following fields:

```
- DB_USER
- DB_PASSWORD
- DB_HOST
- DB_PORT
- DB_NAME
```

## Contributing

- Your branch should have a name that reflects it's purpose.

- Each commit must follow the Commit Conventions and start with either:

```

- Update
- Feature
- Fix
- Refacto

```

## Docs

You can find the current route progress & the back architecture inside the docs folder
