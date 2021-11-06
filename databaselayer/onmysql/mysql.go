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

	password := "aisgamming123#"
	// connect to database
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@/dino", password))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// general query with arguments
	rows, err := db.Query("select * from dino.animals where age > ?", 10)
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
		animals = append(animals, a)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(animals)

	// query single row
	row := db.QueryRow("select * from dino.animals where id = ?", 3)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)

	// // insert row
	// result, err := db.Exec("Insert into dino.animals (animal_type,nickname,zone,age) values ('Carnotosaurus', 'Carno2', ?, ?)", 3, 22)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result.LastInsertId())
	// fmt.Println(result.RowsAffected())

	//update fields which id is 3
	age := 16
	id := 3
	result, err := db.Exec("update dino.animals set age = ? where id = ?", age, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId())
	fmt.Println(result.RowsAffected())

}
