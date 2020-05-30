package handles

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type InternalError struct {
	Message string `json:"message"`
}

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1/pumps").Subrouter()
	api.HandleFunc("/pump-types", getPumpTypes).Methods(http.MethodGet)
	/*api.HandleFunc("/location/{id}", search).Methods(http.MethodGet)*/
	return r
}

func validJson(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func generateErrorMessage(e string) string {
	ie := InternalError{Message: e}
	buf, err := json.Marshal(ie)
	if err != nil {
		return string([]byte(fmt.Sprintf(`{"message": "%s"}`, e)))
	}
	return string(buf)
}
