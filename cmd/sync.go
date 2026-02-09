package cmd

import (
	"Aj-vrod/bicho/internal/database"
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

	err := database.SyncOrgWithDB(ORGANIZATION_PATH)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Sync was completed!")
}
