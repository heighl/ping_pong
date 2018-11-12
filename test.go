package ping_pong

import (
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	c := chi.NewRouter()
	c.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8000", c)
}
