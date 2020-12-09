package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/yeric17/inventory-system/config"
)

func init() {

}

//GetServer inicia el servidor y devuelve un router
func GetServer(router *mux.Router) *http.Server {

	corsOPS := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders: []string{
			"*",
		},
	})

	server := http.Server{
		Addr:         ":" + config.APIPort,
		Handler:      corsOPS.Handler(router),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return &server
}
