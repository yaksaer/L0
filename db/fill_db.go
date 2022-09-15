package db

// FillDB create SQL query and call functions to add information to the DB.
func (dataBase *DataBase) FillDB(data Order) error {
	query := `INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature,
                   customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	err := dataBase.fillOrders(data, query)
	if err != nil {
		return err
	}
	query = `INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size,
	              total_price, nm_id, brand, status)
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	for _, v := range data.Items {
		err = dataBase.fillItem(v, query)
		if err != nil {
			return err
		}
	}
	query = `INSERT INTO deliveries (name, phone, zip, city, address, region, email)
					values ($1, $2, $3, $4, $5, $6, $7)`
	err = dataBase.fillDelivery(data.Delivery, query)
	if err != nil {
		return err
	}
	query = `INSERT INTO payments (transaction, request_id, currency, provider, amount,
	                 payment_dt, bank, delivery_cost, goods_total, custom_fee)
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	err = dataBase.fillPayment(data.Payment, query)
	if err != nil {
		return err
	}
	query = `INSERT INTO order_info (orderID, trackNumber)
					values ($1, $2)`
	err = dataBase.fillOrderInfo(data, query)
	if err != nil {
		return err
	}
	return nil
}
