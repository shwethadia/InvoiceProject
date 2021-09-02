package resourcetypes

import (
	"encoding/json"
	"os"

	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/hyperledger"
	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/models"
)

func Index(clients *hyperledger.Clients) (resourcetypes *models.ResourceTypes, err error) {
	resourcetypes = new(models.ResourceTypes)

	MSPID := os.Getenv("HYPERLEDGER_MSP_ID")
	if len(MSPID) == 0 {
		MSPID = "ibm"
	}

	res, err := clients.Query(MSPID, "mainchannel", "resource_types", "index", [][]byte{
		[]byte("{\"selector\":{ \"active\": { \"$eq\":true } }}"),
	})
	if err != nil {
		return
	}

	if err = json.Unmarshal(res, resourcetypes); err != nil {
		return
	}

	return
}
