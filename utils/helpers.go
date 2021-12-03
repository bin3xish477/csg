package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/bin3xish477/csg/models"
	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

func StringInSlice(str string, slice []string) bool {
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

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func PrintTable(c []*models.Credential) {
	if len(c) == 0 {
		fmt.Printf("[%s-%s] no credentials matched specified options...\n", Yellow, End)
		return
	}

	t := table.NewWriter()

	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.Style().Format.Header = text.FormatTitle

	t.AppendHeader(
		table.Row{
			fmt.Sprintf("%sHost%s", Red, End),
			fmt.Sprintf("%sApp%s", Yellow, End),
			fmt.Sprintf("%sUser%s", Green, End),
			fmt.Sprintf("%sPassword%s", Purple, End),
		})

	for _, cred := range c {
		t.AppendRow([]interface{}{cred.Host, cred.App, cred.User, cred.Password})
	}

	t.Render()
}
