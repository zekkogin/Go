package cache

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muesli/cache2go"
	"log"
	"time"
	"wb-L0/model"
)

type CacheOrder interface {
	CacheAdd(order model.Order)
	GetByID(id int) interface{}
}

type Cache struct {
	cache *cache2go.CacheTable
}

func CacheNew() *Cache {
	cache := cache2go.Cache("OrdersCache")
	return &Cache{cache}
}

func (c Cache) CacheAdd(order model.Order) {
	log.Println("zashel v cache")
	c.cache.Lock()
	defer c.cache.Unlock()
	c.cache.Add(c.cache.Count(), 10*time.Minute, order)
	log.Println("dobavelno v cache")
}

func (c Cache) GetByID() interface{} {
	c.cache.Lock()
	defer c.cache.Unlock()
	order, err := c.cache.Value(c.cache.Count())
	if err != nil {
		return model.Order{}
	} else {
		return order.Data()
	}
}

func (c Cache) LoadFromDatabase(pool pgxpool.Pool) error {

}
