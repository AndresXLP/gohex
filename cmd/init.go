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

var (
	mod     []bool
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes the creation of the files and folders.",
		Long: `Initializes the creation of the files and folders necessary for the 
hexagonal structure of the microservice`,
		Example: `gohex init --path /api/hexa --import github.com/andresxlp/hexa-ms`,
		Run: func(cmd *cobra.Command, args []string) {
			if r := handler.VerifyFilesAndFolderExisting(); r.Err != nil {
				fmt.Println(r.Err)
			} else {
				handler.CreateFolderAndFile()
				if mod[0] {
					r := handler.ExecuteGoModule()
					if r.Err != nil {
						fmt.Println(r.Err.Error())
					} else {
						fmt.Print(enums.SuccessfullyCreated)
						fmt.Print("\n\n")
						fmt.Println(enums.SuccessfullyAndGoModuleTrue)
					}
				} else {
					fmt.Print(enums.SuccessfullyCreated)
					fmt.Print("\n\n")
					fmt.Println(enums.SuccessfullyAndGoModuleFalse)
				}
			}
		},
	}
)

func init() {
	initCmd.Flags().StringVarP(&handler.BasePath, "path", "p", "", "REQUIRED: set base path for router. e.g: /api/hexa")
	initCmd.MarkFlagRequired("path")
	initCmd.Flags().StringVarP(&handler.Module, "import", "i", "", "REQUIRED: set your Package Module in the imports. e.g: github.com/andresxlp/hexa-ms")
	initCmd.MarkFlagRequired("import")
	initCmd.Flags().BoolSliceVar(&mod, "go_module", []bool{true}, "[opcional] true / false - Run go mod init & go mod tidy.")
	rootCmd.AddCommand(initCmd)
}
