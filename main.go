package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rabocse/lenslocked2/controllers"
	"github.com/rabocse/lenslocked2/template"
	"github.com/rabocse/lenslocked2/views"
)

func main() {

	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(template.FS, "home.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(template.FS, "contact.gohtml"))))

	r.Get("/faq", controllers.StaticHandler(views.Must(views.ParseFS(template.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
