package main

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()
	ar := mux.NewRouter()

	mw := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	r.HandleFunc("/api/without-auth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("no auth required\n"))
	}).Methods("GET")

	ar.HandleFunc("/api/with-auth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("auth required\n"))
	}).Methods("GET")

	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(ar))
	r.PathPrefix("/api").Handler(an)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":7080")
}
