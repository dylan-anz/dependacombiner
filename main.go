package main

import (
	"github.com/go-chi/chi"
	"github.com/spf13/cobra"
	"net/http"
)

func main() {
	cmd := &cobra.Command{}

	_ = cmd.Execute()

	router := chi.NewRouter()

	router.Get("/", func(_ http.ResponseWriter, _ *http.Request) {

	})
}
