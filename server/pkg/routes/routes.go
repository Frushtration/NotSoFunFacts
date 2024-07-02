package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"server/pkg/database"
)

var Serve http.Handler

func init() {
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/question", getQuestion).Methods("GET")
	//r.HandleFunc("/question", question).Methods("POST")
	r.HandleFunc("/source", getSource).Methods("GET")
	//r.HandleFunc("/source", source).Methods("POST")

	Serve = r
}

func getQuestion(w http.ResponseWriter, r *http.Request) {
	facts, err := database.GetRandQuestion()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, facts)
}

func getSource(w http.ResponseWriter, r *http.Request) {
	source, err := database.GetRandQuestion()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprint(w, source)
}

// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "home\n")
// }

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
