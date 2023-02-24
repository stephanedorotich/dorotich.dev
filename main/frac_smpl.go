package main

import (
	"math/rand"
	"net/http"
	"strconv"
)

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

func GCF(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func ValidateSol(numbers []int) bool {
	gcd := GCF(numbers[0], numbers[1])
	simplifiedNumerator := numbers[0] / gcd
	simplifiedDenominator := numbers[1] / gcd
	return (simplifiedNumerator == numbers[2] && simplifiedDenominator == numbers[3])
}

func ValidateProblem(w http.ResponseWriter, r *http.Request) {
	// Parse original fraction (n0/d0) and submitted solution (n1/d1)
	// If any value is missing, or any value cannot be converted to an integer,
	// send a HTTP 400 Bad Request response

	var strings []string
	var numbers []int

    strings = append(strings, r.FormValue("n0"))
    strings = append(strings, r.FormValue("d0"))
    strings = append(strings, r.FormValue("n1"))
    strings = append(strings, r.FormValue("d1"))

	for _, str := range strings {
		n, err := strconv.Atoi(str)
		if err != nil {
			http.Error(w, "Invalid form values", http.StatusBadRequest)
			return
		}
		numbers = append(numbers, n)
	}

	// form contents are valid
	v := ValidateSol(numbers)
	if (v) {
		DeliverCorrect(w)
	} else {
		DeliverIncorrect(w, numbers[0], numbers[1])
	}
}

func DeliverProblem(w http.ResponseWriter, n0 int, d0 int) {
	data := struct {
		Numerator int
		Denominator int
	}{n0, d0}
	SendTemplate(w, "frac.html", data)
}

func DeliverCorrect(w http.ResponseWriter) {
	err := templates.ExecuteTemplate(w, "frac_smpl_correct.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeliverIncorrect(w http.ResponseWriter, n int, d int) {
	err := templates.ExecuteTemplate(w, "frac_smpl_incorrect.html", struct {
		Numerator int
		Denominator int
	}{n, d})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ServeHTTP_FracSmpl(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "POST") {
		ValidateProblem(w, r)
		return
	}

	// Deliver new problem
	n, d := RandomFrac();
	DeliverProblem(w, n, d)
	return
}
