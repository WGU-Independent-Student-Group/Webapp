package main

import (
	"net/http"
)

type RequestedMaterialIDs struct {
	Materials []int `json:"materials"`
}

func (app *application) sendRequestedMaterials(w http.ResponseWriter, r *http.Request) {
	var req RequestedMaterialIDs

	err := app.readJSON(w, r, &req)

	if err != nil {
		app.handleDecodeError(w, err)
		return
	}

	// database querying logic here
	//
	// if err = app.writeJSON(w, http.StatusOK, response, nil); err != nil {
	//     app.logger.Error("Error marshalling JSON", "error", err)
	// }

}

func (app *application) sendAllMaterials(w http.ResponseWriter, r *http.Request) {
	var req RequestedMaterialIDs

	err := app.readJSON(w, r, &req)

	if err != nil {
		app.handleDecodeError(w, err)
		return
	}

	// database querying logic here
	//
	// if err = app.writeJSON(w, http.StatusOK, response, nil); err != nil {
	//     app.logger.Error("Error marshalling JSON", "error", err)
	// }

}
