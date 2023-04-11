package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	AssignWebHandles()
	StartServer()
}

// -=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==--==-=-=-=-=-=-=-=--==-=-=-=-=-=-=-=- //

func AssignWebHandles() {
	http.HandleFunc("/file/", FileHandler)
	
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/cv", CVHandler)
	http.HandleFunc("/portfolio", PortfolioHandler)
	// http.HandleFunc("/practice/fraction-simplification", ServeHTTP_FracSmpl);
}

func StartServer() {
	rand.Seed(time.Now().UnixNano())
	if len(os.Args) == 1 {
		go RedirectHTTP()
		fmt.Println("Listening on https://44.232.31.52:443")
		log.Fatal(http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/dorotich.dev/fullchain.pem", "/etc/letsencrypt/live/dorotich.dev/privkey.pem", nil))
	} else if os.Args[1] == "test" {
		// Spin up on Local Host port 1331
		fmt.Println("Listening on https://localhost:1331")
		log.Fatal(http.ListenAndServe(":1331", nil))
	}
}

// -=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==--==-=-=-=-=-=-=-=--==-=-=-=-=-=-=-=- //

func RedirectHTTP() {
	err := http.ListenAndServe(":80", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
	}))
	if err != nil {
		log.Fatalf("Error re-routing HTTP to HTTPS: %v", err)
	}
}

func SendTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// -=-=-=-=-=-=-=-=-=-=-=-=--=-=-=-=-==--==-=-=-=-=-=-=-=--==-=-=-=-=-=-=-=- //

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		SendTemplate(w, "404.html", nil)
        return
    }
	SendTemplate(w, "home.html", nil)
}

func CVHandler(w http.ResponseWriter, r *http.Request) {
	SendTemplate(w, "cv.html", nil)
}

func PortfolioHandler(w http.ResponseWriter, r *http.Request) {
	SendTemplate(w, "portfolio.html", nil)
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
    // Get the requested file name and type from the request URL
	cleanpath := path.Clean(r.URL.Path)
    filename := path.Base(cleanpath)
	filetype := path.Base(path.Dir(cleanpath))

	// fmt.Println("file: " + filetype + "/" + filename)

	// Open the requested file
	file, err := os.Open("data/" + filetype + "/" + filename)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Serve the file
	http.ServeContent(w, r, filename, time.Time{}, file)
}
