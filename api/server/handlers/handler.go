package handlers

import (
	"github.com/ebladrocher/smarthome/system/store"
	"net/http"
)

// Controllers ...
type Handlers struct {
	Index *HandlerIndex
	User  *HandlerUser
}

// NewHandlers ...
func NewHandlers(
	uc store.UseCase,
) *Handlers {

	return &Handlers{
		Index: NewHandlerIndex(),
		User:  NewHandlerUser(uc),
	}
}

// enableCors ...
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

/*func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	setupResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}3

	// process the request...
}
*/
