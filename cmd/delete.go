package cmd

import (
	"fmt"

	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete credential from database",
	Run: func(cmd *cobra.Command, args []string) {

		if cred.ID == 0 {
			fmt.Printf("[%serror%s] must specify the ID of the credential to delete using the `-i` option\n", utils.Red, utils.End)
			fmt.Println("\tuse the `get` command to retrieve credential IDs")
			return
		}

		db.Connect()
		cred.Delete(db.DB)

		fmt.Printf("[%ssuccess%s] deleted credential with ID of : %s%d%s\n",
			utils.Green, utils.End, utils.Blue, cred.ID, utils.End)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().UintVarP(
		&cred.ID, "id", "i", 0, "the id of the credential to delete [required]\nuse the `get` command to retrieve credential IDs",
	)
}
