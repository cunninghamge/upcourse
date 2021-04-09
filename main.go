package main

import (
	"fmt"
	"net/http"
)

func Server(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Nursing 101")
}

func main() {}
