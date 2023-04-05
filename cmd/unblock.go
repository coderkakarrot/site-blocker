/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/areYouLazy/libhosty"

	"github.com/spf13/cobra"
)

// unblockCmd represents the unblock command
var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "Restores previously blocked website's access",
	Long: `The "site-blocker unblock" command in Cobra unblocks previously blocked websites on your computer or network. This command can be used to restore access to previously blocked sites or to make adjustments to the block list. It's an essential tool for anyone who needs to manage website access on their system.
For example:
site-blocker block www.example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		hostname := args[0]
		hf, err := libhosty.Init()
		if err != nil {
			panic(err)
		}
		lines := hf.GetHostsFileLinesByHostname(hostname)
		if len(lines) > 0 {
			fmt.Sprintf("here...")
			for _, line := range lines {
				fmt.Println("Line number", line.Number)
				hf.RemoveHostsFileLineByRow(line.Number)
			}
			err = hf.SaveHostsFile()
		} else {
			fmt.Printf("No entry was found for hostname: %s", hostname)
		}
	},
}

func init() {
	rootCmd.AddCommand(unblockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unblockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unblockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
