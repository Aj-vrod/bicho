package cmd

import (
	"Aj-vrod/bicho/pkg/config"
	"Aj-vrod/bicho/pkg/organization"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

func NewSyncCmd(version string) *cobra.Command {
	return &cobra.Command{
		Use:   "sync",
		Short: "Sync org data to db",
		Run: func(cmd *cobra.Command, args []string) {
			SyncOrgToDB()
		},
		Version: version,
	}
}

func SyncOrgToDB() {
	log.Println("Syncing has started")

	// Connect ith DB
	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal("Failed to load config from env with error: ", err)
	}
	db, err := sql.Open("postgres", cfg.DBDNS)

	// Check connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping DB with error: ", err)
	}
	log.Println("Successfully connected to db")

	defer db.Close()

	// Get list of employees from file
	employees, err := organization.ReadOrgData(ORG_FILE_PATH)

	// Insert employees into db
	for _, e := range employees {
		// Preparing INSERT query
		query, err := db.Prepare("INSERT INTO employees(name, country, job_family, job_title, business_unit, squad, platoon, battalion, start_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);")
		if err != nil {
			log.Fatal("Failed to prepare query with error: ", err)
		}
		defer query.Close()

		result, err := query.Exec(e.Name, e.Country, e.JobFamily, e.JobTitle, e.BusinessUnit, e.Squad, e.Platoon, e.Battalion, e.StartDate)
		if err != nil {
			log.Fatal("Failed to execute query with error: ", err)
		}
		i, err := result.RowsAffected()
		if err != nil {
			log.Fatal("Failed to get list of rows affected with error: ", err)
		}

		log.Printf("Inserted %v row\n", i)
	}
}
