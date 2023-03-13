package main

import (
	"io"
	"net/http"
)

func main() {
	var i ironman
	var w wolverine

	mux := http.NewServeMux() //? her seferinde http.HandleFunc yani http nesnesi yazmak yerine mux nesnesi üzerinden yazdırıyoruz
	mux.Handle("/ironman", i)
	mux.Handle("/wolverine", w)

	http.ListenAndServe(":8080", mux) //http nesnesi üzerinden yapsakdık nil yazıcaktık.
}

type ironman int

// ServeHTTP kafadan verilmiş isim değil. Bu şekilde yazılması gerekiyor.
// HAndle nesne üzerinde ServeHTTP metodunu arar.
func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Iron!")
}

type wolverine int

func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. Wolverine!")
}
