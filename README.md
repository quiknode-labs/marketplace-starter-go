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


7. You can start making HTTP requests using curl or Postman:

```
POST http://localhost:3010/provision
PUT http://localhost:3010/update
DELETE http://localhost:3010/deactivate_endpoint
DELETE http://localhost:3010/deprovision
```

See the [marketplace guide to provisioning](https://www.quicknode.com/guides/marketplace/how-provisioning-works-for-marketplace-partners) for more info.

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