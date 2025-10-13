// Cobraコマンドルートモジュール.
package cmd

import (
	"os"

	"github.com/prs-watch/dcmp/internal"
	"github.com/spf13/cobra"
)

var briefFlag bool // -q, --brief

var rootCmd = &cobra.Command{
	Use:           "dcmp [path] [path] [flags]",
	Short:         "Compare files you pass to dcmp command.",
	Example:       "dcmp hoge.txt fuga.txt",
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := internal.Execute(args[0], args[1], briefFlag)
		if err != nil {
			return err
		}
		return nil
	},
}

/*
Cobraエントリポイント.
*/
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

/*
コマンドに付与するフラグを定義する.
*/
func init() {
	// -q, --brief
	rootCmd.Flags().BoolVarP(&briefFlag, "brief", "q", false, "差分の有無のみ出力します.")
}
