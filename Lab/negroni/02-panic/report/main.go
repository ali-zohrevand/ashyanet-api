package main

import (
	"net/http"

	"fmt"
	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		panic("oh no")
	})

	n := negroni.New()
	recovery := negroni.NewRecovery()
	recovery.PanicHandlerFunc = reportToSentry
	recovery.PrintStack = false
	n.Use(recovery)
	n.UseHandler(mux)

	http.ListenAndServe(":3003", n)
}

func reportToSentry(info *negroni.PanicInformation) {
	// write code here to report error to Sentry
	fmt.Println("Get Panic \n ==================== ")
	fmt.Println(&info)
}
