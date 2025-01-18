package https

import(
	"net/http"
	"fmt"
)

func ListenAndServe(addr string, handler Handler) error

type Handler interface {
	ServeHTTP(ResponseWriter *http.Request)
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}