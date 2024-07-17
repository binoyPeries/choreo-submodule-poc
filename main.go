package main

import (
	"fmt"
	"log"
	"net/http"

	submodule "github.com/binoyPeries/choreo-submodule-poc/submodule"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/greeter", func(w http.ResponseWriter, r *http.Request) {
		helloValue := submodule.Hello()
		fmt.Fprintf(w, "Submodule function executed, response: %s", helloValue)

	})

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
