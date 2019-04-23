package main

import (
	"net/http"

	"github.com/gobuffalo/packr/v2"

	"github.com/go-chi/chi"
)

func main() {

	r := chi.NewRouter()

	assetBox := packr.New("assets", "./assets")

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		msg, _ := assetBox.FindString("message.txt")
		w.Write([]byte(msg))
	})

	http.ListenAndServe(":8080", r)

}
