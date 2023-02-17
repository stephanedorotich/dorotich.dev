package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type frac_smpl_handler struct {
	mu			sync.Mutex // guards n_correct, n_attempt
	n_correct	int
	n_attempt  	int
}

type frac_sol struct {
	n0	int
	d0	int
	n1	int
	d1	int
}

func RandomFrac() (int, int) {
	cf	:= rand.Intn(11)+2 // common factor in [2, 12]
	a 	:= rand.Intn(12)+1 // a, b			in [1, 12] and a != b
	b	:= a
	for b == a {
		b = rand.Intn(12)+1
	}

	if (a < b) {
		return a*cf, b*cf
	} else {
		return b*cf, a*cf
	}
}

func GCD(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func ValidateSol(s * frac_sol) bool {
	gcd := GCD(s.n0, s.d0)
	simplifiedNumerator := s.n0 / gcd
	simplifiedDenominator := s.d0 / gcd
	if (simplifiedNumerator == s.n1 && simplifiedDenominator == s.d1) {
		return true;
	} else {
		return false
	}
}

func (h *frac_smpl_handler) ValidateProblem(w http.ResponseWriter, r *http.Request) {
	var sol frac_sol
	var err error
	var correct string
	var text string
	var hidden string
	
	sol.n0, err = strconv.Atoi(r.PostFormValue("n0"))
	sol.d0, err = strconv.Atoi(r.PostFormValue("d0"))
	sol.n1, err = strconv.Atoi(r.PostFormValue("n1"))
	sol.d1, err = strconv.Atoi(r.PostFormValue("d1"))
	if (err != nil) {
		h.DeliverProblem(w, sol.n0, sol.d0)
		return
	}
	v := ValidateSol(&sol)
	if (v) {
		correct = "correct"
		text = "Correct 😊"
		hidden = "hidden"
	} else {
		correct = "incorrect"
		text = "Nope 😢"
		hidden = ""
	}
	t, _ := template.ParseFiles("templates/frac_smpl_feedback.html")
	t.Execute(w, struct{
		Correct string
		Text string
		Hidden string
		Numerator int
		Denominator int
	}{correct, text, hidden, sol.n0, sol.d0})
	return
}

func (h *frac_smpl_handler)	DeliverProblem(w http.ResponseWriter, n0 int, d0 int) {
	t, _ := template.ParseFiles("templates/frac_smpl.html")
	t.Execute(w, struct {
		Numerator int
		Denominator int
	}{n0, d0})
}

func (h *frac_smpl_handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		h.ValidateProblem(w, r)
		return
	}

	// Deliver new problem
	n0, d0 := RandomFrac();
	h.DeliverProblem(w, n0, d0)
	return
}
