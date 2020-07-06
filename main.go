package main

import (
	"log"
	"net/http"
	"teletraan/handleImage"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ocr", handleImage.OCR)
	log.Println("Listening on 127.0.0.1:4001")
	err := http.ListenAndServe("127.0.0.1:4001", mux)
	if err != nil {
		log.Println(err)
		return
	}
}
