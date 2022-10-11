## Backend environment variables
The following environment variables should be set before running the server.<br />
`POSTGRES_DB` <br /> 
`POSTGRES_USER` <br />
`POSTGRES_PASSWORD` <br />
`POSTGRES_HOST` <br />
`POSTGRES_PORT` <br />
For `/cmd/local/` a `.env.local` file should be created containing the environment variables above. The `Dockerfile` used in the `docker-compose.yml` uses the `/cmd/docker/` directory. For `docker-compose.yml` a `.env` file containing the environment variables above should be created. <br />

## Swagger docs
First install the Swagger CLI called `swag`.
```
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```
Then run `swag init -d cmd/docker/ --pd pkg/controllers/` in `~/jaim-io/backend`.

To see the docs navigate to the host URI + `/api/v1/docs/index.html`.