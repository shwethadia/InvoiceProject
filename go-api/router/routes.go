package router

import (
	HomeHandler "github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/routes/home"
	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/models"
	StatusHandler "github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/routes/status"
)

func GetRoutes() models.Routes {
	return models.Routes{
		models.Route{Name: "Home", Method: "GET", Pattern: "/", HandlerFunc: HomeHandler.Index},
		models.Route{Name: "Status", Method: "GET", Pattern: "/status", HandlerFunc: StatusHandler.Index},
	}
}
