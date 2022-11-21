package main

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Index Page"))

	x := r.URL.Path[1:] //localhost:8080/ sonrasına yazılanları sayfada gösterir
	data := "Merhaba" + x
	w.Write([]byte(data))
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/index", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil) //ilk parametre adres (localhost 8080 portu üzerinde çalışacak)
}

/*
* NOTLAR *
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Merhaba Mars!"))
	})
? HandleFunc(), ne yapacığımızı belirtiyoruz
? ListenAndServe(), yaptığımız işi ayağa kaldırıyoruz

? ResponseWriter, sunucudan nesneye yazma
? ResponseWrite ile istediğimiz veriyi göndeririz, yani bir sınır olmadığı için veriyi byte çevirip göndererek genelleme yaparız.

? Handle'lar bir sayfa gibi çalıştırılır.


*/
