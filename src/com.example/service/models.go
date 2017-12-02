package service

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Snack struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type SnackList struct {
	Snacks []Snack `json:"snackList"`
}

func ListSnacks(db *sql.DB) SnackList {
	sqlQ := "SELECT * from snacks"
	rows, err := db.Query(sqlQ)
	result := SnackList{}
	if err != nil {
		log.Println(err)
	} else {
		defer rows.Close()

		for rows.Next() {
			snack := Snack{}
			err := rows.Scan(&snack.Name, &snack.Quantity)

			if err != nil {
				log.Println(err)
			} else {
				result.Snacks = append(result.Snacks, snack)
			}
		}
	}
	return result
}

func SaveSnack(db *sql.DB, snack Snack) (int64, error) {
	log.Println(snack.Name)
	sqlS := "SELECT * from snacks where name='" + snack.Name + "'"

	result, err := db.Query(sqlS)

	if err != nil {
		log.Panic(err)
	}
	snackR := Snack{}
	if result.Next() {
		result.Scan(&snackR.Name, &snackR.Quantity)
	}
	result.Close()
	log.Printf("Test %v", snackR)
	//Snack literal to check for empty snack
	if (Snack{}) == snackR {
		log.Println("Snack " + snack.Name + " not available in database")

		//Prepare statement replaces ?
		sqlQ := "INSERT INTO snacks VALUES(?,?)"

		stmt, err := db.Prepare(sqlQ)

		if err != nil {
			log.Panic(err)
		}
		ins, err2 := stmt.Exec(snack.Name, snack.Quantity)
		if err2 != nil {
			log.Panic(err2)
		}
		return ins.RowsAffected()
	} else {
		quantity := snackR.Quantity
		sqlU := "UPDATE snacks SET quantity=? where name=?"
		quantity += snack.Quantity

		stmt, err3 := db.Prepare(sqlU)
		if err3 != nil {
			log.Panic(err3)
		}

		up, err4 := stmt.Exec(quantity, snackR.Name)
		if err4 != nil {
			log.Panic(err4)
		}
		return up.RowsAffected()
	}
}

func DeleteSnack(db *sql.DB, name string) (string, error) {
	sqlD := "DELETE FROM snacks WHERE name=?"
	stmt, err := db.Prepare(sqlD)

	if err != nil {
		log.Panic(err)
	}
	_, err = stmt.Exec(name)

	if err != nil {
		log.Panic(err)
	}
	return name, nil
}

func DeleteAll(db *sql.DB) (int64, error) {
	sqlD := "DELETE FROM snacks"
	del, err := db.Exec(sqlD)

	if err != nil {
		log.Panic(err)
	}
	return del.RowsAffected()
}
