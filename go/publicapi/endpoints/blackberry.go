package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/abteilung6/tilmad/go/publicapi/model"
)

func GetBlackberries(w http.ResponseWriter, r *http.Request) {
	blackberries := []model.Blackberry{
		{ID: 1, Name: "Triple Crown"},
		{ID: 2, Name: "Black Satin"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blackberries)
}
