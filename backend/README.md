## Backend environment variables
`POSTGRES_DATABASE` <br /> 
`POSTGRES_ROOT_PASSWORD` <br />
`POSTGRES_USER` <br />
`POSTGRES_PASSWORD` <br />
`POSTGRES_HOST` <br />
`POSTGRES_PORT` <br />
`POSTGRES_PROTOCOL` <br />

## Swagger docs
First install the Swagger CLI called `swag`.
```
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```
Then run `swag init -d cmd --pd pkg/controller` in `~/jaim-io/backend`.

To see the docs navigate to the host URI + `/api/v1/docs/index.html`.