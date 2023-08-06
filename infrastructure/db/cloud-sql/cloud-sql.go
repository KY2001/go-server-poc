package cloudsql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/KY2001/go-server-poc/config"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func InitClient() {
	var err error
	db, err = GetConnectionPool()
	if err != nil {
		log.Fatalf("InitClient: Failed to connect to Cloud SQL: %v\n", err)
	}
}

func CloseClient() {
	if db != nil {
		db.Close()
	}
}

func GetClient() *sql.DB {
	if db == nil {
		InitClient()
		log.Println("GetClient: DB client should be initialized when server starts.")
	}
	return db
}

// Ref: https://github.com/GoogleCloudPlatform/golang-samples/blob/main/cloudsql/mysql/database-sql/connect_connector.go
func GetConnectionPool() (*sql.DB, error) {
	conf := config.NewConfig()

	dbUser := conf.DB.User                                   // e.g. 'my-db-user'
	dbPass := conf.DB.Pass                                   // e.g. 'my-db-password'
	dbName := conf.DB.Name                                   // e.g. 'my-database'
	instanceConnectionName := conf.DB.InstanceConnectionName // e.g. 'project:region:instance'

	usePrivate := conf.DB.PrivateIP

	d, err := cloudsqlconn.NewDialer(context.Background())
	if err != nil {
		return nil, fmt.Errorf("cloudsqlconn.NewDialer: %w", err)
	}
	var opts []cloudsqlconn.DialOption
	if usePrivate != "" {
		opts = append(opts, cloudsqlconn.WithPrivateIP())
	}
	mysql.RegisterDialContext("cloudsqlconn",
		func(ctx context.Context, addr string) (net.Conn, error) {
			return d.Dial(ctx, instanceConnectionName, opts...)
		})

	dbURI := fmt.Sprintf("%s:%s@cloudsqlconn(localhost:3306)/%s?parseTime=true",
		dbUser, dbPass, dbName)

	dbPool, err := sql.Open("mysql", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	// sql.DB の挙動を制御するパラメータ：https://please-sleep.cou929.nu/go-sql-db-connection-pool.html
	dbPool.SetMaxOpenConns(conf.DB.MaxOpenConns)
	dbPool.SetMaxIdleConns(conf.DB.MaxIdleConns)
	dbPool.SetConnMaxLifetime(conf.DB.ConnMaxLifetime)
	dbPool.SetConnMaxIdleTime(conf.DB.ConnMaxIdletime)

	return dbPool, nil
}
