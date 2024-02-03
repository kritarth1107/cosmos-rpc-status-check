package main

import (
	"context"
	"database/sql"
	"os"
	"rpc-status-check/common/logs"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	"github.com/joho/godotenv"
)

func initDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./rpc_status.db")
	if err != nil {
		logs.Log.Error(err.Error())
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS rpc_logs (id INTEGER PRIMARY KEY, endpoint TEXT, status TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP, sexytime TEXT)")
	if err != nil {
		logs.Log.Error(err.Error())
	}
	return db
}

func logSuccess(db *sql.DB, endpoint string) {
	currentTime := time.Now()
	_, err := db.Exec("INSERT INTO rpc_logs (endpoint, status, sexytime) VALUES (?, ?, ?)", endpoint, "success", currentTime)
	if err != nil {
		logs.Log.Error(err.Error())
	}
}
func logFailure(db *sql.DB, endpoint string) {
	currentTime := time.Now()
	_, err := db.Exec("INSERT INTO rpc_logs (endpoint, status, sexytime) VALUES (?, ?, ?)", endpoint, "failure", currentTime)
	if err != nil {
		logs.Log.Error(err.Error())
	}
}

func checkRPCConnections(db *sql.DB) {
	ctx := context.Background()
	successCount := 0
	errorCount := 0
	RPCendpoints := strings.Split(os.Getenv("RPC_ENDPOINTS"), ",")

	for _, endpoint := range RPCendpoints {
		_, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(os.Getenv("ADDRESS_PREFIX")), cosmosclient.WithNodeAddress(endpoint))
		if err != nil {
			logs.Log.Error("Inactive RPC  > " + endpoint)
			logFailure(db, endpoint)
			errorCount++
		} else {
			logs.Log.Debug("Connected RPC > " + endpoint)
			logSuccess(db, endpoint)
			successCount++
		}
	}

	logs.Log.Info("Connected RPC > " + strconv.Itoa(successCount))
	logs.Log.Info("Inactive RPC  > " + strconv.Itoa(errorCount))
}

func main() {

	err := godotenv.Load()
	if err != nil {
		logs.Log.Error("Error loading environment variables")
		os.Exit(0)
	}

	db := initDB()
	defer db.Close()
	checkRPCConnections(db)
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			checkRPCConnections(db)
		}
	}
}
