package cloudfunction

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	if rand.Float64() < 0.1 {
		w.WriteHeader(500)
		return
	}

	p := distuv.Gamma{
		Alpha: 2.0,
		Beta:  0.5,
	}

	sec := p.Rand()
	dur := float64(time.Second) * sec

	log.Printf("waiting %v seconds", dur)

	time.Sleep(time.Duration(dur))

	fmt.Fprint(w, html.EscapeString(fmt.Sprintf("i am slow, it took me %v seconds", dur)))
}
