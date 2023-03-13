package main

import (
	"fmt"
	d "go-mysql/db"
)

type Values struct {
	Name string
	Desc string
}

func main() {
	db := d.GetDB()
	defer d.CloseDB(db)

	value := Values{
		Name: "bilgisayar",
		Desc: "oyun oynarsın",
	}

	_, err := db.Query("INSERT INTO tabloo(name, `desc`) VALUES (?, ?)", value.Name, value.Desc)

	if err != nil {
		panic(err)
	}
	fmt.Print("Successfully inserted")

}
