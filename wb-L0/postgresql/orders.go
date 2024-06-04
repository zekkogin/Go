package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"wb-L0/model"
)

type OrdersDb interface {
	InitTableOrders()
	insertOrder(ctx context.Context, order model.Order) error
}

func (pool Pool) InsertOrder(ctx context.Context, order model.Order) error {
	fmt.Println("Вошел в InsertOrder")
	tx, err := pool.P.Begin(ctx)
	if err != nil {
		fmt.Println("Ошибка при начале транзакции:", err)
		return err
	}
	defer tx.Rollback(ctx)
	_, err = tx.Exec(ctx, `INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		fmt.Println("Ошибка при вставке в orders:", err)
		return err
	}

	_, err = tx.Exec(ctx, `INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, order.OrderUID, order.Delivery.Name, order.Delivery.Phone, order.Delivery.Zip, order.Delivery.City, order.Delivery.Address, order.Delivery.Region, order.Delivery.Email)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, `INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`, order.OrderUID, order.Payment.Transaction, order.Payment.RequestID, order.Payment.Currency, order.Payment.Provider, order.Payment.Amount, order.Payment.PaymentDt, order.Payment.Bank, order.Payment.DeliveryCost, order.Payment.GoodsTotal, order.Payment.CustomFee)
	if err != nil {
		return err
	}

	for _, val := range order.Items {
		_, err = tx.Exec(ctx, `INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`, order.OrderUID, val.ChrtID, val.TrackNumber, val.Price, val.Rid, val.Name, val.Sale, val.Size, val.TotalPrice, val.NmID, val.Brand, val.Status)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return err
}

func (pool Pool) InitTableOrders() {
	db := stdlib.OpenDBFromPool(pool.P)
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to close DB: %s", err)
		}
	}()

	migrations, err := goose.CollectMigrations("postgresql/migrations", 0, goose.MaxVersion)
	if err != nil {
		fmt.Errorf("Failed to collect first migration ")
	}

	if len(migrations) == 0 {
		log.Fatal("No migrations found, create init table first sql file in migrations dir")
	}
	err = goose.UpTo(db, "postgresql/migrations", migrations[0].Version)

	if err != nil {
		fmt.Errorf("Error occurred in goose.UpTo -> init-table.sql: %v", err)
	}
}
