package handler

import (
	"net/http"
	"strings"
	"fmt"
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

type rootHandler struct {
}

func rootPanicRecover(w http.ResponseWriter) {
	if r := recover(); r != nil {
		text := fmt.Sprintf("Panic is recovered: %+v", r)
		fmt.Println(text)
		http.Error(w, text, http.StatusInternalServerError)
	}
}

func (this *rootHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	defer rootPanicRecover(w)
	fmt.Printf("Request - %s:%s", r.Method, r.URL.Path)
	//	keeper := store.GetAccountKeeper()

	list := strings.Split(r.URL.Path, "/")
	if len(list) > 2 {
		switch list[1] {
		//case hnd_ObjectAPI, hnd_ManagerAPI :
		//	hnd := GetAuthHandler(this.store)
		//	hnd.ServeHTTP(w, r)
		//case hnd_TicketAPI :
		//	hnd := getTicketApiHandler(this.store)
		//	hnd.ServeHTTP(w, r)
		//case hnd_SecureAPI :
		//	hnd := getSecureApiHandler(this.store)
		//	hnd.ServeHTTP(w, r)
		default:
			http.Error(w, "Bad request", http.StatusBadRequest )
		}
	} else {
		http.Error(w, "Bad request", http.StatusBadRequest )
	}
}




