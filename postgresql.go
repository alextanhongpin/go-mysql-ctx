package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lib/pq"
)

func NewDatabase(dsn string) *sql.DB {
	connector, err := pq.NewConnector(dsn)
	if err != nil {
		log.Fatal(err)
	}
	return sql.OpenDB(connector)
}

func main() {
	dsn := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println(dsn)
	db := NewDatabase(dsn)

	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var b []byte
	err := db.QueryRowContext(ctx, "SELECT pg_sleep(10)").Scan(&b)
	if err != nil {
		if err == context.Canceled || err == context.DeadlineExceeded {
			fmt.Println("error", ctx.Err())
		} else {
			log.Fatal(err)
		}
	}
	fmt.Println("got", string(b))
}
