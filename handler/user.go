package handler

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func UserIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
