package main

import (
	"fmt"
	"log"
	"net/http"

	sameOrgSubmodule "github.com/binoyPeries/choreo-submodule-poc/same-org-submodule"
	submodule "github.com/binoyPeries/choreo-submodule-poc/submodule"

	// pvtsubmodule "github.com/binoyPeries/choreo-submodule-poc/submodule-pvt"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/greeter", func(w http.ResponseWriter, r *http.Request) {
		helloValue := submodule.Hello()
		// workValue := pvtsubmodule.Work()
		ownSubmoduleValue := sameOrgSubmodule.HelloWorld()
		fmt.Fprintf(w, "Submodules function executed, responses are : %s && %s", helloValue, ownSubmoduleValue)

	})

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
