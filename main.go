package main

import (
	"fmt"
	"net/http"
	"os"

	"math/rand"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func buildRequestHandler(rng *rand.Rand, minDelay int64, maxDelay int64) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		num := rng.Float64()

		var threshold float64
		var err error

		threshStr := os.Getenv("SUCCESS_RATE")
		if threshold, err = strconv.ParseFloat(threshStr, 64); err != nil || (threshold < 0 || threshold > 1) {
			threshold = 0.9
		}

		fmt.Printf("The threshold is %v and the random number is %v\n", threshold, num)

		delay := minDelay
		if maxDelay > 0 {
			delay += rng.Int63n(maxDelay - minDelay)
		}

		fmt.Printf("Adding a randomly selected %v ms delay\n", delay)
		time.Sleep(time.Duration(delay) * time.Millisecond)

		if num <= threshold {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
			w.Write([]byte("The request was a success 😁"))
			return
		}

		w.WriteHeader(500)
		w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
		w.Write([]byte("There was an unexpected error 😭"))
		return
	}
}

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	minDelay, err := strconv.ParseInt(os.Getenv("MIN_DELAY_MS"), 10, 64)
	if err != nil || minDelay < 0 {
		minDelay = 0
	}

	maxDelay, err := strconv.ParseInt(os.Getenv("MAX_DELAY_MS"), 10, 64)
	if err != nil || maxDelay < minDelay {
		maxDelay = minDelay
	}

	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(buildRequestHandler(rng, minDelay, maxDelay))

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
