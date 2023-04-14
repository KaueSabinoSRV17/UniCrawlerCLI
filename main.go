package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Uni",
	Short: "Uni - Uma CLI da faculdade",

	Run: func(cmd *cobra.Command, args []string) {
		n1, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("oi")
		}
		n2, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("oi")
		}
		fmt.Println(n1 + n2)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
