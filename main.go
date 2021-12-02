package main

import (
	"fmt"
	"os"

	"github.com/bin3xish477/csg/cmd"
	"github.com/bin3xish477/csg/db"
	"github.com/bin3xish477/csg/utils"
)

const (
	red    = "\u001b[91m"
	green  = "\u001b[32m"
	blue   = "\u001b[94m"
	yellow = "\u001b[33m"
	purple = "\u001b[35m"
	end    = "\u001b[0m"
	bold   = "\u001b[1m"
	underL = "\u001b[4m"
)

// runs before main() and checks to see if vault database
// has been created, if it hasn't, one will be created
func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	csgFolder := fmt.Sprintf("%s/.csg", homeDir)
	vault := fmt.Sprintf("%s/vault.db", csgFolder)
	if !utils.FileExists(csgFolder) {
		os.MkdirAll(csgFolder, 0700)
	}
	if !utils.FileExists(vault) {
		db.CreateDB(csgFolder, vault)
	}
}

func main() {
	cmd.Execute()
	//db, err := sql.Open("sqlite3", vault)
	//if err != nil {
	//log.Fatalln("unable to open connection to sqlite db...")
	//}
	//defer db.Close()
}
