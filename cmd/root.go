/*
Copyright © 2022 Andres Puello apuello1025@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Example: `gohex init --path /api/gohex --module github.com/andresxlp/gohex`,
	Args:    cobra.ExactArgs(1),
	Long: `Gohex creates the files and folders to implement the hexagonal
architecture in a microservice with Go.`,
	Use: "gohex",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Use: "no-help", Hidden: true})
}
