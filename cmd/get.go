package cmd

import (
	"fmt"

	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/models"
	"github.com/bin3xish477/csg/utils"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get credential from database",
	Run: func(cmd *cobra.Command, args []string) {
		if cred.Host == "" && cred.App == "" && cred.User == "" {
			fmt.Printf("%serror%s: must specify at least one filter with `-i`, `-a`, OR `-u` options\n", utils.Red, utils.End)
		}
		var creds []*models.Credential
		db.Connect()
		db.DB.Where("host = ? OR app = ? OR user = ?", cred.Host, cred.App, cred.User).Find(&creds)
		utils.PrintTable(creds)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(
		&cred.Host, "host", "i", "", "the hostname/IP this credential is associated with",
	)
	getCmd.Flags().StringVarP(
		&cred.App, "app", "a", "", "the application this credential is associated with",
	)
	getCmd.Flags().StringVarP(
		&cred.User, "user", "u", "", "the username associated with this new credential",
	)
}
