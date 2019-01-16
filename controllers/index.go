package controllers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello, World!"))
}
func Status(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("UP"))
}
