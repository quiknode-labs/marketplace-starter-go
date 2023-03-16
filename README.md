# Quicknode Marketplace Starter Code - Go

This repo is an example of how to build a [QuickNode Marketplace](https://quicknode.com/marketplace) add-on using Go and PostgreSQL

It implements the 4 provisioning routes that a partner needs to [integrate with Marketplace](https://www.quicknode.com/guides/quicknode-products/marketplace/how-provisioning-works-for-marketplace-partners/), as well as the required Healthcheck route.

It also has support for:

- [RPC methods](https://www.quicknode.com/guides/quicknode-products/marketplace/how-to-create-an-rpc-add-on-for-marketplace/) via a `POST /rpc` route
- [A dashboard view](https://www.quicknode.com/guides/quicknode-products/marketplace/how-sso-works-for-marketplace-partners/) with Single Sign On using JSON Web Tokens (JWT).


## Getting Started

To install and run the application locally:

1. Clone this repo.
2. Create a postgresql databse called `marketplace-starter-go`:

```bash
createdb marketplace-starter-go
```

3. Copy the `.env.example` to `.env` and update the `DB_URL` to one that matches your local postgresql DB.

```bash
cp .env.example .env
```

4. Run migrations:

```bash
go run migrate/migrate.go
```

5. Build the code:

```bash
go build
```

6. Start the web server by running the executable:

```bash
./marketplace-starter-go
```

## Routes

The application has 4 provisioning routes protected by HTTP Basic Auth:

- `POST /provision`
- `PUT /update`
- `DELETE /deactivate`
- `DELETE /deprovision`

It has a public healthcheck route that returns 200 if the service and the database is up and running:

- `GET /healthcheck`

It has a dashboard that can be accessed using Single Sign On with JSON Web Token (JWT):

- `GET /dashboard?jwt=foobar`

It has an JSON RPC route:

- `POST /rpc`

## Testing with qn-marketplace-cli

You can use the [qn-marketplace-cli](https://github.com/quiknode-labs/qn-marketplace-cli) tool to quickly test your add-on while developing it.

To obtain a basic auth string, you can use Go or your language of choice with your username and password, as such:

```go
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "username:password"
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedData)
}
```

For the commands below, the `--basic-auth` flag is the Base64 encoding of `username:password`.
You need to make sure to replace that with your valid credentials (as defined in your `.env` file).

Provisioning:

```sh
./qn-marketplace-cli pudd --base-url http://localhost:3010 --basic-auth dXNlcm5hbWU6cGFzc3dvcmQ=
```

SSO:

Below, make sure that the `jwt-secret` matches `QN_SSO_SECRET` in `.env` file.

```
./qn-marketplace-cli sso --url http://localhost:3010/provision  --basic-auth dXNlcm5hbWU6cGFzc3dvcmQ= --jwt-secret jwt-secret --email jon@example.com --name jon --org QuickNode
```

RPC:

```sh
./qn-marketplace-cli rpc --url http://localhost:3010/provision --rpc-url http://localhost:3010/rpc --rpc-method qn_test --rpc-params "[\"abc\"]" --basic-auth dXNlcm5hbWU6cGFzc3dvcmQ=
```

Healthcheck:

```sh
./qn-marketplace-cli healthcheck --url http://localhost:3010/healthcheck
```


## LICENSE

MIT