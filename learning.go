package main

import (
	"learning/packages/routes"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	err := http.ListenAndServe("", router)
	if err != nil {
		panic(err)
	}
}
