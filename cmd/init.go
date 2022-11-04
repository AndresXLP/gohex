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
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes the creation of the files and folders.",
		Long: `Initializes the creation of the files and folders necessary for the 
hexagonal structure of the microservice`,
		Example: `gohex init --path /api/gohex --module github.com/andresxlp/gohex`,
		Run: func(cmd *cobra.Command, args []string) {
			handler.StarCreate(&app.Service{})
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			handler.StartGoModule(&app.Service{})
			fmt.Println(`Initial Hexagonal Struct Created Successfully.

Don't forget to set the environment variables:
SERVER_HOST
SERVER_PORT`)

		},
	}
)

func init() {
	initCmd.Flags().StringVarP(&handler.BasePath, "path", "p", "", "REQUIRED: set base path for router. e.g: /api/gohex")
	initCmd.MarkFlagRequired("path")
	initCmd.Flags().StringVarP(&handler.Module, "module", "m", "", "REQUIRED: set your Module Go. e.g: github.com/andresxlp/gohex")
	initCmd.MarkFlagRequired("module")
	rootCmd.AddCommand(initCmd)
}
