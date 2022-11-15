/*
Copyright © 2022 Andres Puello apuello1025@gmail.com
*/
package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"regexp"

	"github.com/andresxlp/gohex/cmd/enums"
	"github.com/andresxlp/gohex/internal/infra/handler"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes the creation of the files and folders.",
		Long: `Initializes the creation of the files and folders necessary for the 
hexagonal structure of the microservice`,
		Example: `gohex init`,
		Run: func(cmd *cobra.Command, args []string) {
			module := projectInGOPATH()
			handler.CreateFolderAndFile(module)
			r := handler.ExecuteGoModule()
			if r.Err != nil {
				fmt.Println(r.Err.Error())
			} else {
				fmt.Print(enums.SuccessfullyCreated)
				fmt.Print("\n\n")
				fmt.Println(enums.SuccessfullyAndGoModuleTrue)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}

func projectInGOPATH() string {
	var (
		res   string
		err   error
		yn, _ = regexp.Compile(`^[ynYN]{1}$`)
		n, _  = regexp.Compile(`^[nN]{1}$`)
	)

	validate := func(res string) error {
		if !yn.MatchString(res) {
			exec.Command("clear").Run()
			return errors.New("invalid option only [Y/n]")
		}
		return nil
	}
	template := &promptui.PromptTemplates{
		Prompt:  "{{.}}",
		Valid:   "{{.|green}}",
		Invalid: "{{.|red}}",
		Success: "{{.|bold}}",
	}

	prompt := promptui.Prompt{
		Label:     `your project is inside the GOPATH? - e.g: "$GOPATH/src/github.com/andresxlp/hexa-ms" [y/n] `,
		Validate:  validate,
		Templates: template,
	}
	res, err = prompt.Run()
	if err != nil {
		fmt.Printf("error in prompt %v\n", err.Error())
	}
	module := ""
	exec.Command("clear").Run()
	if n.MatchString(res) {
		prompt = promptui.Prompt{
			Label: "type your module project ",
		}
		res, err = prompt.Run()
		module = res
	}
	return module
}
