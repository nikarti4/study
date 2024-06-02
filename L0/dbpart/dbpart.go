package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = ""
	dbname = "wbl0"
)


func main() {

	// конфигурации БД
	//conf := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	conf := "postgres://postgres:@localhost:5432/tests?sslmode=disable"
	// попытка открытия БД
	db, err := sql.Open("postgres", conf)
	defer db.Close()

	// если не удалось открыть - вывести лог
	CheckFatal(err)

	// проверка работы БД, лог если всё плохо
	err = db.Ping()

	CheckFatal(err)

	fmt.Println("Connected!")

	CreateTestsTabel(db)


	// insert := `INSERT INTO "test" (id, name) values(1, 'Jonh')`

	// _, err = db.Exec(insert)
	// CheckFatal(err)

	// rows, err := db.Query(`SELECT * FROM payment`)
	// CheckFatal(err)

	// defer rows.Close()

	// for rows.Next() {

	// 	var name string
	// 	var age int

	// 	err = rows.Scan(&name, &age)
	// 	CheckFatal(err)
	 
	// 	fmt.Println(name, age)
	// }

	// CheckFatal(err)

}

func CheckFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTestsTabel(db *sql.DB) {
	q := `CREATE TABLE IF NOT EXISTS tests (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL
		)`
	
	_, err := db.Exec(q)

	CheckFatal(err)
}
