package routers

import (
	"fmt"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	fmt.Printf("Start InitRoutes")

	router = SetUserRoute(router)
	return router
}
