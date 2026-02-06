package cmd

import (
	"Aj-vrod/bicho/pkg/config"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

func NewApiCmd(version string) *cobra.Command {

	return &cobra.Command{
		Use:   "api",
		Short: "Amazing and luxurious CLI",
		Run: func(cmd *cobra.Command, args []string) {
			startServer()
		},
		Version: version,
	}
}

func startServer() {
	cfg, err := config.LoadFromEnv()
	fmt.Printf("Server listening in port %v...\n", cfg.HTTPPort)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", rootHandler)

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTPPort), r)
	if err != nil {
		fmt.Println("shutting down because of error below")
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("it has reached roothandler")
	w.Write([]byte("welcome"))
}
