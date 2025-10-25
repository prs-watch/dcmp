// Style control module for printing.
package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
)

var (
	HEADER = color.New(color.FgBlack).Add(color.Bold) // Header style
	BEFORE = color.New(color.FgRed)                   // Before style
	AFTER  = color.New(color.FgGreen)                 // After style
)

/*
ApplyColorMode configures the color output mode for printing.
Parameters:
  - colorMode: "auto" (terminal detection), "never" (no color), or "always" (force color)

Returns an error if an invalid color mode is specified.
*/
func ApplyColorMode(colorMode string) error {
	switch colorMode {
	case "auto":
		if !isatty.IsTerminal(os.Stdout.Fd()) {
			color.NoColor = true
		}
	case "never":
		color.NoColor = true
	case "always":
		color.NoColor = false
	default:
		return fmt.Errorf("invalid color mode: %s", colorMode)
	}
	return nil
}

/*
PrintChange outputs Change information detected in C/A/D analysis.
Parameters:
  - bls: Before line start
  - ble: Before line end
  - bct: Before content lines
  - als: After line start
  - ale: After line end
  - act: After content lines
*/
func PrintChange(bls int, ble int, bct []string, als int, ale int, act []string) {

	HEADER.Printf("%d-%dc%d-%d\n", bls, ble, als, ale)
	BEFORE.Printf("%s\n", strings.Join(bct, "\n"))
	fmt.Printf("------------\n")
	AFTER.Printf("%s\n", strings.Join(act, "\n"))
	fmt.Printf("\n")
}

/*
PrintAdd outputs Add information detected in C/A/D analysis.
Parameters:
  - auis: After unprocessed index start
  - auie: After unprocessed index end
  - act: After content lines
*/
func PrintAdd(auis int, auie int, act []string) {
	HEADER.Printf("0a%d-%d\n", auis, auie)
	AFTER.Printf("%s\n", strings.Join(act, "\n"))
	fmt.Printf("\n")
}

/*
PrintDelete outputs Delete information detected in C/A/D analysis.
Parameters:
  - buis: Before unprocessed index start
  - buie: Before unprocessed index end
  - bct: Before content lines
*/
func PrintDelete(buis int, buie int, bct []string) {
	HEADER.Printf("%d-%dd0\n", buis, buie)
	BEFORE.Printf("%s\n", strings.Join(bct, "\n"))
	fmt.Printf("\n")
}

/*
PrintBrief outputs a message when differences are detected with -q, --brief option.
*/
func PrintBrief() {
	fmt.Printf("Files differ\n")
}

/*
PrintIdentical outputs a message when files are identical with -s, --report-identical-files option.
*/
func PrintIdentical() {
	fmt.Printf("Files are identical\n")
}
