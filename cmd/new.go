/*
Copyright © 2022 Andres Puello apuello1025@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/andresxlp/gohex/cmd/enums"
	"github.com/andresxlp/gohex/internal/infra/handler"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new",
		Short: "Initializes the creation of the files and folders.",
		Long: `Initializes the creation of the files and folders necessary for the 
hexagonal structure of the microservice`,
		Example: `gohex new 'package-name'`,
		Run: func(command *cobra.Command, args []string) {
			module := args[0]
			handler.CreateFolderAndFile(module)
			r := handler.ExecuteGoModule(module)
			if r.Err != nil {
				fmt.Println(r.Err.Error())
			} else {
				fmt.Print(enums.SuccessfullyCreated)
				fmt.Print("\n\n")
				fmt.Println(enums.SuccessfullyAndGoModuleTrue)
			}
		},
	}
}
