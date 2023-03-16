# marketplace-starter-go

This repo is an example of how to build a [QuickNode Marketplace](https://quicknode.com/marketplace) add-on using Go and PostgreSQL


## Getting Started


First, create a postgresql databse called `marketplace-starter-go`:

```bash
createdb marketplace-starter-go
```

Then, copy the `.env.example` to `.env` and update the `DB_URL` to one that matches your local postgresql DB.

```bash
cp .env.example .env
```

Now you need to run migrations:

```bash
go run migrate/migrate.go
```

Then run CompileDaemon which will monitor changes to your file as you develop and rebuild code as needed:

```bash
CompileDaemon -command="./marketplace-starter-go"
```

Then you can start making HTTP requests using curl or Postman:

```
POST http://localhost:3010/provision
PUT http://localhost:3010/update
DELETE http://localhost:3010/deactivate_endpoint
DELETE http://localhost:3010/deprovision
```

See the [marketplace guide to provisioning](https://www.quicknode.com/guides/marketplace/how-provisioning-works-for-marketplace-partners) for more info.

## How this project was created (for the record)

```bash
mkdir marketplace-starter-go
cd marketplace-starter-go
go mod init
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
go get github.com/joho/godotenv
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

## Testing with qn-marketplace-cli

You can test using the [qn-marketplace-cli](https://github.com/quiknode-labs/qn-marketplace-cli):

### Testing Provisioning:

```sh
./qn-marketplace-cli pudd --base-url http://localhost:3010 --basic-auth dXNlcm5hbWU6cGFzc3dvcmQ=
```


### Testing RPC:

```sh
./qn-marketplace-cli rpc --url http://localhost:3010/provision --rpc-url http://localhost:3010/rpc --rpc-method qn_test --rpc-params "[\"abc\"]" --basic-auth dXNlcm5hbWU6cGFzc3dvcmQ=
```

### Testing Healthcheck:

```sh
./qn-marketplace-cli healthcheck --url http://localhost:3010/healthcheck
```

### Testing Single Sign On (SSO):

Below, make sure that the `jwt-secret` matches `QN_SSO_SECRET` in `.env` file.

```
./qn-marketplace-cli sso --url http://localhost:3010/provision  --basic-auth dXNlcm5hbWU6cGFzc3dvcmQ= --jwt-secret jwt-secret --email jon@example.com --name jon --org QuickNode
```

## TODO

- Make sure it works if JSON RPC request uses an integer in ID


## LICENSE

MIT