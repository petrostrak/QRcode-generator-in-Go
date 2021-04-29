package main

import (
	"image/png"
	"log"
	"net/http"
	"text/template"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

type Page struct {
	Title string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	p := Page{
		Title: "QR code Generator",
	}

	t, err := template.ParseFiles("generator.html")
	if err != nil {
		log.Fatal("cannot parse html file", err)
	}

	t.Execute(w, p)
}

func CodePage(w http.ResponseWriter, r *http.Request) {
	dataString := r.FormValue("datastring")

	qrCode, err := qr.Encode(dataString, qr.L, qr.Auto)
	if err != nil {
		log.Fatal("cannot generate QR code from given text")
	}

	qrCode, err = barcode.Scale(qrCode, 512, 512)
	if err != nil {
		log.Fatal("cannot scale generated QR code image")
	}

	png.Encode(w, qrCode)
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/generator/", CodePage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
