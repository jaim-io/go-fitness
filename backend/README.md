## Backend environment variables
`MYSQL_DATABASE` <br /> 
`MYSQL_ROOT_PASSWORD` <br />
`MYSQL_USER` <br />
`MYSQL_PASSWORD` <br />
`MYSQL_HOST` <br />
`MYSQL_PORT` <br />
`MYSQL_PROTOCOL` <br />

## Swagger docs
First install the Swagger CLI called `swag`.
```
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
```
Then run `swag init -d cmd --pd pkg/controller` in `~/jaim-io/backend`.

To see the docs navigate to the host URI + `/api/v1/docs/index.html`.