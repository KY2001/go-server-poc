package cloudsql

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"github.com/KY2001/go-server-poc/config"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/go-sql-driver/mysql"
)

// Ref: https://github.com/GoogleCloudPlatform/golang-samples/blob/main/cloudsql/mysql/database-sql/connect_connector.go
func GetConnectionPool() (*sql.DB, error) {
	conf := config.NewConfig()

	dbUser := conf.DB.DBUser                                // e.g. 'my-db-user'
	dbPass := conf.DB.DBPass                                // e.g. 'my-db-password'
	dbName := conf.DB.DBName                                // e.g. 'my-database'
	instanceConnectionName := conf.DB.InstanceConnectonName // e.g. 'project:region:instance'

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
	return dbPool, nil
}
