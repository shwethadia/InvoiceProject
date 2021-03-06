package resources

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/models"

	ResourcesModel "github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/models/v1/resources"
	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/hyperledger"
)

func Update(clients *hyperledger.Clients) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var opts ResourcesModel.UpdateOpts
		var resource models.Resource
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		if err := decoder.Decode(&resource); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if r.Method == "PUT" {
			opts.Replace = true
		}

		updatedResource, err := ResourcesModel.Update(clients, id, &resource, &opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(updatedResource)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
