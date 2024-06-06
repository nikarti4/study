package dbpart

import (
	"database/sql"
	"fmt"

	"L0/common"
	"L0/model"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "new_pass"
	dbname   = "my_test"
)

func ConnectDB() *sql.DB {
	// конфигурации БД
	conf := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// попытка открытия БД
	db, err := sql.Open("postgres", conf)

	// если не удалось открыть - вывести лог
	common.CheckFatal(err)

	// проверка работы БД, лог если всё плохо
	err = db.Ping()
	common.CheckFatal(err)

	fmt.Println("Connected to DB!")

	return db
}

func CreateDeliveryTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS delivery (
		name VARCHAR(100) NOT NULL,
		phone VARCHAR(100) PRIMARY KEY,
		zip VARCHAR(100) NOT NULL,
		city VARCHAR(100) NOT NULL,
		address VARCHAR(100) NOT NULL,
		region VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL
	)`

	_, err := db.Exec(q)
	common.CheckFatal(err)
	//fmt.Println("Created table delivery!")
}

func CreatePaymentTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS payment (
		transaction VARCHAR(100) PRIMARY KEY,
		request_id VARCHAR(100),
		currency VARCHAR(100) NOT NULL,
		provider VARCHAR(100) NOT NULL,
		amount BIGINT NOT NULL,
		payment_dt BIGINT NOT NULL,
		bank VARCHAR(100) NOT NULL,
		delivery_cost BIGINT NOT NULL,
		goods_total BIGINT NOT NULL,
		custom_fee BIGINT NOT NULL
	)`

	_, err := db.Exec(q)
	common.CheckFatal(err)
	//fmt.Println("Created table payment!")
}

func CreateItemsTable(db *sql.DB) {

	q := `CREATE TABLE IF NOT EXISTS items (
		chrt_id BIGINT PRIMARY KEY,
		track_number VARCHAR(100) NOT NULL,
		price BIGINT NOT NULL,
		rid VARCHAR(100) NOT NULL,
		name VARCHAR(100) NOT NULL,
		sale BIGINT NOT NULL,
		size VARCHAR(100) NOT NULL,
		total_price BIGINT NOT NULL,
		nm_id BIGINT NOT NULL,
		brand VARCHAR(100) NOT NULL,
		status BIGINT NOT NULL
	)`

	_, err := db.Exec(q)
	common.CheckFatal(err)
	//fmt.Println("Created table items!")
}

func CreateOrdersTable(db *sql.DB) {

	q := `CREATE TABLE IF NOT EXISTS orders (
			order_uid VARCHAR(100) PRIMARY KEY,
			track_number VARCHAR(100),
			entry VARCHAR(100) NOT NULL,
			locale VARCHAR(100) NOT NULL,
			internal_signature VARCHAR(100),
			customer_id VARCHAR(100) NOT NULL,
			delivery_service VARCHAR(100) NOT NULL,
			shardkey VARCHAR(100) NOT NULL,
			sm_id BIGINT NOT NULL,
			date_created VARCHAR(100) NOT NULL,
			oof_shard VARCHAR(100) NOT NULL
		)`

	_, err := db.Exec(q)
	common.CheckFatal(err)
	//fmt.Println("Created table orders!")
}

func CreateOrdersItemsTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS orders_items (
		order_uid VARCHAR(100),
		chrt_id BIGINT
	)`

	_, err := db.Exec(q)
	common.CheckFatal(err)
}

func CreateOrdersDeliveryTable(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS orders_delivery (
		order_uid VARCHAR(100),
		phone VARCHAR(100)
	)`

	_, err := db.Exec(q)
	common.CheckFatal(err)
}

func CreateAllTables(db *sql.DB) {
	CreateDeliveryTable(db)
	CreatePaymentTable(db)
	CreateItemsTable(db)
	CreateOrdersTable(db)
	CreateOrdersItemsTable(db)
	CreateOrdersDeliveryTable(db)
	fmt.Println("All tables are good!")
}

func InsertOrder(db *sql.DB, order model.Order_t) {
	insert := `INSERT INTO orders (
			order_uid,
			track_number,
			entry,
			locale,
			internal_signature,
			customer_id,
			delivery_service,
			shardkey,
			sm_id,
			date_created,
			oof_shard 
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11 
		)`

	_, err := db.Exec(insert,
		order.Order_uid, order.Track_number, order.Entry, order.Locale,
		order.Internal_signature, order.Customer_id, order.Delivery_service,
		order.Shardkey, order.Sm_id, order.Date_created, order.Oof_shard,
	)

	//common.CheckFatal(err)
	if err != nil {
		fmt.Printf("INSERT in orders FAIL - %v\n", err)
	} else {

		fmt.Println("INSERT in orders complete!")
	}

	// now here
	InsertDelivery(db, order.Delivery)
	InsertPayment(db, order.Payment)
	InsertItems(db, order.Items)
	InsertOrdersItems(db, order)
	InsertOrdersDelivery(db, order)
}

func InsertDelivery(db *sql.DB, delivery model.Delivery_t) {
	insert := `INSERT INTO delivery (
			name,
			phone,
			zip,
			city,
			address,
			region,
			email
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)`

	_, err := db.Exec(insert,
		delivery.Name, delivery.Phone, delivery.Zip, delivery.City,
		delivery.Address, delivery.Region, delivery.Email,
	)

	//common.CheckFatal(err)
	if err != nil {
		fmt.Printf("INSERT in delivery FAIL - %v\n", err)
	} else {
		fmt.Println("INSERT in delivery complete!")
	}
}

func InsertPayment(db *sql.DB, payment model.Payment_t) {
	insert := `INSERT INTO payment (
			transaction,
			request_id,
			currency,
			provider,
			amount,
			payment_dt,
			bank,
			delivery_cost,
			goods_total,
			custom_fee
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
		)`

	_, err := db.Exec(insert,
		payment.Transaction, payment.Request_id, payment.Currency, payment.Provider,
		payment.Amount, payment.Payment_dt, payment.Bank, payment.Delivery_cost,
		payment.Goods_total, payment.Custom_fee,
	)

	//common.CheckFatal(err)
	if err != nil {
		fmt.Printf("INSERT in payment FAIL - %v\n", err)
	} else {
		fmt.Println("INSERT in payment complete!")
	}
}

func InsertItems(db *sql.DB, items []model.Items_t) {
	insert := `INSERT INTO items (
			chrt_id,
			track_number,
			price,
			rid,
			name,
			sale,
			size,
			total_price,
			nm_id,
			brand,
			status
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		)`

	for _, item := range items {
		_, err := db.Exec(insert,
			item.Chrt_id, item.Track_number, item.Price, item.Rid, item.Name, item.Sale,
			item.Size, item.Total_price, item.Nm_id, item.Brand, item.Status,
		)

		//common.CheckFatal(err)
		if err != nil {
			fmt.Printf("INSERT in items FAIL - %v\n", err)
		} else {
			fmt.Println("INSERT in items complete!")
		}
	}
}

func InsertOrdersItems(db *sql.DB, order model.Order_t) {
	insert := `INSERT INTO orders_items (
		order_uid,
		chrt_id
	) VALUES (
		$1, $2
	)`

	//for _, elem := range order {

		for _, item := range order.Items {
			_, err := db.Exec(insert,
				order.Order_uid, item.Chrt_id,
			)
			if err != nil {
				fmt.Printf("INSERT in orders_items FAIL - %v\n", err)
			} else {
				fmt.Println("INSERT in orders_items complete!")
			}
		}

	//}
}


func InsertOrdersDelivery(db *sql.DB, order model.Order_t) {
	insert := `INSERT INTO orders_delivery (
		order_uid,
		phone
	) VALUES (
		$1, $2
	)`

	//for _, elem := range order {
			_, err := db.Exec(insert,
			order.Order_uid, order.Delivery.Phone,
		)

		if err != nil {
			fmt.Printf("INSERT in orders_delivery FAIL - %v\n", err)
		} else {
			fmt.Println("INSERT in orders_delivery complete!")
		}
	//}


}


func GetOrderByID(db *sql.DB, order_uid string) (*model.Order_t, error) {
	rows, err := db.Query(`
		SELECT * 
		FROM orders
		WHERE order_uid = $1`, order_uid)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	order := &model.Order_t{}

	for rows.Next() {

		err = rows.Scan(
			&order.Order_uid, &order.Track_number, &order.Entry, &order.Locale,
			&order.Internal_signature, &order.Customer_id, &order.Delivery_service,
			&order.Shardkey, &order.Sm_id, &order.Date_created, &order.Oof_shard,
		)

		if err != nil {
			return nil, err
		}
	}

	return order, nil
}

func GetAllOrders(db *sql.DB) ([]model.Order_t, error) {
	var orders []model.Order_t

	rows, err := db.Query(`
	SELECT * 
	FROM orders`)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	order := model.Order_t{}

	for rows.Next() {

		err = rows.Scan(
			&order.Order_uid, &order.Track_number, &order.Entry, &order.Locale,
			&order.Internal_signature, &order.Customer_id, &order.Delivery_service,
			&order.Shardkey, &order.Sm_id, &order.Date_created, &order.Oof_shard,
		)

		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}
