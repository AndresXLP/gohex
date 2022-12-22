/*
Copyright © 2022 Andres Puello apuello1025@gmail.com
*/
package cmd

import (
	"github.com/spf13/cobra"
)

func GohexCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "gohex",
		Example: `gohex new 'package-name'`,
		Args:    cobra.MinimumNArgs(1),
		Long: `Gohex creates the files and folders to implement the hexagonal
architecture in a microservice with Go.`,
		ValidArgs:         []string{"new"},
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	}
	cmd.AddCommand(NewCmd())
	return cmd
}
