package server

import (
	"net/http"

	"github.com/goware/cors"
	"github.com/harkce/kvstorage/keyval"
	"github.com/julienschmidt/httprouter"
)

func Router() http.Handler {
	router := httprouter.New()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         86400,
	})

	kvHandler := keyval.Handler{}
	router.POST("/store/:key", kvHandler.Save)
	router.GET("/get/:key", kvHandler.Retrieve)
	router.DELETE("/remove/:key", kvHandler.Delete)

	return cors.Handler(router)
}
