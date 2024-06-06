package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"log"
	"time"
	"wb-L0/cache"
	"wb-L0/nats"
	psql "wb-L0/postgresql"
)

type App struct {
	C    *cache.Cache
	Pool psql.Pool
	Sc   stan.Conn
	R    *gin.Engine
}

//func (a App) GetDB() *pgxpool.Pool {
//	return a.Db
//}

func NewApp() App {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return App{
		C:    cache.CacheNew(),
		Pool: initializeDB(ctx),
		Sc:   initializeStanConn(ctx),
		R:    initializeHttp(ctx),
	}
}

func initializeDB(ctx context.Context) psql.Pool {
	pool, err := psql.NewClient(ctx, 5)
	if err != nil {
		log.Fatalf("Error occurred with create new pgx.pool: %s", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	return psql.Pool{P: pool}
}

func initializeStanConn(ctx context.Context) stan.Conn {
	sc := nats.Connect(ctx, "sub")
	return sc
}

func initializeHttp(ctx context.Context) *gin.Engine {
	r := gin.Default()
	return r
}
