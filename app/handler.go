package app

import (
	"encoding/json"
	"net/http"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 * 10 // 10 MB

type InferSchemaResponse struct {
	FileName string      `json:"fileName"`
	Schema   interface{} `json:"schema"`
}

func InferSchemaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	inferSchemaResponse := InferSchemaResponse{}
	// restricting file size to 10 MB
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if multipartParseErr := r.ParseMultipartForm(MAX_UPLOAD_SIZE); multipartParseErr != nil {
		http.Error(w, "max file size (10 mb) exceeded", http.StatusBadRequest)
		return
	}

	for key, _ := range r.MultipartForm.File {
		fileheader := r.MultipartForm.File[key]
		inferSchemaResponse.FileName = fileheader[0].Filename

		file, err := fileheader[0].Open()
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		data, err := InferDataSchema(file)
		if err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		inferSchemaResponse.Schema = data

		break

	}
	res, marshalErr := json.Marshal(inferSchemaResponse)
	if marshalErr != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
