package cmd

import (
	"fmt"

	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/utils"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get credential/s from database",
	Run: func(cmd *cobra.Command, args []string) {

		if cred.Host == "" && cred.App == "" && cred.User == "" {
			fmt.Printf("[%serror%s] must specify at least one filter with `-i`, `-a`, OR `-u` options\n", utils.Red, utils.End)
			return
		}

		db.Connect()
		creds := cred.Get(db.DB)

		utils.PrintTable(creds)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(
		&cred.Host, "host", "i", "", "the hostname/IP to filter for (use '*' to retrieve all credentials)",
	)
	getCmd.Flags().StringVarP(
		&cred.App, "app", "a", "", "the application to filter for",
	)
	getCmd.Flags().StringVarP(
		&cred.User, "user", "u", "", "the username to filter for",
	)
}
