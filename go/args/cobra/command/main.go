package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var strp string
var intp int
var boolp bool

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "A simple flags experimentation command, build with Cobra.",
	Run:  flagsFunc,
}

func init() {
	rootCmd.Flags().StringVarP(&strp, "string", "s", "foo", "a string")
	rootCmd.Flags().IntVarP(&intp, "number", "n", 42, "an integer")
	rootCmd.Flags().BoolVarP(&boolp, "boolean", "b", false, "a boolean")
}

func flagsFunc(cmd *cobra.Command, args []string) {
	fmt.Println("string:", strp)
	fmt.Println("int:", intp)
	fmt.Println("boolean:", boolp)
	fmt.Println("args:", args)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
