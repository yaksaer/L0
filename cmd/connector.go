package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
	"os"
	"wb/db"
)

type connections struct {
	Cache    *cache.Cache
	NatsConn stan.Conn
	Database *db.DataBase
}

// StartConnections start connections to all necessary services: NATS, DB, Cache.
func StartConnections() *connections {
	var conn connections
	conn.dbConnect()
	//Creating new cache object
	conn.Cache = cache.New(-1, -1)
	conn.Database.StartCache(conn.Cache)
	//Connecting to NATS
	conn.natsConnect()
	return &conn
}

// Stop close connections with NATS and DB.
func (conn *connections) Stop() {
	conn.NatsConn.Close()
	conn.Database.DB.Close(conn.Database.Ctx)
}

// dbConnect connect to DB and return DB object.
func (conn *connections) dbConnect() {
	var dataBase db.DataBase
	err := error(nil)
	dataBase.Ctx = context.Background()
	dataBase.DB, err = pgx.Connect(dataBase.Ctx, db.URL)
	if err != nil {
		fmt.Println("Can't connect to database: ", err)
		os.Exit(1)
	}
	conn.Database = &dataBase
}

// natsConnect connect to NATS and subscribe to its channel. Processing data received from NATS channel.
func (conn *connections) natsConnect() {
	connection, err := stan.Connect("test-cluster", "client-1")
	if err != nil {
		fmt.Printf("Can't connect to NATS = %v\n", err)
		return
	}
	//Subscribe on NATS server
	_, err = connection.Subscribe("service", func(m *stan.Msg) {
		var data db.Order
		err := json.Unmarshal(m.Data, &data)
		if err != nil || data.OrderUID == "" || data.Payment == (db.Payment{}) ||
			data.Items[0] == (db.Items{}) || data.Delivery == (db.Delivery{}) {
			fmt.Println("Wrong json")
			return
		}
		//If correct JSON have been received
		fmt.Println("Receive correct json")
		err = conn.Database.FillDB(data) //FillingDB
		if err != nil {
			fmt.Printf("Error occured while DB filling: %v\n", err)
			return
		}
		fmt.Println("Json successfully adeed to DB")
		conn.Database.StartCache(conn.Cache) //Add new data in cache
		fmt.Println("Json successfully added to cache")
	})
	if err != nil {
		fmt.Printf("failed to subscription: %w", err)
	}

	conn.NatsConn = connection
}
