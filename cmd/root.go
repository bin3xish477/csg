package cmd

import (
	"github.com/bin3xish477/csg/models"
	"github.com/spf13/cobra"
)

var (
	cred models.Credential
)

var rootCmd = &cobra.Command{
	Use:     "csg",
	Version: "1.0",
	Short:   "Store/organize credentials found during a CTF/engagement",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		return
	}
}
