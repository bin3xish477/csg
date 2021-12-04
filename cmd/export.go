package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/utils"
	"github.com/spf13/cobra"
)

var fileName string

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export all credentials from database to CSV file",
	Run: func(cmd *cobra.Command, args []string) {

		db.Connect()
		creds := cred.Export(db.DB)

		now := time.Now()
		timeStamp := fmt.Sprintf("%d%02d%02d", now.Year(), now.Month(), now.Day())
		newFileName := strings.Replace(fileName, ".", fmt.Sprintf("-%s.", timeStamp), -1)
		f, err := os.Create(newFileName)
		if err != nil {
			fmt.Printf("[%serror%s] could not create file %s%s%s\n",
				utils.Red, utils.End, utils.Yellow, newFileName, utils.End)
			return
		}
		defer f.Close()

		cw := csv.NewWriter(f)

		for _, cred := range creds {
			err = cw.Write([]string{cred.Host, cred.App, cred.User, cred.Password})
			if err != nil {
				fmt.Printf("[%serror%s] could not write to file %s%s%s\n",
					utils.Red, utils.End, utils.Yellow, newFileName, utils.End)
				return
			}
		}

		fmt.Printf("[%ssuccess%s] exported all credential to file: %s%s%s\n",
			utils.Green, utils.End, utils.Blue, newFileName, utils.End)
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVarP(
		&fileName, "file", "f", "creds.csv", "file name to save credentials to",
	)
}
