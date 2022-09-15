package db

import (
	"fmt"
	"github.com/jackc/pgconn"
)

// fillOrders add information in table orders.
func (dataBase *DataBase) fillOrders(data Order, query string) error {
	err := dataBase.DB.QueryRow(dataBase.Ctx, query, data.OrderUID, data.TrackNumber, data.Entry, data.Locale,
		data.InternalSignature, data.CustomerID, data.DeliveryService, data.Shardkey, data.SmID,
		data.DateCreated, data.OofShard).Scan()
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		fmt.Println(pgErr)
		return pgErr
	}
	return nil
}

// fillItem add information in table items.
func (dataBase *DataBase) fillItem(item Items, query string) error {
	err := dataBase.DB.QueryRow(dataBase.Ctx, query, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name,
		item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status).Scan()
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		fmt.Println(pgErr)
		return pgErr
	}
	return nil
}

// fillPayment добавляет данные в таблицу payments
func (dataBase *DataBase) fillPayment(paym Payment, query string) error {
	err := dataBase.DB.QueryRow(dataBase.Ctx, query, paym.Transaction, paym.RequestID, paym.Currency, paym.Provider,
		paym.Amount, paym.PaymentDt, paym.Bank, paym.DeliveryCost, paym.GoodsTotal, paym.CustomFee).Scan()
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		fmt.Println(pgErr)
		return pgErr
	}
	return nil
}

// fillDelivery add information in table deliveries.
func (dataBase *DataBase) fillDelivery(deliv Delivery, query string) error {
	err := dataBase.DB.QueryRow(dataBase.Ctx, query, deliv.Name, deliv.Phone, deliv.Zip, deliv.City, deliv.Address,
		deliv.Region, deliv.Email).Scan()
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		fmt.Println(pgErr)
		return pgErr
	}
	return nil
}

// fillOrderInfo add information in table order_info.
func (dataBase *DataBase) fillOrderInfo(data Order, query string) error {
	err := dataBase.DB.QueryRow(dataBase.Ctx, query, data.OrderUID, data.TrackNumber).Scan()
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		fmt.Println(pgErr)
		return pgErr
	}
	return nil
}
