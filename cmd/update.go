package cmd

import (
	"fmt"

	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update credential in database",
	Run: func(cmd *cobra.Command, args []string) {

		if cred.ID == 0 {
			fmt.Printf("[%serror%s] must a valid ID... use the `get` command to retrieve credential ID\n", utils.Red, utils.End)
			return
		}

		if cred.Host == "" && cred.App == "" && cred.User == "" && cred.Password == "" {
			fmt.Printf("[%serror%s] must specify at least one of the following options: `-i`, `-a`, `-u`, or `-p`\n", utils.Red, utils.End)
			return
		}

		db.Connect()
		cred.Update(cred.ID, db.DB)

		fmt.Printf("[%ssuccess%s] updated credential with ID of %s%d%s\n",
			utils.Green, utils.End, utils.Blue, cred.ID, utils.End)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().UintVarP(
		&cred.ID, "id", "d", 0, "the id of the credential to update [required]",
	)
	updateCmd.Flags().StringVarP(
		&cred.Host, "host", "i", "", "the new hostname/IP to set",
	)
	updateCmd.Flags().StringVarP(
		&cred.App, "app", "a", "", "the new application to set",
	)
	updateCmd.Flags().StringVarP(
		&cred.User, "user", "u", "", "the new username to set",
	)
	updateCmd.Flags().StringVarP(
		&cred.Password, "password", "p", "", "the new password to set",
	)

	updateCmd.MarkFlagRequired("id")
}
