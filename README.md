# token-dash

token-dash is a marketplace add-on that allows you to create a visual dashboard with widgets that monitor ERC-20 tokens.


## Getting Started


First, create a postgresql databse called `token-dash`:

```bash
createdb token-dash
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
CompileDaemon -command="./token-dash"
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
mkdir token-dash
cd token-dash
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
- Store the Params in the database
- Look at the dashboard here for widget ideas: https://dune.com/ilemi/Token-Overview-Metrics
