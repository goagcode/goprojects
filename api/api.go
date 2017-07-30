package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if resJSON, err := json.Marshal(data); err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(resJSON)
	}
}

func errJSON(w http.ResponseWriter, err string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("{ error: %s }", err)))
}

// GetPage gets a single page from the API
func GetPage(w http.ResponseWriter, r *http.Request) {
	pageId := mux.Vars()["pageId"]
	data, err := cms.GetPage(pageId)
	if err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, data)
}

// CreatePage creates a new post or page
func CreatePage(w http.ResponseWriter, r *http.Request) {
	page := new(cms.Page)
	if data, err := ioutil.ReadAll(r.Body); err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.Unmarshal(data, page)
	if id, err := cms.CreatePage(page); err != nil {
		errJSON(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]int{
		"user_id": id,
	})
}
