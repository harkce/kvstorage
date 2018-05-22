package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/harkce/kvstorage/keyval"
	"github.com/harkce/kvstorage/server"
	"github.com/subosito/gotenv"
)

func main() {
	log.Println("Starting kvstorage...")
	gotenv.Load(os.Getenv("GOPATH") + "/src/github.com/harkce/kvstorage/.env")

	keyval.InitKeyVal()

	router := server.Router()
	port := os.Getenv("PORT")

	log.Println(fmt.Sprintf("kvstorage started @:%s", port))

	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
