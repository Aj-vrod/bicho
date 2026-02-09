package cmd

import (
	"Aj-vrod/bicho/internal/database"
	"Aj-vrod/bicho/pkg/config"
	"Aj-vrod/bicho/pkg/organization"
	"encoding/json"
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
	if err != nil {
		log.Fatal("Failed to load config from env with error: ", err)
	}

	fmt.Printf("Server listening in port %v...\n", cfg.HTTPPort)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get(ROOT_PATH, rootHandler)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get(ROOT_PATH, rootHandler)
		r.Get(ORGANIZATION_PATH, OrganizationHandler)
	})

	err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTPPort), r)
	if err != nil {
		log.Fatal("Failed to start the server with error: ", err)
	}
}

func rootHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func OrganizationHandler(w http.ResponseWriter, _ *http.Request) {
	// Get latest org data
	org, err := database.GetEmployees()
	if err != nil {
		log.Fatal("Failed to read org data with error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	// Parse data into tree structure
	treeOrg, err := organization.ProcessOrgData(org)
	if err != nil {
		log.Fatal("Failed to parse org data with error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	data, err := json.Marshal(treeOrg)
	if err != nil {
		log.Fatal("Failed to marshall org data with error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
