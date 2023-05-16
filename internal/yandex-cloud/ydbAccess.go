package ydbAccess

import (
	"context"
	"database/sql"
	"log"
	"os"

	ev "github.com/ydb-platform/ydb-go-sdk-auth-environ"
	"github.com/ydb-platform/ydb-go-sdk/v3"
)

func ConnectToYDB() *sql.DB {
	ctx := context.Background()

	// See many ydb.Option's for configure driver https://pkg.go.dev/github.com/ydb-platform/ydb-go-sdk/v3#Option
	nativeDriver, _ := ydb.Open(ctx, os.Getenv("DB_STR"), ev.WithEnvironCredentials(ctx))

	// See ydb.ConnectorOption's for configure connector https://pkg.go.dev/github.com/ydb-platform/ydb-go-sdk/v3#ConnectorOption
	connector, _ := ydb.Connector(nativeDriver)
	db := sql.OpenDB(connector)

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected to DB")

	return db
}
