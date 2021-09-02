package resources

import (
	"encoding/json"
	"net/http"

	ResourcesModel "github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/models/v1/resources"
	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/hyperledger"
)

func Index(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		resources, err := ResourcesModel.Index(clients)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		packet, err := json.Marshal(resources)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(packet)
	}
}
