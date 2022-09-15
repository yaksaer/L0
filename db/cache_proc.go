package db

import (
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/patrickmn/go-cache"
)

// getAllID receive all recorded order_uid from DB and return them as []string.
func (dataBase *DataBase) getAllID() []string {
	query := `SELECT array_agg(order_uid) FROM orders`
	var buf []string
	err := dataBase.DB.QueryRow(dataBase.Ctx, query).Scan(&buf)
	pgEr, ok := err.(*pgconn.PgError)
	if ok {
		fmt.Println(pgEr)
	}
	return buf
}

// getAllData creates SQL query and return structure with all information about an order.
func (dataBase *DataBase) getAllData(orderID string) Order {
	query := `SELECT o.*, to_jsonb(p.*) AS "payment", 
       			(SELECT jsonb_agg(to_jsonb(i.*)) FROM items i WHERE i.track_number=o.track_number) AS "items", 
				to_jsonb(deliv) AS "delivery"
				FROM orders o
    			INNER JOIN payments p ON o.order_uid=p.transaction
    			INNER JOIN items i ON i.track_number=o.track_number
    			JOIN (
    			    SELECT d.Name, d.Phone, d.Zip, d.City, d.Address, d.Region, d.Email 
    			    FROM deliveries d
    			    WHERE d.id=(
    			        SELECT oi.DelID
    			        FROM order_info oi
    			        WHERE oi.orderID=$1
    			    )
    			) AS deliv ON TRUE
         		WHERE o.order_uid=$1;`
	var orderBuf Order
	err := dataBase.DB.QueryRow(dataBase.Ctx, query, orderID).Scan(&orderBuf.OrderUID,
		&orderBuf.TrackNumber, &orderBuf.Entry, &orderBuf.Locale, &orderBuf.InternalSignature,
		&orderBuf.CustomerID, &orderBuf.DeliveryService, &orderBuf.Shardkey, &orderBuf.SmID,
		&orderBuf.DateCreated, &orderBuf.OofShard, &orderBuf.Payment, &orderBuf.Items, &orderBuf.Delivery)
	pgEr, ok := err.(*pgconn.PgError)
	if ok {
		fmt.Println(pgEr)
	}
	return orderBuf
}

// StartCache connect to DB and write all recorder order's information in cache.
func (dataBase *DataBase) StartCache(newCache *cache.Cache) {
	//Recieving all recorded order_uid from DB
	idBuf := dataBase.getAllID()
	if len(idBuf) == 0 {
		return
	}
	for _, v := range idBuf {
		//Get al information about order
		data := dataBase.getAllData(v)
		//Add data to cache
		newCache.Set(v, data, -1)
	}
}
