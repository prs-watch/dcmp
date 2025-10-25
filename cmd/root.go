// Cobra command root module.
package cmd

import (
	"os"

	"github.com/prs-watch/dcmp/internal"
	"github.com/spf13/cobra"
)

var briefFlag bool               // -q, --brief
var identicalFlag bool           // -s, --report-identical-files
var ignoreBlankFlag bool         // -B, --ignore-blank-lines
var ignoreCaseFlag bool          // -i, --ignore-case
var ignoreSpaceFlag bool         // -b, --ignore-space-change
var ignoreAllSpaceFlag bool      // -w, --ignore-all-space
var colorMode string             // --color
var ignoreCrFlag bool            // --strip-trailing-cr
var ignoreMatchingLines []string // -I, --ignore-matching-lines
var expandTabsFlag bool          // -t, --expand-tabs

var rootCmd = &cobra.Command{
	Use:           "dcmp [path] [path] [flags]",
	Short:         "Compare files you pass to dcmp command.",
	Example:       "dcmp hoge.txt fuga.txt",
	SilenceErrors: true,
	SilenceUsage:  true,
	Args:          cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		err := internal.Execute(
			args[0],
			args[1],
			briefFlag,
			identicalFlag,
			ignoreBlankFlag,
			ignoreCaseFlag,
			ignoreSpaceFlag,
			ignoreAllSpaceFlag,
			colorMode,
			ignoreCrFlag,
			ignoreMatchingLines,
			expandTabsFlag,
		)
		if err != nil {
			return err
		}
		return nil
	},
}

/*
Cobra entry point.
*/
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

/*
Define flags to be attached to the command.
*/
func init() {
	// -q, --brief
	rootCmd.Flags().BoolVarP(&briefFlag, "brief", "q", false, "Output to stdout only when file differences exist.")
	// -s, --report-identical-files
	rootCmd.Flags().BoolVarP(&identicalFlag, "report-identical-files", "s", false, "Output to stdout only when files are identical.")
	// -B, --ignore-blank-lines
	rootCmd.Flags().BoolVarP(&ignoreBlankFlag, "ignore-blank-lines", "B", false, "Ignore blank lines when comparing files.")
	// -i, --ignore-case
	rootCmd.Flags().BoolVarP(&ignoreCaseFlag, "ignore-case", "i", false, "Ignore case differences when comparing files.")
	// -b, --ignore-space-change
	rootCmd.Flags().BoolVarP(&ignoreSpaceFlag, "ignore-space-change", "b", false, "Ignore whitespace changes when comparing files.")
	// -w, --ignore-all-space
	rootCmd.Flags().BoolVarP(&ignoreAllSpaceFlag, "ignore-all-space", "w", false, "Ignore all whitespace when comparing files.")
	// --color
	rootCmd.Flags().StringVarP(&colorMode, "color", "", "auto", "Control colored output with auto/always/never. Default is auto.")
	// --strip-trailing-cr
	rootCmd.Flags().BoolVarP(&ignoreCrFlag, "strip-trailing-cr", "", false, "Ignore trailing CR when comparing files.")
	// -I, --ignore-matching-lines
	rootCmd.Flags().StringArrayVarP(&ignoreMatchingLines, "ignore-matching-lines", "I", []string{}, "Ignore lines matching the specified regex when comparing files.")
	// -t, --expand-tabs
	rootCmd.Flags().BoolVarP(&expandTabsFlag, "expand-tabs", "t", false, "Replace tabs with spaces before comparison.")
}
