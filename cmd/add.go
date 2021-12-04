package cmd

import (
	"fmt"

	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/utils"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add credential to database",
	Run: func(cmd *cobra.Command, args []string) {

		db.Connect()
		cred.Save(db.DB)

		fmt.Printf("[%ssuccess%s] added credential to host: %s%s%s\n",
			utils.Green, utils.End, utils.Blue, cred.Host, utils.End)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(
		&cred.Host, "host", "i", "", "the hostname/IP this credential is associated with [required]",
	)
	addCmd.Flags().StringVarP(
		&cred.App, "app", "a", "", "the application this credential is associated with",
	)
	addCmd.Flags().StringVarP(
		&cred.User, "user", "u", "", "the username associated with this new credential",
	)
	addCmd.Flags().StringVarP(
		&cred.Password, "password", "p", "", "the password associated with this new credential [required]",
	)

	addCmd.MarkFlagRequired("host")
	addCmd.MarkFlagRequired("password")
}
