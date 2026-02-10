package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/rabocse/lenslocked2/controllers"
	"github.com/rabocse/lenslocked2/views"
)

func main() {

	r := chi.NewRouter()

	tpl, err := views.Parse(filepath.Join("template/home.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("template/contact.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("template/faq.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
