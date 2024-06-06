SELECT json_build_object(
               'order_uid', o.order_uid,
               'track_number', o.track_number,
               'entry', o.entry,
               'delivery', json_build_object(
                       'name', d.name,
                       'phone', d.phone,
                       'zip', d.zip,
                       'city', d.city,
                       'address', d.address,
                       'region', d.region,
                       'email', d.email
                           ),
               'payment', json_build_object(
                       'transaction', p.transaction,
                       'request_id', p.request_id,
                       'currency', p.currency,
                       'provider', p.provider,
                       'amount', p.amount,
                       'payment_dt', p.payment_dt,
                       'bank', p.bank,
                       'delivery_cost', p.delivery_cost,
                       'goods_total', p.goods_total,
                       'custom_fee', p.custom_fee
                          ),
               'items', json_agg(
                       json_build_object(
                               'chrt_id', i.chrt_id,
                               'track_number', i.track_number,
                               'price', i.price,
                               'rid', i.rid,
                               'name', i.name,
                               'sale', i.sale,
                               'size', i.size,
                               'total_price', i.total_price,
                               'nm_id', i.nm_id,
                               'brand', i.brand,
                               'status', i.status
                       )
                        ),
               'locale', o.locale,
               'internal_signature', o.internal_signature,
               'customer_id', o.customer_id,
               'delivery_service', o.delivery_service,
               'shardkey', o.shardkey,
               'sm_id', o.sm_id,
               'date_created', o.date_created,
               'oof_shard', o.oof_shard
       ) AS orders_json
FROM orders AS o
         INNER JOIN delivery AS d ON o.order_uid = d.order_uid
         INNER JOIN payment AS p ON o.order_uid = p.order_uid
         INNER JOIN items AS i ON o.order_uid = i.order_uid
WHERE o.order_uid = $1
GROUP BY
    o.order_uid, o.track_number, o.entry, d.name, d.phone, d.zip, d.city, d.address,
    d.region, d.email, p.transaction, p.request_id, p.currency, p.provider, p.amount,
    p.payment_dt, p.bank, p.delivery_cost, p.goods_total, p.custom_fee, i.chrt_id,
    i.track_number, i.price, i.rid, i.name, i.sale, i.size, i.total_price, i.nm_id,
    i.brand, i.status, o.locale, o.internal_signature, o.customer_id, o.delivery_service,
    o.shardkey, o.sm_id, o.date_created, o.oof_shard;