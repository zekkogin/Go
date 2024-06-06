package cache

import (
	"context"
	"github.com/muesli/cache2go"
	"time"
	"wb-L0/model"
	"wb-L0/postgresql"
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
	c.cache.Lock()
	defer c.cache.Unlock()
	c.cache.Add(c.cache.Count(), 10*time.Minute, order)
}

func (c Cache) GetByID(order_uid string) interface{} {
	c.cache.Lock()
	defer c.cache.Unlock()
	order, err := c.cache.Value(order_uid)
	if err != nil {
		return model.Order{}
	} else {
		return order.Data()
	}
}

func (c Cache) LoadFromDatabase(pool postgresql.Pool) {
	if pool.DbIsEmpty() {
		return
	}
	var models []model.Order
	models = pool.GetRows(context.Background())
	l := len(models)
	// TODO: сделать неполную загрузку в кэш
	for i := 0; i < l; i++ {
		var m model.Order
		m = models[i]
		c.CacheAdd(m)
	}
}
