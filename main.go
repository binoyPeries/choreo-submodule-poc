package main

import (
	"fmt"

	submodule "github.com/choreo-poc/submodule-repo"
)

func main() {
	submodule.Hello()
	fmt.Println("Submodule function executed")

}
