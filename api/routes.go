package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/materials/list", sendRequestedMaterials)
	return mux
}
