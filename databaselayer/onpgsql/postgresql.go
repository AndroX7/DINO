package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/dino?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connect")
	defer db.Close()

	// general query with arguments
	rows, err := db.Query("select * from animals where age > $1 ", 5) // $ instead of ?
	handlerwos(rows, err)

	// query a single row
	row := db.QueryRow("select * from animals where age > $1", 5)
	a := animal{}
	err = row.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)

	// // insert a row
	// result, err := db.Exec("Insert into animals(animal_type,nickname,zone,age) values('Carnotaurus','Carno',$1,$2)", 3, 22)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(result.LastInsertId()) // not supported here
	// fmt.Println(result.RowsAffected())

	// //update a row
	// res, err := db.Exec("update animals set age = $1 where id = $2", 16, 4)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res.LastInsertId())
	// fmt.Println(res.RowsAffected())

	//another way to update by using sql
	// var id int
	// var nickname string
	// db.QueryRow("update animals set age = $1 where id = $2 returning nickname", 34, 4).Scan(&id)
	// fmt.Println(id)

	// prepare queries to use them multiple times, this also improves performance
	fmt.Println("Prepared Statement")
	stmt, err := db.Prepare("select * from animals where age > $1")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	// try with age > 5
	rows, err = stmt.Query(5)
	handlerwos(rows, err)

	// try with age > 10
	rows, err = stmt.Query(10)
	handlerwos(rows, err)

	testTransaction(db)
}

func handlerwos(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animalType, &a.nickname, &a.zone, &a.age)
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
}

func testTransaction(db *sql.DB) {
	fmt.Println("Transactions...")
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("select * from animals where age > $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(15)
	handlerwos(rows, err)
	rows, err = stmt.Query(30)
	handlerwos(rows, err)
	results, err := tx.Exec("update animals set age = $1 where id = $2", 999, 4)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(results.LastInsertId()) // not supported here
	fmt.Println(results.RowsAffected())
	err = tx.Commit() // use to commit the changes and it won't be rollback when the rollback called because it's already commited
	if err != nil {
		log.Fatal(err)
	}
}
