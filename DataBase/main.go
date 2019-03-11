package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DBConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DBConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DBConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS erson(
		name STRING,
		age INT)`
	_, err := DBConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// cmd = "INSERT INTO person (name,age) VALUES (?,?)"
	// _, err = DBConnection.Exec(cmd, "Mike", 22)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DBConnection.Exec(cmd, 32, "Mike")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// sqlの中身を抽出する
	// cmd = "SELECT * FROM person"
	// rows, _ := DBConnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// cmd = "SELECT * FROM person where age = ?"
	// row := DBConnection.QueryRow(cmd, 20)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err == sql.ErrNoRows {
	// 	log.Println("No Row")
	// } else {
	// 	log.Println(err)
	// }
	// fmt.Println(p.Name, p.Age)

	// cmd = "DELETE FROM person WHERE name = ?"
	// _, err = DBConnection.Exec(cmd, "Boby")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}
