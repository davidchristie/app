package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kelseyhightower/envconfig"
)

const (
	accessToken = "<ACCESS_TOKEN>"
	code        = "<CODE>"
)

type config struct {
	Port string `required:"true"`
}

func main() {
	config := config{}
	envconfig.MustProcess("", &config)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/authorize", authorizeHandler())
	router.Post("/access_token", accessTokenHandler())
	router.Get("/github/emails", serveJSONFile("./services/mock-oauth/data/github/emails.json"))
	router.Get("/github/user", serveJSONFile("./services/mock-oauth/data/github/user.json"))
	router.Get("/google/user", serveJSONFile("./services/mock-oauth/data/google/user.json"))
	fmt.Printf("Port: %v\n", config.Port)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%v", config.Port),
		Handler:      router,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func accessTokenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{"access_token":"%s"}`, accessToken)
	}
}

func authorizeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := r.URL.Query().Get("state")
		redirectURI := r.URL.Query().Get("redirect_uri")
		url := fmt.Sprintf("%s?state=%s&code=%s", redirectURI, state, code)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}

func serveJSONFile(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(data)
	}
}
