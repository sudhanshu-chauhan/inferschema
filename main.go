package main

import (
	"net/http"
	"sudhanshu-exercise/app"
)

func main() {
	router := app.GetRouter()
	http.ListenAndServe(":8000", router)

}
