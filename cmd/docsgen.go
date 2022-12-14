package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	"github.com/kloudlite/kl/lib/common"
	"github.com/spf13/cobra"
)

func code(str string, lang string) string {
	return fmt.Sprintf("```%s\n%s\n```", lang, str)
}

func getOptions(cmd *cobra.Command) string {

	return fmt.Sprintf(`### Options

%s`,
		code(
			fmt.Sprintf("%s%s",
				func() string {
					if strings.TrimSpace(cmd.Flags().FlagUsages()) == "" {
						return ""
					}
					return cmd.Flags().FlagUsages()
				}(), func() string {
					if cmd.Parent() != nil {
						return fmt.Sprintf("  -h, --help   help for %s", cmd.Name())
					}
					return ""
				}()),
			""))

}

func getSynopsis(cmd *cobra.Command) string {
	if cmd.Long != "" {
		return fmt.Sprintf(`### Synopsis

%s`, code(cmd.Long, ""))
	}
	return ""
}

func commandsList(cmd *cobra.Command) string {

	result := "### SEE ALSO\n"

	if cmd.Parent() != nil {
		parent := cmd.Parent()

		result = fmt.Sprintf("%s\n* [%s](%s.md)  - %s",
			result, parent.CommandPath(),
			strings.ReplaceAll(parent.CommandPath(), " ", "_"),
			parent.Short)
	}

	for _, c := range cmd.Commands() {
		result = fmt.Sprintf("%s\n* [%s](%s.md)  - %s",
			result, c.CommandPath(),
			strings.ReplaceAll(c.CommandPath(), " ", "_"),
			c.Short)
	}

	return result
}

func generateDocs(cmd *cobra.Command, dir string) error {

	fileContent := fmt.Sprintf(`## %s

%s

%s

%s

%s

###### Auto generated by kl CLI on %s
`, cmd.CommandPath(), cmd.Short, getSynopsis(cmd), getOptions(cmd), commandsList(cmd),
		fmt.Sprintf("%d-%s-%d", time.Now().Day(), time.Now().Month(), time.Now().Year()))

	// fmt.Println(fileContent)
	if err := ioutil.WriteFile(path.Join(dir, strings.ReplaceAll(cmd.CommandPath(), " ", "_")+".md"), []byte(fileContent), 0644); err != nil {
		return err
	}

	for _, c := range cmd.Commands() {
		generateDocs(c, dir)
	}

	return nil
}

func runDocGen(cmd *cobra.Command, _ []string) {

	if _, er := os.Stat("./docs"); errors.Is(er, os.ErrNotExist) {
		err := os.MkdirAll("./docs", 0644)
		if err != nil {
			common.PrintError(err)
			return
		}

	} else {
		err := os.RemoveAll("./docs")

		if err != nil {
			common.PrintError(err)
			return
		}

		err = os.MkdirAll("./docs", os.ModePerm)
		if er != nil {
			common.PrintError(err)
			return
		}
	}

	// if err := doc.GenMarkdownTree(cmd, "./docs"); err != nil {
	// 	common.PrintError(err)
	// }

	if err := generateDocs(cmd, "./docs"); err != nil {
		common.PrintError(err)
	}
}
