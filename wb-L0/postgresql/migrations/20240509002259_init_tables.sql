-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
                        order_uid VARCHAR(255) PRIMARY KEY,
                        track_number VARCHAR(255),
                        entry VARCHAR(255),
                        locale VARCHAR(10),
                        internal_signature VARCHAR(255),
                        customer_id VARCHAR(255),
                        delivery_service VARCHAR(255),
                        shardkey VARCHAR(255),
                        sm_id INTEGER,
                        date_created TIMESTAMP,
                        oof_shard VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS delivery (
                          order_uid VARCHAR(255),
                          name VARCHAR(255),
                          phone VARCHAR(255),
                          zip VARCHAR(255),
                          city VARCHAR(255),
                          address VARCHAR(255),
                          region VARCHAR(255),
                          email VARCHAR(255),
                          FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);

CREATE TABLE IF NOT EXISTS payment (
                         order_uid VARCHAR(255),
                         transaction VARCHAR(255),
                         request_id VARCHAR(255),
                         currency VARCHAR(10),
                         provider VARCHAR(255),
                         amount INTEGER,
                         payment_dt INTEGER,
                         bank VARCHAR(255),
                         delivery_cost INTEGER,
                         goods_total INTEGER,
                         custom_fee INTEGER,
                         FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);

CREATE TABLE IF NOT EXISTS items (
                       order_uid VARCHAR(255),
                       chrt_id INTEGER,
                       track_number VARCHAR(255),
                       price INTEGER,
                       rid VARCHAR(255),
                       name VARCHAR(255),
                       sale INTEGER,
                       size INTEGER,
                       total_price INTEGER,
                       nm_id INTEGER,
                       brand VARCHAR(255),
                       status INTEGER,
                       FOREIGN KEY (order_uid) REFERENCES orders(order_uid)
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE delivery;
DROP TABLE payment;
DROP TABLE items;
DROP TABLE orders;
-- +goose StatementEnd