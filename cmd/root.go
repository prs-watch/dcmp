// Cobraコマンドルートモジュール.
package cmd

import (
	"os"

	"github.com/prs-watch/dcmp/internal"
	"github.com/spf13/cobra"
)

var briefFlag bool       // -q, --brief
var identicalFlag bool   // -s, --report-identical-files
var ignoreBlankFlag bool // -B, --ignore-blank-lines
var ignoreCaseFlag bool  // -i, --ignore-case
var ignoreSpaceFlag bool // -b, --ignore-space-change

var rootCmd = &cobra.Command{
	Use:           "dcmp [path] [path] [flags]",
	Short:         "Compare files you pass to dcmp command.",
	Example:       "dcmp hoge.txt fuga.txt",
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := internal.Execute(args[0], args[1], briefFlag, identicalFlag, ignoreBlankFlag, ignoreCaseFlag, ignoreSpaceFlag)
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
	rootCmd.Flags().BoolVarP(&briefFlag, "brief", "q", false, "ファイル差分が存在する場合のみ標準出力.")
	// -s, --report-identical-files
	rootCmd.Flags().BoolVarP(&identicalFlag, "report-identical-files", "s", false, "ファイルが同一の場合のみ標準出力.")
	// -B, --ignore-blank-lines
	rootCmd.Flags().BoolVarP(&ignoreBlankFlag, "ignore-blank-lines", "B", false, "空行を無視してファイル比較を実行.")
	// -i, --ignore-case
	rootCmd.Flags().BoolVarP(&ignoreCaseFlag, "ignore-case", "i", false, "大文字小文字を無視してファイル比較を実行.")
	// -b, --ignore-space-change
	rootCmd.Flags().BoolVarP(&ignoreSpaceFlag, "ignore-space-change", "b", false, "空白文字を無視してファイル比較を実行.")
}
