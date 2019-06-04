package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewDSN() string {
	cfg := mysql.NewConfig()
	cfg.User = "john"
	cfg.Passwd = "123456"
	cfg.DBName = "test"
	cfg.Params = map[string]string{
		"charset": "utf8mb4",
	}
	cfg.ParseTime = true
	cfg.Collation = "utf8mb4_general_ci"
	return cfg.FormatDSN()
}

func NewDatabase() *sql.DB {
	dsn := NewDSN()
	fmt.Println(dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	return db
}

func main() {
	db := NewDatabase()
	defer db.Close()
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(1)
	// db.SetConnMaxLifetime(0)

	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// This only works with Postgresql, not Mysql
	// result ,err := pool.ExecContext(ctx, stmt)

	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30
	var connectionID string
	err = conn.QueryRowContext(ctx, "SELECT CONNECTION_ID()").Scan(&connectionID)
	if err != nil {
		log.Fatal(err)
	}
	// Return the connection back to the pool.
	var b []byte
	err = conn.QueryRowContext(ctx, "SELECT SLEEP(10)").Scan(&b)
	if err != nil {
		if err == context.Canceled || err == context.DeadlineExceeded {
			fmt.Println("cancelling", rr)
			kill(db, connectionID)
		}
	}
	log.Println(string(b))

}

func kill(db *sql.DB, connID string) {
	db.Exec("KILL QUERY ?", connID)
}
