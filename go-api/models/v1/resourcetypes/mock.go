package resourcetypes

import (
	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/models"
)

var mockResourceTypes models.ResourceTypes

func init() {
	iron, _ := models.NewResourceType("Iron")
	copper, _ := models.NewResourceType("Copper")
	platinum, _ := models.NewResourceType("Platinum")

	mockResourceTypes = models.ResourceTypes{
		*iron,
		*copper,
		*platinum,
	}
}
