package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	details "github.com/cjavadevw/go-microservices/details"

	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) { // Health End point for time
	log.Println("Checking Health Status")
	response := map[string]string{
		"Status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) { // Health End point for time
	log.Println("Serving the home page")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and Running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) { // Health End point for time
	log.Println("Fetching the Details")
	hostname, err := details.GetHostname()
	if err != nil {
		panic(err)
	}

	IP, _ := details.GetIP()
	fmt.Println(hostname, IP)

	response := map[string]string{
		"host": hostname,
		"ip":   IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler) // Health End point for time
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Println("Server has started.....")
	log.Fatal(http.ListenAndServe(":80", r))

}

/* USING gorillar mux
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

	})

	http.ListenAndServe(":80", r)
}
*/
/*package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
}
func main() {
	http.HandleFunc("/", rootHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Web server has started")
	http.ListenAndServe(":80", nil)
}
*/
/*package main

import (
	"fmt"

	geo "github.com/cjavadevw/go-microservices/geometry"

	"rsc.io/quote"
)

/*func rectProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := 2 * (length + width)
	return area, perimeter
}*/
/*
func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func main() {
	var x int = 8
	y, z := 2, 3
	var isWorking bool = false
	fmt.Println("Hello World")
	fmt.Println(quote.Go())
	fmt.Println(x, y, z, isWorking)

	a, p := rectProps(5, 6)
	fmt.Printf("The area is %f and the perimeter is %f\n", a, p)

	//var daysOftheMonth map[string]int
	//daysOftheMonth["Jan"] = 31

	var daysOftheMonth = map[string]int{"Jan": 31, "Feb": 28}
	fmt.Println(daysOftheMonth)

	area := geo.Area(5, 6)
	fmt.Println(area)
	diag := geo.Diagonal(4, 3)
	fmt.Println(diag)
}
*/
