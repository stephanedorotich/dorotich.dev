package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func RedirectHTTP() {
	err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
	}))
	if err != nil {
		log.Fatalf("Error re-routing HTTP to HTTPS: %v", err)
	}
}

func main() {
    rand.Seed(time.Now().UnixNano())

	http.Handle("/practice/fraction-simplification", new(frac_smpl_handler));
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	go RedirectHTTP()

	fmt.Println("Listening on http://44.232.31.52:443")
	err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/dorotich.dev/fullchain.pem", "/etc/letsencrypt/live/dorotich.dev/privkey.pem", nil)
	if err != nil {
		log.Fatalf("Error listening for HTTPS traffic: %v", err)
	}
}
