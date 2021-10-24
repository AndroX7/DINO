package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type animal struct {
	id         int
	animalType string
	nickname   string
	zone       int
	age        int
}

func main() {

	// connect to database
	db, err := sql.Open("mysql", "root:password@/dino")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// general query with arguments
	rows, err := db.Query("select * from Dino.animals where age > ?", 10)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err = rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		animals = append(animals, a)
	}

	fmt.Println(animals)

	row := db.QueryRow("select * from Dino.animals where id = ?", 2)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)
}
