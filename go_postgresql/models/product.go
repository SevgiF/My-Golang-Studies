package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //* go'da standart libraryde bulunan database/sql kütüphanesiyle birlikte kullanılıcak  (yardımcı araç gibi)
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "9m61t1A-"
	dbname   = "productsdb"
)

var db *sql.DB

func init() {
	var err error
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode = disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", connString)
	// db.SetMaxIdleConns(5) //?inaktif zamanda boşta bulunabilecek connection sayısını veriririz. (conn'da çalıştıracak sorgu kalmadıysa yoksa o idle (boş) duruma düşer)
	// db.SetMaxOpenConns(10) //?herhangi bir zamanda eşzamanlı kaç tane açık connection olacağı
	// db.SetConnMaxIdleTime(1 * time.Second) //?bir connection ne kadar süre idle durumda kalabilir
	// db.SetConnMaxLifetime(30 * time.Second) //?bir connection ne kadar süreliğine açık kalabilir

	if err != nil {
		log.Fatal(err)
	}

}

type Product struct {
	ID                 int
	Title, Description string
	Price              float32
}

func InsertProduct(data Product) {
	result, err := db.Exec("INSERT INTO products(title, description, price) VALUES($1, $2, $3)", data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Etkilenen Kayıt Sayısı(%d)", rowsAffected)
}

func UpdateProduct(data Product) {
	result, err := db.Exec("UPDATE products SET title=$2, description=$3, price=$4 WHERE id=$1", data.ID, data.Title, data.Description, data.Price)
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Etkilenen Kayıt Sayısı(%d)", rowsAffected)
}

func GetProduct() {
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Records Found!")
			return
		}
		log.Fatal(err)
	}
	defer rows.Close()

	var products []*Product
	for rows.Next() {
		prd := &Product{}
		err := rows.Scan(&prd.ID, &prd.Title, &prd.Description, &prd.Price)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, prd) //* döngünün içerisinden elde edilen prd'yi products dizisine ekleyecek
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, pr := range products { //* index ve değerini gönderecek. index istemiyorum
		fmt.Printf("%d, %s , %s, $%.2f\n", pr.ID, pr.Title, pr.Description, pr.Price)
	}
}

func GetProductByID(id int) {
	var product string
	err := db.QueryRow("SELECT title FROM products WHERE id=$1", id).Scan(&product)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No product with that ID.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("Product is %s\n", product)
	}
}
