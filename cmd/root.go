/*
Copyright © 2023 NAME HERE <pratik.coderx@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "site-blocker",
	Short: "Blocks and unblock access to site.",
	Long: `This CLI enables user to block and unblock access to site.
Often time we get distracted by the social media & OTT platform, using this cli tool we can restrict access to site.
This tool just points passed website to 127.0.0.1
For example: 
On running "site-blocker example.com" command
Create below entry it the /etc/hosts
127.0.0.1 example.com www.example.com`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.site-blocker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
