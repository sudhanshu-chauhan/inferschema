package app

import "github.com/gorilla/mux"

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/schema", InferSchemaHandler).Methods("Post")
	return router
}
