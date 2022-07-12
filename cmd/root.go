/*
Copyright © 2022 h0tw4t3r.eth

*/
package cmd

import (
	"os"

	"github.com/mevspace/ABISignatureBruteforcer/bruteforce"
	"github.com/spf13/cobra"
)

var functionName string
var target string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ABISignatureBruteforcer",
	Short: "Bruteforce your way to cheap smart-contract calls",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if err := validateFunctionName(functionName); err != nil {
			panic(err)
		}

		if err := validateTarget(target); err != nil {
			panic(err)
		}

		bruteforce.StartBruteforcing(functionName, target[2:])
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ABISignatureBruteforcer.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&functionName, "name", "s", "", "name of your function that needs to be bruteforced. use * as placeholder [hello_world_*(uint,address)]")
	rootCmd.Flags().StringVarP(&target, "target", "t", "", "target signature, use * as wildcard for any byte. i.e: 0x000000**")
	rootCmd.MarkFlagRequired("name")
	rootCmd.MarkFlagRequired("target")
}
