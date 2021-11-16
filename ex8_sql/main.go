package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "gophercises"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)

	//db, err := sql.Open("postgres", psqlInfo)
	//must(err)
	//resetDB(db, dbname)
	//db.Close()

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	defer db.Close()
	must(err)

	must(createPhoneNumberTable(db))

	_, err = enterPhoneNumber(db, "1234567890")
	_, err = enterPhoneNumber(db, "123 456 7891")
	_, err = enterPhoneNumber(db, "(123) 456 7892")
	_, err = enterPhoneNumber(db, "(123) 456-7893")
	_, err = enterPhoneNumber(db, "123-456-7894")
	_, err = enterPhoneNumber(db, "123-456-7890")
	_, err = enterPhoneNumber(db, "1234567892")
	_, err = enterPhoneNumber(db, "(123)456-7892")

	must(CleanPhoneDB(db))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func resetDB(db *sql.DB, dbName string) error {
	_, err := db.Exec(`DROP DATABASE IF EXISTS ` + dbName)
	if err != nil {
		return err
	}

	_, err = db.Exec(`CREATE DATABASE ` + dbName)
	if err != nil {
		return err
	}

	return nil
}

func createPhoneNumberTable(db *sql.DB) error {
	dropStatement := `DROP TABLE IF EXISTS phone_numbers`
	createStatement := `CREATE TABLE IF NOT EXISTS phone_numbers (
	  id SERIAL,
	  phoneNumber VARCHAR(255)
	)`

	_, err := db.Exec(dropStatement)
	_, err = db.Exec(createStatement)
	return err
}

func enterPhoneNumber(db *sql.DB, number string) (int, error) {
	statement := `INSERT INTO phone_numbers (phoneNumber) VALUES ($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, number).Scan(&id)

	if err != nil {
		return -1, err
	}
	return id, err
}

func CleanPhone(number string) string {
	re := regexp.MustCompile("[0-9]+")

	output := strings.Join(re.FindAllString(number, -1), "")
	return output
}

func CleanPhoneDB(db *sql.DB) error {
	updStatement := `UPDATE phone_numbers SET phoneNumber = $1 WHERE id = $2;`
	delStatement := `DELETE FROM phone_numbers WHERE id = $1;`
	qStatement := `SELECT id, phoneNumber FROM phone_numbers`
	rows, err := db.Query(qStatement)
	if err != nil {
		return err
	}
	defer rows.Close()

	phoneList := make([]string, 0)

	for rows.Next() {
		var id int
		var phoneNumber string
		if err := rows.Scan(&id, &phoneNumber); err != nil {
			return err
		}

		cleanPhone := CleanPhone(phoneNumber)

		if contains(phoneList, cleanPhone) {
			_, err = db.Exec(delStatement, id)
			if err != nil {
				return err
			}

		} else {
			_, err = db.Exec(updStatement, cleanPhone, id)
			if err != nil {
				return err
			}

			phoneList = append(phoneList, cleanPhone)
		}

	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//func RemoveDuplicate (db *sql.DB) error {
//	delStatement := `DELETE FROM phone_numbers id = $1;`
//	qStatement := `SELECT id, phoneNumber FROM phone_numbers`
//	rows, err := db.Query(qStatement)
//	if err != nil {
//		return err
//	}
//	defer rows.Close()
//
//	for rows.Next() {
//		var id int
//		var phoneNumber string
//		if err := rows.Scan(&id, &phoneNumber); err != nil {
//			return err
//		}
//
//		_, err = db.Exec(delStatement, id)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
