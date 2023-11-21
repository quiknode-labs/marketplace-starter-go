package main

import (
	"github.com/quiknode-labs/marketplace-starter-go/initializers"
	"github.com/quiknode-labs/marketplace-starter-go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Account{})
	initializers.DB.AutoMigrate(&models.Endpoint{})
	initializers.DB.AutoMigrate(&models.RpcRequest{})
	initializers.DB.Exec("CREATE INDEX IF NOT EXISTS idx_accounts_on_quicknode_id_deleted_at_id ON accounts (quicknode_id, deleted_at, id DESC)")
	initializers.DB.Exec("CREATE INDEX IF NOT EXISTS idx_rpc_requests_on_account_id_deleted_at_created_at ON rpc_requests (account_id, deleted_at, created_at DESC)")
}
