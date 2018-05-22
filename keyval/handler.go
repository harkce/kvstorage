package keyval

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct{}

func (h *Handler) Save(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	raw, _ := ioutil.ReadAll(r.Body)
	val := string(raw)

	keyVal.SetDefault(key, val)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, val)
	return
}

func (h *Handler) Retrieve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	val, found := keyVal.Get(key)
	if !found {
		val = ""
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, val)
	return
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")
	keyVal.Delete(key)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
	return
}
