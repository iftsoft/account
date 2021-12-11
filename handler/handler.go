package handler

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	hnd_BaseAPI        = "api"
	hnd_AccuntAPI      = "account"
	hnd_OperationAPI   = "operation"
	hnd_TransactionAPI = "transaction"
)

func ApiHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get root handler of API
		hnd := &rootHandler{}
		// Serve HTTP request
		hnd.ServeHTTP(w, r)
	})
}

///////////////////////////////////////////////////////////////////////
// Root API handler
type rootHandler struct {
	baseHandler // Common functions inheritance
}

func rootPanicRecover(w http.ResponseWriter) {
	if r := recover(); r != nil {
		text := fmt.Sprintf("Panic is recovered: %+v", r)
		fmt.Println(text)
		http.Error(w, text, http.StatusInternalServerError)
	}
}

// Serve HTTP request
func (this *rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer rootPanicRecover(w)
	fmt.Printf("Request - %s:%s\n", r.Method, r.URL.Path)

	list := strings.Split(r.URL.Path, "/")
	if len(list) > 2 {
		if list[1] != hnd_BaseAPI {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		switch list[2] {
		case hnd_AccuntAPI:
			hnd := getAccountApiHandler()
			hnd.ServeHTTP(w, r)
		default:
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
