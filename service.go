package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/3rdparty", ThirdPartyHandler)
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "hello world")
}

func ThirdPartyHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	res, err := http.Get("http://us-central1-linkerd-external-demo.cloudfunctions.net/slow")
	end := time.Now()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dur := end.Sub(start)

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, fmt.Sprintf(`{"duration": %d}`, dur))
}
