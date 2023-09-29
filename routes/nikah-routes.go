package routes

import (
	"github.com/gorilla/mux"
	"github.com/josepnapitupulu/API_Create_Nikah/controllers"
)

var RegisterSidisRoutes = func(router *mux.Router) {
	// router.HandleFunc("/", controllers.Index).Methods("GET")
	// router.HandleFunc("/jemaatBaru", controllers.Create).Methods("POST")
	router.HandleFunc("/nikah/", controllers.CreateNikah).Methods("POST")
	router.HandleFunc("/nikah/", controllers.GetNikah).Methods("GET")
	router.HandleFunc("/nikah/{nikahId}", controllers.GetNikahById).Methods("GET")
	router.HandleFunc("/nikah/{nikahId}", controllers.UpdateNikah).Methods("PUT")
	router.HandleFunc("/nikah/{nikahId}", controllers.DeleteNikah).Methods("DELETE")
}
