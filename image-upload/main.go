package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/nfnt/resize"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "9m61t1A-"
	dbname   = "productmanagement"
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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Uploading File")
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("fileUp.html")
		t.Execute(w, nil)
	}

	//*1. parse input
	r.ParseMultipartForm(10 << 20)

	//*2. retrieve file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//*scaling
	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	m := resize.Resize(480, 854, img, resize.Lanczos3)

	out, err := os.Create("test_resized2.jpeg")
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer out.Close()

	// write new image to file
	var opt jpeg.Options
	opt.Quality = 80

	jpeg.Encode(out, m, nil)

	//*3. write temporrary file on our server
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	//!database
	fbytes := base64.StdEncoding.EncodeToString(fileBytes)
	result, err := db.Exec("CALL save_images($1,$2)", fbytes, "fidan")
	if err != nil {
		log.Fatal(err)
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("Etkilenen Kayıt Sayısı(%d)", rowsAffected)

	//*4. return result
	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func main() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":2020", nil)
}
