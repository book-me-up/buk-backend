package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDB(ctx context.Context) *sqlx.DB {
	log.Println("Connecting to Postgres...")

	jsonCtx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*5))
	defer cancel()

	connStr := "postgres://postgres:postgres_pass@postgres:5432/buk_db?sslmode=disable"
	json, err := sqlx.ConnectContext(jsonCtx, "postgres", connStr)
	if err != nil {
		log.Fatal(err, "Postgres connection failed...")
	}

	if err := json.Ping(); err != nil {
		log.Fatal("error: ", err.Error())
	} else {
		log.Println("Postgres connected !")
	}

	log.Println("Postgres connected !")

	return json
}
