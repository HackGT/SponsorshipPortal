package sample

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Sample_1(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Write([]byte(fmt.Sprintf("Sample_1 handler, param: %v", vars["param"])))
}

func Sample_2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sample controller, Sample_2, catch-all GET handler!"))
}

func Load(r *mux.Router) {
	r.HandleFunc("/test/{param}", Sample_1)
	r.Methods("GET").HandlerFunc(Sample_2)
}
