package api

import (
	"encoding/json"
	"net/http"
)

// AllPages returns all pages
func AllPages(w http.ResponseWriter, r *http.Request) {
	data, err := cms.GetPages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, data)
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if resJSON, err := json.Marshal(data); err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(resJSON)
}
