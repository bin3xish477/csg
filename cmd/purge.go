package cmd

import (
	"fmt"
	"strings"

	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/utils"
	"github.com/spf13/cobra"
)

var (
	toPurge string
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "purge credentials by host",
	Run: func(cmd *cobra.Command, args []string) {

		var confirm string
		if toPurge == "*" {
			fmt.Printf("[%sATTENTION%s] you are about to delete all credentials.\n", utils.Red, utils.End)
			fmt.Printf("[%sconfirm%s] would you like to proceed (y/n): ", utils.Blue, utils.End)
			fmt.Scanln(&confirm)
			if strings.Replace(confirm, `\n`, "", -1) == "y" {
				db.Connect()
				cred.Purge(db.DB, toPurge)

				fmt.Printf("[%ssuccess%s] purged credentials for host: %s%s%s\n",
					utils.Green, utils.End, utils.Blue, toPurge, utils.End)
			} else {
				fmt.Printf("[%scancelled%s] operation\n", utils.Green, utils.End)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)

	purgeCmd.Flags().StringVarP(
		&toPurge, "target", "t", "*",
		fmt.Sprintf("specify a hostname/IP to delete all credentials for \n[%swarning%s] if left blank, all credentials will be purged!!", utils.Red, utils.End),
	)
}
