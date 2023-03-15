# qn-go-add-on

This repo is an example of how to build a [QuickNode Marketplace](https://quicknode.com/marketplace) add-on using Go and PostgreSQL


## Getting Started


First, create a postgresql databse called `qn-go-add-on`:

```bash
createdb qn-go-add-on
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
CompileDaemon -command="./qn-go-add-on"
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
mkdir qn-go-add-on
cd qn-go-add-on
go mod init
go get github.com/githubnemo/CompileDaemon
go install github.com/githubnemo/CompileDaemon
go get github.com/joho/godotenv
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```


## TODO

- Make sure it works if JSON RPC request uses an integer in ID
- Add Basic Auth to provisioning endpoints
- Add Dashboard with SSO using JWT
- Finish RPC to validate against DB before responding
