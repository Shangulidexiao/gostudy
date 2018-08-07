package ch7

import (
	"database/sql"
	"fmt"

	//mysql dirver
	_ "github.com/go-sql-driver/mysql"
)

type Db struct {
	db *sql.DB
}
type phone struct {
	name  string
	price int
	color string
}

func phoneList() {
	db, err := sql.Open("mysql", "go:go1234@/phone")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM phone")
	if err != nil {
		panic(err.Error())
	}
	for rows.Next() {
		var name, color string
		var id, price, delete int
		rows.Scan(&id, &name, &price, &color, &delete)
		fmt.Println(id, name, price, color, delete)
	}
}

func insertPhone() {
	db, err := sql.Open("mysql", "go:go1234@/phone")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtInsert, err := db.Prepare("INSERT INTO phone (`name`,`price`,`color`) values(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()
	stmtInsert.Exec("iphonex", 780000, "银色")
	stmtInsert.Exec("iphonex", 780000, "玫瑰金")
	stmtInsert.Exec("iphonex", 780000, "灰色")

}
