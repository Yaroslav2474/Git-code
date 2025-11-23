package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func main() {

	db, _ := sql.Open("sqlite", "./HUI.db")

	defer db.Close()

	table := `
	CREATE TABLE IF NOT EXISTS heroes(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL, 
	hp INTEGER, 
	lvl INTEGER,
	gold INTEGER);`

	_, err := db.Exec(table)

	if err != nil {
		fmt.Print(err)
	}

	savesql := `
	INSERT INTO heroes(name, hp, lvl, gold)
	VALUES(?,?,?,?)`

	result, err := db.Exec(savesql, "Ярик", 100, 1, 50)

	if err != nil {
		fmt.Print(err)
	}

	id, _ := result.LastInsertId()

	fmt.Printf("ID героя %d\n", id)

	rows, _ := db.Query("SELECT id, name, hp, lvl, gold FROM heroes")
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var hp, lvl, gold int
		err = rows.Scan(&id, &name, &hp, &lvl, &gold)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Printf("ID: %d, Имя: %s, HP: %d, Уровень: %d, Золото: %d\n", id, name, hp, lvl, gold)
	}

}
