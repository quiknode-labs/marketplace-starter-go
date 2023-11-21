package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/quiknode-labs/marketplace-starter-go/initializers"
	"github.com/quiknode-labs/marketplace-starter-go/models"
	"github.com/quiknode-labs/marketplace-starter-go/types"
	"gorm.io/gorm"
)

func Dashboard(c *gin.Context) {
	session := sessions.Default(c)
	tokenString := c.Request.URL.Query().Get("jwt")
	jwtSecret := os.Getenv("QN_SSO_SECRET")

	if session.Get("quicknodeId") != nil {
		fmt.Printf("Dashboard: found existing session for user: %s\n", session.Get("userName"))
	} else {
		// Parse the token string
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// Define the secret key used to sign the token
			return []byte(jwtSecret), nil
		})

		// Handle any parsing errors
		if err != nil {
			fmt.Println("[1] Error parsing token:", err)
			c.String(http.StatusUnauthorized, "Unauthorized - Error parsing token: %v", err)
			return
		}

		// Verify the token
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Printf("JWT Claims: %v", claims)

			quicknodeId := claims["quicknode_id"].(string)
			userName := claims["name"].(string)
			organizationName := claims["organization_name"].(string)
			email := claims["email"].(string)
			session.Set("quicknodeId", quicknodeId)
			session.Set("email", email)
			session.Set("userName", userName)
			session.Set("organizationName", organizationName)
			session.Save()

			fmt.Printf("Token is valid and verified for user: %s\n", userName)
		} else {
			c.String(http.StatusUnauthorized, "Unauthorized - Failed JWT verification")
			return
		}
	}

	// find the account
	var account models.Account
	findAccountResult := initializers.DB.Where("quicknode_id = ?", session.Get("quicknodeId")).Last(&account)
	if findAccountResult.Error != nil {
		c.String(http.StatusNotFound, "Could not find account with quicknode-id = : %v", session.Get("quicknodeId"))
		return
	}
	log.Println("Found account with ID =", account.ID)

	// Find the RPC requests
	// Pagination parameters from the client
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	var count int64
	initializers.DB.Table("rpc_requests").Where("account_id = ?", account.ID).Count(&count)
	offset := (page - 1) * limit

	var rpcRequests []models.RpcRequest
	findRpcRequestsResult := initializers.DB.Where("account_id = ?", account.ID).Order("created_at desc").Offset(offset).Limit(limit).Find(&rpcRequests)
	if findRpcRequestsResult.Error != nil {
		c.String(http.StatusNotFound, "Error fetching RPC requests for account with ID = : %v", account.ID)
		return
	}
	log.Println("Found", len(rpcRequests), "RPC requests")

	options := GetChartOptions()
	chartOptionsJson, err := json.Marshal(options)
	if err != nil {
		fmt.Println("Error marshaling chart options to JSON:", err)
		return
	}

	chartData, err := RpcRequestLogsMetricsForChart(initializers.DB, int(account.ID))
	log.Print("chartData", chartData)
	if err != nil {
		fmt.Println("Error fetch data for chart", err)
		return
	}
	chartDataJSON, err := json.Marshal(chartData)
	if err != nil {
		fmt.Println("Error marshaling chart data to JSON:", err)
		return
	}

	c.HTML(http.StatusOK, "dash.gohtml", gin.H{
		"AccountId":        c.Param("id"),
		"ActiveTab":        "dashboard",
		"UserName":         session.Get("userName"),
		"OrganizationName": session.Get("organizationName"),
		"Email":            session.Get("email"),
		"RpcRequests":      rpcRequests,
		"Paginate":         false,
		"ChartDataJSON":    string(chartDataJSON),
		"ChartOptionsJSON": string(chartOptionsJson),
		"Count":            count,
		"Page":             page,
		"PreviousPage":     page - 1,
		"NextPage":         page + 1,
		"Pages":            createPageRange(1, int(count/int64(limit))+1),
		"Limit":            limit,
		"FirstPage":        count == 0 || page == 1,
		"LastPage":         count == 0 || page == int(count/int64(limit))+1,
	})
}

func RequestsIndex(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get("quicknodeId") == nil {
		c.String(http.StatusUnauthorized, "Unauthorized - No session found")
		return
	}

	// find the account
	var account models.Account
	findAccountResult := initializers.DB.Where("quicknode_id = ?", session.Get("quicknodeId")).Last(&account)
	if findAccountResult.Error != nil {
		c.String(http.StatusNotFound, "Could not find account with quicknode-id = : %v", session.Get("quicknodeId"))
		return
	}

	// Find the RPC requests
	// Pagination parameters from the client
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	var count int64
	initializers.DB.Table("rpc_requests").Where("account_id = ?", account.ID).Count(&count)
	offset := (page - 1) * limit
	var rpcRequests []models.RpcRequest
	findRpcRequestsResult := initializers.DB.Where("account_id = ?", account.ID).Order("created_at desc").Offset(offset).Limit(limit).Find(&rpcRequests)
	if findRpcRequestsResult.Error != nil {
		c.String(http.StatusNotFound, "Error fetching RPC requests for account with ID = : %v", account.ID)
		return
	}

	log.Println("Found", count, "RPC requests", "page", page, "limit", limit)

	c.HTML(http.StatusOK, "requests.gohtml", gin.H{
		"AccountId":        c.Param("id"),
		"ActiveTab":        "requests",
		"UserName":         session.Get("userName"),
		"OrganizationName": session.Get("organizationName"),
		"Email":            session.Get("email"),
		"RpcRequests":      rpcRequests,
		"Paginate":         (count > int64(limit)),
		"Count":            count,
		"Page":             page,
		"PreviousPage":     page - 1,
		"NextPage":         page + 1,
		"Pages":            createPageRange(1, int(count/int64(limit))+1),
		"Limit":            limit,
		"FirstPage":        count == 0 || page == 1,
		"LastPage":         count == 0 || page == int(count/int64(limit))+1,
	})
}

// Helper methods
func createPageRange(firstPage, lastPage int) []int {
	var pages []int
	for i := firstPage; i <= lastPage; i++ {
		pages = append(pages, i)
	}
	return pages
}

func GetChartOptions() types.ChartOptions {
	options := types.ChartOptions{}

	options.Plugins.Legend.Display = true
	options.Plugins.Legend.Position = "bottom"
	options.Plugins.Title.Display = true
	options.Plugins.Title.Text = "Requests in the last 14 days"
	options.Plugins.Title.Align = "start"
	options.Plugins.Title.Weight = "normal"
	options.Plugins.Title.Padding.Top = 10
	options.Plugins.Title.Padding.Bottom = 30
	options.Plugins.Tooltip.Mode = "nearest"
	options.Plugins.Tooltip.Intersect = true

	options.Scales.X.Stacked = true
	options.Scales.X.Ticks.Major.Enabled = true
	options.Scales.X.Ticks.Major.FontStyle = "bold"
	options.Scales.X.Ticks.Source = "data"
	options.Scales.X.Ticks.AutoSkip = true
	options.Scales.X.Ticks.MaxRotation = 0
	options.Scales.X.Ticks.SampleSize = 100
	options.Scales.X.Ticks.BackdropPadding = 10

	options.Scales.Y.Stacked = true

	options.Layout.Padding = 30

	return options
}

// The GetRpcRequestsByMethod is not being used now but it's a good example of how to use GORM to query the database
// so i'm leaving it here. Also, it could be helpful to build an API that returns metrics to the customer in JSON.
func GetRpcRequestsByMethod(db *gorm.DB, accountID int) ([]types.RpcMethodCallsByDay, error) {
	var results []types.RpcMethodCallsByDay
	query := `
	SELECT 
	  method_name, 
	  date_trunc('day', created_at) AS truncated_date, 
	  COUNT(id) 
	FROM 
	  rpc_requests 
	WHERE 
	  account_id = ? 
	  AND created_at >= ? 
	GROUP BY 
	  method_name, truncated_date 
	ORDER BY 
	  truncated_date, method_name
	`

	if err := db.Raw(query, accountID, time.Now().AddDate(0, 0, -14)).Scan(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// FILLME - Update colors below based on your method names
var MethodNameColors = map[string]string{
	"qn_test":   "#0ad193",
	"qn_family": "#009fd1",
	"qn_gettx":  "#924fe7",
}

func RpcRequestLogsMetricsForChart(db *gorm.DB, accountID int) (map[string]interface{}, error) {
	var requestLogs []models.RpcRequest
	now := time.Now()
	sevenDaysAgo := now.AddDate(0, 0, -14)

	// Fetch the logs
	if err := db.Where("account_id = ? AND created_at >= ?", accountID, sevenDaysAgo).Order("created_at").Find(&requestLogs).Error; err != nil {
		return nil, err
	}

	// Group logs by method name
	requestLogsGroupedByMethod := make(map[string][]models.RpcRequest)
	for _, log := range requestLogs {
		methodName := strings.ToLower(log.MethodName)
		requestLogsGroupedByMethod[methodName] = append(requestLogsGroupedByMethod[methodName], log)
	}
	log.Print("requestLogsGroupedByMethod", requestLogsGroupedByMethod)

	// Create labels and datasets
	labels := []string{}
	for i := 0; i <= 13; i++ {
		labels = append(labels, now.AddDate(0, 0, i-13).Format("2006-01-02"))
	}

	datasets := []map[string]interface{}{}
	for methodName, color := range MethodNameColors {
		methodName = strings.ToLower(methodName)
		log.Print("methodName", methodName)

		if logs, exists := requestLogsGroupedByMethod[methodName]; exists {
			data := []int{}
			for _, label := range labels {
				count := 0
				for _, log := range logs {
					if log.CreatedAt.Format("2006-01-02") == label {
						count++
					}
				}
				data = append(data, count)
			}

			dataset := map[string]interface{}{
				"label":           methodName,
				"data":            data,
				"backgroundColor": color,
				"maxBarThickness": 10,
			}

			datasets = append(datasets, dataset)
		}
	}

	result := map[string]interface{}{
		"labels":   labels,
		"datasets": datasets,
	}

	return result, nil
}
