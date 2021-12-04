package cmd

import (
	"fmt"

	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/utils"
	"github.com/spf13/cobra"
)

var (
	toPurge string
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "purge all credentials",
	Run: func(cmd *cobra.Command, args []string) {

		db.Connect()
		cred.Purge(db.DB, toPurge)

		fmt.Printf("[%ssuccess%s] purged credentials for host: %s%s%s\n",
			utils.Green, utils.End, utils.Blue, toPurge, utils.End)
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)

	purgeCmd.Flags().StringVarP(
		&toPurge, "target", "t", "*",
		fmt.Sprintf("specify a hostname/IP to delete all credentials for \n[%swarning%s] if left blank, all credentials will be purged!!", utils.Red, utils.End),
	)
}
