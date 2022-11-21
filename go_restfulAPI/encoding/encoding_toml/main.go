package main

import (
	"database/sql"
	. "encodingYAML/encoding_toml/models"
	"fmt"

	"github.com/BurntSushi/toml"
	_ "github.com/lib/pq"
)

func main() {
	// Example 1 (without database)
	var conf Config
	if _, err := toml.DecodeFile("./configurations/config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", conf)

	//Example 2 (with database)
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode = disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Password, conf.Database.Database)
	db, err := sql.Open("postgres", connString) //postgres'e göre bağlantıyı kullanacak
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var count int
	err = db.QueryRow("SELECT 5+5").Scan(&count)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)

}
