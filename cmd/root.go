package cmd

import (
	"os"

	"github.com/prs-watch/dcmp/internal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "dcmp [path] [path]",
	Short:   "Compare files you pass to dcmp command.",
	Example: "dcmp hoge.txt fuga.txt",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := internal.Execute(args[0], args[1])
		if err != nil {
			return err
		}
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// nope
}
