package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
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

	if len(os.Args) == 1 {
		// Spin up HTTPS server (port 443)
		// Spint up HTTP redirect server (port 80)
		go RedirectHTTP()
		fmt.Println("Listening on https://44.232.31.52:443")
		log.Fatal(http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/dorotich.dev/fullchain.pem", "/etc/letsencrypt/live/dorotich.dev/privkey.pem", nil))
	} else if os.Args[1] == "test" {
		// Spin up on Local Host port 1331
		fmt.Println("Listening on https://localhost:1331")
		log.Fatal(http.ListenAndServe(":1331", nil))
	}
}
