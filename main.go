package main

import (
	"inferschema/app"
	"net/http"
)

func main() {
	router := app.GetRouter()
	http.ListenAndServe(":8000", router)

}
