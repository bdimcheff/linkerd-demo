package cloudfunction

import (
	"fmt"
	"html"
	"math/rand"
	"net/http"
	"time"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if rand.Float64() < 0.1 {
		w.WriteHeader(500)
		return
	}

	time.Sleep(time.Duration(rand.Int31n(10)) * time.Second)

	fmt.Fprint(w, html.EscapeString("i'm slow"))
}
