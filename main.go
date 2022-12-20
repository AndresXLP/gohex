/*
Copyright © 2022 Andres Puello apuello1025@gmail.com

*/
package main

import (
	"github.com/andresxlp/gohex/cmd"
	"github.com/spf13/cobra"
)

func main() {
	cobra.CheckErr(cmd.GohexCmd().Execute())
}
