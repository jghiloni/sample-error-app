package main

import (
	"fmt"
	"net/http"
	"os"

	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
)

func doRequest(w http.ResponseWriter, r *http.Request) {
	num := rand.Float64()

	var threshold float64
	var err error

	threshStr := os.Getenv("ERROR_THRESHOLD")
	if threshold, err = strconv.ParseFloat(threshStr, 64); err != nil || (threshold < 0 || threshold > 1) {
		threshold = 0.9
	}

	fmt.Printf("The threshold is %v and the random number is %v\n", threshold, num)

	if num <= threshold {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		w.Write([]byte("The request was a success 😁"))
		return
	} else {
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		w.Write([]byte("There was an unexpected error 😭"))
	}
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(doRequest)

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
