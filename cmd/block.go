package cmd

import (
	"fmt"
	"github.com/areYouLazy/libhosty"
	"github.com/spf13/cobra"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Blocks access to distracting websites",
	Long: `It blocks access to specific websites on your computer, helping improve productivity. You can easily specify FQDN to block distracting sites, such as social media or OTT websites.
For example:
site-blocker block www.example.com`,
	Run: func(cmd *cobra.Command, args []string) {
		hostname := args[0]
		hf, err := libhosty.Init()
		if err != nil {
			panic(err)
		}
		// Check host entry already exists in file or not.
		_, line := hf.GetHostsFileLineByHostname(hostname)
		if line != nil {
			if line.IsCommented {
				hf.UncommentHostsFileLineByRow(line.Number)
			} else {
				fmt.Printf("%s is already blocked\n", hostname)
			}
		} else {
			hf.AddHostsFileLineRaw("127.0.0.1", hostname, "")
			hf.AddEmptyFileLine()
		}
		err = hf.SaveHostsFile()
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)
}
