package cloudfunction

import (
	"fmt"
	"html"
	"net/http"
	"time"
)

// HelloWorld prints the JSON encoded "message" field in the body
// of the request or "Hello, World!" if there isn't one.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	fmt.Fprint(w, html.EscapeString("i'm slow"))
}
