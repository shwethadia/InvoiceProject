package resources

import (
	"net/http"

	"github.com/gorilla/mux"

	ResourcesModel "github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/models/v1/resources"
	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/hyperledger"
)

func Destroy(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id := vars["id"]

		if err := ResourcesModel.Destroy(clients, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Success"))
	}
}
