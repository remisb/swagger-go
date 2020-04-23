package main

import (
	"github.com/go-chi/chi"
	"github.com/swaggo/http-swagger" // http-swagger middleware
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	"net/http"
	"os"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Mathematic server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://jivalabs.com/support
// @contact.email bauzys@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1


func main() {
	if os.Getenv("DEBUG") != "" {

	}

	//cfg := struct {
	//	configFile string
	//}

	port := ":8080"
	handler := chi.NewRouter()
	handler.Get("/info", infoHandler)
	handler.Route("/api/v1", func(r chi.Router) {
		r.Get("/users", listUsers)
	})
	handler.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost" + port + "/swagger/doc.json")),
	)
	handler.Get("/api/v1/info", infoHandler)
	http.ListenAndServe(port, handler)
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test output"))
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List of articles"))
}
