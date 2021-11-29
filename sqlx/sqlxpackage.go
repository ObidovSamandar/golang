package sqlxpackage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE person (
    first_name text,
    last_name text,
    email text
);

CREATE TABLE place (
    country text,
    city text NULL,
    telcode integer
)`

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

func ConnectionAndQueiries() {

	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=mylocaldatabase sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	// gin

	// db.MustExec(schema)
	// tx := db.MustBegin()
	// tx.Exec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Karim", "Karimov", "karim@karimov.net")
	// tx.Exec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Olim", "Olimov", "karim@karimov.net")
	// tx.Exec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "USA", "New York", "111111111")
	// tx.Exec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "USA", "New York", "111111111")

	// postgresqlUpdate := "UPDATE person SET email=$2 WHERE first_name=$1"

	// postgresqlDelete := "DELETE FROM person WHERE first_name=$1"

	// tx.Commit()

	p := []Person{}

	// db.Exec(postgresqlUpdate, "Olim", "olim@olim.net")
	// db.Exec(postgresqlDelete, "Olim")

	err = db.Select(&p, "SELECT * FROM person")

	fmt.Println(p)

}
