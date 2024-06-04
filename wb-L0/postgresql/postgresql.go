package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"
)

type Pool struct {
	P *pgxpool.Pool
}

func NewClient(ctx context.Context, maxAttempts int) (pool *pgxpool.Pool, err error) {

	err = DoWithTries(func() error {
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbName := os.Getenv("DB_NAME")

		dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
		url := fmt.Sprintf("%s%s", dbURL, "?pool_max_conns=10&pool_min_conns=1&pool_max_conn_lifetime=30m&pool_health_check_period=1m")
		fmt.Println(url)
		cfg, err := pgxpool.ParseConfig(url)
		if err != nil {
			log.Fatalf("Unable to parse config: %v\n", err)
		}
		pool, err = pgxpool.NewWithConfig(ctx, cfg)
		if err != nil {
			log.Fatalf("Unable to newwithconfig: %v\n", err)
		}
		pool.Config()
		if err != nil {
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		fmt.Errorf("Error occurred: %s", err)
	}

	return pool, nil
}

func DoWithTries(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}
		return nil
	}
	return
}
