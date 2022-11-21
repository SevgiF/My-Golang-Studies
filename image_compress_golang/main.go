package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nfnt/resize"
	"github.com/nickalie/go-webpbin"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hi u!")
	image := r.FormValue("image")
	idx := strings.Index(image, ";base64,")
	if idx < 0 {
		panic("InvalidImage")
	}
	imgType := image[11:idx]
	log.Println(imgType)

	unbased, err := base64.StdEncoding.DecodeString(image[idx+8:])
	if err != nil {
		panic("Cannot decode b64")
	}
	// file, err := os.Create("images/example1.png")
	// defer file.Close()
	// _, err1 := file.Write(unbased)
	// if err1 != nil {
	// 	panic("error")
	// }
	x := bytes.NewReader(unbased)

	switch imgType {
	case "png":
		img, err := png.Decode(x)
		if err != nil {
			panic("bad png")
		}
		//m := resize.Resize(320, 180, im, resize.Lanczos3)
		//mc, _ := pngquant.Compress(m, "1")

		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, img, nil); err != nil {
			panic("unable to decode jpeg")
		}
		newReader := bytes.NewReader(buf.Bytes())
		newImg, err := jpeg.Decode(newReader)
		if err != nil {
			panic("bad png to jpeg")
		}

		m := resize.Resize(480, 270, newImg, resize.Lanczos3)

		f, err := os.OpenFile("images/example2.jpeg", os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			panic("cannot open file")
		}
		defer f.Close()
		// f.Write(buf.Bytes())

		var opt jpeg.Options
		opt.Quality = 50
		jpeg.Encode(f, m, &opt)

		file, err := os.Create("images/image.webp")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		webpbin.Encode(file, m)

		// enc := png.Encoder{CompressionLevel: png.BestCompression}
		// enc.Encode(f, img)
	case "jpeg":
		img, err := jpeg.Decode(x)
		if err != nil {
			panic("bad jpeg")
		}

		m := resize.Resize(480, 854, img, resize.Lanczos3)

		f, err := os.OpenFile("images/example4.jpeg", os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			panic("cannot open file")
		}
		defer f.Close()

		var opt jpeg.Options
		opt.Quality = 50
		jpeg.Encode(f, m, &opt)

		file, err := os.Create("images/imageJPG.webp")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		webpbin.Encode(file, m)
	}

}

func getFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi g!")
}

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/upload", uploadFile)
	r.HandleFunc("/get", getFile)
	http.ListenAndServe(":2022", r)
}
