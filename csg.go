package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
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

var (
	vault string
	db    *sql.DB

	csgHelp = fmt.Sprintf(`usage: %scsg%s [command] [args]...

  [%sCommands%s]
    • %sadd%s - add credential to vault
    • %sdel%s - delete credential to vault
    • %sget%s - get credential from vault`,
		purple, end, red, end, yellow, end, yellow, end,
		yellow, end)

	addCmdHelp = fmt.Sprintf(`usage: %scsg%s %sadd%s [args]...

  [%sArgs%s]
    --%shost%s • the host this credential was discovered on (IP Address)
    --%sapp%s  • the application this credential belongs to
    --%suser%s • the username of this new credential
    --%spass%s • the password of this new credential (%srequired%s)`,
		purple, end, green, end, red, end, yellow, end,
		yellow, end, yellow, end, yellow, end, blue, end)

	delCmdHelp = fmt.Sprintf(`usage: %scsg%s %sdel%s [args]...

  [%sArgs%s]
    --%shost%s • delete credentials for a specific host
    --%sapp%s  • delete credentials for a specific application
    --%suser%s • delete credentials for a specific user`,
		purple, end, green, end, red, end, yellow, end,
		yellow, end, yellow, end)

	getCmdHelp = fmt.Sprintf(`usage: %scsg%s %sget%s [args]...

  [%sArgs%s]
    --%shost%s • the host to retrieve credentials for (%srequired%s)
    --%sapp%s  • the specific application to retrieve credentials for
    --%suser%s • the username to retrieve credentials for`,
		purple, end, green, end, red, end, yellow, end, blue, end,
		yellow, end, yellow, end)
)

func stringInSlice(str string, slice []string) bool {
	var temp []string
	for _, s := range slice {
		temp = append(temp, strings.ToLower(s))
	}

	for _, s := range temp {
		if s == str {
			return true
		}
	}
	return false
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func makeDB(csgFolder string) {
	f, err := os.Create(vault)
	if err != nil {
		log.Printf("unable to create directory in folder: %s", csgFolder)
	}
	defer f.Close()

	err = os.Chmod(vault, 0755)
	if err != nil {
		panic(err)
	}

	stmt := "CREATE TABLE `vault` ( `uid` INTEGER PRIMARY KEY AUTOINCREMENT, `host` VARCHAR(64) NULL, `app` VARCHAR(64) NULL, `user` VARCHAR(64) NULL, `passwd` VARCHAR(64) NULL);"
	fmt.Printf("executing \n%s\n", stmt)
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatalln("unable to create database")
	}
}

// addCred will add a credential to the vault
func addCred(host, app, user, passwd string) {
}

// delCred will delete a credential from the vault
func delCred(host, app, user, passwd string) {
}

// getCred will retrieve a credential from the vault
func getCred(host, app, user, passwd string) {
}

// runs before main() and checks to see if vault database
// has been created, if it hasn't, one will be created
func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	csgFolder := fmt.Sprintf("%s/.csg", homeDir)
	vault = fmt.Sprintf("%s/vault.db", csgFolder)
	if !fileExists(csgFolder) {
		os.MkdirAll(csgFolder, 0700)
	}
	if !fileExists(vault) {
		makeDB(vault)
	}
}

func main() {
	db, err := sql.Open("sqlite3", vault)
	if err != nil {
		log.Fatalln("unable to open connection to sqlite db...")
	}
	defer db.Close()

	//addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	//delCmd := flag.NewFlagSet("del", flag.ExitOnError)
	//getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	//addArgsHost := addCmd.String("host", "nil", "the host this credential was discovered on")
	//addArgsApp := addCmd.String("app", "nil", "the application this credential is for")
	//addArgsUser := addCmd.String("user", "nil", "the host this credential works on")
	//addArgsPass := addCmd.String("pass", "nil", "the host this credential works on")

	//delArgsHost := delCmd.Int("host", 0, "delete a host by id")

	//getArgsHost := getCmd.String("host", "nil", "the host to get credentials for")

	args := os.Args[1:]

	if len(args) < 1 || (args[0] != "add" && args[0] != "del" && args[0] != "get") {
		fmt.Println(csgHelp)
		return
	}

	if len(args) >= 2 {
		switch args[0] {
		}
	} else {
		switch args[0] {
		case "add":
			fmt.Println(addCmdHelp)
		case "del":
			fmt.Println(delCmdHelp)
		case "get":
			fmt.Println(getCmdHelp)
		}
	}
}
