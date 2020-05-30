package handles

import (
	"encoding/json"
	"github.com/rmukubvu/pumpdata/store"
	"net/http"
)

func getPumpTypes(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := store.FetchPumpTypes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(generateErrorMessage("no types loaded")))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
