package main

import (
	"os"

	"github.com/bin3xish477/csg/cmd"
	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/utils"
)

func init() {
	if !utils.FileExists(db.Folder) {
		os.MkdirAll(db.Folder, 0700)
	}
	if !utils.FileExists(db.Vault) {
		db.CreateDB()
	}
}

func main() {
	cmd.Execute()
}
