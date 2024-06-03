package dbpart

import (
	"database/sql"
	"fmt"
	"log"

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
	CheckFatal(err)

	// проверка работы БД, лог если всё плохо
	err = db.Ping()
	CheckFatal(err)

	fmt.Println("Connected to DB!")

	return db
}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
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
	CheckFatal(err)
	fmt.Println("Created table delivery!")
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
	CheckFatal(err)
	fmt.Println("Created table payment!")
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
	CheckFatal(err)
	fmt.Println("Created table items!")
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
	CheckFatal(err)
	fmt.Println("Created table orders!")
}

func CreateAllTables(db *sql.DB) {
	CreateDeliveryTable(db)
	CreatePaymentTable(db)
	CreateItemsTable(db)
	CreateOrdersTable(db)
}

/*
func SelectAll(db *sql.DB, table string) {
	param := fmt.Sprintf("SELECT * FROM %s", table)
	rows, err := db.Query(`SELECT * FROM table`)
	CheckFatal(err)

	defer rows.Close()

	for rows.Next() {

	 	var name string
	 	var age int

	 	err = rows.Scan(&age, &name)
	 	CheckFatal(err)

	 	fmt.Println(age, name)
	}

	CheckFatal(err)

	fmt.Println("Selected from %v", table)

	//insert := `INSERT INTO "tests" (id, name) values(2, 'Ann')`

	//_, err = db.Exec(insert)
	//CheckFatal(err)

	//fmt.Println("Inserted!")
}
*/
