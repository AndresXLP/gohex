/*
Copyright © 2022 Andres Puello apuello1025@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/andresxlp/gohex/internal/app"
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
		Example: `gohex init --path /api/hexa --module github.com/andresxlp/hexa-ms`,
		Run: func(cmd *cobra.Command, args []string) {
			handler.ProgressBarFolderAndFile()
			handler.StarCreate(&app.Service{})
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Print(`Initial Hexagonal Struct Created Successfully.

Don't forget to set the environment variables:
SERVER_HOST
SERVER_PORT

`)
			if mod[0] {
				handler.ProgressBarGoModule()
				handler.StartGoModule(&app.Service{})
			} else {
				fmt.Println("Don't forget execute go mod init and go mod tidy")
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
