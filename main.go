package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":7777", MakeRouter())
}
