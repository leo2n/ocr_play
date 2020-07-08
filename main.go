package main

import (
	"log"
	"net/http"
	"teletraan/handleImage"
	"teletraan/public"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.ico", public.FaviconHandler)
	mux.HandleFunc("/index", handleImage.Index)
	mux.HandleFunc("/ocr", handleImage.OCR)
	mux.HandleFunc("/query", handleImage.QueryOCR)
	log.Println("Listening on 127.0.0.1:4001")
	err := http.ListenAndServe("127.0.0.1:4001", mux)
	if err != nil {
		log.Println(err)
		return
	}
}
